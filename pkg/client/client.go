package client

import (
	"context"
	"fmt"

	grpc2 "github.com/NpoolPlatform/go-service-framework/pkg/grpc"

	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/gasfeeder"

	constant "github.com/NpoolPlatform/gas-feeder/pkg/message/const"
)

func do(ctx context.Context, fn func(_ctx context.Context, cli npool.GasFeederClient) (cruder.Any, error)) (cruder.Any, error) {
	conn, err := grpc2.GetGRPCConn(constant.ServiceName, grpc2.GRPCTAG)
	if err != nil {
		return nil, fmt.Errorf("fail get gas feeder connection: %v", err)
	}
	defer conn.Close()

	cli := npool.NewGasFeederClient(conn)

	return fn(ctx, cli)
}

func GetCoinGasOnly(ctx context.Context, conds cruder.FilterConds) (*npool.CoinGas, error) {
	info, err := do(ctx, func(_ctx context.Context, cli npool.GasFeederClient) (cruder.Any, error) {
		// DO RPC CALL HERE WITH conds PARAMETER
		return &npool.CoinGas{}, nil
	})
	if err != nil {
		return nil, fmt.Errorf("fail get gas feeder: %v", err)
	}
	return info.(*npool.CoinGas), nil
}
