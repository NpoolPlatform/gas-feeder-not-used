package coingas

import (
	"context"
	"fmt"
	"time"

	constant "github.com/NpoolPlatform/gas-feeder/pkg/const"
	"github.com/NpoolPlatform/go-service-framework/pkg/price"

	"github.com/NpoolPlatform/gas-feeder/pkg/db"
	"github.com/NpoolPlatform/gas-feeder/pkg/db/ent"
	"github.com/NpoolPlatform/gas-feeder/pkg/db/ent/coingas"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/gasfeeder"

	"github.com/google/uuid"
)

type CoinGas struct {
	*db.Entity
}

func New(ctx context.Context, tx *ent.Tx) (*CoinGas, error) {
	e, err := db.NewEntity(ctx, tx)
	if err != nil {
		return nil, fmt.Errorf("fail create entity: %v", err)
	}

	return &CoinGas{
		Entity: e,
	}, nil
}

func (s *CoinGas) rowToObject(row *ent.CoinGas) *npool.CoinGas {
	return &npool.CoinGas{
		ID:                  row.ID.String(),
		GasCoinTypeID:       row.GasCoinTypeID.String(),
		CoinTypeID:          row.CoinTypeID.String(),
		FeedingTID:          row.FeedingTid.String(),
		DepositThresholdLow: price.DBPriceToVisualPrice(row.DepositThresholdLow),
		DepositAmount:       price.DBPriceToVisualPrice(row.DepositAmount),
		OnlineScale:         row.OnlineScale,
	}
}

func (s *CoinGas) Create(ctx context.Context, in *npool.CoinGas) (*npool.CoinGas, error) {
	var info *ent.CoinGas
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		info, err = s.Tx.CoinGas.Create().
			SetCoinTypeID(uuid.MustParse(in.CoinTypeID)).
			SetGasCoinTypeID(uuid.MustParse(in.GasCoinTypeID)).
			SetDepositThresholdLow(price.VisualPriceToDBPrice(in.GetDepositThresholdLow())).
			SetDepositAmount(price.VisualPriceToDBPrice(in.GetDepositAmount())).
			SetOnlineScale(in.GetOnlineScale()).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("fail create CoinGas: %v", err)
	}

	return s.rowToObject(info), nil
}

func (s *CoinGas) CreateBulk(ctx context.Context, in []*npool.CoinGas) ([]*npool.CoinGas, error) {
	rows := []*ent.CoinGas{}
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		bulk := make([]*ent.CoinGasCreate, len(in))
		for i, info := range in {
			bulk[i] = s.Tx.CoinGas.Create().
				SetCoinTypeID(uuid.MustParse(info.CoinTypeID)).
				SetGasCoinTypeID(uuid.MustParse(info.GasCoinTypeID)).
				SetDepositThresholdLow(price.VisualPriceToDBPrice(info.DepositThresholdLow)).
				SetDepositAmount(price.VisualPriceToDBPrice(info.DepositAmount)).
				SetOnlineScale(info.GetOnlineScale())
		}
		rows, err = s.Tx.CoinGas.CreateBulk(bulk...).Save(_ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("fail create CoinGass: %v", err)
	}

	infos := []*npool.CoinGas{}
	for _, row := range rows {
		infos = append(infos, s.rowToObject(row))
	}

	return infos, nil
}

func (s *CoinGas) Row(ctx context.Context, id uuid.UUID) (*npool.CoinGas, error) {
	var info *ent.CoinGas
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		info, err = s.Tx.CoinGas.Query().Where(coingas.ID(id)).Only(_ctx)
		return err
	})

	if err != nil {
		return nil, fmt.Errorf("fail get CoinGas: %v", err)
	}

	return s.rowToObject(info), nil
}

func (s *CoinGas) Update(ctx context.Context, in *npool.CoinGas) (*npool.CoinGas, error) {
	var info *ent.CoinGas
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		info, err = s.Tx.CoinGas.UpdateOneID(uuid.MustParse(in.GetID())).
			SetCoinTypeID(uuid.MustParse(in.GetCoinTypeID())).
			SetGasCoinTypeID(uuid.MustParse(in.GetGasCoinTypeID())).
			SetFeedingTid(uuid.MustParse(in.GetFeedingTID())).
			SetDepositThresholdLow(price.VisualPriceToDBPrice(in.GetDepositThresholdLow())).
			SetDepositAmount(price.VisualPriceToDBPrice(in.GetDepositAmount())).
			SetOnlineScale(in.GetOnlineScale()).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("fail update CoinGas: %v", err)
	}

	return s.rowToObject(info), nil
}

func (s *CoinGas) Rows(ctx context.Context, conds cruder.Conds, offset, limit int) ([]*npool.CoinGas, int, error) {
	rows := []*ent.CoinGas{}
	var total int

	err := db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		stm, err := s.queryFromConds(conds)
		if err != nil {
			return fmt.Errorf("fail construct stm: %v", err)
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return fmt.Errorf("fail count CoinGas: %v", err)
		}

		rows, err = stm.Order(ent.Desc(coingas.FieldUpdatedAt)).Offset(offset).Limit(limit).All(_ctx)
		if err != nil {
			return fmt.Errorf("fail query CoinGas: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, 0, fmt.Errorf("fail get CoinGas: %v", err)
	}

	infos := []*npool.CoinGas{}
	for _, row := range rows {
		infos = append(infos, s.rowToObject(row))
	}

	return infos, total, nil
}

//nolint
func (s *CoinGas) queryFromConds(conds cruder.Conds) (*ent.CoinGasQuery, error) {
	stm := s.Tx.CoinGas.Query()
	for k, v := range conds {
		switch k {
		case constant.FieldID:
			id, err := cruder.AnyTypeUUID(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid ID: %v", err)
			}
			stm = stm.Where(coingas.ID(id))
		case constant.FieldGasCoinTypeID:
			id, err := cruder.AnyTypeUUID(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid GasCoinTypeID: %v", err)
			}
			stm = stm.Where(coingas.GasCoinTypeID(id))
		case constant.FieldCoinTypeID:
			cointypeid, err := cruder.AnyTypeUUID(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid CoinTypeID: %v", err)
			}
			stm = stm.Where(coingas.CoinTypeID(cointypeid))
		case constant.FieldDepositThresholdLow:
			_depositThresholdLow, err := cruder.AnyTypeFloat64(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid DepositThresholdLow: %v", err)
			}
			depositThresholdLow := price.VisualPriceToDBPrice(_depositThresholdLow)

			switch v.Op {
			case cruder.EQ:
				stm = stm.Where(coingas.DepositThresholdLowEQ(depositThresholdLow))
			case cruder.GT:
				stm = stm.Where(coingas.DepositThresholdLowGT(depositThresholdLow))
			case cruder.LT:
				stm = stm.Where(coingas.DepositThresholdLowLT(depositThresholdLow))
			}
		case constant.FieldDepositAmount:
			_depositAmount, err := cruder.AnyTypeFloat64(v.Val)
			if err != nil {
				return nil, fmt.Errorf("invalid DepositAmount: %v", err)
			}
			depositAmount := price.VisualPriceToDBPrice(_depositAmount)

			switch v.Op {
			case cruder.EQ:
				stm = stm.Where(coingas.DepositAmountEQ(depositAmount))
			case cruder.GT:
				stm = stm.Where(coingas.DepositAmountGT(depositAmount))
			case cruder.LT:
				stm = stm.Where(coingas.DepositAmountLT(depositAmount))
			}
		default:
			return nil, fmt.Errorf("invalid CoinGas field")
		}
	}

	return stm, nil
}

func (s *CoinGas) RowOnly(ctx context.Context, conds cruder.Conds) (*npool.CoinGas, error) {
	var info *ent.CoinGas

	err := db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		stm, err := s.queryFromConds(conds)
		if err != nil {
			return fmt.Errorf("fail construct stm: %v", err)
		}

		info, err = stm.Only(_ctx)
		if err != nil {
			return fmt.Errorf("fail query CoinGas: %v", err)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get CoinGas: %v", err)
	}

	return s.rowToObject(info), nil
}

func (s *CoinGas) Count(ctx context.Context, conds cruder.Conds) (uint32, error) {
	var err error
	var total int

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		stm, err := s.queryFromConds(conds)
		if err != nil {
			return fmt.Errorf("fail construct stm: %v", err)
		}

		total, err = stm.Count(_ctx)
		if err != nil {
			return fmt.Errorf("fail check CoinGass: %v", err)
		}

		return nil
	})
	if err != nil {
		return 0, fmt.Errorf("fail count CoinGass: %v", err)
	}

	return uint32(total), nil
}

func (s *CoinGas) Exist(ctx context.Context, id uuid.UUID) (bool, error) {
	var err error
	exist := false

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		exist, err = s.Tx.CoinGas.Query().Where(coingas.ID(id)).Exist(_ctx)
		return err
	})
	if err != nil {
		return false, fmt.Errorf("fail check CoinGas: %v", err)
	}

	return exist, nil
}

func (s *CoinGas) ExistConds(ctx context.Context, conds cruder.Conds) (bool, error) {
	var err error
	exist := false

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		stm, err := s.queryFromConds(conds)
		if err != nil {
			return fmt.Errorf("fail construct stm: %v", err)
		}

		exist, err = stm.Exist(_ctx)
		if err != nil {
			return fmt.Errorf("fail check congases: %v", err)
		}

		return nil
	})
	if err != nil {
		return false, fmt.Errorf("fail check congases: %v", err)
	}

	return exist, nil
}

func (s *CoinGas) Delete(ctx context.Context, id uuid.UUID) (*npool.CoinGas, error) {
	var info *ent.CoinGas
	var err error

	err = db.WithTx(ctx, s.Tx, func(_ctx context.Context) error {
		info, err = s.Tx.CoinGas.UpdateOneID(id).
			SetDeletedAt(uint32(time.Now().Unix())).
			Save(_ctx)
		return err
	})
	if err != nil {
		return nil, fmt.Errorf("fail delete CoinGas: %v", err)
	}

	return s.rowToObject(info), nil
}
