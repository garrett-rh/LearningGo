package v1_maps

import (
	"errors"
)

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word] // second val is a boolean to indicate if key is found
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
