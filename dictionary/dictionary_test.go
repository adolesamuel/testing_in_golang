package dict

import "testing"

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func asssertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assetDefinition(t testing.TB, d Dictionary, word, definition string) {
	t.Helper()

	got, err := d.Search(word)
	if err != nil {
		t.Fatal("shold find added word:", err)
	}
	assertStrings(t, got, definition)
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")

		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		if got == nil {
			t.Fatal("expected to get an error")
		}

		asssertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {

		dictionary := Dictionary{}

		word := "test"
		definition := "this is just a test"

		dictionary.Add(word, definition)

		assetDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definiition := "this is just a test"
		dictionary := Dictionary{word: definiition}
		err := dictionary.Add(word, "new test")

		asssertError(t, err, ErroWordExists)
		assetDefinition(t, dictionary, word, definiition)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		asssertError(t, err, nil)
		assetDefinition(t, dictionary, word, newDefinition)

	})
	t.Run("new word", func(t *testing.T) {

		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		asssertError(t, err, ErrWordDoesNotExist)

	})

}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}

		err := dictionary.Delete(word)

		asssertError(t, err, nil)

		_, err = dictionary.Search(word)

		asssertError(t, err, ErrNotFound)
	})

	t.Run("non-existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}

		err := dictionary.Delete(word)

		asssertError(t, err, ErrWordDoesNotExist)
	})
}
