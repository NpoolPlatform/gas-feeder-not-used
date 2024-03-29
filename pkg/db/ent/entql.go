// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/gas-feeder/pkg/db/ent/coingas"
	"github.com/NpoolPlatform/gas-feeder/pkg/db/ent/deposit"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 2)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   coingas.Table,
			Columns: coingas.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: coingas.FieldID,
			},
		},
		Type: "CoinGas",
		Fields: map[string]*sqlgraph.FieldSpec{
			coingas.FieldCreatedAt:           {Type: field.TypeUint32, Column: coingas.FieldCreatedAt},
			coingas.FieldUpdatedAt:           {Type: field.TypeUint32, Column: coingas.FieldUpdatedAt},
			coingas.FieldDeletedAt:           {Type: field.TypeUint32, Column: coingas.FieldDeletedAt},
			coingas.FieldCoinTypeID:          {Type: field.TypeUUID, Column: coingas.FieldCoinTypeID},
			coingas.FieldGasCoinTypeID:       {Type: field.TypeUUID, Column: coingas.FieldGasCoinTypeID},
			coingas.FieldFeedingTid:          {Type: field.TypeUUID, Column: coingas.FieldFeedingTid},
			coingas.FieldDepositThresholdLow: {Type: field.TypeUint64, Column: coingas.FieldDepositThresholdLow},
			coingas.FieldDepositAmount:       {Type: field.TypeUint64, Column: coingas.FieldDepositAmount},
			coingas.FieldOnlineScale:         {Type: field.TypeInt32, Column: coingas.FieldOnlineScale},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   deposit.Table,
			Columns: deposit.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: deposit.FieldID,
			},
		},
		Type: "Deposit",
		Fields: map[string]*sqlgraph.FieldSpec{
			deposit.FieldCreatedAt:     {Type: field.TypeUint32, Column: deposit.FieldCreatedAt},
			deposit.FieldUpdatedAt:     {Type: field.TypeUint32, Column: deposit.FieldUpdatedAt},
			deposit.FieldDeletedAt:     {Type: field.TypeUint32, Column: deposit.FieldDeletedAt},
			deposit.FieldAccountID:     {Type: field.TypeUUID, Column: deposit.FieldAccountID},
			deposit.FieldTransactionID: {Type: field.TypeUUID, Column: deposit.FieldTransactionID},
			deposit.FieldDepositAmount: {Type: field.TypeUint64, Column: deposit.FieldDepositAmount},
		},
	}
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (cgq *CoinGasQuery) addPredicate(pred func(s *sql.Selector)) {
	cgq.predicates = append(cgq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the CoinGasQuery builder.
func (cgq *CoinGasQuery) Filter() *CoinGasFilter {
	return &CoinGasFilter{cgq}
}

// addPredicate implements the predicateAdder interface.
func (m *CoinGasMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the CoinGasMutation builder.
func (m *CoinGasMutation) Filter() *CoinGasFilter {
	return &CoinGasFilter{m}
}

// CoinGasFilter provides a generic filtering capability at runtime for CoinGasQuery.
type CoinGasFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *CoinGasFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *CoinGasFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(coingas.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *CoinGasFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(coingas.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *CoinGasFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(coingas.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *CoinGasFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(coingas.FieldDeletedAt))
}

// WhereCoinTypeID applies the entql [16]byte predicate on the coin_type_id field.
func (f *CoinGasFilter) WhereCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(coingas.FieldCoinTypeID))
}

// WhereGasCoinTypeID applies the entql [16]byte predicate on the gas_coin_type_id field.
func (f *CoinGasFilter) WhereGasCoinTypeID(p entql.ValueP) {
	f.Where(p.Field(coingas.FieldGasCoinTypeID))
}

// WhereFeedingTid applies the entql [16]byte predicate on the feeding_tid field.
func (f *CoinGasFilter) WhereFeedingTid(p entql.ValueP) {
	f.Where(p.Field(coingas.FieldFeedingTid))
}

// WhereDepositThresholdLow applies the entql uint64 predicate on the deposit_threshold_low field.
func (f *CoinGasFilter) WhereDepositThresholdLow(p entql.Uint64P) {
	f.Where(p.Field(coingas.FieldDepositThresholdLow))
}

// WhereDepositAmount applies the entql uint64 predicate on the deposit_amount field.
func (f *CoinGasFilter) WhereDepositAmount(p entql.Uint64P) {
	f.Where(p.Field(coingas.FieldDepositAmount))
}

// WhereOnlineScale applies the entql int32 predicate on the online_scale field.
func (f *CoinGasFilter) WhereOnlineScale(p entql.Int32P) {
	f.Where(p.Field(coingas.FieldOnlineScale))
}

// addPredicate implements the predicateAdder interface.
func (dq *DepositQuery) addPredicate(pred func(s *sql.Selector)) {
	dq.predicates = append(dq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the DepositQuery builder.
func (dq *DepositQuery) Filter() *DepositFilter {
	return &DepositFilter{dq}
}

// addPredicate implements the predicateAdder interface.
func (m *DepositMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the DepositMutation builder.
func (m *DepositMutation) Filter() *DepositFilter {
	return &DepositFilter{m}
}

// DepositFilter provides a generic filtering capability at runtime for DepositQuery.
type DepositFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *DepositFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql [16]byte predicate on the id field.
func (f *DepositFilter) WhereID(p entql.ValueP) {
	f.Where(p.Field(deposit.FieldID))
}

// WhereCreatedAt applies the entql uint32 predicate on the created_at field.
func (f *DepositFilter) WhereCreatedAt(p entql.Uint32P) {
	f.Where(p.Field(deposit.FieldCreatedAt))
}

// WhereUpdatedAt applies the entql uint32 predicate on the updated_at field.
func (f *DepositFilter) WhereUpdatedAt(p entql.Uint32P) {
	f.Where(p.Field(deposit.FieldUpdatedAt))
}

// WhereDeletedAt applies the entql uint32 predicate on the deleted_at field.
func (f *DepositFilter) WhereDeletedAt(p entql.Uint32P) {
	f.Where(p.Field(deposit.FieldDeletedAt))
}

// WhereAccountID applies the entql [16]byte predicate on the account_id field.
func (f *DepositFilter) WhereAccountID(p entql.ValueP) {
	f.Where(p.Field(deposit.FieldAccountID))
}

// WhereTransactionID applies the entql [16]byte predicate on the transaction_id field.
func (f *DepositFilter) WhereTransactionID(p entql.ValueP) {
	f.Where(p.Field(deposit.FieldTransactionID))
}

// WhereDepositAmount applies the entql uint64 predicate on the deposit_amount field.
func (f *DepositFilter) WhereDepositAmount(p entql.Uint64P) {
	f.Where(p.Field(deposit.FieldDepositAmount))
}
