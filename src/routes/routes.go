package routes

import (
	"net/http"

	"github.com/agnos/hospital-middleware/patient"
	"github.com/agnos/hospital-middleware/pkg"
	"github.com/agnos/hospital-middleware/staff"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HospitalRouter(hospital string, db *gorm.DB) http.Handler {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	engine.Use(Middleware(hospital))
	RegisterRoutes(engine)
	engine.Use(pkg.WrapperError())
	return engine
}

func RegisterRoutes(engine *gin.Engine) {
	api := engine.Group("/api/v1")
	api.GET("", HelthCheck)
	staffGroup := api.Group("/staff")
	patientGroup := api.Group("/patient")
	staffGroup.POST("/register", staff.Register)
	staffGroup.POST("/login", staff.Login)
	// Patient routes
	patientGroup.GET(":id", patient.GetPatientByPassportOrNationalID)
}

func HelthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HOSPITAL " + c.GetString("hospital"),
	})
}

func Middleware(hospital string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("hospital", hospital)
		lang := c.GetHeader("Accept-Language")
		if lang == "" {
			lang = "all_lang"
		}
		c.Set("lang", lang)
		c.Next()
	}
}
