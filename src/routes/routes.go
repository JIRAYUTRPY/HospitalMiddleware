package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HospitalRouter(hospital string) http.Handler {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(func(c *gin.Context) {
		c.Set("hospital", hospital)
		c.Next()
	})
	RegisterRoutes(engine)
	engine.Use(ErrorHandler())
	return engine
}

func RegisterRoutes(engine *gin.Engine) {
	api := engine.Group("/api/v1")
	staff := api.Group("/staff")
	patient := api.Group("/patient")
	api.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "HOSPITAL " + c.GetString("hospital"),
		})
	})
	staff.POST("/register", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	staff.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	patient.POST("/register", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
	patient.POST("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})
}
