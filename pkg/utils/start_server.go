package utils

import (
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

func StartServerWithGracefulShutdown(app *fiber.App) {
	idleChan := make(chan struct{})

	go func() {
		sigintChan := make(chan os.Signal, 1)
		signal.Notify(sigintChan, os.Interrupt)
		<-sigintChan

		if err := app.Shutdown(); err != nil {
			log.Printf("Cannot shutdown the server right now! Reason: %v", err)
		}

		close(idleChan)
	}()

	connectionAddr, _ := ConnectionUrlBuilder("fiber")

	if err := app.Listen(connectionAddr); err != nil {
		log.Printf("Couldn't start the server! Reason: %v", err)
	}

	<- idleChan
}