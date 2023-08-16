package single

import (
	"fmt"
	"sync"
)

type single struct {
}

var singleInstance *single
var lock = &sync.Mutex{}

func GetInstance() *single {
	lock.Lock()
	defer lock.Unlock()

	if singleInstance != nil {
		fmt.Println("Single instance already created.")
		return singleInstance
	}

	fmt.Println("Creating single instance now.")
	singleInstance = &single{}

	return singleInstance
}
