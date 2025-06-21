package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/agnos/hospital-middleware/config"
	"github.com/agnos/hospital-middleware/models"
	"github.com/agnos/hospital-middleware/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestAuthMiddleware(t *testing.T) {

	t.Run("invalid token", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		db.AutoMigrate(&models.StaffModel{})
		hospitalRouter := routes.HospitalRouter("A", db, &config.JWTConfig{
			AccessSecret: "test",
		})
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("GET", "/api/v1/patient/1234567890123", nil)
		hospitalRouter.ServeHTTP(w, c.Request)
		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
