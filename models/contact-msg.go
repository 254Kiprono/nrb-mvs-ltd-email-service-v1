package models

import (
	"time"
)

// ContactMessage represents a submission from the 'Send Us a Message' form.
type ContactMessage struct {
	ID          uint   `gorm:"primaryKey"`
	FullName    string `gorm:"column:full_name"`
	PhoneNumber string `gorm:"column:phone_number"`
	Email       string `gorm:"column:email"`
	Message     string `gorm:"column:message;type:text"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
