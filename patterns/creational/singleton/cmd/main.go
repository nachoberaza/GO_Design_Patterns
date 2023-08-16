package main

import (
	"fmt"
	"github.com/baldurdevs/go_desing_patterns/patterns/creational/singleton/single"
)

func main() {
	for i := 0; i < 10; i++ {
		go single.GetInstance()
	}

	_, err := fmt.Scanln()
	if err != nil {
		return
	}
}
