package dictionary

import "testing"

func TestDictionaty(t *testing.T) {

	mydictionary := Dictionary{"test": "this is just a test"}

	t.Run("test contains word", func(t *testing.T) {
		got, _ := mydictionary.Search("test")
		want := "this is just a test"

		AssertStrings(t, got, want)
	})
	t.Run("test word doesnt exist", func(t *testing.T) {
		_, err := mydictionary.Search("otherword")

		want := "unable to find word"

		if err == nil {
			t.Fatal("expected an error")
		}

		AssertStrings(t, err.Error(), want)

	})
}

func AssertStrings(t *testing.T, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("expected %s, got %s", want, got)
	}
}
