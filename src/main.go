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

	db.AutoMigrate(&models.PatientModel{}, &models.StaffModel{})

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
