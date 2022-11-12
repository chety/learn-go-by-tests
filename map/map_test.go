package mapInfo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"chety": "Best hacker", "rodik": "hav hav", "mirko": "boac boac"}
	t.Run("should return values for existing keys", func(t *testing.T) {
		for key, value := range dictionary {
			actual, _ := dictionary.Search(key)
			expected := value
			assert.Equal(t, expected, actual)
		}
	})

	t.Run("unknown words should throw NotFoundError", func(t *testing.T) {
		_, err := dictionary.Search("unknown word")
		if err == nil {
			t.Fatal("expected  error but got none")
		}
		assert.Equal(t, NotFoundError, err)
	})
}

func TestAdd(t *testing.T) {
	key, value := "JS", "Most underrated programming language"
	t.Run("should add a new key", func(t *testing.T) {
		dictionary := Dictionary{}
		_, err := dictionary.Add(key, value)
		assert.NoError(t, err)
		actual, _ := dictionary.Search(key)
		assert.Equal(t, value, actual)
	})

	t.Run("should throw WordAlreadyExistsError if key is already in the map", func(t *testing.T) {
		dictionary := Dictionary{key: value}
		_, err := dictionary.Add(key, value)
		assert.EqualError(t, err, WordAlreadyExistsError.Error())
	})
}

func TestUpdate(t *testing.T) {
	t.Run("should update an existing key", func(t *testing.T) {
		key, value := "GO", "Most fruitful programming language"
		dictionary := Dictionary{key: value}

		_, err := dictionary.Update(key, "Most frustrated programming language")

		assert.NoError(t, err)
		actual, _ := dictionary.Search(key)
		assert.Equal(t, "Most frustrated programming language", actual)
	})

	t.Run("should return NotFoundError if key is not in the dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		_, err := dictionary.Update("key", "value")
		assert.EqualError(t, err, NotFoundError.Error())
	})
}

func TestDelete(t *testing.T) {
	t.Run("should delete an existing key", func(t *testing.T) {
		dictionary := Dictionary{"key": "value"}
		_, err := dictionary.Delete("key")
		assert.NoError(t, err)
		_, searchErr := dictionary.Search("key")
		assert.EqualError(t, searchErr, NotFoundError.Error())
	})

	t.Run("should return NotFoundError if key does not exist", func(t *testing.T) {
		dictionary := Dictionary{}
		_, err := dictionary.Delete("key")
		assert.EqualError(t, err, NotFoundError.Error())
	})
}

//func assertString(t *testing.T, expected, actual string) {
//	t.Helper()
//	assert.Equal(t, expected, actual)
//}
