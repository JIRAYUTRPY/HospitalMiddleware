package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/agnos/hospital-middleware/config"
	"github.com/agnos/hospital-middleware/routes"
	_ "github.com/lib/pq"
	"golang.org/x/sync/errgroup"
)

var (
	server errgroup.Group
)

func main() {

	config := config.DatabaseConfig{}
	ctx := context.Background()
	db, err := sql.Open("postgres", config.GetConnectionString())

	if err != nil {
		panic(err)
	}
	defer db.Close()

	config.Migration(ctx)

	serverA := http.Server{
		Addr:         ":3000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      routes.HospitalRouter("A"),
		ErrorLog:     log.Default(),
	}
	serverB := http.Server{
		Addr:         ":3001",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      routes.HospitalRouter("B"),
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
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")

}
