package helloworld

import "testing"

func TestMain(t *testing.T) {
	t.Run("test with input", func(t *testing.T) {
		got := Hello("atul", "")
		want := "Hello, atul"
		compare(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		compare(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("atul", "French")
		want := "Bonjour, atul"
		compare(t, got, want)
	})

	t.Run("test without input", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		compare(t, got, want)
	})
}

func compare(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got = %q, want = %q", got, want)
	}
}
