package models

import "time"

type Client struct {
    ClientID        int       `db:"client_id"`
    FullName        string    `db:"full_name"`
    PassportNumber  string    `db:"passport_number"`
    BirthDate       time.Time `db:"birth_date"`
    Gender          string    `db:"gender"`
    Phone           string    `db:"phone"`
    Email           string    `db:"email"`
    HealthInfo      string    `db:"health_info"`
}