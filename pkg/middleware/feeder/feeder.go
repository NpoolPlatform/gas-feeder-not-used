package feeder

import (
	"context"
	"fmt"
	"time"

	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	billingcli "github.com/NpoolPlatform/cloud-hashing-billing/pkg/client"
	billingpb "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	coininfopb "github.com/NpoolPlatform/message/npool/coininfo"
	coininfocli "github.com/NpoolPlatform/sphinx-coininfo/pkg/client"

	coingascrud "github.com/NpoolPlatform/gas-feeder/pkg/crud/coingas"
	depositcrud "github.com/NpoolPlatform/gas-feeder/pkg/crud/deposit"
	npool "github.com/NpoolPlatform/message/npool/gasfeeder"

	sphinxproxypb "github.com/NpoolPlatform/message/npool/sphinxproxy"
	sphinxproxycli "github.com/NpoolPlatform/sphinx-proxy/pkg/client"

	accountlock "github.com/NpoolPlatform/staker-manager/pkg/middleware/account"

	constant "github.com/NpoolPlatform/gas-feeder/pkg/const"

	"github.com/google/uuid"
)

func withCoinGasCRUD(ctx context.Context, fn func(schema *coingascrud.CoinGas) error) error {
	schema, err := coingascrud.New(ctx, nil)
	if err != nil {
		return err
	}
	return fn(schema)
}

func withDepositCRUD(ctx context.Context, fn func(schema *depositcrud.Deposit) error) error {
	schema, err := depositcrud.New(ctx, nil)
	if err != nil {
		return err
	}
	return fn(schema)
}

type Feeder struct {
	coins        []*coininfopb.CoinInfo
	accounts     []*billingpb.GoodPayment
	gases        []*npool.CoinGas
	coinsettings []*billingpb.CoinSetting
	addresses    map[string]string
}

func (f *Feeder) GetCoin(coinTypeID string) (*coininfopb.CoinInfo, error) {
	for _, coin := range f.coins {
		if coinTypeID == coin.ID {
			return coin, nil
		}
	}
	return nil, fmt.Errorf("invalid coin type id")
}

func (f *Feeder) GetGasAccountID(coinTypeID string) (string, error) {
	for _, setting := range f.coinsettings {
		if setting.CoinTypeID == coinTypeID {
			return setting.GasProviderAccountID, nil
		}
	}
	return "", fmt.Errorf("invalid coin setting")
}

//nolint
func (f *Feeder) FeedGas(ctx context.Context, gas *npool.CoinGas) error {
	invalid := 0
	ignore := 0
	insufficient := 0

	for _, acc := range f.accounts {
		if acc.PaymentCoinTypeID != gas.CoinTypeID {
			continue
		}

		to, ok := f.addresses[acc.AccountID]
		if !ok {
			account, err := billingcli.GetAccount(ctx, acc.AccountID)
			if err != nil || account == nil {
				invalid++
				continue
			}
			to = account.Address
			f.addresses[acc.AccountID] = to
		}

		coin, err := f.GetCoin(gas.CoinTypeID)
		if err != nil || coin == nil {
			return fmt.Errorf("fail get coin: %v %v", err, gas.CoinTypeID)
		}

		balance, err := sphinxproxycli.GetBalance(ctx, &sphinxproxypb.GetBalanceRequest{
			Name:    coin.Name,
			Address: to,
		})
		if err != nil || balance == nil {
			return fmt.Errorf("fail check balance: %v", err)
		}
		if balance.Balance <= coin.ReservedAmount {
			continue
		}

		coin, err = f.GetCoin(gas.GasCoinTypeID)
		if err != nil || coin == nil {
			return fmt.Errorf("fail get gas coin: %v %v", err, gas.GasCoinTypeID)
		}

		exist := false
		err = withDepositCRUD(ctx, func(schema *depositcrud.Deposit) error {
			exist, err = schema.ExistConds(ctx, cruder.NewConds().
				WithCond(constant.FieldAccountID, cruder.EQ, acc.AccountID))
			return err
		})
		if err != nil {
			return fmt.Errorf("fail create deposit: %v", err)
		}

		if exist {
			balance, err = sphinxproxycli.GetBalance(ctx, &sphinxproxypb.GetBalanceRequest{
				Name:    coin.Name,
				Address: to,
			})
			if err != nil || balance == nil {
				return fmt.Errorf("fail check balance: %v", err)
			}

			if gas.DepositThresholdLow < balance.Balance {
				ignore++
				continue
			}
		}

		gasAccountID, err := f.GetGasAccountID(gas.GasCoinTypeID)
		if err != nil {
			return fmt.Errorf("invalid gas coin type: %v: %v :%v", err, gasAccountID, gas.GasCoinTypeID)
		}

		from, ok := f.addresses[gasAccountID]
		if !ok {
			account, err := billingcli.GetAccount(ctx, gasAccountID)
			if err != nil || account == nil {
				return fmt.Errorf("fail get gas account %v: %v", gasAccountID, err)
			}
			from = account.Address
			f.addresses[acc.AccountID] = from
		}

		balance, err = sphinxproxycli.GetBalance(ctx, &sphinxproxypb.GetBalanceRequest{
			Name:    coin.Name,
			Address: from,
		})
		if err != nil || balance == nil {
			return fmt.Errorf("fail check balance: %v", err)
		}

		if balance.Balance <= coin.ReservedAmount+gas.DepositAmount {
			insufficient++
			continue
		}

		for {
			err = accountlock.Lock(gasAccountID)
			if err != nil {
				logger.Sugar().Infof("wait for %v gas account %v: %v", coin.Name, gasAccountID, err)
				time.Sleep(30 * time.Second)
				continue
			}
			break
		}

		transaction, err := billingcli.CreateTransaction(ctx, &billingpb.CoinAccountTransaction{
			AppID:              uuid.UUID{}.String(),
			UserID:             uuid.UUID{}.String(),
			GoodID:             uuid.UUID{}.String(),
			FromAddressID:      gasAccountID,
			ToAddressID:        acc.AccountID,
			CoinTypeID:         gas.GasCoinTypeID,
			Amount:             gas.DepositAmount,
			Message:            fmt.Sprintf("transfer gas at %v", time.Now()),
			ChainTransactionID: uuid.New().String(),
		})
		if err != nil {
			return fmt.Errorf("fail create transaction: %v", err)
		}

		err = withDepositCRUD(ctx, func(schema *depositcrud.Deposit) error {
			_, err := schema.Create(ctx, &npool.Deposit{
				AccountID:     acc.GetAccountID(),
				TransactionID: transaction.GetID(),
				DepositAmount: gas.GetDepositAmount(),
			})
			return err
		})
		if err != nil {
			return fmt.Errorf("fail create deposit: %v", err)
		}
	}

	logger.Sugar().Infof("feed gas invalid %v ignore %v insufficient %v", invalid, ignore, insufficient)
	return nil
}

func (f *Feeder) FeedAll(ctx context.Context) error {
	for _, gas := range f.gases {
		err := f.FeedGas(ctx, gas)
		if err != nil {
			logger.Sugar().Errorf("fail feed gas: %v", err)
		}
	}
	return nil
}

const (
	GasFeedInterval = 4 * time.Hour
)

func feeder(ctx context.Context) {
	coins, err := coininfocli.GetCoinInfos(ctx, cruder.NewFilterConds())
	if err != nil {
		logger.Sugar().Errorf("fail get coininfos: %v", err)
		return
	}

	accounts, err := billingcli.GetGoodPayments(ctx, cruder.NewFilterConds())
	if err != nil {
		logger.Sugar().Errorf("fail get good payments: %v", err)
		return
	}

	gases := []*npool.CoinGas{}
	err = withCoinGasCRUD(ctx, func(schema *coingascrud.CoinGas) error {
		gases, _, err = schema.Rows(ctx, cruder.NewConds(), 0, 0)
		return err
	})
	if err != nil {
		logger.Sugar().Errorf("fail get coin gases: %v", err)
		return
	}

	settings, err := billingcli.GetCoinSettings(ctx)
	if err != nil {
		logger.Sugar().Errorf("fail get coin settings: %v", err)
		return
	}

	_feeder := &Feeder{
		coins:        coins,
		accounts:     accounts,
		gases:        gases,
		coinsettings: settings,
		addresses:    map[string]string{},
	}
	err = _feeder.FeedAll(ctx)
	if err != nil {
		logger.Sugar().Errorf("fail feed gases: %v: %v", err, accounts)
	}
}

func Run() {
	ticker := time.NewTicker(GasFeedInterval)
	ctx := context.Background()

	feeder(ctx)
	for range ticker.C {
		feeder(ctx)
	}
}
