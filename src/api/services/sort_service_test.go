package services

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	elements := []int{7, 9, 5, 3, 0, 4, 6, 8, 2}

	fmt.Println(elements)

	Sort(elements)

	fmt.Println(elements)

	if elements[0] != 9 {
		t.Error("first element should be 9")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("last element should be 0")
	}
}
