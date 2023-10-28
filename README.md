<h1 align="center"> 
	API Para Envio de Email em Golang
</h1>

### API Para Envio de Email.

Essa é uma API de envio de email que recebe do corpo da requisição as informações básicas para isso, conforme o JSON especificado abaixo, através do endpoint: `http://localhost:8088/send-email`.  
Além disso, salva o email enviado em uma coleção de históricos de email no MongoDB e o devolve na resposta com um Status 200 OK.

**Entrada:**
```
{
"from": "seu_email@gmail.com",
"to": "destinatario@gmail.com",
"subject": "Assunto do e-mail",
"body": "Corpo do e-mail"
}
```

**Saída:**
```
{
"ID": "Id do email salvo no histórico",
"From": "seu_email@gmail.com",
"To": "destinatario@gmail.com",
"Subject": "Assunto do e-mail",
"Body": "Corpo do e-mail",
"SentTime": "Data e hora de envio"
}
```

### 🛠 Utilizando:

As seguintes tecnologias foram utilizadas no projeto:

- ![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)
- ![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)

Além dos seguintes padrões e recursos:

- SOLID;
- SMTP do Gmail: smtp.gmail.com;
- Chi Chi;
- Validator do go-playground.