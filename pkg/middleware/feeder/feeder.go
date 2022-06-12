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

type account struct {
	coinTypeID  string
	accountID   string
	amountScale int
}

type Feeder struct {
	coins        []*coininfopb.CoinInfo
	accounts     []*account
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
	lowBalance := 0
	transferred := 0

	for _, acc := range f.accounts {
		if acc.coinTypeID != gas.CoinTypeID {
			continue
		}

		to, ok := f.addresses[acc.accountID]
		if !ok {
			account, err := billingcli.GetAccount(ctx, acc.accountID)
			if err != nil || account == nil {
				invalid++
				continue
			}
			to = account.Address
			f.addresses[acc.accountID] = to
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
			return fmt.Errorf("fail check coin %v | %v balance: %v", coin.Name, to, err)
		}
		if balance.Balance <= coin.ReservedAmount {
			lowBalance++
			continue
		}

		coin, err = f.GetCoin(gas.GasCoinTypeID)
		if err != nil || coin == nil {
			return fmt.Errorf("fail get gas coin: %v %v", err, gas.GasCoinTypeID)
		}

		exist := false
		err = withDepositCRUD(ctx, func(schema *depositcrud.Deposit) error {
			exist, err = schema.ExistConds(ctx, cruder.NewConds().
				WithCond(constant.FieldAccountID, cruder.EQ, acc.accountID))
			return err
		})
		if err != nil {
			return fmt.Errorf("fail exist deposit: %v", err)
		}

		if exist {
			balance, err = sphinxproxycli.GetBalance(ctx, &sphinxproxypb.GetBalanceRequest{
				Name:    coin.Name,
				Address: to,
			})
			if err != nil || balance == nil {
				logger.Sugar().Errorf("fail check gas %v | %v | %v balance: %v", coin.Name, to, acc.accountID, err)
				continue
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
			f.addresses[gasAccountID] = from
		}

		balance, err = sphinxproxycli.GetBalance(ctx, &sphinxproxypb.GetBalanceRequest{
			Name:    coin.Name,
			Address: from,
		})
		if err != nil || balance == nil {
			return fmt.Errorf("fail check balance: %v", err)
		}

		amount := gas.DepositAmount * float64(acc.amountScale)
		if balance.Balance <= coin.ReservedAmount+amount {
			insufficient++
			continue
		}

		for {
			err = accountlock.Lock(gasAccountID)
			if err != nil {
				logger.Sugar().Infof("wait for %v gas account %v: %v", coin.Name, gasAccountID, err)
				time.Sleep(5 * time.Minute)
				continue
			}
			break
		}

		transaction, err := billingcli.CreateTransaction(ctx, &billingpb.CoinAccountTransaction{
			AppID:              uuid.UUID{}.String(),
			UserID:             uuid.UUID{}.String(),
			GoodID:             uuid.UUID{}.String(),
			FromAddressID:      gasAccountID,
			ToAddressID:        acc.accountID,
			CoinTypeID:         gas.GasCoinTypeID,
			Amount:             amount,
			Message:            fmt.Sprintf("transfer gas at %v scale %v", time.Now(), acc.amountScale),
			ChainTransactionID: uuid.New().String(),
		})
		if err != nil {
			return fmt.Errorf("fail create transaction: %v", err)
		}

		transferred++

		err = withDepositCRUD(ctx, func(schema *depositcrud.Deposit) error {
			_, err := schema.Create(ctx, &npool.Deposit{
				AccountID:     acc.accountID,
				TransactionID: transaction.ID,
				DepositAmount: amount,
			})
			return err
		})
		if err != nil {
			return fmt.Errorf("fail create deposit: %v", err)
		}
	}

	logger.Sugar().Infof("feed gas invalid %v ignore %v insufficient %v low balance %v transferred %v coin %v gas coin %v",
		invalid, ignore, insufficient, lowBalance, transferred, gas.CoinTypeID, gas.GasCoinTypeID)
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

func (f *Feeder) update(ctx context.Context) error {
	coins, err := coininfocli.GetCoinInfos(ctx, cruder.NewFilterConds())
	if err != nil {
		return fmt.Errorf("fail get coininfos: %v", err)
	}

	gases := []*npool.CoinGas{}
	err = withCoinGasCRUD(ctx, func(schema *coingascrud.CoinGas) error {
		gases, _, err = schema.Rows(ctx, cruder.NewConds(), 0, 0)
		return err
	})
	if err != nil {
		return fmt.Errorf("fail get coin gases: %v", err)
	}

	settings, err := billingcli.GetCoinSettings(ctx)
	if err != nil {
		return fmt.Errorf("fail get coin settings: %v", err)
	}

	f.coins = coins
	f.gases = gases
	f.coinsettings = settings

	return nil
}

const (
	PaymentAmountScale = 1
	OnlineAmountScale  = 20
)

func (f *Feeder) paymentFeeder(ctx context.Context) {
	payments, err := billingcli.GetGoodPayments(ctx, cruder.NewFilterConds())
	if err != nil {
		logger.Sugar().Errorf("fail get good payments: %v", err)
		return
	}

	accounts := []*account{}
	for _, payment := range payments {
		accounts = append(accounts, &account{
			accountID:   payment.AccountID,
			coinTypeID:  payment.PaymentCoinTypeID,
			amountScale: PaymentAmountScale,
		})
	}

	f.accounts = accounts

	err = f.FeedAll(ctx)
	if err != nil {
		logger.Sugar().Errorf("fail feed gases: %v", err)
	}
}

func (f *Feeder) onlineFeeder(ctx context.Context) {
	accounts := []*account{}
	for _, setting := range f.coinsettings {
		accounts = append(accounts, &account{
			accountID:   setting.UserOnlineAccountID,
			coinTypeID:  setting.CoinTypeID,
			amountScale: OnlineAmountScale,
		})
	}

	f.accounts = accounts

	err := f.FeedAll(ctx)
	if err != nil {
		logger.Sugar().Errorf("fail feed gases: %v", err)
	}
}

const (
	PaymentGasFeedInterval = 4 * time.Hour
	HotGasFeedInterval     = 10 * time.Minute
)

func Run() {
	paymentTicker := time.NewTicker(PaymentGasFeedInterval)
	onlineTicker := time.NewTicker(HotGasFeedInterval)

	ctx := context.Background()
	_feeder := &Feeder{
		addresses: map[string]string{},
	}

	err := _feeder.update(ctx)
	if err != nil {
		logger.Sugar().Errorf("fail update feeder: %v", err)
		return
	}

	_feeder.paymentFeeder(ctx)

	for {
		select {
		case <-paymentTicker.C:
			err := _feeder.update(ctx)
			if err != nil {
				logger.Sugar().Errorf("fail update feeder: %v", err)
				continue
			}
			_feeder.paymentFeeder(ctx)
		case <-onlineTicker.C:
			err := _feeder.update(ctx)
			if err != nil {
				logger.Sugar().Errorf("fail update feeder: %v", err)
				continue
			}
			_feeder.onlineFeeder(ctx)
		}
	}
}
