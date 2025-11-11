package models

import (
	"time"
)

// QuoteRequest
type QuoteRequest struct {
	ID              uint   `gorm:"primaryKey"`
	ServiceType     string `gorm:"column:service_type"`
	PickupLocation  string `gorm:"column:pickup_location;type:text"`
	DropoffLocation string `gorm:"column:dropoff_location;type:text"`
	MovingDate      string `gorm:"column:moving_date"`
	IsDateFlexible  bool   `gorm:"column:is_date_flexible"`
	FullName        string `gorm:"column:full_name"`
	PhoneNumber     string `gorm:"column:phone_number"`
	Email           string `gorm:"column:email"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
