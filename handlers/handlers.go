package handlers

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io"
	"mail-application/email"
	"mail-application/models"
	"net/http"
)

func SendMailHandler(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Falha ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}

	var emailRequest models.EmailRequest
	if err := json.Unmarshal(requestBody, &emailRequest); err != nil {
		http.Error(w, "Falha na decodificação do JSON", http.StatusBadRequest)
		return
	}

	validate := validator.New()
	if err := validate.Struct(emailRequest); err != nil {
		http.Error(w, "Campos obrigatórios ausentes no corpo da solicitação", http.StatusBadRequest)
		return
	}

	emailSaved, err := email.SendEmail(emailRequest)
	if err != nil {
		http.Error(w, "Falha durante o processamento do email", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response, err := json.Marshal(emailSaved)
	if err != nil {
		http.Error(w, "Falha durante na codificação da resposta", http.StatusBadRequest)
		return
	}
	_, err = w.Write(response)
	if err != nil {
		http.Error(w, "Falha ao escrever resposta", http.StatusBadRequest)
		return
	}
}
