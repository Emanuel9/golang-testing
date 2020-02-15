package sort

import (
	"testing"
)

func TestBubbleSortOrderDesc(t *testing.T) {
	// Init
	elements := []int{9, 7, 5, 3, 1, 2, 4, 6, 8, 0}

	// Execution
	BubbleSort(elements)

	//Validation
	if elements[0] != 9 {
		t.Error("first element should be 9")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("last element should be 0")
	}

}

func TestBubbleSortAlreadySorted(t *testing.T) {
	//Init
	elements := []int{5, 3, 2, 1}

	//Execution
	BubbleSort(elements)
}

func getElements(n int) []int {
	result := make([]int, n)
	j := 0
	for i := n - 1; i > 0; i-- {
		result[j] = i
		j++
	}

	return result
}

func BenchmarkBubbleSort(b *testing.B) {
	elements := getElements(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	elements := getElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
