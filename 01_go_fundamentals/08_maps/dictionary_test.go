package mymaps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("Known Word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("Unknown Word", func(t *testing.T) {
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertErrors(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("New Word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("Existing Word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new test")

		assertErrors(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Existing Word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertDefinition(t, dictionary, word, newDefinition)
		assertNoErrors(t, err)
	})

	t.Run("None Existent Word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertErrors(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"

	dictionary := Dictionary{word: "this is just a test"}

	dictionary.Delete(word)

	_, err := dictionary.Search(word)

	assertErrors(t, err, ErrNotFound)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("Expected %q got %q", want, got)
	}
}

func assertErrors(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertNoErrors(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("Expected no error got %q", got)
	}
}
