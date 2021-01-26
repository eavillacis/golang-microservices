package cmd

import (
	"context"

	"github.com/eavillacis/velociraptor/services/catalog/api"
	"github.com/eavillacis/velociraptor/services/catalog/conf"
	"cloud.google.com/go/pubsub"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	// Import PostgreSQL Database Driver
	_ "github.com/lib/pq"
)

func init() {
	rootCmd.AddCommand(serveCmd)
} 

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start a new Catalog REST server",
	Long:  `Start a new Catalog REST server`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	config, err := conf.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to load configuration: %+v", err)
	}

	conn, err := gorm.Open("postgres", config.DB.URL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %+v", err)
	}
	defer conn.Close()

	redis := redis.NewClient(&redis.Options{
		Addr:     config.ErrorTracker.URL,
		Password: "",
		DB:       0,
	})
	_, err = redis.Ping().Result()
	if err != nil {
		log.Fatalf("Failed to connect to redis database: %+v", err)
	}

	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx, config.ProjectID)
	if err != nil {
		log.Fatalf("Failed to generate a new pubsub client: %+v", err)
	}

	api := api.NewAPI(conn, pubsubClient, config, redis)

	api.ListenAndServe()
}
