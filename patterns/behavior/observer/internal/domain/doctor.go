package domain

import "fmt"

func NewDoctor() *Doctor {
	return &Doctor{}
}

type Doctor struct{}

func (d *Doctor) Notify(data interface{}) {
	fmt.Printf("A doctor has been called for %s",
		data.(string))
}
