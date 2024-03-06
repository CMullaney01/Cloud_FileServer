package main

import (
	"backend/api/middlewares"
	"backend/api/routes"
	"backend/infrastructure/mongomgmt"
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {

	initViper()

	// clean-up pending
	ctx := context.Background()
	interval := 1 * time.Minute // clean-up every minute
	fileManager := mongomgmt.NewFileManager()
	go fileManager.StartPendingUploadsCleanup(ctx, interval)

	// setup http calls
	app := fiber.New(fiber.Config{
		AppName:      "My File Server",
		ServerHeader: "Fiber",
	})
	middlewares.InitFiberMiddlewares(app, routes.InitPublicRoutes, routes.InitProtectedRoutes)

	var listenIp = viper.GetString("ListenIP")
	var listenPort = viper.GetString("ListenPort")

	log.Printf("will listen on %v:%v", listenIp, listenPort)

	err := app.Listen(fmt.Sprintf("%v:%v", listenIp, listenPort))
	log.Fatal(err)
}

func initViper() {
	// Use .env.local file
	viper.SetConfigName(".env.local")
	viper.AddConfigPath(".")
	viper.SetConfigType("env")

	// Look for environment variables prefixed with "Auth"
	viper.SetEnvPrefix("Auth")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Try reading from .env.local
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Unable to read config file: %s", err)
	} else {
		log.Println("Config file loaded successfully")
	}
}
