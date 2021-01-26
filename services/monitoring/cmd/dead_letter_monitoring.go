package cmd

import (
	"context"

	"github.com/eavillacis/velociraptor/pkg/subscriptions"
	"github.com/eavillacis/velociraptor/services/monitoring/conf"
	"cloud.google.com/go/pubsub"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deadLetterMonitoringCmd)
}

var deadLetterMonitoringCmd = &cobra.Command{
	Use:   "sub:dead-letter-monitoring",
	Short: "Start a new Dead Letter Monitoring Subscription Listener",
	Long:  "Start a new Dead Letter Monitoring Subscription Listener",
	Run: func(cmd *cobra.Command, args []string) {
		startDeadLetterMonitoring()
	},
}

func startDeadLetterMonitoring() {
	config, err := conf.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Failed to load configuration: %+v", err)
	}

	ctx := context.Background()

	pubsubClient, err := pubsub.NewClient(ctx, config.ProjectID)
	if err != nil {
		log.Fatalf("Failed to generate a new pubsub client: %+v", err)
	}

	subscription := pubsubClient.Subscription(subscriptions.DeadLetterMonitoring)

	log.Println("Dead Letter Monitoring Subscription started...")

	err = subscription.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		log.WithFields(log.Fields{"data": string(msg.Data)}).Info("Received dead letter monitoring message")

		// If necessary process each message in an appropriate way

		msg.Ack()
	})

	if err != nil {
		log.Fatalf("error receiving dead letter monitoring messages: %+v", err)
	}
}
