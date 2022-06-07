package feeder

import (
	"context"
	"time"

	logger "github.com/NpoolPlatform/go-service-framework/pkg/logger"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	billingcli "github.com/NpoolPlatform/cloud-hashing-billing/pkg/client"
	billingpb "github.com/NpoolPlatform/message/npool/cloud-hashing-billing"

	coininfopb "github.com/NpoolPlatform/message/npool/coininfo"
	coininfocli "github.com/NpoolPlatform/sphinx-coininfo/pkg/client"

	coingascrud "github.com/NpoolPlatform/gas-feeder/pkg/crud/coingas"
	npool "github.com/NpoolPlatform/message/npool/gasfeeder"
)

func withCRUD(ctx context.Context, fn func(schema *coingascrud.CoinGas) error) error {
	schema, err := coingascrud.New(ctx, nil)
	if err != nil {
		return err
	}
	return fn(schema)
}

type Feeder struct {
	coins    []*coininfopb.CoinInfo
	accounts []*billingpb.GoodPayment
	gases    []*npool.CoinGas
}

func (f *Feeder) FeedAll(ctx context.Context) error {
	return nil
}

func Run(ctx context.Context) error {
	ticker := time.NewTicker(12 * time.Hour)
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	for range ticker.C {
		coins, err := coininfocli.GetCoinInfos(ctx, cruder.NewFilterConds())
		if err != nil {
			logger.Sugar().Errorf("fail get coininfos: %v", err)
			continue
		}

		accounts, err := billingcli.GetGoodPayments(ctx, cruder.NewFilterConds())
		if err != nil {
			logger.Sugar().Errorf("fail get good payments: %v", err)
			continue
		}

		gases := []*npool.CoinGas{}
		err = withCRUD(ctx, func(schema *coingascrud.CoinGas) error {
			gases, _, err = schema.Rows(ctx, cruder.NewConds(), 0, 0)
			return err
		})
		if err != nil {
			logger.Sugar().Errorf("fail get coin gases: %v", err)
			continue
		}

		_feeder := &Feeder{
			coins:    coins,
			accounts: accounts,
			gases:    gases,
		}
		err = _feeder.FeedAll(ctx)
		if err != nil {
			logger.Sugar().Errorf("fail feed gases: %v", err)
		}
	}

	// Check balance
	// If low than threshold, wait gas account release, lock gas account, create transaction
	// Record to deposit
	return nil
}
