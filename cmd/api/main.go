package main

import (
	"context"
	_ "github.com/Jumaniyozov/go-rest-template/api/docs"
	"github.com/Jumaniyozov/go-rest-template/internal/app"
	"log"
	"os/signal"
	"syscall"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer stop()

	application := app.New()

	if err := application.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
