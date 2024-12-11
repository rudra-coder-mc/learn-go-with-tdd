package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

// Test for ArraySum function
func TestArrays(t *testing.T) {
	t.Run("run ArraySum", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}

		got := ArraySum(nums)
		want := 15

		if got != want {
			t.Errorf("expected %d but got %d", want, got) // Use %d for integers
		}
	})

	t.Run("run ArraySumAll", func(t *testing.T) {
		numbers1 := []int{1, 2, 3}
		numbers2 := []int{4, 5}

		got := ArraySumAll(numbers1, numbers2)
		want := []int{6, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v but got %v", want, got) // Use %v for slices
		}
	})
}

// Example for ArraySum function
func ExampleArraySum() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(ArraySum(nums))
	// Output: 15
}
