package jobseeker

import (
	"net/http"

	globaltypes "example.com/agenagn/global-types"
	"example.com/agenagn/initializers"
	"example.com/agenagn/models"
	"example.com/agenagn/utils"
	"github.com/gin-gonic/gin"
)


func Login(ctx *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&credentials); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch user data from the database based on the email
	var user models.JobSeeker
	result := initializers.DB.Where("email_address = ?", credentials.Email).First(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check if the provided password matches the stored hash
	if !utils.ComparePasswords(user.Password, credentials.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// If everything is correct, generate and send the token
	var response globaltypes.AuthResponse
	response.ID = int(user.ID)
	response.Role = "job_seeker"
	utils.SendToken(ctx, "job_seeker", response)
}

