package usecases

import (
	"strings"

	"ws4/models"
	"ws4/repositiories"
)


func FilterCharacter (filter, value string) ([]models.Character, error) {
	characters, err := repositiories.GetCharacters()
	if err != nil {
		return nil, err
	}

	return filterFactory(characters, filter, value), nil
}

func filterFactory(characters []models.Character, filter, value string) []models.Character {
	switch filter {
	case "species":
		return filterBySpecies(characters, value)
	default:
		return characters
	}
}

func filterBySpecies(characters []models.Character, filterValue string) []models.Character {
	filterdCharacters := make([]models.Character, 0)
	for _, ch := range characters {
		if strings.ToLower(ch.Species) == strings.ToLower(filterValue) {
			filterdCharacters = append(filterdCharacters, ch)
		}
	}
	return filterdCharacters
}
