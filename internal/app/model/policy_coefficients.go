package models

import "time"

type PolicyCoefficient struct {
    PolicyID       int       `db:"policy_id" json:"policy_id"`
    CoefficientID  int       `db:"coefficient_id" json:"coefficient_id"`
    AppliedValue   float64   `db:"applied_value" json:"applied_value"`
    CreatedAt      time.Time `db:"created_at" json:"created_at"`
}
