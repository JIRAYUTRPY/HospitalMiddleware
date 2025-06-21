package core

import (
	"net/http"

	"github.com/agnos/hospital-middleware/config"
	"github.com/agnos/hospital-middleware/models"
	"github.com/agnos/hospital-middleware/pkg"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Staff struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func Register(c *gin.Context) {
	var request models.RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"reason": "Invalid request",
		})
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("email = ? AND staff_hn = ?", request.Email, c.GetString("hospital")).First(&models.StaffModel{}).Error; err == nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"reason": "Email already exists",
		})
		return
	}
	hashedPassword, err := pkg.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"reason": "Failed to hash password",
		})
		return
	}
	hn := c.GetString("hospital")
	staff := models.StaffModel{
		Email:        request.Email,
		HashPassword: hashedPassword,
		FirstNameTh:  request.FirstNameTh,
		MiddleNameTh: request.MiddleNameTh,
		LastNameTh:   request.LastNameTh,
		FirstNameEn:  request.FirstNameEn,
		MiddleNameEn: request.MiddleNameEn,
		LastNameEn:   request.LastNameEn,
		BirthDate:    request.BirthDate,
		Gender:       request.Gender,
		PhoneNumber:  request.PhoneNumber,
		StaffHN:      hn,
	}
	if err := db.Create(&staff).Error; err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"reason": "Failed to create staff",
		})
		return
	}
	jwtConfig := c.MustGet("jwt_config").(config.JWTConfig)
	accessToken, err := pkg.GenerateTokens(staff.ID, jwtConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"reason": "Failed to generate tokens",
		})
		return
	}
	c.JSON(http.StatusOK, models.RegisterResponse{
		ID:          staff.ID,
		AccessToken: accessToken,
	})
}

func Login(c *gin.Context) {
	var request models.LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, map[string]any{
			"reason": err.Error(),
		})
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	var staff models.StaffModel
	if err := db.Where("email = ? AND staff_hn = ?", request.Email, c.GetString("hospital")).First(&staff).Error; err != nil {
		c.JSON(http.StatusNotFound, map[string]any{
			"reason": "Staff not found",
		})
		return
	}
	if err := pkg.VerifyPassword(request.Password, staff.HashPassword); err != nil {
		c.JSON(http.StatusUnauthorized, map[string]any{
			"reason": "Invalid password",
		})
		return
	}
	jwtConfig := c.MustGet("jwt_config").(config.JWTConfig)
	accessToken, err := pkg.GenerateTokens(staff.ID, jwtConfig)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]any{
			"reason": "Failed to generate tokens",
		})
		return
	}
	c.JSON(http.StatusOK, models.LoginResponse{
		ID:          staff.ID,
		AccessToken: accessToken,
	})
}
