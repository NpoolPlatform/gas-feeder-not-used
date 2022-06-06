package client

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	constant "github.com/NpoolPlatform/gas-feeder/pkg/const"
	"github.com/NpoolPlatform/libent-cruder/pkg/cruder"

	// nolint:nolintlint
	testinit "github.com/NpoolPlatform/gas-feeder/pkg/test-init"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/google/uuid"
)

func TestClient(t *testing.T) {
	runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION"))
	if err == nil && runByGithubAction {
		return
	}

	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}

	_, err = GetGasFeederInfoOnly(context.Background(),
		cruder.NewFilterConds().
			WithCond(constant.FieldID, cruder.EQ, structpb.NewStringValue(uuid.UUID{}.String())))
	if err != nil {
		t.Fatal(err)
	}
	// Here won't pass test due to we always test with localhost
}
