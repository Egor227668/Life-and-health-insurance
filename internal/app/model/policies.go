package models

import "time"

type Policy struct {
    PolicyID        int       db:"policy_id"
    PolicyNumber    string    db:"policy_number"
    Client          Client    db:"client"
    Agent           Agent     db:"agent"
    InsuranceType   InsuranceType db:"insurance_type"
    StartDate       time.Time db:"start_date"
    EndDate         time.Time db:"end_date"
    Status          string    db:"status"
    CoverageAmount  float64   db:"coverage_amount"
    Premium         float64   db:"premium"
}