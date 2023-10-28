package models

type EmailRequest struct {
	From    string `json:"from" validate:"required"`
	To      string `json:"to" validate:"required"`
	Subject string `json:"subject" validate:"required"`
	Body    string `json:"body" validate:"required"`
}
