package models

import "time"

type Email struct {
	ID       string `bson:"_id,omitempty"`
	From     string
	To       string
	Subject  string
	Body     string
	SentTime time.Time
}
