package maps

var (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word  because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictonary map[string]string

func (d Dictonary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictonary) Add(word, defintion string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = defintion
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
