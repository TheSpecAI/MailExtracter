package routes

import (
	"github.com/gin-gonic/gin"
	"www.github.com/NirajSalunke/server/controllers"
)

func LoadMailRoutes(r *gin.RouterGroup) {
	r.GET("/", controllers.GetAllMails)
	r.POST("/", controllers.GetMail)
}
