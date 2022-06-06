package schema

import "entgo.io/ent"

// CoinGas holds the schema definition for the CoinGas entity.
type CoinGas struct {
	ent.Schema
}

// Fields of the CoinGas.
func (CoinGas) Fields() []ent.Field {
	return nil
}

// Edges of the CoinGas.
func (CoinGas) Edges() []ent.Edge {
	return nil
}
