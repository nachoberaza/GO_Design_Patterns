package main

import (
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/Adapter/internal/adapters"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/Adapter/internal/domain"
	"github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/Adapter/internal/services"
)

func main() {
	client := &domain.Client{}
	mac := &services.Mac{}

	client.InsertLightningConnectorComputer(mac)

	windowsMachine := &services.Windows{}
	windowsMachineAdapter := &adapters.WindowsAdapter{
		WindowMachine: windowsMachine,
	}

	client.InsertLightningConnectorComputer(windowsMachineAdapter)
}
