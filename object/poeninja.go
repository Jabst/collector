package object

import (
	"consumer/collector/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Endpoints lists all the endpoints provided in the json file
type Endpoints struct {
	PoeNinja []map[string]string `json:"poe_ninja"`
}

type Resources struct {
	Items map[string]Resource
}

type Resource struct {
	Value string
}

type Dictionary struct {
	Key []string
}

type NinjaDatabase struct {
	Products   Target
	Dictionary Dictionary
}

type CategoryInformation struct {
	Lines []Data `json:"lines"`
}

type Data struct {
	Name  string  `json:"name"`
	CVal  float64 `json:"chaosValue"`
	ExVal float64 `json:"exaltedValue"`
}

type Target struct {
	Items map[string]Price
}

type Price struct {
	ExVal float64
	CVal  float64
}

const endpointsFilename = "./config/endpoints.json"

// InitializeInMemory initializes the required poe-ninja database entities
func InitializeInMemory() (NinjaDatabase, error) {
	endpoints, err := utils.OpenFile(endpointsFilename)
	if err != nil {
		return NinjaDatabase{}, err
	}

	defer endpoints.Close()

	resources, err := bootstrapData(endpoints)
	if err != nil {
		return NinjaDatabase{}, err
	}

	values, dictionary := initEntities(resources)

	products := Target{
		Items: make(map[string]Price),
	}

	for _, key := range dictionary.Key {
		information, err := utils.DoHTTPRequest(values.Items[key].Value)
		if err != nil {
			return NinjaDatabase{}, err
		}
		var dataStruct CategoryInformation

		if len(information) == 0 {
			log.Println(fmt.Sprintf("Failed to fetch data for: %s", key))
			continue
		}

		if err := json.Unmarshal(information, &dataStruct); err != nil {
			return NinjaDatabase{}, err
		}

		for _, elem := range dataStruct.Lines {
			products.Items[elem.Name] = Price{
				ExVal: elem.ExVal,
				CVal:  elem.CVal,
			}

		}
	}

	return NinjaDatabase{
		Products:   products,
		Dictionary: dictionary,
	}, nil
}

// Reads the file and unmarshal the byte data into the struct
func bootstrapData(fileDescriptor *os.File) (Endpoints, error) {
	var endpoint Endpoints

	byteValue, err := ioutil.ReadAll(fileDescriptor)
	if err != nil {
		return Endpoints{}, err
	}

	if err := json.Unmarshal(byteValue, &endpoint); err != nil {
		return Endpoints{}, err
	}

	return endpoint, nil
}

func initEntities(resources Endpoints) (Resources, Dictionary) {
	var values = Resources{
		Items: make(map[string]Resource),
	}

	var dictionary = Dictionary{
		Key: make([]string, 0, 0),
	}

	for _, resource := range resources.PoeNinja {

		for key, value := range resource {
			values.Items[key] = Resource{
				Value: value,
			}
			dictionary.Key = append(dictionary.Key, key)
		}
	}

	return values, dictionary
}
