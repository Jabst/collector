package object

import (
	"consumer/collector/utils"
	"encoding/json"
	"io/ioutil"
	"os"
)

// ConfigDiscordBots lists all the endpoints provided in the json file
type ConfigDiscordBots struct {
	DiscordBots []BotAddress `json:"discord_bots"`
}

type BotAddress struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

type DiscordBots struct {
	DiscordBots []Bots
}

type Bots struct {
	Name    string
	Address string
}

const discordBotsFilename = "./config/discord.json"

// InitializeDiscordBots initializes the required discord endpoints entities
func InitializeDiscordBots() (DiscordBots, error) {
	endpoints, err := utils.OpenFile(discordBotsFilename)
	if err != nil {
		return DiscordBots{}, err
	}

	defer endpoints.Close()

	var dataInModel DiscordBots

	dataInModel.DiscordBots = make([]Bots, 0)

	discordBots, err := bootstrapDiscordBotsData(endpoints)
	if err != nil {
		return DiscordBots{}, err
	}

	for _, elem := range discordBots.DiscordBots {
		dataInModel.DiscordBots = append(dataInModel.DiscordBots, Bots{Name: elem.Name, Address: elem.Address})
	}

	return dataInModel, nil
}

// Reads the file and unmarshal the byte data into the struct
func bootstrapDiscordBotsData(fileDescriptor *os.File) (ConfigDiscordBots, error) {
	var endpoint ConfigDiscordBots

	byteValue, err := ioutil.ReadAll(fileDescriptor)
	if err != nil {
		return ConfigDiscordBots{}, err
	}

	if err := json.Unmarshal(byteValue, &endpoint); err != nil {
		return ConfigDiscordBots{}, err
	}

	return endpoint, nil
}
