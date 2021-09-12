package main

import "fmt"

/* Se enviaran notificaciones que pueden ser SMS o EMAILS*/

type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

/*Type especificos*/
type SMSNotification struct{}

/* Receiver Function */
func (SMSNotification) SendNotification() {
	fmt.Println("Enviando notificación vía SMS")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct{}

/* Receiver Function */
func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

type EmailNotification struct{}

/* Receiver Function*/
func (EmailNotification) SendNotification() {
	fmt.Println("Enviando notificación vía EMAIL")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct{}

/* Receiver Function */
func (EmailNotificationSender) GetSenderMethod() string {
	return "EMAIL"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "GMAIL"
}

/* Maneja el type de notificacion que queremos enviar*/
func getNotificacionFactory(typeNotification string) (INotificationFactory, error) {
	if typeNotification == "SMS" {
		return &SMSNotification{}, nil
	}

	if typeNotification == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("No se proporcionado el type de notificación correcto.")
}

/* Se encargará de enviar la notificación de acuerdo al typo */
func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func getChannel(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderChannel())
}

func main() {
	smsFactory, _ := getNotificacionFactory("SMS")
	emailFactory, _ := getNotificacionFactory("Email")

	sendNotification(smsFactory)
	sendNotification(emailFactory)
}
