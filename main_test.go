package localmax

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindLocalMax(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []int
		expected []int
	}{
		{"empty", []int{}, []int{}},
		{"simple peak", []int{0, 1, 0}, []int{1}},
		{"plateau peak", []int{0, 1, 1, 1, 0}, []int{1}},
		{"staircase down", []int{5, 5, 5, 4, 4, 4, 3, 3, 3}, []int{}},
		{"staircase spikes", []int{5, 7, 5, 4, 4, 4, 3, 9, 3}, []int{7, 9}},
		{"up-n-down", []int{0, 2, 4, 5, 5, 5, 3, 3, 3, 1, 1, 1, 0, 0, 0}, []int{5}},
	}

	for _, testCase := range testCases {
		actualOutput := FindLocalMax(testCase.input)
		require.Equal(t, actualOutput, testCase.expected, "V1 '"+testCase.desc+"' test failed.")
	}
}
