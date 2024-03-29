// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CoinGasColumns holds the columns for the "coin_gas" table.
	CoinGasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "coin_type_id", Type: field.TypeUUID, Unique: true},
		{Name: "gas_coin_type_id", Type: field.TypeUUID},
		{Name: "feeding_tid", Type: field.TypeUUID},
		{Name: "deposit_threshold_low", Type: field.TypeUint64},
		{Name: "deposit_amount", Type: field.TypeUint64},
		{Name: "online_scale", Type: field.TypeInt32, Default: 1},
	}
	// CoinGasTable holds the schema information for the "coin_gas" table.
	CoinGasTable = &schema.Table{
		Name:       "coin_gas",
		Columns:    CoinGasColumns,
		PrimaryKey: []*schema.Column{CoinGasColumns[0]},
	}
	// DepositsColumns holds the columns for the "deposits" table.
	DepositsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "created_at", Type: field.TypeUint32},
		{Name: "updated_at", Type: field.TypeUint32},
		{Name: "deleted_at", Type: field.TypeUint32},
		{Name: "account_id", Type: field.TypeUUID},
		{Name: "transaction_id", Type: field.TypeUUID},
		{Name: "deposit_amount", Type: field.TypeUint64},
	}
	// DepositsTable holds the schema information for the "deposits" table.
	DepositsTable = &schema.Table{
		Name:       "deposits",
		Columns:    DepositsColumns,
		PrimaryKey: []*schema.Column{DepositsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CoinGasTable,
		DepositsTable,
	}
)

func init() {
}
