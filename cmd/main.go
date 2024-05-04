package main

import (
	"fmt"
	"github.com/1337Bart/improve-yourself/internal/db"
	handlerData "github.com/1337Bart/improve-yourself/internal/handlers/data"
	handlerLogin "github.com/1337Bart/improve-yourself/internal/handlers/login"
	handlerSettings "github.com/1337Bart/improve-yourself/internal/handlers/settings"
	"github.com/1337Bart/improve-yourself/internal/routes"
	"github.com/1337Bart/improve-yourself/internal/service/data"
	"github.com/1337Bart/improve-yourself/internal/service/login"
	"github.com/1337Bart/improve-yourself/internal/service/settings"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/joho/godotenv"
)

func main() {
	env := godotenv.Load()
	if env != nil {
		panic("Error loading .env file")
	}

	dbUrl := os.Getenv("DATABASE_URL")
	dbConn, err := db.InitDB(dbUrl)
	if err != nil {
		fmt.Printf("error connecting to database: %s\n", err)
		panic("error connecting to db")
	}

	loginService := login.NewLoginService(dbConn)
	settingsService := settings.NewSettingsService(dbConn)
	dataService := data.NewDataService(dbConn)

	loginHandler := handlerLogin.NewHandler(loginService, settingsService, dataService)
	settingsHandler := handlerSettings.NewHandler(settingsService)
	dataHandler := handlerData.NewHandler(settingsService, dataService)

	port := os.Getenv("PORT")
	if port == "" {
		port = ":4000"
	} else {
		port = ":" + port
	}

	app := fiber.New(fiber.Config{
		IdleTimeout: 5 * time.Second,
	})

	app.Use(compress.New())

	routes.SetRoutes(app, loginHandler, settingsHandler, dataHandler)

	go func() {
		if err := app.Listen(port); err != nil {
			log.Panic(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c

	app.Shutdown()
	fmt.Println("Shutting down server")
}
