package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/tkircsi/production-ready-api/internal/database"
	transportHTTP "github.com/tkircsi/production-ready-api/internal/transport/http"
)

type App struct{}

func main() {
	fmt.Println("App is started")
	app := App{}

	_, err := database.NewDatabse()
	if err != nil {
		log.Fatal("DB error")
	}
	app.Run()
}

func (app *App) Run() {
	fmt.Println("Setting up our app")
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	srv := http.Server{
		Addr:    ":5000",
		Handler: handler.Router,
	}
	idleClose := make(chan interface{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		fmt.Println("\rHTTP server is shutting down....")
		if err := srv.Shutdown(context.Background()); err != nil {
			fmt.Printf("HTTP server shutdown error: %v", err)
		}
		close(idleClose)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("HTTP server error: %v", err)
	}
	<-idleClose
}
