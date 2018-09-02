package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCase(t *testing.T) {
	input := []int{2, 7, 11, 15}
	target := 9

	expected := []int{0, 1}

	output := twoSum(input, target)

	if !reflect.DeepEqual(expected, output) {
		t.Errorf("Expected: %v, Output: %v\n", expected, output)
	}

	fmt.Printf("Expected: %v, Output: %v\n", expected, output)
}
