package cmd

import (
	"context"
	"encoding/json"

	"github.com/eavillacis/velociraptor/pkg/subscriptions"
	"github.com/eavillacis/velociraptor/services/monitoring/conf"
	"github.com/eavillacis/velociraptor/services/monitoring/worker"
	"cloud.google.com/go/pubsub"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	queueMessages "github.com/eavillacis/velociraptor/pkg/queue/messages"
	monitoringModels "github.com/eavillacis/velociraptor/services/monitoring/models"
)

func init() {
	rootCmd.AddCommand(mongoMonitoring)
}

var mongoMonitoring = &cobra.Command{
	Use:   "sub:mongo-monitoring",
	Short: "Start a new mongo Monitoring Subscription Listener",
	Long:  "Start a new mongo Monitoring Subscription Listener",
	Run: func(cmd *cobra.Command, args []string) {
		startMongoMonitoring()
	},
}

func startMongoMonitoring() {
	config, err := conf.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to load configuration: %+v", err)
	}

	ctx := context.Background()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+config.Mongo.User+":"+config.Mongo.Password+"@"+config.Mongo.Host+":"+config.Mongo.Port))
	if err != nil {
		log.Fatalf("Failed to connect to Mongo service: %+v", err)
	}

	pubsubClient, err := pubsub.NewClient(ctx, config.ProjectID)
	if err != nil {
		log.Fatalf("Failed to generate a new pubsub client: %+v", err)
	}

	subscription := pubsubClient.Subscription(subscriptions.Monitoring)

	log.Println("Mongo Monitoring Subscription started...")

	err = subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		log.WithFields(log.Fields{"data": string(msg.Data)}).Info("Received mongo monitoring message")

		var data queueMessages.LogProvider

		if err := json.Unmarshal(msg.Data, &data); err != nil {
			log.WithFields(log.Fields{"data": data}).WithError(err).Error("error unmarshalling json")
			msg.Ack()
			return
		}

		db := config.Mongo.DB
		job := worker.NewMongoProcessJob(client, data, db)

		if err := job.Run(); err != nil {

			log.WithFields(log.Fields{"body": ""}).WithError(err).Error("error running mongo monitoring job")
			deadLetterMessage := &monitoringModels.DeadLetterMessage{
				Data:         string(msg.Data),
				Subscription: subscriptions.Monitoring,
				Message:      "failed to process message on mongo monitoring",
			}

			// If the amount of times that this message has failed to be processed correctly, then publish to dead letter topic.
			messageID, err := publishDeadLetterMessage(ctx, pubsubClient, deadLetterMessage)
			if err != nil {
				log.WithFields(log.Fields{"message": msg}).WithError(err).Error("error publishing message to dead letters topic from mongo monitoring")
				msg.Nack()
				return
			}

			log.WithFields(log.Fields{"message_id": messageID}).Info("message published to dead letters topic from mongo monitoring")
			msg.Ack()
			return
		}

		msg.Ack()
	})

	if err != nil {
		log.Fatalf("error receiving dead letter monitoring messages: %+v", err)
	}
}
