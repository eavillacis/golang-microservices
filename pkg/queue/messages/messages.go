package messages

import (
	"time"
)

// LogProvider ...
type LogProvider struct {
	IsError        bool        `bson:"is_error"`
	LogCreatedAt   time.Time   `bson:"log_datetime"`
	LogLocation    string      `bson:"log_location"`
	LogDescription string      `bson:"log_description"`
	LogReference   string      `bson:"log_reference"`
	LogRequest     interface{} `bson:"log_request"`
	LogResponse    interface{} `bson:"log_response"`
}

