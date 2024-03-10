package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"homework/internal/handler"
	"homework/internal/service"
	"homework/internal/storage"
	"homework/pkg/server"
	"log"
	"os"
)

func main() {
	repo := storage.NewDeviceStorage()
	serviceApp := service.NewService(repo)

	handlerApp := handler.NewHandler(serviceApp)

	serverApp := new(server.Server)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error while loading env var: %s", err.Error())
	}
	addr := os.Getenv("ADDR")
	if addr == "" {
		addr = "localhost"
	}
	port := os.Getenv("PORT")
	if port == "" {
		addr = "8080"
	}
	fmt.Printf("Starting server at %s:%s/\n", addr, port)

	if err := serverApp.Run(addr, port, handlerApp.InitRoutes()); err != nil {
		log.Fatalf("error while running the server: %s", err.Error())
	}
}
