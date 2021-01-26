package worker

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/eavillacis/velociraptor/pkg/errors"
	
	queueMessages "github.com/eavillacis/velociraptor/pkg/queue/messages"
	log "github.com/sirupsen/logrus"
)

// MongoProcessJob struct
type MongoProcessJob struct {
	client *mongo.Client
	data   queueMessages.LogProvider
	db     string
}

// NewMongoProcessJob Job that handle mongo logs
func NewMongoProcessJob(client *mongo.Client, data queueMessages.LogProvider, db string) *MongoProcessJob {
	return &MongoProcessJob{
		client: client,
		data:   data,
		db:     db,
	}
}

// Run Runnable command
func (j *MongoProcessJob) Run() error {

	log.Info("init mongo worker...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := j.client.Database(j.db).Collection(j.data.LogReference)

	jsonData, err := bson.Marshal(j.data)
	if err != nil {
		errors.Wrap(err, "error unmarshaling json")
	}

	res, err := collection.InsertOne(ctx, jsonData)
	if err != nil {
		return errors.Wrap(err, "error retrieving line item")
	}

	id := res.InsertedID

	log.Info("inserted_id: ", id)

	return nil
}
