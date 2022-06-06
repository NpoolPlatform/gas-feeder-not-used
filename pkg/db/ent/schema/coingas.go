package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/gas-feeder/pkg/db/mixin"
	"github.com/google/uuid"
)

// CoinGas holds the schema definition for the CoinGas entity.
type CoinGas struct {
	ent.Schema
}

func (CoinGas) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the CoinGas.
func (CoinGas) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("coin_type_id", uuid.UUID{}).Unique(),
		field.UUID("gas_coin_type_id", uuid.UUID{}),
		field.Uint64("deposit_threshold"),
	}
}
