package domain

import (
	"fmt"

	"github.com/nachoberaza/GO_Desing_Patterns/patterns/structural/Adapter/internal/client_interfaces"
)

type Client struct {
}

func (c *Client) InsertLightningConnectorComputer(com client_interfaces.Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}
