package api

import (
	"consumer/collector/utils"
	"fmt"
	"log"
)

func (api *API) ShipData(message []byte) {
	for _, elem := range api.DiscordBots.DiscordBots {
		log.Println(fmt.Sprintf("Sending data to %s", elem.Name))

		_, err := utils.DoPostHTTPRequest(elem.Address, message)
		if err != nil {
			panic(err)
		}
	}
}
