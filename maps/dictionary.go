package main

// Dictionary for storing definitions for words
type Dictionary map[string]string

// DictionaryErr are errors that can happen while using the dictionary
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	// ErrNotFound means the defintion could not be found for the given word
	ErrNotFound = DictionaryErr("could not find the word you were looking for")

	// ErrWordExists means the caller is trying to add a word that is already in the dictioanry
	ErrWordExists = DictionaryErr("could not add word because it already exists")

	// ErrWordDoesNotExist means the call is trying to update a word that does not exist in the dictionary
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// Search finds a given word in the dictionary if it exists.
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add puts the supplied word and definition into the dictionary, if it doesn't
// already exist.
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

// Update replaces the defintion for the word with a new definition, but only
// if the word is already in the dictionary.
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

// Delete removes the word from the dictionary if it exists, otherwise it has
// no effect.
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
