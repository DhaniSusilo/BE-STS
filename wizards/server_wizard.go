package wizards

import (
	"learning-backend/shared/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterServer(router *gin.Engine) {
	// Users
	users := router.Group("/users")
	{
		users.POST("/login", Http.Login)
		users.POST("/register", Http.SignUp)

		users.Use(middlewares.AuthMiddleware())
	}

	member := router.Group("/members")
	{
		member.POST("/add",Http.AddMember)
	}

	admin := router.Group("/admins")
	{
		admin.POST("/dashboard",Http.GetDashboardData)
		admin.POST("/rekap",Http.GetRekapitulasi)
		admin.GET("/all",Http.GetAllUser)
	}
}
