package core

import (
	"net/http"
	"regexp"

	"github.com/agnos/hospital-middleware/models"
	"github.com/agnos/hospital-middleware/pkg"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func validateID(id string) (bool, string) {
	if id == "" {
		return false, "ID is required"
	}
	if !regexp.MustCompile(`^[0-9]+$`).MatchString(id) || len(id) != 13 {
		return false, "ID must be a number and 13 digits"
	}

	return true, ""
}

func GetPatientByPassportOrNationalID(c *gin.Context) {
	id := c.Param("id")
	valid, reason := validateID(id)
	if !valid {
		c.JSON(http.StatusBadRequest, map[string]any{
			"reason": reason,
		})
		return
	}
	hospital := c.GetString("hospital")
	lang := c.GetString("lang")
	db := c.MustGet("db").(*gorm.DB)
	var patient models.PatientModel
	result := db.Where("(national_id LIKE ? OR passport_id LIKE ?) AND patent_hn = ?", id+"%", id+"%", hospital).First(&patient)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, map[string]any{
				"reason": "Patient not found",
			})
			return
		}
		c.Error(result.Error)
		return
	}

	if !pkg.CheckLang(lang) {
		c.JSON(http.StatusOK, patient.PatientResponseDTOAllLang())
		return
	}
	c.JSON(http.StatusOK, patient.PatientResponseDTO(lang))
}
