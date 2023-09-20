package implementations

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/template_method/internal/templates"
)

type Sms struct {
	Otp templates.Otp
}

func (s *Sms) GenRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *Sms) SaveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *Sms) GetMessage(otp string) string {
	return "SMS OTP for login is " + otp
}

func (s *Sms) SendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}
