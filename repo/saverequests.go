package repo

import (
	"email-service/models"

	"gorm.io/gorm"
)

type ContactRepository interface {
	SaveMessage(message *models.ContactMessage) error
	SaveQuoteRequest(quote *models.QuoteRequest) error
}

type contactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) ContactRepository {
	return &contactRepository{db: db}
}

func (r *contactRepository) SaveMessage(message *models.ContactMessage) error {
	return r.db.Create(message).Error
}

func (r *contactRepository) SaveQuoteRequest(quote *models.QuoteRequest) error {
	return r.db.Create(quote).Error
}
