package main

import "testing"

// %v: Default format. The value's default format is used.
// %s: Formats the value as a string. If the value is not already a string, it is converted using the String() method of the value's type or by calling strconv.Itoa().
// %q: Formats the value as a double-quoted string, safely escaping non-printable characters with escape sequences like \n, \t, etc.
// %x: Formats the value as a hexadecimal number with lowercase letters.
// %X: Formats the value as a hexadecimal number with uppercase letters.
// %T: Formats the value's type.
// %p: Formats the value's memory address.
// %b: Formats the integer value as a binary number.
// %d: Formats the integer value as a decimal number.
// %o: Formats the integer value as an octal number.
// %U: Formats the Unicode format of the value.

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
			t.Fatal("expected to get an error.")
		}
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")

	assertDefinition(t, dictionary, "test", "this is just a test")
}

// "short variable declaration," where you can declare multiple variables of the same type in a single line
func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if definition != got {
		t.Errorf("got %q want %q", got, definition)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %v want %s given, %q", got, want, "test")

	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
