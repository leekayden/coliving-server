package models

import "time"

type User struct {
	Id               uint64 `gorm:"primaryKey"`
	FirstName        string
	LastName         string
	Location         string // e.g. hougang
	Email            string
	Password         string
	ReceiveAddEmails bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
	// AccountType Enum? // tenant, landlord, admin
}
