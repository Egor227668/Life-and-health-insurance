package models

import "time"

type Damage struct {
    DamageID        int       `db:"damage_id"`
    InsuranceClaim  InsuranceClaim `db:"insurance_claim"`
    DamageType      string    `db:"damage_type"`
    Severity        string    `db:"severity"`
    ICDCode         string    `db:"icd_code"`
    DiagnosisDate   time.Time `db:"diagnosis_date"`
}