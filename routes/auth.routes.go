package routes

import (
	"github.com/gin-gonic/gin"
	"www.github.com/NirajSalunke/server/controllers"
)

func LoadAuthRoutes(r *gin.Engine) {
	r.GET("/login", controllers.HandleLogin)
	r.GET("/oauth2callback", controllers.HandleCallBack)
}
