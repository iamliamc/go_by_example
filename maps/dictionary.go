package main

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// var dictionary = map[string]string{}
// OR
// var dictionary = make(map[string]string)
// never var m map[string]string

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
}
