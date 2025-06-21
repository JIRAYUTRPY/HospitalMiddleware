package models

import "time"

type StaffModel struct {
	ID           int       `gorm:"primaryKey;autoIncrement"`
	FirstNameTh  string    `gorm:"type:varchar(255);column:first_name_th"`
	MiddleNameTh string    `gorm:"type:varchar(255);column:middle_name_th"`
	LastNameTh   string    `gorm:"type:varchar(255);column:last_name_th"`
	FirstNameEn  string    `gorm:"type:varchar(255);column:first_name_en"`
	MiddleNameEn string    `gorm:"type:varchar(255);column:middle_name_en"`
	LastNameEn   string    `gorm:"type:varchar(255);column:last_name_en"`
	BirthDate    string    `gorm:"type:date;column:birth_date"`
	Gender       string    `gorm:"type:varchar(10);column:gender"`
	PhoneNumber  string    `gorm:"type:varchar(255);column:phone_number"`
	Email        string    `gorm:"type:varchar(255);column:email;unique_index"`
	HashPassword string    `gorm:"type:varchar(255);column:hash_password"`
	StaffHN      string    `gorm:"type:varchar(10);column:staff_hn"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type RegisterRequest struct {
	Email        string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
	FirstNameTh  string `json:"first_name_th" binding:"required"`
	MiddleNameTh string `json:"middle_name_th" binding:"required"`
	LastNameTh   string `json:"last_name_th" binding:"required"`
	FirstNameEn  string `json:"first_name_en" binding:"required"`
	MiddleNameEn string `json:"middle_name_en" binding:"required"`
	LastNameEn   string `json:"last_name_en" binding:"required"`
	BirthDate    string `json:"birth_date" binding:"required"`
	Gender       string `json:"gender" binding:"required"`
	PhoneNumber  string `json:"phone_number" binding:"required"`
}

type RegisterResponse struct {
	ID          int    `json:"id"`
	AccessToken string `json:"access_token"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	ID          int    `json:"id"`
	AccessToken string `json:"access_token"`
}
