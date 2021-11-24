package repositiories

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"

	"ws4/models"
)

func GetCharacters() ([]models.Character, error) {
	f, err := os.Open("/Users/david.castillo/code/wizeline/w4/data/data.json")
	defer f.Close()

	if err != nil {
		return nil, errors.New("database not found")
	}

	jsonBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.New("database data error")
	}

	var characters []models.Character
	err = json.Unmarshal(jsonBytes, &characters)
	if err != nil {
		return nil, errors.New("parsing error")
	}

	return characters, nil
}
