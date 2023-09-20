package implementations

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/template_method/internal/templates"
)

type Email struct {
	Otp templates.Otp
}

func (s *Email) GenRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("EMAIL: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *Email) SaveOTPCache(otp string) {
	fmt.Printf("EMAIL: saving otp: %s to cache\n", otp)
}

func (s *Email) GetMessage(otp string) string {
	return "EMAIL OTP for login is " + otp
}

func (s *Email) SendNotification(message string) error {
	fmt.Printf("EMAIL: sending email: %s\n", message)
	return nil
}
