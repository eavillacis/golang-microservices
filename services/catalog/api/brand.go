package api

import (
	"github.com/eavillacis/velociraptor/pkg/topics"
	"encoding/json"
	"net/http"
	"context"
	"strings"
	"time"

	"github.com/eavillacis/velociraptor/pkg/errors"
	"github.com/eavillacis/velociraptor/pkg/httputils"
	"github.com/eavillacis/velociraptor/services/catalog/models"
	"cloud.google.com/go/pubsub"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	
	log "github.com/sirupsen/logrus"
	queueMessages "github.com/eavillacis/velociraptor/pkg/queue/messages"
)

var (
	errInvalidBrandID   = errors.New("invalid brand id")
	errInvalidBrandName = errors.New("brand name already exists")
)

type brandBody struct {
	Name string `json:"name" binding:"max=40,required"`
}

// CreateBrand creates a new brand
func (a *API) CreateBrand(c *gin.Context) {

	body := &brandBody{}
	logData := queueMessages.LogProvider{
		IsError:        true,
		LogCreatedAt:   time.Now(),
		LogDescription: "Create Brand",
		LogLocation:    "velociraptor.services.catalog.brands - create",
		LogReference:   "create_brand",
		LogRequest:     body,
	}

	if err := c.ShouldBindJSON(body); err != nil {
		logData.LogResponse = "Invalid Body"
		publishLogMessage(&logData, a.config.ProjectID)
		c.JSON(http.StatusBadRequest, httputils.ErrorResponse(errors.InvalidBody(err)))
		return
	}

	//Check if brand exists
	var brandByName models.Brand
	query := a.db
	query = query.Where("name ilike ?", body.Name)
	if result := query.First(&brandByName); result.Error != nil {
		if result.RecordNotFound() {
			brandID, err := uuid.NewV4()
			if err != nil {
				
				logData.LogResponse = "Error generating brand uuid"
				publishLogMessage(&logData, a.config.ProjectID)

				log.WithError(err).Error("error generating brand uuid")
				c.JSON(http.StatusInternalServerError, httputils.ErrorResponse(errors.InternalError(err)))
				return
			}

			brand := models.Brand{
				ID:          brandID,
				Name:        strings.TrimSpace(body.Name),
			}

			if result := a.db.Create(&brand); result.Error != nil {
				log.WithFields(log.Fields{
					"name": brand.Name,
				}).WithError(result.Error).Error("error creating brand")
				c.JSON(http.StatusInternalServerError, httputils.ErrorResponse(errors.InternalError(result.Error)))
				return
			}

			logData.LogResponse = brand
			publishLogMessage(&logData, a.config.ProjectID)

			c.JSON(http.StatusCreated, brand)
			return
		}

		log.WithFields(log.Fields{
			"brand_name": body.Name,
		}).WithError(result.Error).Error("Error retrieving brand")

		logData.LogResponse = "error retrieving brand"
		publishLogMessage(&logData, a.config.ProjectID)

		c.JSON(http.StatusInternalServerError, httputils.ErrorResponse(errors.InternalError(errors.Wrap(result.Error, "error retrieving brand"))))
		return
	}

	logData.LogResponse = "Brand already exists"
	publishLogMessage(&logData, a.config.ProjectID)

	c.JSON(http.StatusConflict, httputils.ErrorResponse(errors.InvalidBody(errInvalidBrandName)))
	return
}

func publishLogMessage(message *queueMessages.LogProvider, projectID string) (string, error) {
	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to generate a new pubsub client: %+v", err)
	}
	t := pubsubClient.Topic(topics.Monitoring)

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
