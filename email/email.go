package email

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"mail-application/database"
	"mail-application/models"
	"net/smtp"
	"time"
)

func SendEmail(emailRequest models.EmailRequest) (models.Email, error) {
	auth := smtp.PlainAuth("", smtpUsername, smtpAppPassword, smtpServer)

	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s\r\n", emailRequest.To, emailRequest.Subject, emailRequest.Body)

	err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, emailRequest.From, []string{emailRequest.To}, []byte(message))
	if err != nil {
		log.Println("Erro ao enviar e-mail: ", err)
		return models.Email{}, err
	}

	log.Println("E-mail enviado com sucesso para: ", emailRequest.To)

	emailSaved, err := saveEmailSent(emailRequest)
	if err != nil {
		log.Println("Erro ao salvar email: ", err)
		return models.Email{}, err
	}

	return emailSaved, nil
}

func saveEmailSent(emailRequest models.EmailRequest) (models.Email, error) {
	client, err := database.ConnectToMongoDB()
	if err != nil {
		log.Println("Erro ao conectar no database: ", err)
		return models.Email{}, err
	}

	collection := client.Database("email_history").Collection("email_history")

	emailToSave := models.Email{
		From:     emailRequest.From,
		To:       emailRequest.To,
		Subject:  emailRequest.Subject,
		Body:     emailRequest.Body,
		SentTime: time.Now(),
	}

	emailInserted, err := collection.InsertOne(context.Background(), emailToSave)
	if err != nil {
		log.Println("Erro ao inserir no database: ", err)
		return models.Email{}, err
	}

	var emailSaved models.Email
	filter := bson.M{"_id": emailInserted.InsertedID}
	err = collection.FindOne(context.Background(), filter).Decode(&emailSaved)
	if err != nil {
		log.Println("Erro ao recuperar email salvo:  ", err)
		return models.Email{}, err
	}

	log.Println(emailSaved)
	log.Println("Email salvo no histórico com sucesso! Id: ", emailSaved.ID)

	return emailSaved, nil
}
