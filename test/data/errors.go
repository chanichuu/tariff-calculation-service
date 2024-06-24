package data

import (
	"errors"
	"fmt"
)

func FieldValidationError(keyValuePairs [][2]string) error {
	detailString := ""

	for _, keyValuePair := range keyValuePairs {
		detailString += fmt.Sprintf("Invalid value: %s", keyValuePair[0])
	}
	return errors.New(detailString)
}
