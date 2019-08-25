package maps

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")

type Dictonary map[string]string

func (d Dictonary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictonary) Add(word, defintion string) {
	d[word] = defintion
}
