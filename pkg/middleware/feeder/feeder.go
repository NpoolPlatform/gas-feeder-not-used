package feeder

import (
	"context"
	"fmt"
	"time"

	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	billingcli "github.com/NpoolPlatform/cloud-hashing-billing/pkg/client"
	billingconst "github.com/NpoolPlatform/cloud-hashing-billing/pkg/const"
	billingpb "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	coininfopb "github.com/NpoolPlatform/message/npool/coininfo"
	coininfocli "github.com/NpoolPlatform/sphinx-coininfo/pkg/client"

	coingascrud "github.com/NpoolPlatform/gas-feeder/pkg/crud/coingas"
	depositcrud "github.com/NpoolPlatform/gas-feeder/pkg/crud/deposit"
	npool "github.com/NpoolPlatform/message/npool/gasfeeder"

	sphinxproxypb "github.com/NpoolPlatform/message/npool/sphinxproxy"
	sphinxproxycli "github.com/NpoolPlatform/sphinx-proxy/pkg/client"

	depositcli "github.com/NpoolPlatform/account-middleware/pkg/client/deposit"
	depositpb "github.com/NpoolPlatform/message/npool/account/mw/v1/deposit"

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
	coinTypeID    string
	accountID     string
	address       string
	onlineAccount bool
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
	ignore := 0
	insufficient := 0
	lowBalance := 0
	transferred := 0

	for _, acc := range f.accounts {
		if acc.coinTypeID != gas.CoinTypeID {
			continue
		}

		to := acc.address

		coin, err := f.GetCoin(gas.CoinTypeID)
		if err != nil || coin == nil {
			return fmt.Errorf("fail get coin: %v %v", err, gas.CoinTypeID)
		}

		toCoinBalance := 0.0

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

		toCoinBalance = balance.Balance

		coin, err = f.GetCoin(gas.GasCoinTypeID)
		if err != nil || coin == nil {
			return fmt.Errorf("fail get gas coin: %v %v", err, gas.GasCoinTypeID)
		}

		toGasBalance := 0.0
		balance, err = sphinxproxycli.GetBalance(ctx, &sphinxproxypb.GetBalanceRequest{
			Name:    coin.Name,
			Address: to,
		})
		if err != nil || balance == nil {
			logger.Sugar().Errorf("fail check gas %v | %v | %v balance: %v", coin.Name, to, acc.accountID, err)
			continue
		}

		toGasBalance = balance.Balance

		if gas.DepositThresholdLow < balance.Balance {
			ignore++
			continue
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

		scale := int32(1)
		if acc.onlineAccount && gas.OnlineScale > 0 {
			scale = gas.OnlineScale
		}

		amount := gas.DepositAmount * float64(scale)

		logger.Sugar().Infow(
			"FeedAll",
			"From", from,
			"To", to,
			"CoinName", coin.Name,
			"GasProviderBalance", balance.Balance,
			"Scale", scale,
			"DepositThresholdLow", gas.DepositThresholdLow,
			"DepositAmount", gas.DepositAmount,
			"TargetGasAmount", amount,
			"GasBalance", toGasBalance,
			"CoinBalance", toCoinBalance,
			"AccountID", acc.accountID,
		)

		if balance.Balance <= amount {
			insufficient++
			continue
		}

		transaction, err := billingcli.CreateTransaction(ctx, &billingpb.CoinAccountTransaction{
			AppID:              uuid.UUID{}.String(),
			UserID:             uuid.UUID{}.String(),
			GoodID:             uuid.UUID{}.String(),
			FromAddressID:      gasAccountID,
			ToAddressID:        acc.accountID,
			CoinTypeID:         gas.GasCoinTypeID,
			Amount:             amount,
			Message:            fmt.Sprintf("transfer gas at %v scale %v", time.Now(), scale),
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

		gas.FeedingTID = transaction.ID
		err = withCoinGasCRUD(ctx, func(schema *coingascrud.CoinGas) error {
			_, err = schema.Update(ctx, gas)
			return err
		})
		if err != nil {
			return fmt.Errorf("fail add feeding tid: %v", err)
		}

		break
	}

	logger.Sugar().Infof("feed gas ignore %v insufficient %v low balance %v transferred %v coin %v gas coin %v",
		ignore, insufficient, lowBalance, transferred, gas.CoinTypeID, gas.GasCoinTypeID)
	return nil
}

func (f *Feeder) FeedAll(ctx context.Context) error {
	for _, gas := range f.gases {
		tx, err := billingcli.GetTransaction(ctx, gas.FeedingTID)
		if err != nil {
			logger.Sugar().Errorw("FeedAll", "CoinTypeID", gas.CoinTypeID)
			continue
		}
		if tx != nil {
			switch tx.State {
			case billingconst.CoinTransactionStateSuccessful:
			case billingconst.CoinTransactionStateFail:
			default:
				continue
			}
		}

		err = f.FeedGas(ctx, gas)
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

func (f *Feeder) getAddress(ctx context.Context, accountID string) (string, error) {
	to, ok := f.addresses[accountID]
	if !ok {
		account, err := billingcli.GetAccount(ctx, accountID)
		if err != nil {
			return "", err
		}
		if account == nil {
			return "", fmt.Errorf("invalid account")
		}
		to = account.Address
		f.addresses[accountID] = to
	}

	return to, nil
}

func (f *Feeder) paymentFeeder(ctx context.Context) {
	logger.Sugar().Infow("paymentFeeder", "Start", "...")

	payments, err := billingcli.GetGoodPayments(ctx, cruder.NewFilterConds())
	if err != nil {
		logger.Sugar().Errorf("fail get good payments: %v", err)
		return
	}

	accounts := []*account{}
	for _, payment := range payments {
		address, err := f.getAddress(ctx, payment.AccountID)
		if err != nil {
			logger.Sugar().Errorw("paymentFeeder", "AccountID", payment.AccountID, "error", err)
			return
		}

		accounts = append(accounts, &account{
			accountID:     payment.AccountID,
			coinTypeID:    payment.PaymentCoinTypeID,
			address:       address,
			onlineAccount: false,
		})
	}

	f.accounts = accounts

	err = f.FeedAll(ctx)
	if err != nil {
		logger.Sugar().Errorf("fail feed gases: %v", err)
	}

	logger.Sugar().Infow("paymentFeeder", "Done", "...")
}

func (f *Feeder) onlineFeeder(ctx context.Context) {
	logger.Sugar().Infow("onlineFeeder", "Start", "...")

	accounts := []*account{}
	for _, setting := range f.coinsettings {
		address, err := f.getAddress(ctx, setting.UserOnlineAccountID)
		if err != nil {
			logger.Sugar().Errorw("onlineFeeder", "AccountID", setting.UserOnlineAccountID, "error", err)
			continue
		}

		accounts = append(accounts, &account{
			accountID:     setting.UserOnlineAccountID,
			coinTypeID:    setting.CoinTypeID,
			address:       address,
			onlineAccount: true,
		})
	}

	f.accounts = accounts

	err := f.FeedAll(ctx)
	if err != nil {
		logger.Sugar().Errorf("fail feed gases: %v", err)
	}

	logger.Sugar().Infow("onlineFeeder", "Done", "...")
}

func (f *Feeder) depositFeeder(ctx context.Context) {
	accounts := []*account{}
	offset := int32(0)
	limit := int32(1000) //nolint

	logger.Sugar().Infow("depositFeeder", "Start", "...")

	for {
		accs, err := depositcli.GetAccounts(ctx, &depositpb.Conds{}, offset, limit)
		if err != nil {
			logger.Sugar().Infow("depositFeeder", "error", err)
			return
		}
		if len(accs) == 0 {
			logger.Sugar().Infow("depositFeeder", "Done", "...")
			return
		}

		for _, acc := range accs {
			accounts = append(accounts, &account{
				accountID:     acc.AccountID,
				coinTypeID:    acc.CoinTypeID,
				address:       acc.Address,
				onlineAccount: false,
			})
		}

		f.accounts = accounts

		err = f.FeedAll(ctx)
		if err != nil {
			logger.Sugar().Errorf("fail feed gases: %v", err)
		}

		offset += limit
	}
}

const (
	GasFeedInterval = 1 * time.Minute
)

func Run() {
	ticker := time.NewTicker(GasFeedInterval)

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
	_feeder.onlineFeeder(ctx)
	_feeder.depositFeeder(ctx)

	for range ticker.C {
		err := _feeder.update(ctx)
		if err != nil {
			logger.Sugar().Errorf("fail update feeder: %v", err)
			continue
		}
		_feeder.paymentFeeder(ctx)

		err = _feeder.update(ctx)
		if err != nil {
			logger.Sugar().Errorf("fail update feeder: %v", err)
			continue
		}
		_feeder.onlineFeeder(ctx)

		err = _feeder.update(ctx)
		if err != nil {
			logger.Sugar().Errorf("fail update feeder: %v", err)
			continue
		}
		_feeder.depositFeeder(ctx)
	}
}
