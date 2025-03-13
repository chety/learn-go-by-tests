package mapInfo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"chety": "Best hacker", "rodik": "hav hav", "mirko": "boac boac"}

	t.Run("should return values for existing keys", func(t *testing.T) {
		for key, value := range dictionary {
			got, err := dictionary.Search(key)
			assert.Equal(t, value, *got)
			assert.Nil(t, err)
		}
	})

	t.Run("unknown words should throw NotFoundError", func(t *testing.T) {
		value, err := dictionary.Search("unknown word")
		if err == nil {
			t.Fatal("expected  error but got none")
		}
		assert.Nil(t, value)
		assert.Equal(t, NotFoundError, err)
	})
}

func TestAdd(t *testing.T) {
	key, value := "JS", "Most underrated programming language"

	t.Run("should add a new key", func(t *testing.T) {
		dictionary := Dictionary{}
		result, err := dictionary.Add(key, value)
		assert.True(t, result)
		assert.NoError(t, err)

		actual, err2 := dictionary.Search(key)
		assert.Equal(t, value, *actual)
		assert.Nil(t, err2)
	})

	t.Run("should throw WordAlreadyExistsError if key is already in the map", func(t *testing.T) {
		dictionary := Dictionary{key: value}
		result, err := dictionary.Add(key, value)

		assert.False(t, result)
		assert.EqualError(t, err, WordAlreadyExistsError.Error())
	})
}

func TestUpdate(t *testing.T) {

	t.Run("should update an existing key", func(t *testing.T) {
		key, value := "GO", "Most fruitful programming language"
		dictionary := Dictionary{key: value}

		result, err := dictionary.Update(key, "Most frustrated programming language")
		assert.NoError(t, err)
		assert.True(t, result)

		actual, _ := dictionary.Search(key)
		assert.Equal(t, "Most frustrated programming language", *actual)
	})

	t.Run("should return NotFoundError if key is not in the dictionary", func(t *testing.T) {
		dictionary := Dictionary{}
		result, err := dictionary.Update("key", "value")
		assert.False(t, result)
		assert.EqualError(t, err, NotFoundError.Error())
	})
}

func TestDelete(t *testing.T) {
	t.Run("should delete an existing key", func(t *testing.T) {
		dictionary := Dictionary{"key": "value"}
		result, err := dictionary.Delete("key")
		assert.True(t, result)
		assert.NoError(t, err)

		_, searchErr := dictionary.Search("key")
		assert.EqualError(t, searchErr, NotFoundError.Error())
	})

	t.Run("should return NotFoundError if key does not exist", func(t *testing.T) {
		dictionary := Dictionary{}
		result, err := dictionary.Delete("key")
		assert.False(t, result)
		assert.EqualError(t, err, NotFoundError.Error())
	})
}

//func assertString(t *testing.TB, expected, actual string) {
//	t.Helper()
//	assert.Equal(t, expected, actual)
//}
