package routes

import (
	"net/http"
	"strings"

	"github.com/agnos/hospital-middleware/config"
	"github.com/agnos/hospital-middleware/core"
	"github.com/agnos/hospital-middleware/pkg"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HospitalRouter(hospital string, db *gorm.DB, jwtConfig *config.JWTConfig) http.Handler {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	engine.Use(Middleware(hospital, jwtConfig))
	RegisterRoutes(engine, jwtConfig)
	engine.Use(pkg.WrapperError())
	return engine
}

func RegisterRoutes(engine *gin.Engine, jwtConfig *config.JWTConfig) {
	api := engine.Group("/api/v1")
	api.GET("", HelthCheck)
	staffGroup := api.Group("/staff")
	patientGroup := api.Group("/patient")
	staffGroup.POST("/register", core.Register)
	staffGroup.POST("/login", core.Login)
	// Patient routes
	patientGroup.Use(AuthMiddleware())
	staffGroup.Use(AuthMiddleware())
	patientGroup.GET(":id", core.GetPatientByPassportOrNationalID)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		rawToken := c.GetHeader("Authorization")
		if rawToken == "" {
			c.JSON(http.StatusUnauthorized, map[string]any{
				"reason": "Unauthorized",
			})
			c.Abort()
			return
		}
		token := strings.TrimPrefix(rawToken, "Bearer ")
		jwtConfig := c.MustGet("jwt_config").(config.JWTConfig)
		claims, err := pkg.ValidateToken(token, jwtConfig)
		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]any{
				"reason": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Set("staff_id", claims.AccountID)
		c.Next()
	}
}

func HelthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]any{
		"message": "HOSPITAL " + c.GetString("hospital"),
	})
}

func Middleware(hospital string, jwtConfig *config.JWTConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("hospital", hospital)
		lang := c.GetHeader("Accept-Language")
		if lang == "" {
			lang = "all_lang"
		}
		c.Set("lang", lang)
		c.Set("jwt_config", *jwtConfig)
		c.Next()
	}
}
