package coingas

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"testing"

	testinit "github.com/NpoolPlatform/gas-feeder/pkg/test-init"
	cruder "github.com/NpoolPlatform/libent-cruder/pkg/cruder"
	npool "github.com/NpoolPlatform/message/npool/gasfeeder"

	constant "github.com/NpoolPlatform/gas-feeder/pkg/const"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	if err := testinit.Init(); err != nil {
		fmt.Printf("cannot init test stub: %v\n", err)
	}
}

//nolint
func TestCRUD(t *testing.T) {
	coingas := npool.CoinGas{
		CoinTypeID:          uuid.New().String(),
		GasCoinTypeID:       uuid.New().String(),
		DepositThresholdLow: float64(100.1),
		DepositAmount:       float64(1.1),
		CollectingTID:       uuid.UUID{}.String(),
	}

	schema, err := New(context.Background(), nil)
	assert.Nil(t, err)

	info, err := schema.Create(context.Background(), &coingas)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			coingas.ID = info.ID
		}
		assert.Equal(t, info, &coingas)
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)
	coingas.DepositAmount = float64(9.9)

	info, err = schema.Update(context.Background(), &coingas)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &coingas)
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)

	info, err = schema.Row(context.Background(), uuid.MustParse(info.ID))
	if assert.Nil(t, err) {
		assert.Equal(t, info, &coingas)
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)

	infos, total, err := schema.Rows(context.Background(),
		cruder.NewConds().WithCond(constant.FieldID, cruder.EQ, info.ID),
		0, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, 1)
		assert.Equal(t, infos[0], &coingas)
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)

	info, err = schema.RowOnly(context.Background(),
		cruder.NewConds().WithCond(constant.FieldID, cruder.EQ, info.ID))
	if assert.Nil(t, err) {
		assert.Equal(t, info, &coingas)
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)

	exist, err := schema.ExistConds(context.Background(),
		cruder.NewConds().WithCond(constant.FieldID, cruder.EQ, info.ID),
	)
	if assert.Nil(t, err) {
		assert.Equal(t, exist, true)
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)

	coingas.DepositThresholdLow = 888.88

	info, err = schema.Update(context.Background(), &coingas)

	if assert.Nil(t, err) {
		assert.Equal(t, info, &coingas)
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)

	coingas1 := &npool.CoinGas{
		CoinTypeID:          uuid.New().String(),
		GasCoinTypeID:       uuid.New().String(),
		DepositThresholdLow: float64(11111),
		DepositAmount:       float64(11111),
	}
	coingas2 := &npool.CoinGas{
		CoinTypeID:          uuid.New().String(),
		GasCoinTypeID:       uuid.New().String(),
		DepositThresholdLow: float64(22222),
		DepositAmount:       float64(222),
	}

	infos, err = schema.CreateBulk(context.Background(), []*npool.CoinGas{coingas1, coingas2})
	if assert.Nil(t, err) {
		assert.Equal(t, len(infos), 2)
		assert.NotEqual(t, infos[0].ID, uuid.UUID{}.String())
		assert.NotEqual(t, infos[1].ID, uuid.UUID{}.String())
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)

	count, err := schema.Count(context.Background(),
		cruder.NewConds().WithCond(constant.FieldID, cruder.EQ, info.ID),
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, uint32(1))
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)

	info, err = schema.Delete(context.Background(), uuid.MustParse(info.ID))
	if assert.Nil(t, err) {
		assert.Equal(t, info, &coingas)
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)

	count, err = schema.Count(context.Background(),
		cruder.NewConds().WithCond(constant.FieldID, cruder.EQ, info.ID),
	)
	if assert.Nil(t, err) {
		assert.Equal(t, count, uint32(0))
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)

	_, err = schema.Row(context.Background(), uuid.MustParse(info.ID))
	assert.NotNil(t, err)
}
