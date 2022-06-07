package deposit

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
	deposit := npool.Deposit{
		AccountID:     uuid.New().String(),
		DepositAmount: float64(1.1),
	}

	schema, err := New(context.Background(), nil)
	assert.Nil(t, err)

	info, err := schema.Create(context.Background(), &deposit)
	if assert.Nil(t, err) {
		if assert.NotEqual(t, info.ID, uuid.UUID{}.String()) {
			deposit.ID = info.ID
		}
		deposit.CreatedAt = info.CreatedAt
		assert.Equal(t, info, &deposit)
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)
	deposit.DepositAmount = float64(9.9)

	info, err = schema.Update(context.Background(), &deposit)
	if assert.Nil(t, err) {
		assert.Equal(t, info, &deposit)
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)

	info, err = schema.Row(context.Background(), uuid.MustParse(info.ID))
	if assert.Nil(t, err) {
		assert.Equal(t, info, &deposit)
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)

	infos, total, err := schema.Rows(context.Background(),
		cruder.NewConds().WithCond(constant.FieldID, cruder.EQ, info.ID),
		0, 0)
	if assert.Nil(t, err) {
		assert.Equal(t, total, 1)
		assert.Equal(t, infos[0], &deposit)
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)

	info, err = schema.RowOnly(context.Background(),
		cruder.NewConds().WithCond(constant.FieldID, cruder.EQ, info.ID))
	if assert.Nil(t, err) {
		assert.Equal(t, info, &deposit)
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

	deposit.DepositAmount = 888.88

	info, err = schema.Update(context.Background(), &deposit)

	if assert.Nil(t, err) {
		assert.Equal(t, info, &deposit)
	}

	schema, err = New(context.Background(), nil)
	assert.Nil(t, err)

	deposit1 := &npool.Deposit{
		AccountID:     uuid.New().String(),
		DepositAmount: float64(11111),
	}
	deposit2 := &npool.Deposit{
		AccountID:     uuid.New().String(),
		DepositAmount: float64(222),
	}

	infos, err = schema.CreateBulk(context.Background(), []*npool.Deposit{deposit1, deposit2})
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
		assert.Equal(t, info, &deposit)
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
