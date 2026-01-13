package models

import "time"

type InsuranceCompany struct {
    CompanyID   int       `db:"company_id"`
    Name        string    `db:"name"`
    INN         string    `db:"inn"`
    LegalAddress string   `db:"legal_address"`
    Phone       string    `db:"phone"`
    Email       string    `db:"email"`
}