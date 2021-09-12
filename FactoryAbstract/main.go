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
