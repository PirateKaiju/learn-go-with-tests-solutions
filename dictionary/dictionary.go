package dictionary

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("unable to find word")

func (d Dictionary) Search(key string) (string, error) {
	definition, found := d[key]

	if !found {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d *Dictionary) Add(key string, value string) error {
	return nil
}
