package utils

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func StartServerWithGracefulShutdown(a *fiber.App) {
	idelConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := a.Shutdown(); err != nil {
			log.Printf("Server is shutting down %v", err)
		}

		close(idelConnsClosed)
	}()
	log.Printf("server url, %v", os.Getenv("SERVER_URL"))
	if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Server is not running, %v", err)
	}
	<-idelConnsClosed
}

func StartServer(a *fiber.App) {
	if err := a.Listen(os.Getenv("SERVER_URL")); err != nil {
		log.Printf("Server is not running, %v", err)
	}
}
