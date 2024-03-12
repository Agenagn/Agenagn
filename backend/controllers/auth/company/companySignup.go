package company

import (
	"net/http"

	globaltypes "example.com/agenagn/global-types"
	"example.com/agenagn/initializers"
	"example.com/agenagn/models"
	"example.com/agenagn/utils"
	"github.com/gin-gonic/gin"
)




func CompanySignUp(c *gin.Context) {
	user := models.Company{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/*---------------------hash password---------------------------*/

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = hashedPassword

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if result.RowsAffected == 0 {
		// Rows were not affected, indicating that the record was not created
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	var response globaltypes.AuthResponse
	response.ID = int(user.ID)
	response.Role = "company"

	utils.SendToken(c, "company", response)
}






