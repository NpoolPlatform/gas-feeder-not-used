package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/gas-feeder/pkg/db/mixin"
	"github.com/google/uuid"
)

// Deposit holds the schema definition for the Deposit entity.
type Deposit struct {
	ent.Schema
}

func (Deposit) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Fields of the Deposit.
func (Deposit) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Unique(),
		field.UUID("account_id", uuid.UUID{}),
		field.UUID("transaction_id", uuid.UUID{}),
		field.Uint64("deposit_amount"),
	}
}
