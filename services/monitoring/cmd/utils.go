package cmd

import (
	"context"
	"encoding/json"

	"github.com/eavillacis/velociraptor/pkg/topics"
	"cloud.google.com/go/pubsub"

	monitoringModels "github.com/eavillacis/velociraptor/services/monitoring/models"
)

// publishDeadLetterMessage ...
func publishDeadLetterMessage(ctx context.Context, pubsubClient *pubsub.Client, message *monitoringModels.DeadLetterMessage) (string, error) {
	t := pubsubClient.Topic(topics.DeadLetters)

	data, err := json.Marshal(message)
	if err != nil {
		// json Marshal should never fail since the message is created from a struct
		return "", nil
	}

	result := t.Publish(ctx, &pubsub.Message{Data: data})

	messageID, err := result.Get(ctx)
	if err != nil {
		return "", err
	}

	return messageID, nil
}
