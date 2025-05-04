package mock

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type EmailDatabase map[string][]string
type PhoneDatabase map[string][]string

func loadJSON(file string, target interface{}) error {
	path := filepath.Join("internal", "data_api", "mock", file)
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, target)
}

func LoadMocks() (emails EmailDatabase, phones PhoneDatabase, err error) { //revisar ponteiros
	err = loadJSON("emails_by_cnpj.json", &emails)
	if err != nil {
		return nil, nil, fmt.Errorf("error loading json: %v", err)
	}

	err = loadJSON("phones_by_cnpj.json", &phones)
	if err != nil {
		return nil, nil, fmt.Errorf("error loading json: %v", err)
	}

	return emails, phones, nil
}
