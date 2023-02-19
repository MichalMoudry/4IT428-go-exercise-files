package main

import "fmt"

type EmailProtocol string

type Email struct {
	Sender    string
	Recipient string
	Message   string
}

const (
	SMTP EmailProtocol = "SMTP"
	IMAP EmailProtocol = "IMAP"
	POP3 EmailProtocol = "POP3"
)

type EmailRepository struct {
	sentEmails []Email
}

func (emailRepo *EmailRepository) Save(email Email) bool {
	emailRepo.sentEmails = append(emailRepo.sentEmails, email)
	return true
}

func (emailRepo *EmailRepository) GetSentEmails() []Email {
	return emailRepo.sentEmails
}

type IEmailRepository interface {
	Save(email Email) bool
	GetSentEmails() []Email
}

type EmailService struct {
	Repository IEmailRepository
}

func (emailService EmailService) SendWithSMTP(email Email) bool {
	return true
}

func (emailService EmailService) SendWithIMAP(email Email) bool {
	return true
}

func (emailService EmailService) SendWithPOP3(email Email) bool {
	return true
}

func (emailService EmailService) Send(email Email, protocol EmailProtocol) bool {
	var result bool
	switch protocol {
	case SMTP:
		result = emailService.SendWithSMTP(email)
	case IMAP:
		result = emailService.SendWithIMAP(email)
	case POP3:
		result = emailService.SendWithPOP3(email)
	default:
		result = emailService.SendWithSMTP(email)
	}
	if result {
		emailService.Repository.Save(email)
	}
	return result
}

func main() {
	emailRepo := EmailRepository{sentEmails: []Email{}}
	emailService := EmailService{Repository: &emailRepo}

	email := Email{Sender: "test@test.com", Recipient: "test1@test1.com", Message: "Test email content"}
	if emailService.Send(email, POP3) {
		fmt.Println(fmt.Sprintf("Email was sent. %s was used.", POP3))
	} else {
		fmt.Println("Unable to send email.")
	}
}
