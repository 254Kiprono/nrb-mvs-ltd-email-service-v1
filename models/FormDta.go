package models

type ContactForm struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Phone   string `json:"phone" binding:"required"`
	Service string `json:"service" binding:"required"`
	Message string `json:"message" binding:"required"`
}

