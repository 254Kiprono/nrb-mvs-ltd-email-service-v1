package controller

import (
	"email-service/models"
	"email-service/repo"
	"email-service/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Contact Email Handler
func SendContactEmail(c *gin.Context, db *gorm.DB) {
	var form models.ContactForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contactRepo := repo.NewContactRepository(db)

	message := models.ContactMessage{
		FullName:    form.Name,
		PhoneNumber: form.Phone,
		Email:       form.Email,
		Message:     form.Message,
	}

	// Save to DB
	if err := contactRepo.SaveMessage(&message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save message to database"})
		return
	}

	// Send Email Notification
	emailContent := fmt.Sprintf(
		"New Contact Message\n\nName: %s\nEmail: %s\nPhone: %s\nService: %s\n\nMessage:\n%s",
		form.Name,
		form.Email,
		form.Phone,
		form.Service,
		form.Message,
	)

	if err := utils.SendEmail("New Contact Form Submission", form.Email, form.Name, emailContent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contact form submitted successfully and saved to database"})
}

// Quote Request Email Handler
func SendQuoteEmail(c *gin.Context, db *gorm.DB) {
	var form models.QuoteRequest

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contactRepo := repo.NewContactRepository(db)

	quote := models.QuoteRequest{
		ServiceType:     form.ServiceType,
		PickupLocation:  form.PickupLocation,
		DropoffLocation: form.DropoffLocation,
		FullName:        form.FullName,
		PhoneNumber:     form.PhoneNumber,
		Email:           form.Email,
		IsDateFlexible:  form.IsDateFlexible,
	}

	// Parse and assign moving date if provided
	if form.MovingDate != "" {
		if parsedDate, err := time.Parse("2006-01-02", form.MovingDate); err == nil {
			quote.MovingDate = parsedDate.Format("2006-01-02")
		}
	}

	// Save to DB
	if err := contactRepo.SaveQuoteRequest(&quote); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save quote request to database"})
		return
	}

	// Send Email
	emailContent := fmt.Sprintf(
		"New Quote Request\n\n"+
			"Name: %s\n"+
			"Email: %s\n"+
			"Phone: %s\n"+
			"Service Type: %s\n"+
			"Pickup Location: %s\n"+
			"Dropoff Location: %s\n"+
			"Moving Date: %v\n"+
			"Flexible Date: %v",
		form.FullName,
		form.Email,
		form.PhoneNumber,
		form.ServiceType,
		form.PickupLocation,
		form.DropoffLocation,
		form.MovingDate,
		form.IsDateFlexible,
	)

	if err := utils.SendEmail("New Quote Request - "+form.ServiceType, form.Email, form.FullName, emailContent); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Quote request submitted successfully and saved to database"})
}
