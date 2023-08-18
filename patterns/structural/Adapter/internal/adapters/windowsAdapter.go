package adapters

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/Adapter/internal/services"
)

type WindowsAdapter struct {
	WindowMachine *services.Windows
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.WindowMachine.InsertIntoUSBPort()
}
