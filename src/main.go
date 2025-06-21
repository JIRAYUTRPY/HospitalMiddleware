package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/agnos/hospital-middleware/config"
	"github.com/agnos/hospital-middleware/models"
	"github.com/agnos/hospital-middleware/routes"
	_ "github.com/lib/pq"
	"golang.org/x/sync/errgroup"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	server errgroup.Group
)

func main() {

	config := config.DatabaseConfig{}
	db, err := gorm.Open(postgres.Open(config.GetConnectionString()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	jwtConfig := config.GetJWTConfig()
	nationalID := "1234567891234"
	passportID := "1234567891235"
	firstNameTh := "สมชาย"
	firstNameEn := "Somchai"
	lastNameTh := "สมหญิง"
	lastNameEn := "Somying"
	phoneNumber := "0812345678"
	email := "somchai@example.com"
	patentHN := "A"
	patentHN2 := "B"
	middleNameTh := "สมหญิง"
	middleNameEn := "Somying"

	patientSeed := models.PatientModel{
		ID:           uint(1),
		NationalID:   &nationalID,
		FirstNameTh:  &firstNameTh,
		MiddleNameTh: &middleNameTh,
		LastNameTh:   &lastNameTh,
		FirstNameEn:  &firstNameEn,
		MiddleNameEn: &middleNameEn,
		LastNameEn:   &lastNameEn,
		BirthDate:    "2000-01-01",
		Gender:       "M",
		PhoneNumber:  &phoneNumber,
		Email:        &email,
		PatentHN:     patentHN,
	}

	patientSeed2 := models.PatientModel{
		ID:           uint(2),
		PassportID:   &passportID,
		FirstNameTh:  &firstNameTh,
		MiddleNameTh: &middleNameTh,
		LastNameTh:   &lastNameTh,
		FirstNameEn:  &firstNameEn,
		MiddleNameEn: &middleNameEn,
		LastNameEn:   &lastNameEn,
		BirthDate:    "2000-01-01",
		Gender:       "M",
		PhoneNumber:  &phoneNumber,
		Email:        &email,
		PatentHN:     patentHN2,
	}

	db.AutoMigrate(&models.PatientModel{}, &models.StaffModel{})
	db.Create(&patientSeed)
	db.Create(&patientSeed2)

	if err != nil {
		panic(err)
	}

	serverA := http.Server{
		Addr:         ":3000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      routes.HospitalRouter("A", db, &jwtConfig),
		ErrorLog:     log.Default(),
	}

	serverB := http.Server{
		Addr:         ":3001",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      routes.HospitalRouter("B", db, &jwtConfig),
		ErrorLog:     log.Default(),
	}

	server.Go(func() error {
		return serverA.ListenAndServe()
	})
	server.Go(func() error {
		return serverB.ListenAndServe()
	})

	if err := server.Wait(); err != nil {
		log.Fatal(err)
	}

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := serverA.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}

	if err := serverB.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}

	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")
}
