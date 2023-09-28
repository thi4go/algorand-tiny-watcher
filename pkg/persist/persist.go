package persist

import (
	"algorand-tiny-watcher/model"
	"encoding/json"
	"os"
)

const addressesFile = "addresses.json"

func SaveAddresses(addresses map[string]*model.AccountState) error {
	file, err := os.Create(addressesFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(addresses)
}

func LoadAddresses() (map[string]*model.AccountState, error) {
	file, err := os.Open(addressesFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	defer file.Close()

	var addresses map[string]*model.AccountState
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&addresses)
	if err != nil {
		return nil, err
	}
	return addresses, nil
}
