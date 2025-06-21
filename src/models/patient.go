package models

import (
	"time"

	"github.com/agnos/hospital-middleware/pkg"
)

type PatientModel struct {
	ID           uint      `gorm:"primaryKey;autoIncrement"`
	NationalID   *string   `gorm:"column:national_id;type:varchar(255);not null;default:''"`
	PassportID   *string   `gorm:"column:passport_id;type:varchar(255);not null;default:''"`
	FirstNameTh  *string   `gorm:"column:first_name_th;type:varchar(255)"`
	MiddleNameTh *string   `gorm:"column:middle_name_th;type:varchar(255)"`
	LastNameTh   *string   `gorm:"column:last_name_th;type:varchar(255)"`
	FirstNameEn  *string   `gorm:"column:first_name_en;type:varchar(255)"`
	MiddleNameEn *string   `gorm:"column:middle_name_en;type:varchar(255)"`
	LastNameEn   *string   `gorm:"column:last_name_en;type:varchar(255)"`
	BirthDate    string    `gorm:"column:birth_date;type:date"`
	Gender       string    `gorm:"column:gender;type:varchar(10)"`
	PhoneNumber  *string   `gorm:"column:phone_number;type:varchar(255)"`
	Email        *string   `gorm:"column:email;type:varchar(255)"`
	PatentHN     string    `gorm:"column:patent_hn;type:varchar(10);not null;default:'A'"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

type GetByPNIDResponse struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	MiddleName  string `json:"middle_name"`
	LastName    string `json:"last_name"`
	BirthDate   string `json:"birth_date"`
	Gender      string `json:"gender"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type GetByPNIDResponseAllLang struct {
	ID           string `json:"id"`
	FirstNameTh  string `json:"first_name_th"`
	MiddleNameTh string `json:"middle_name_th"`
	LastNameTh   string `json:"last_name_th"`
	FirstNameEn  string `json:"first_name_en"`
	MiddleNameEn string `json:"middle_name_en"`
	LastNameEn   string `json:"last_name_en"`
	BirthDate    string `json:"birth_date"`
	Gender       string `json:"gender"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email"`
}

func (p *PatientModel) PatientResponseDTOAllLang() GetByPNIDResponseAllLang {
	var PNDResponse GetByPNIDResponseAllLang
	if p.NationalID != nil {
		PNDResponse.ID = *p.NationalID
	} else {
		PNDResponse.ID = *p.PassportID
	}
	PNDResponse.FirstNameTh = pkg.IsNullReturnString(p.FirstNameTh)
	PNDResponse.MiddleNameTh = pkg.IsNullReturnString(p.MiddleNameTh)
	PNDResponse.LastNameTh = pkg.IsNullReturnString(p.LastNameTh)
	PNDResponse.FirstNameEn = pkg.IsNullReturnString(p.FirstNameEn)
	PNDResponse.MiddleNameEn = pkg.IsNullReturnString(p.MiddleNameEn)
	PNDResponse.LastNameEn = pkg.IsNullReturnString(p.LastNameEn)
	PNDResponse.BirthDate = p.BirthDate
	PNDResponse.Gender = p.Gender
	PNDResponse.PhoneNumber = pkg.IsNullReturnString(p.PhoneNumber)
	PNDResponse.Email = pkg.IsNullReturnString(p.Email)
	return PNDResponse
}

func (p *PatientModel) PatientResponseDTO(lang string) GetByPNIDResponse {
	var PNDResponse GetByPNIDResponse
	if p.NationalID != nil {
		PNDResponse.ID = *p.NationalID
	} else {
		PNDResponse.ID = *p.PassportID
	}
	if pkg.IsTH(lang) {
		PNDResponse.FirstName = *p.FirstNameTh
		PNDResponse.MiddleName = *p.MiddleNameTh
		PNDResponse.LastName = *p.LastNameTh
	} else {
		PNDResponse.FirstName = *p.FirstNameEn
		PNDResponse.MiddleName = *p.MiddleNameEn
		PNDResponse.LastName = *p.LastNameEn
	}
	PNDResponse.BirthDate = p.BirthDate
	PNDResponse.Gender = p.Gender
	PNDResponse.PhoneNumber = pkg.IsNullReturnString(p.PhoneNumber)
	PNDResponse.Email = pkg.IsNullReturnString(p.Email)
	return PNDResponse
}
