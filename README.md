<h1 align="center"> 
    API Para Envio de Email em Golang
</h1>

### API Para Envio de Email

Esta é uma API de envio de email que recebe informações básicas no corpo da requisição em um formato JSON. Ela também salva o email enviado em uma coleção de históricos de email no MongoDB e devolve o email na resposta com um Status 200 OK.

**Entrada** (<font color="red">todos os campos são obrigatórios</font>):
```json
{
    "from": "seu_email@gmail.com",
    "to": "destinatario@gmail.com",
    "subject": "Assunto do e-mail",
    "body": "Corpo do e-mail"
}
```

**Saída:**
```json
{
    "ID": "ID do email salvo no histórico",
    "From": "seu_email@gmail.com",
    "To": "destinatario@gmail.com",
    "Subject": "Assunto do e-mail",
    "Body": "Corpo do e-mail",
    "SentTime": "Data e hora de envio"
}

```

**config.go:**
```go
package email

const (
	smtpServer      = "smtp.gmail.com"
	smtpPort        = 587
	smtpUsername    = "example@gmail.com"
	smtpAppPassword = "exampleAppPassword"
)
```
O arquivo **[config.go](/email/config.go)**, do diretório **[/email](/email)**, exibido acima deve ter as informações genéricas alteradas conforme as informações do SMTP do Gmail do utilizador da API.

### 🛠 Utilizando:

As seguintes tecnologias foram utilizadas no projeto:

- ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
- ![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)

Além dos seguintes padrões e recursos:

- SOLID;
- [SMTP do Gmail](https://support.google.com/a/answer/176600?hl=pt#:~:text=filtrar%20mensagens%20suspeitas.-,O%20nome%20de%20dom%C3%ADnio%20totalmente%20qualificado%20do%20servi%C3%A7o%20SMTP%20%C3%A9,Protocolos%20SSL%20e%20TLS);
- [Chi](https://github.com/go-chi/chi)
- [Validator do go-playground](https://pkg.go.dev/github.com/go-playground/validator/v10)