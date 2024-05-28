package mydicts

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("could not find the word you are looking for")
	ErrCantUpdate = errors.New("cannot update word because it does not exist")
	ErrWordExists = errors.New("cannot add word because it already exists")
	ErrCantDelete = errors.New("cannot delete word because it does not exist")
)

// Search for a word in the dictionary
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", ErrNotFound
}

// Add a new word to the dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	if err == nil {
		return ErrWordExists
	}
	if err == ErrNotFound {
		d[word] = definition
		return nil
	}
	return err
}

// Update a word in the dictionary
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	if err == nil {
		d[word] = definition
		return nil
	}
	if err == ErrNotFound {
		return ErrCantUpdate
	}
	return err
}

// Delete a word from the dictionary
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	if err == nil {
		delete(d, word)
		return nil
	}
	if err == ErrNotFound {
		return ErrCantDelete
	}
	return err
}