// Code generated by entc, DO NOT EDIT.

package coingas

import (
	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/gas-feeder/pkg/db/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v uint32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v uint32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v uint32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// CoinTypeID applies equality check predicate on the "coin_type_id" field. It's identical to CoinTypeIDEQ.
func CoinTypeID(v uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// GasCoinTypeID applies equality check predicate on the "gas_coin_type_id" field. It's identical to GasCoinTypeIDEQ.
func GasCoinTypeID(v uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGasCoinTypeID), v))
	})
}

// DepositThreshold applies equality check predicate on the "deposit_threshold" field. It's identical to DepositThresholdEQ.
func DepositThreshold(v uint64) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDepositThreshold), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v uint32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v uint32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...uint32) predicate.CoinGas {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinGas(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...uint32) predicate.CoinGas {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinGas(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v uint32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v uint32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v uint32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v uint32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v uint32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v uint32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...uint32) predicate.CoinGas {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinGas(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...uint32) predicate.CoinGas {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinGas(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v uint32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v uint32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v uint32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v uint32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v uint32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v uint32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...uint32) predicate.CoinGas {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinGas(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...uint32) predicate.CoinGas {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinGas(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDeletedAt), v...))
	})
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v uint32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v uint32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v uint32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDeletedAt), v))
	})
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v uint32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDeletedAt), v))
	})
}

// CoinTypeIDEQ applies the EQ predicate on the "coin_type_id" field.
func CoinTypeIDEQ(v uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDNEQ applies the NEQ predicate on the "coin_type_id" field.
func CoinTypeIDNEQ(v uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDIn applies the In predicate on the "coin_type_id" field.
func CoinTypeIDIn(vs ...uuid.UUID) predicate.CoinGas {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinGas(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDNotIn applies the NotIn predicate on the "coin_type_id" field.
func CoinTypeIDNotIn(vs ...uuid.UUID) predicate.CoinGas {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinGas(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCoinTypeID), v...))
	})
}

// CoinTypeIDGT applies the GT predicate on the "coin_type_id" field.
func CoinTypeIDGT(v uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDGTE applies the GTE predicate on the "coin_type_id" field.
func CoinTypeIDGTE(v uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLT applies the LT predicate on the "coin_type_id" field.
func CoinTypeIDLT(v uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCoinTypeID), v))
	})
}

// CoinTypeIDLTE applies the LTE predicate on the "coin_type_id" field.
func CoinTypeIDLTE(v uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCoinTypeID), v))
	})
}

// GasCoinTypeIDEQ applies the EQ predicate on the "gas_coin_type_id" field.
func GasCoinTypeIDEQ(v uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGasCoinTypeID), v))
	})
}

// GasCoinTypeIDNEQ applies the NEQ predicate on the "gas_coin_type_id" field.
func GasCoinTypeIDNEQ(v uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGasCoinTypeID), v))
	})
}

// GasCoinTypeIDIn applies the In predicate on the "gas_coin_type_id" field.
func GasCoinTypeIDIn(vs ...uuid.UUID) predicate.CoinGas {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinGas(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGasCoinTypeID), v...))
	})
}

// GasCoinTypeIDNotIn applies the NotIn predicate on the "gas_coin_type_id" field.
func GasCoinTypeIDNotIn(vs ...uuid.UUID) predicate.CoinGas {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinGas(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGasCoinTypeID), v...))
	})
}

// GasCoinTypeIDGT applies the GT predicate on the "gas_coin_type_id" field.
func GasCoinTypeIDGT(v uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGasCoinTypeID), v))
	})
}

// GasCoinTypeIDGTE applies the GTE predicate on the "gas_coin_type_id" field.
func GasCoinTypeIDGTE(v uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGasCoinTypeID), v))
	})
}

// GasCoinTypeIDLT applies the LT predicate on the "gas_coin_type_id" field.
func GasCoinTypeIDLT(v uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGasCoinTypeID), v))
	})
}

// GasCoinTypeIDLTE applies the LTE predicate on the "gas_coin_type_id" field.
func GasCoinTypeIDLTE(v uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGasCoinTypeID), v))
	})
}

// DepositThresholdEQ applies the EQ predicate on the "deposit_threshold" field.
func DepositThresholdEQ(v uint64) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDepositThreshold), v))
	})
}

// DepositThresholdNEQ applies the NEQ predicate on the "deposit_threshold" field.
func DepositThresholdNEQ(v uint64) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDepositThreshold), v))
	})
}

// DepositThresholdIn applies the In predicate on the "deposit_threshold" field.
func DepositThresholdIn(vs ...uint64) predicate.CoinGas {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinGas(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDepositThreshold), v...))
	})
}

// DepositThresholdNotIn applies the NotIn predicate on the "deposit_threshold" field.
func DepositThresholdNotIn(vs ...uint64) predicate.CoinGas {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CoinGas(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDepositThreshold), v...))
	})
}

// DepositThresholdGT applies the GT predicate on the "deposit_threshold" field.
func DepositThresholdGT(v uint64) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDepositThreshold), v))
	})
}

// DepositThresholdGTE applies the GTE predicate on the "deposit_threshold" field.
func DepositThresholdGTE(v uint64) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDepositThreshold), v))
	})
}

// DepositThresholdLT applies the LT predicate on the "deposit_threshold" field.
func DepositThresholdLT(v uint64) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDepositThreshold), v))
	})
}

// DepositThresholdLTE applies the LTE predicate on the "deposit_threshold" field.
func DepositThresholdLTE(v uint64) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDepositThreshold), v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CoinGas) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CoinGas) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CoinGas) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		p(s.Not())
	})
}
