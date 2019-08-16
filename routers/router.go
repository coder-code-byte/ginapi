package routers

import (
	"ginapi/http/middleware/jwt"
	"ginapi/http/middleware/middlehttp"
	"ginapi/http/service"
	v1 "ginapi/http/service/v1"
	"ginapi/pkg/setting"
	"net/http"

	"github.com/gin-gonic/gin"
)

// InitRouter this is for InitRouter
func InitRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	gin.SetMode(setting.RunMode)
	router.Use(middlehttp.LogIP)
	router.Use(middlehttp.NoCache)
	router.Use(middlehttp.Options)
	router.Use(middlehttp.Secure)
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "The incorrect API route.",
		})
	})
	router.GET("/api/token", service.Auth)
	apiv1 := router.Group("api")
	apiv1.Use(jwt.JWT())
	{
		apiv1.POST("/test", v1.Test)
		apiv1.POST("/test2", v1.TestPost)
	}
	return router
}
