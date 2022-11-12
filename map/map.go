package mapInfo

type Dictionary map[string]string
type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

const (
	NotFoundError          = DictionaryError("word not found")
	WordAlreadyExistsError = DictionaryError("word already exists")
)

func (d Dictionary) Search(key string) (string, error) {
	if value, found := d[key]; found {
		return value, nil
	}
	return "", NotFoundError
}

func (d Dictionary) Add(key, value string) (bool, error) {
	_, found := d[key]
	if found {
		return false, WordAlreadyExistsError
	}
	d[key] = value
	return true, nil
}

func (d Dictionary) Update(key, value string) (bool, error) {
	if _, err := d.Search(key); err != nil {
		return false, NotFoundError
	}
	d[key] = value
	return true, nil
}

func (d Dictionary) Delete(key string) (bool, error) {
	if _, err := d.Search(key); err != nil {
		return false, NotFoundError
	}
	delete(d, key)
	return true, nil
}
