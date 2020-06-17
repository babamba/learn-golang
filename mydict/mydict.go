package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errWordExist  = errors.New("That word already exist")
	errCantUpdate = errors.New("Can't update non-existing word")
	errCantDelete = errors.New("Can't Delete Word Not Found")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word th the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word) // 단어가 존재하지않는다면.

	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExist
	}
	return nil

	// if err == errNotFound {
	// 	d[word] = def
	// } else if err === nil {
	// 	return errWordExist
	// }
	// return nil
}

// Update word
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word) // 단어가 존재하지않는다면.
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete word
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errCantDelete
	}
	return nil
}
