package publish

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/eavillacis/velociraptor/pkg/errors"
	"cloud.google.com/go/pubsub"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

//Now is a wrapper of Publish method to publish some data in pubsub service.
func Now(c *gin.Context, ps *pubsub.Client, topic string, data interface{}) *errors.Fault {
	ctx := context.Background()
	t := ps.Topic(topic)
	dataJSON, err := json.Marshal(data)
	if err != nil {
		bug := errors.Create(http.StatusInternalServerError, err.Error()).StackTrace()
		return bug
	}
	trackingID := c.Request.Header.Get("Service-Traceid")
	attributes := map[string]string{
		"service_traceid": trackingID,
	}
	message := &pubsub.Message{
		Data:       dataJSON,
		Attributes: attributes,
	}

	outcome := t.Publish(ctx, message)
	messageID, err := outcome.Get(ctx)
	if err != nil {
		bug := errors.Create(http.StatusInternalServerError, err.Error()).StackTrace()
		return bug
	}
	log.WithFields(
		log.Fields{
			"topic":           topic,
			"message_id":      messageID,
			"service_traceid": trackingID,
		},
	).Info("pubsub information")
	return nil
}
