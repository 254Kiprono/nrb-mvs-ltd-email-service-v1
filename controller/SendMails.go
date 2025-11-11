package controller

import (
	"email-service/models"
	"email-service/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Contact Email Handler
func SendContactEmail(c *gin.Context) {
	var form models.ContactForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emailContent := fmt.Sprintf(
		"New Contact Message\n\nName: %s\nEmail: %s\nPhone: %s\nService: %s\n\nMessage:\n%s",
		form.Name,
		form.Email,
		form.Phone,
		form.Service,
		form.Message,
	)

	err := utils.SendEmail("New Contact Form Submission", form.Email, form.Name, emailContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact form submitted successfully"})
}

// Quote Request Email Handler
func SendQuoteEmail(c *gin.Context) {
	var form models.QuoteForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	emailContent := fmt.Sprintf(
		"New Quote Request\n\n"+
			"Name: %s\n"+
			"Email: %s\n"+
			"Phone: %s\n"+
			"Project Type: %s\n"+
			"Budget: %s\n"+
			"Timeline: %s\n\n"+
			"Project Description:\n%s",
		form.Name,
		form.Email,
		form.Phone,
		form.ProjectType,
		form.Budget,
		form.Timeline,
		form.Description,
	)

	err := utils.SendEmail("New Quote Request - "+form.ProjectType, form.Email, form.Name, emailContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Quote request submitted successfully"})
}
