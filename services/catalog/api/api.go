package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-redis/redis"

	"github.com/eavillacis/velociraptor/services/catalog/conf"
	
	"cloud.google.com/go/pubsub"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// API ...
type API struct { 
	db              *gorm.DB	
	pubsub          *pubsub.Client
	config          *conf.Configuration
	route           *gin.RouterGroup
	handler         *gin.Engine
	redis           *redis.Client
}

// NewAPI mounts all routes
func NewAPI(conn *gorm.DB, pubsubClient *pubsub.Client, config *conf.Configuration, redis *redis.Client) *API {
	api := &API{
		db:     conn,
		config: config,
		redis:  redis,
	}

	handler := gin.New()
	handler.Use(gin.Recovery())
	// handler.Use(httputils.Logger())

	routesV1 := handler.Group("/v1")

	api.handler = handler
	api.route = routesV1
	api.pubsub = pubsubClient

	routesV1.GET("/catalog/status", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Catalog Service OK!"})
	})
	api.initRoutes()

	return api
}

// ListenAndServe starts the API
// It has graceful shutdown
func (a *API) ListenAndServe() {
	host := fmt.Sprintf(":%d", a.config.Port)

	srv := &http.Server{
		Addr:    host,
		Handler: a.handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	log.Printf("Catalog Server listening on address %s\n", host)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
