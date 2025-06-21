package tests

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/agnos/hospital-middleware/config"
	"github.com/agnos/hospital-middleware/models"
	"github.com/agnos/hospital-middleware/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestRegister(t *testing.T) {

	t.Run("valid request", func(b *testing.T) {
		gin.SetMode(gin.TestMode)
		db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		db.AutoMigrate(&models.StaffModel{})
		hospitalRouter := routes.HospitalRouter("A", db, &config.JWTConfig{
			AccessSecret: "test",
		})
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/api/v1/staff/register", strings.NewReader(`{"email": "test@test.com", "password": "test"}`))
		hospitalRouter.ServeHTTP(w, c.Request)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("request success", func(b *testing.T) {
		gin.SetMode(gin.TestMode)
		db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		db.AutoMigrate(&models.StaffModel{})
		hospitalRouter := routes.HospitalRouter("A", db, &config.JWTConfig{
			AccessSecret: "test",
		})
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/api/v1/staff/register", strings.NewReader(`{"email": "test@test.com", "password": "test", "first_name_th": "test", "middle_name_th": "test", "last_name_th": "test", "first_name_en": "test", "middle_name_en": "test", "last_name_en": "test", "birth_date": "1990-01-01", "gender": "M", "phone_number": "0812345678"}`))
		hospitalRouter.ServeHTTP(w, c.Request)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
