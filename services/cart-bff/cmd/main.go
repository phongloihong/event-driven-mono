package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/phongloihong/event-driven-mono/libs/http/ginServer"
	"github.com/phongloihong/event-driven-mono/services/cart-bff/contexts"
	"github.com/phongloihong/event-driven-mono/services/cart-bff/routes"
)

const (
	serverPort         = ":8080"
	shutdownTimeout    = 5 * time.Second
	serviceContextTime = 30 * time.Second
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), serviceContextTime)
	defer cancel()

	servicesCtx := contexts.NewServiceContext(ctx)
	srv := setupServer(servicesCtx)

	// Start server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// Shutdown with timeout
	ctx, cancel = context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}

func setupServer(servicesCtx *contexts.ServiceContext) *http.Server {
	ginCfg := ginServer.GinConfig{
		Port: serverPort[1:],
	}
	r := ginServer.NewHttpServer(&ginCfg)
	r.Use(gin.Recovery())
	routes.InitRoutes(r, servicesCtx)

	return &http.Server{
		Addr:    serverPort,
		Handler: r,
	}
}
