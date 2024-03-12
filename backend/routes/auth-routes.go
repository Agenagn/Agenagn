package routes

import (
	job_seeker "example.com/agenagn/controllers/auth/job-seeker"
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.Engine) {
    authGroup := router.Group("/auth")
    {
        job_seekerGroup := authGroup.Group("/job-seeker")
        {
            job_seekerGroup.POST("/signup", job_seeker.SignUp)
            job_seekerGroup.POST("/login", job_seeker.Login)			
        }
    }
}
