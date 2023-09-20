package main

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/template_method/internal/implementations"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/behavior/template_method/internal/templates"
)

func main() {

	smsOTP := &implementations.Sms{}
	o := templates.Otp{
		IOtp: smsOTP,
	}

	if err := o.GenAndSendOTP(4); err != nil {
		panic("error genning and sending otp by sms")
	}

	fmt.Println("")
	emailOTP := &implementations.Email{}
	o = templates.Otp{
		IOtp: emailOTP,
	}

	if err := o.GenAndSendOTP(4); err != nil {
		panic("error genning and sending otp by email")
	}

}
