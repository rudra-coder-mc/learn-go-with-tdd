package iteration

import (
	"fmt"
	"testing"
)

func TestRepete(t *testing.T) {
	repeted := Repeat("a", 5)
	expected := "aaaaa"

	if repeted != expected {
		t.Errorf("expacted %q but got %q", expected, repeted)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeat := Repeat("ram ", 5)
	fmt.Println(repeat)
}
