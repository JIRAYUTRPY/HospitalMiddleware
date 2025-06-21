package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/agnos/hospital-middleware/core"
	"github.com/agnos/hospital-middleware/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	id           = "1234567890123"
	firstNameTh  = "จอห์น"
	middleNameTh = ""
	lastNameTh   = "ดอย"
	firstNameEn  = "John"
	middleNameEn = "Doe"
	lastNameEn   = "Doe"
	birthDate    = "2000-01-01"
	phoneNumber  = "0812345678"
	email        = "john.doe@example.com"
	gender       = "M"
	hn           = "A"
)

func TestGetPatientByPassportOrNationalID(t *testing.T) {
	t.Run("valid id", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		assert.NoError(t, err)
		db.AutoMigrate(&models.PatientModel{})
		db.Create(&models.PatientModel{
			NationalID:   &id,
			ID:           1,
			FirstNameTh:  &firstNameTh,
			MiddleNameTh: &middleNameTh,
			LastNameTh:   &lastNameTh,
			FirstNameEn:  &firstNameEn,
			MiddleNameEn: &middleNameEn,
			LastNameEn:   &lastNameEn,
			BirthDate:    birthDate,
			Gender:       gender,
			PhoneNumber:  &phoneNumber,
			Email:        &email,
			PatentHN:     hn,
		})
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("db", db)
		c.Set("hospital", "A")
		c.Params = gin.Params{
			{Key: "id", Value: "1234567890123"},
		}
		core.GetPatientByPassportOrNationalID(c)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "John")
	})

	t.Run("valid hospital name", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
		assert.NoError(t, err)
		db.AutoMigrate(&models.PatientModel{})

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Set("db", db)
		db.Create(&models.PatientModel{
			NationalID:   &id,
			ID:           1,
			FirstNameTh:  &firstNameTh,
			MiddleNameTh: &middleNameTh,
			LastNameTh:   &lastNameTh,
			FirstNameEn:  &firstNameEn,
			MiddleNameEn: &middleNameEn,
			LastNameEn:   &lastNameEn,
			BirthDate:    birthDate,
			Gender:       gender,
			PhoneNumber:  &phoneNumber,
			Email:        &email,
			PatentHN:     hn,
		})
		c.Set("hospital", "B")
		c.Params = gin.Params{
			{Key: "id", Value: id},
		}
		core.GetPatientByPassportOrNationalID(c)
		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Contains(t, w.Body.String(), "Patient not found")
	})
}
