package models

import "time"

type Beneficiary struct {
    BeneficiaryID   int       `db:"beneficiary_id"`
    FullName        string    `db:"full_name"`
    Relationship    string    `db:"relationship"`
    PassportNumber  string    `db:"passport_number"`
    Policy          Policy    `db:"policy"`
}