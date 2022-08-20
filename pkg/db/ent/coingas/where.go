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

// FeedingTid applies equality check predicate on the "feeding_tid" field. It's identical to FeedingTidEQ.
func FeedingTid(v uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFeedingTid), v))
	})
}

// DepositThresholdLow applies equality check predicate on the "deposit_threshold_low" field. It's identical to DepositThresholdLowEQ.
func DepositThresholdLow(v uint64) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDepositThresholdLow), v))
	})
}

// DepositAmount applies equality check predicate on the "deposit_amount" field. It's identical to DepositAmountEQ.
func DepositAmount(v uint64) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDepositAmount), v))
	})
}

// OnlineScale applies equality check predicate on the "online_scale" field. It's identical to OnlineScaleEQ.
func OnlineScale(v int32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOnlineScale), v))
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

// FeedingTidEQ applies the EQ predicate on the "feeding_tid" field.
func FeedingTidEQ(v uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldFeedingTid), v))
	})
}

// FeedingTidNEQ applies the NEQ predicate on the "feeding_tid" field.
func FeedingTidNEQ(v uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldFeedingTid), v))
	})
}

// FeedingTidIn applies the In predicate on the "feeding_tid" field.
func FeedingTidIn(vs ...uuid.UUID) predicate.CoinGas {
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
		s.Where(sql.In(s.C(FieldFeedingTid), v...))
	})
}

// FeedingTidNotIn applies the NotIn predicate on the "feeding_tid" field.
func FeedingTidNotIn(vs ...uuid.UUID) predicate.CoinGas {
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
		s.Where(sql.NotIn(s.C(FieldFeedingTid), v...))
	})
}

// FeedingTidGT applies the GT predicate on the "feeding_tid" field.
func FeedingTidGT(v uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldFeedingTid), v))
	})
}

// FeedingTidGTE applies the GTE predicate on the "feeding_tid" field.
func FeedingTidGTE(v uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldFeedingTid), v))
	})
}

// FeedingTidLT applies the LT predicate on the "feeding_tid" field.
func FeedingTidLT(v uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldFeedingTid), v))
	})
}

// FeedingTidLTE applies the LTE predicate on the "feeding_tid" field.
func FeedingTidLTE(v uuid.UUID) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldFeedingTid), v))
	})
}

// DepositThresholdLowEQ applies the EQ predicate on the "deposit_threshold_low" field.
func DepositThresholdLowEQ(v uint64) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDepositThresholdLow), v))
	})
}

// DepositThresholdLowNEQ applies the NEQ predicate on the "deposit_threshold_low" field.
func DepositThresholdLowNEQ(v uint64) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDepositThresholdLow), v))
	})
}

// DepositThresholdLowIn applies the In predicate on the "deposit_threshold_low" field.
func DepositThresholdLowIn(vs ...uint64) predicate.CoinGas {
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
		s.Where(sql.In(s.C(FieldDepositThresholdLow), v...))
	})
}

// DepositThresholdLowNotIn applies the NotIn predicate on the "deposit_threshold_low" field.
func DepositThresholdLowNotIn(vs ...uint64) predicate.CoinGas {
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
		s.Where(sql.NotIn(s.C(FieldDepositThresholdLow), v...))
	})
}

// DepositThresholdLowGT applies the GT predicate on the "deposit_threshold_low" field.
func DepositThresholdLowGT(v uint64) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDepositThresholdLow), v))
	})
}

// DepositThresholdLowGTE applies the GTE predicate on the "deposit_threshold_low" field.
func DepositThresholdLowGTE(v uint64) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDepositThresholdLow), v))
	})
}

// DepositThresholdLowLT applies the LT predicate on the "deposit_threshold_low" field.
func DepositThresholdLowLT(v uint64) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDepositThresholdLow), v))
	})
}

// DepositThresholdLowLTE applies the LTE predicate on the "deposit_threshold_low" field.
func DepositThresholdLowLTE(v uint64) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDepositThresholdLow), v))
	})
}

// DepositAmountEQ applies the EQ predicate on the "deposit_amount" field.
func DepositAmountEQ(v uint64) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDepositAmount), v))
	})
}

// DepositAmountNEQ applies the NEQ predicate on the "deposit_amount" field.
func DepositAmountNEQ(v uint64) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDepositAmount), v))
	})
}

// DepositAmountIn applies the In predicate on the "deposit_amount" field.
func DepositAmountIn(vs ...uint64) predicate.CoinGas {
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
		s.Where(sql.In(s.C(FieldDepositAmount), v...))
	})
}

// DepositAmountNotIn applies the NotIn predicate on the "deposit_amount" field.
func DepositAmountNotIn(vs ...uint64) predicate.CoinGas {
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
		s.Where(sql.NotIn(s.C(FieldDepositAmount), v...))
	})
}

// DepositAmountGT applies the GT predicate on the "deposit_amount" field.
func DepositAmountGT(v uint64) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDepositAmount), v))
	})
}

// DepositAmountGTE applies the GTE predicate on the "deposit_amount" field.
func DepositAmountGTE(v uint64) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDepositAmount), v))
	})
}

// DepositAmountLT applies the LT predicate on the "deposit_amount" field.
func DepositAmountLT(v uint64) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDepositAmount), v))
	})
}

// DepositAmountLTE applies the LTE predicate on the "deposit_amount" field.
func DepositAmountLTE(v uint64) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDepositAmount), v))
	})
}

// OnlineScaleEQ applies the EQ predicate on the "online_scale" field.
func OnlineScaleEQ(v int32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldOnlineScale), v))
	})
}

// OnlineScaleNEQ applies the NEQ predicate on the "online_scale" field.
func OnlineScaleNEQ(v int32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldOnlineScale), v))
	})
}

// OnlineScaleIn applies the In predicate on the "online_scale" field.
func OnlineScaleIn(vs ...int32) predicate.CoinGas {
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
		s.Where(sql.In(s.C(FieldOnlineScale), v...))
	})
}

// OnlineScaleNotIn applies the NotIn predicate on the "online_scale" field.
func OnlineScaleNotIn(vs ...int32) predicate.CoinGas {
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
		s.Where(sql.NotIn(s.C(FieldOnlineScale), v...))
	})
}

// OnlineScaleGT applies the GT predicate on the "online_scale" field.
func OnlineScaleGT(v int32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldOnlineScale), v))
	})
}

// OnlineScaleGTE applies the GTE predicate on the "online_scale" field.
func OnlineScaleGTE(v int32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldOnlineScale), v))
	})
}

// OnlineScaleLT applies the LT predicate on the "online_scale" field.
func OnlineScaleLT(v int32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldOnlineScale), v))
	})
}

// OnlineScaleLTE applies the LTE predicate on the "online_scale" field.
func OnlineScaleLTE(v int32) predicate.CoinGas {
	return predicate.CoinGas(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldOnlineScale), v))
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
