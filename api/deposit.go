//nolint
package api

import (
	"context"
	"fmt"

<<<<<<< HEAD
	constant "github.com/NpoolPlatform/gas-feeder/pkg/const"
	crud "github.com/NpoolPlatform/gas-feeder/pkg/crud/deposit"
=======
	crud "github.com/NpoolPlatform/gas-feeder/pkg/crud/deposit"
	constant "github.com/NpoolPlatform/gas-feeder/pkg/db/ent/deposit"
>>>>>>> api done
	ccoin "github.com/NpoolPlatform/gas-feeder/pkg/message/const"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/gasfeeder"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func checkFeildsInDeposit(info *npool.Deposit) error {
	_, err := uuid.Parse(info.GetAccountID())
	if err != nil {
		logger.Sugar().Errorf("parse AccountID: %s invalid", info.GetAccountID())
		return status.Error(codes.InvalidArgument, "AccountID invalid")
	}
	return nil
}

func (s *Server) CreateDeposit(ctx context.Context, in *npool.CreateDepositRequest) (*npool.CreateDepositResponse, error) {
	info := in.GetInfo()
	if err := checkFeildsInDeposit(info); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.CreateDepositResponse{}, status.Error(codes.Internal, err.Error())
	}
	cInfo, err := schema.Create(ctx, &npool.Deposit{
		AccountID:     info.GetAccountID(),
		DepositAmount: info.GetDepositAmount(),
	})
	if err != nil {
		logger.Sugar().Errorf("fail create Deposit error %v", err)
		return &npool.CreateDepositResponse{}, status.Error(codes.Internal, "internal server error")
	}

	return &npool.CreateDepositResponse{
		Info: cInfo,
	}, nil
}

func (s *Server) UpdateDeposit(ctx context.Context, in *npool.UpdateDepositRequest) (*npool.UpdateDepositResponse, error) {
	info := in.GetInfo()
	if err := checkFeildsInDeposit(info); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.UpdateDepositResponse{}, status.Error(codes.Internal, err.Error())
	}
	updateInfo, err := schema.Update(ctx, in.GetInfo())
	if err != nil {
		logger.Sugar().Errorf("fail update CoinGas: %v", err)
		return &npool.UpdateDepositResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.UpdateDepositResponse{
		Info: updateInfo,
	}, nil
}

func (s *Server) GetDeposit(ctx context.Context, in *npool.GetDepositRequest) (*npool.GetDepositResponse, error) {
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
		return &npool.GetDepositResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.Row(ctx, uuid.MustParse(in.GetID()))

	if err != nil {
		logger.Sugar().Errorf("fail get CoinGas: %v", err)
		return &npool.GetDepositResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDepositResponse{
		Info: info,
	}, nil
}

func depositCondsToConds(conds cruder.FilterConds) (cruder.Conds, error) {
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
		case constant.FieldAccountID:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetStringValue())
		case constant.FieldDepositAmount:
			newConds = newConds.WithCond(k, v.Op, v.Val.GetNumberValue())
		default:
			return nil, fmt.Errorf("invalid Deposit field")
		}
	}
	return newConds, nil
}

func (s *Server) GetDepositOnly(ctx context.Context, in *npool.GetDepositOnlyRequest) (*npool.GetDepositOnlyResponse, error) {
	newConds, err := depositCondsToConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("invalid  Conds fields: %v", err)
		return &npool.GetDepositOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	ctx, cancel := context.WithTimeout(ctx, ccoin.GrpcTimeout)
	defer cancel()

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.GetDepositOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	info, err := schema.RowOnly(ctx, newConds)
	if err != nil {
		logger.Sugar().Errorf("fail get Deposit: %v", err)
		return &npool.GetDepositOnlyResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDepositOnlyResponse{
		Info: info,
	}, nil
}

func (s *Server) GetDeposites(ctx context.Context, in *npool.GetDepositesRequest) (*npool.GetDepositesResponse, error) {
	newConds, err := depositCondsToConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("invalid  Conds: %v", err)
		return &npool.GetDepositesResponse{}, status.Error(codes.Internal, err.Error())
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
		logger.Sugar().Errorf("fail get Deposit: %v", err)
		return &npool.GetDepositesResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.GetDepositesResponse{
		Infos: infos,
		Total: int32(total),
	}, nil
}

func (s *Server) ExistDeposit(ctx context.Context, in *npool.ExistDepositRequest) (*npool.ExistDepositResponse, error) {
	id, err := uuid.Parse(in.GetID())
	if err != nil {
		return &npool.ExistDepositResponse{}, fmt.Errorf("invalid deposit ID: %v", err)
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.ExistDepositResponse{}, status.Error(codes.Internal, err.Error())
	}

	exist, err := schema.Exist(ctx, id)
	if err != nil {
		logger.Sugar().Errorf("fail check deposit: %v", err)
		return &npool.ExistDepositResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistDepositResponse{
		Result: exist,
	}, nil
}

func (s *Server) ExistDepositConds(ctx context.Context, in *npool.ExistDepositCondsRequest) (*npool.ExistDepositCondsResponse, error) {
	conds, err := depositCondsToConds(in.GetConds())
	if err != nil {
		logger.Sugar().Errorf("invalid deposit fields: %v", err)
		return &npool.ExistDepositCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	if len(conds) == 0 {
		logger.Sugar().Errorf("empty deposit fields: %v", err)
<<<<<<< HEAD
<<<<<<< HEAD
		return &npool.ExistDepositCondsResponse{}, status.Error(codes.Internal, "empty deposit fields")
=======
		return &npool.ExistDepositCondsResponse{}, status.Error(codes.Internal, "empty stock fields")
>>>>>>> api done
=======
		return &npool.ExistDepositCondsResponse{}, status.Error(codes.Internal, "empty deposit fields")
>>>>>>> fix version_test.go
	}

	schema, err := crud.New(ctx, nil)
	if err != nil {
		logger.Sugar().Errorf("fail create schema entity: %v", err)
		return &npool.ExistDepositCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	exist, err := schema.ExistConds(ctx, conds)
	if err != nil {
		logger.Sugar().Errorf("fail check deposit: %v", err)
		return &npool.ExistDepositCondsResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &npool.ExistDepositCondsResponse{
		Result: exist,
	}, nil
}

func (s *Server) DeleteDeposit(ctx context.Context, in *npool.DeleteDepositRequest) (*npool.DeleteDepositResponse, error) {
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
		logger.Sugar().Errorf("delete Deposit: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &npool.DeleteDepositResponse{
		Info: deletedInfo,
	}, nil
}
