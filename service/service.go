package service

import "fmt"

type MessageService interface {
	SendChargeNotification(int) bool
}

type SMSService struct{}

type MyService struct {
	MessageService MessageService
}

func (sms SMSService) SendChargeNotification(value int) bool {
	fmt.Println("Sending production charge notification")
	return true
}

func (ms MyService) ChargeCustomer(value int) error {
	ms.MessageService.SendChargeNotification(value)
	fmt.Printf("Charging the customer for value %d\n", value)
	return nil
}
