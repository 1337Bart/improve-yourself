package main

import (
	"fmt"
	"github.com/1337Bart/improve-yourself/db"
	"github.com/1337Bart/improve-yourself/routes"
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
	err := db.InitDB(dbUrl)
	if err != nil {
		fmt.Printf("error connecting to database: %s\n", err)
		panic("error connecting to db")
	}

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

	routes.SetRoutes(app)

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
