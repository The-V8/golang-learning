package maps

const (
	ErrNotFound          = DictionaryErr("could not find the specific item")
	ErrTermExists        = DictionaryErr("term already exists")
	ErrTermDoesNotExists = DictionaryErr("term not available")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(term string) (string, error) {
	definition, ok := d[term]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(term, definition string) error {
	_, err := d.Search(term)

	switch err {
	case ErrNotFound:
		d[term] = definition
	case nil:
		return ErrTermExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(term, newDefinition string) error {
	_, err := d.Search(term)

	switch err {
	case nil:
		d[term] = newDefinition
	case ErrNotFound:
		return ErrTermDoesNotExists
	}

	return err
}

func (d Dictionary) Delete(term string) {
	delete(d, term)
}
