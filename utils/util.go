package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func DoHTTPRequest(requestURL string) ([]byte, error) {
	resp, err := http.Get(requestURL)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}

func DoHTTPRequestCookie(requestURL string) ([]byte, error) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", requestURL, nil)

	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	log.Println(string(body))
	return body, nil
}

func DoPostHTTPRequest(requestURL string, bodyReq interface{}) ([]byte, error) {
	bodyInBytes, err := json.Marshal(bodyReq)
	if err != nil {
		return nil, err
	}
	resp, err := http.Post(requestURL, "application/json", bytes.NewBuffer(bodyInBytes))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func CreateFile(fileName string) (*os.File, error) {
	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}

	return file, err
}

func WriteToFile(fd *os.File, outputData []byte) error {
	_, err := fd.WriteString(string(outputData))
	if err != nil {
		return err
	}

	return nil
}

func OpenFile(fileName string) (*os.File, error) {
	endpoints, err := os.Open(fileName)

	return endpoints, err
}

// BreakdownInto will map an array into a 2D array of [][sliceSize] length
func BreakdownInto(container []string, sliceSize int) [][]string {

	var ret [][]string
	var acc []string
	var tracker int = 0
	for _, content := range container {
		acc = append(acc, content)
		if tracker == sliceSize {
			ret = append(ret, acc)
			tracker = 0
			acc = []string{}
		}

		tracker++
	}

	return ret
}
