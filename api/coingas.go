//nolint
package api

import (
	"context"
	"fmt"

	crud "github.com/NpoolPlatform/gas-feeder/pkg/crud/coingas"
	constant "github.com/NpoolPlatform/gas-feeder/pkg/db/ent/coingas"
	ccoin "github.com/NpoolPlatform/gas-feeder/pkg/message/const"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/gasfeeder"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func checkFeildsInCoinGas(info *npool.CoinGas) error {
	_, err := uuid.Parse(info.GetCoinTypeID())
	if err != nil {
		logger.Sugar().Errorf("parse CoinTypeID: %s invalid", info.GetCoinTypeID())
		return status.Error(codes.InvalidArgument, "CoinTypeID invalid")
	}

	_, err = uuid.Parse(info.GetGasCoinTypeID())
	if err != nil {
		logger.Sugar().Errorf("parse GasCoinTypeID: %s invalid", info.GetCoinTypeID())
		return status.Error(codes.InvalidArgument, "GasCoinTypeID invalid")
	}

	return nil
}

func (s *Server) CreateCoinGas(ctx context.Context, in *npool.CreateCoinGasRequest) (*npool.CreateCoinGasResponse, error) {
	info := in.GetInfo()
	if err := checkFeildsInCoinGas(info); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.CreateCoinGasResponse{}, status.Error(codes.Internal, err.Error())
	}
	cInfo, err := schema.Create(ctx, &npool.CoinGas{
		CoinTypeID:       info.GetCoinTypeID(),
		GasCoinTypeID:    info.GetGasCoinTypeID(),
		DepositThreshold: info.GetDepositThreshold(),
	})
	if err != nil {
		logger.Sugar().Errorf("fail create CoinGas error %v", err)
		return &npool.CreateCoinGasResponse{}, status.Error(codes.Internal, "internal server error")
	}

	return &npool.CreateCoinGasResponse{
		Info: cInfo,
	}, nil
}

func (s *Server) UpdateCoinGas(ctx context.Context, in *npool.UpdateCoinGasRequest) (*npool.UpdateCoinGasResponse, error) {
	info := in.GetInfo()
	if err := checkFeildsInCoinGas(info); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.UpdateCoinGasResponse{}, status.Error(codes.Internal, err.Error())
	}
	updateInfo, err := schema.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update CoinGas: %v", err)
		return &npool.UpdateCoinGasResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateCoinGasResponse{
		Info: updateInfo,
	}, nil
}

func (s *Server) GetCoinGas(ctx context.Context, in *npool.GetCoinGasRequest) (*npool.GetCoinGasResponse, error) {
	_, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorf("parse ID: %s invalid", in.GetID())
		return nil, status.Error(codes.InvalidArgument, "ID invalid")
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.GetCoinGasResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.Row(ctx, uuid.MustParse(in.GetID()))

	if err != nil {
		logger.Sugar().Errorf("fail get CoinGas: %v", err)
		return &npool.GetCoinGasResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinGasResponse{
		Info: info,
	}, nil
}

func coinGasCondsToConds(conds cruder.FilterConds) (cruder.Conds, error) {
	newConds := cruder.NewConds()

	for k, v := range conds {
		switch v.Op {
		case cruder.EQ:
		case cruder.GT:
		case cruder.LT:
		case cruder.LIKE:
		default:
			return nil, fmt.Errorf("invalid filter condition op")
		}

		switch k {
		case constant.FieldID:
			fallthrough //nolint
		case constant.FieldCoinTypeID:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		case constant.FieldGasCoinTypeID:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		case constant.FieldDepositThreshold:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetNumberValue())
		default:
			return nil, fmt.Errorf("invalid CoinGas field")
		}
	}
	return newConds, nil
}

func (s *Server) GetCoinGasOnly(ctx context.Context, in *npool.GetCoinGasOnlyRequest) (*npool.GetCoinGasOnlyResponse, error) {
	newConds, err := coinGasCondsToConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("invalid  Conds fields: %v", err)
		return &npool.GetCoinGasOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.GetCoinGasOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.RowOnly(ctx, newConds)
	if err != nil {
		logger.Sugar().Errorf("fail get CoinGas: %v", err)
		return &npool.GetCoinGasOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinGasOnlyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetCoinGases(ctx context.Context, in *npool.GetCoinGasesRequest) (*npool.GetCoinGasesResponse, error) {
	newConds, err := coinGasCondsToConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("invalid  Conds: %v", err)
		return &npool.GetCoinGasesResponse{}, status.Error(codes.Internal, err.Error())
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	infos, total, err := schema.Rows(ctx, newConds, int(in.GetOffset()), int(in.GetLimit()))
	if err != nil {
		logger.Sugar().Errorf("fail get CoinGas: %v", err)
		return &npool.GetCoinGasesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetCoinGasesResponse{
		Infos: infos,
		Total: int32(total),
	}, nil
}

func (s *Server) ExistCoinGas(ctx context.Context, in *npool.ExistCoinGasRequest) (*npool.ExistCoinGasResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistCoinGasResponse{}, fmt.Errorf("invalid coingas ID: %v", err)
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.ExistCoinGasResponse{}, status.Error(codes.Internal, err.Error())
	}

	exist, err := schema.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check coingas: %v", err)
		return &npool.ExistCoinGasResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCoinGasResponse{
		Result: exist,
	}, nil
}

func (s *Server) ExistCoinGasConds(ctx context.Context, in *npool.ExistCoinGasCondsRequest) (*npool.ExistCoinGasCondsResponse, error) {
	conds, err := coinGasCondsToConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("invalid coingas fields: %v", err)
		return &npool.ExistCoinGasCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	if len(conds) == 0 {
		logger.Sugar().Errorf("empty coingas fields: %v", err)
		return &npool.ExistCoinGasCondsResponse{}, status.Error(codes.Internal, "empty coingas fields")
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.ExistCoinGasCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	exist, err := schema.ExistConds(ctx, conds)
	if err != nil {
		logger.Sugar().Errorf("fail check coingas: %v", err)
		return &npool.ExistCoinGasCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistCoinGasCondsResponse{
		Result: exist,
	}, nil
}

func (s *Server) DeleteCoinGas(ctx context.Context, in *npool.DeleteCoinGasRequest) (*npool.DeleteCoinGasResponse, error) {
	_, err := uuid.Parse(in.GetID())
	if err != nil {
		logger.Sugar().Errorf("parse ID: %s invalid", in.GetID())
		return nil, status.Error(codes.InvalidArgument, "ID invalid")
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}
	deletedInfo, err := schema.Delete(ctx, uuid.MustParse(in.GetID()))
	if err != nil {
		logger.Sugar().Errorf("delete CoinGas: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteCoinGasResponse{
		Info: deletedInfo,
	}, nil
}
