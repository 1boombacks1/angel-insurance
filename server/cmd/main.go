package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/1boombacks1/botInsurance/config"
	"github.com/1boombacks1/botInsurance/internal/controllers"
	"github.com/1boombacks1/botInsurance/internal/services"
	"github.com/1boombacks1/botInsurance/utils"
	"github.com/gofiber/fiber/v2"
)

// var (
// 	pathToFont = "utils/font/FiraCode-Medium.ttf"
// 	fontSize   = 100.0
// )

func main() {
	// Configuration
	cfg, err := config.NewConfig(".env")
	if err != nil {
		log.Fatalf("Config err: %s", err)
	}

	// Database
	// log.Printf("Initializing postgres...")
	// conn, _ := postgres.NewPGXPool(cfg.Uri)
	// defer conn.Close()

	// // Repositories
	// log.Print("Initializing repositories...")
	// repos := repositories.NewRepositories(conn)

	// Services
	log.Print("Initializing services...")
	photoHandler := utils.NewInsurancePhotoHandler(cfg.ApiKey, cfg.PathToFont, cfg.PathToPhoto, 100.0)
	services := services.NewServices(photoHandler)

	// Fiber router
	log.Print("Initializing handlers and routes...")
	app := fiber.New()
	controllers.NewRouter(app, services)

	serverNotify := make(chan error)
	go func() {
		// Fiber server
		log.Print("Starting fiber server...")
		serverNotify <- app.Listen(fmt.Sprintf(":%s", cfg.Server.Port))
		close(serverNotify)
	}()

	// Waiting signal
	log.Print("Configuring graceful shutdown...")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case s := <-quit:
		log.Print("app - Main - signal: " + s.String())
	case err = <-serverNotify:
		log.Print(fmt.Errorf("app - Main - server notify: %w", err))
	}

	// Graceful shutdown
	log.Print("Gracefully shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = app.ShutdownWithContext(ctx); err != nil {
		log.Print(fmt.Errorf("app - Main - app.ShutdownWithContext: %w", err))
	}
	// path := filepath.FromSlash("D:/CODING/GO/hackaton/photos")
	// createDirIfNotExist(path)

	// files, err := os.ReadDir(path)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// apiKey := "1ff57fff7a8f4e27f3661ffa89a430fc9b297252"
	// photoHandler := utils.NewInsurancePhotoHandler(apiKey, pathToFont, path, fontSize)

	// for _, file := range files {
	// 	f, err := os.Open(filepath.Join(path, file.Name()))
	// 	if err != nil {
	// 		log.Printf("Failed to load file: %v", err)
	// 		continue
	// 	}
	// 	// photoHandler.RegisterMetadata(f)
	// 	_, err = photoHandler.IsCorrectResolution(f)
	// 	if err != nil {
	// 		log.Printf("Failed to define resolution is correct or not: %v", err)
	// 		continue
	// 	}
	// }

	// id, err := repos.Client.CreateClient(context.Background(), models.Client{
	// 	LastName:   "Николаев",
	// 	FirstName:  "Яков",
	// 	Patronymic: "Алексеевич",
	// 	Phone:      "+79141043058",
	// 	LinkToChat: "https://t.me/boombacks",
	// 	Login:      "boombacks",
	// 	Password:   "password",
	// })

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// client, err := repos.Client.GetClientById(context.Background(), id)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(client)
	// }
}

func createDirIfNotExist(path string) {
	os.Mkdir(path, 0666)
}
