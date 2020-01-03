package object

import (
	"consumer/collector/utils"
	"encoding/json"
	"io/ioutil"
	"os"
)

type Accounts struct {
	Accounts []struct {
		AccountName string `json:"account_name"`
		League      string `json:"league"`
	} `json:"accounts`
}

type TrackAccounts struct {
	Accounts []TrackAccount
}

type TrackAccount struct {
	AccountName string
	League      string
}

const accountsFilename = "./config/accounts.json"

// Initializes the accounts to be tracked in memory
func InitializeAccounts() (TrackAccounts, error) {

	accountFileDescriptor, err := utils.OpenFile(accountsFilename)
	if err != nil {
		return TrackAccounts{}, nil
	}

	defer accountFileDescriptor.Close()

	accounts, err := bootstrapAccountData(accountFileDescriptor)
	if err != nil {
		return TrackAccounts{}, nil
	}

	ret := TrackAccounts{
		Accounts: make([]TrackAccount, 0),
	}

	for _, acc := range accounts.Accounts {
		var tmp = TrackAccount{
			AccountName: acc.AccountName,
			League:      acc.League,
		}
		ret.Accounts = append(ret.Accounts, tmp)
	}

	return ret, nil

}

func bootstrapAccountData(fileDescriptor *os.File) (Accounts, error) {
	var accounts Accounts

	byteValue, err := ioutil.ReadAll(fileDescriptor)
	if err != nil {
		return Accounts{}, nil
	}

	if err := json.Unmarshal(byteValue, &accounts); err != nil {
		return Accounts{}, nil
	}

	return accounts, nil
}
