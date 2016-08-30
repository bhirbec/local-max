package localmax

import (
	"testing"
)

/* Unit tests */

func TestMax(t *testing.T) {
	input := []int{3, 4, 2}
	expected := []int{4}
	runTest(t, input, expected)

	input = []int{0, 7, 2, 15, 0}
	expected = []int{7, 15}
	runTest(t, input, expected)
}

func TestPlateau(t *testing.T) {
	input := []int{3, 4, 4, 4, 2}
	expected := []int{4}
	runTest(t, input, expected)

	input = []int{3, 3, 3, 4, 4, 4, 2, 2, 2, 2}
	expected = []int{4}
	runTest(t, input, expected)
}

func TestBothPlateauAndMax(t *testing.T) {
	input := []int{0, 7, 7, 7, 2, 1, 4, 0}
	expected := []int{7, 4}
	runTest(t, input, expected)
}

func TestSeveralPlateau(t *testing.T) {
	input := []int{0, 7, 7, 7, 2, 3, 3, 3, 0}
	expected := []int{7, 3}
	runTest(t, input, expected)
}

func TestFlatInput(t *testing.T) {
	input := []int{0, 0, 0, 0}
	expected := []int{}
	runTest(t, input, expected)
}

func TestEmptyInput(t *testing.T) {
	runTest(t, []int{}, []int{})
}

func runTest(t *testing.T, input, expected []int) {
	output := FindLocalMaxV1(input)
	checkResult(t, output, expected)

	output = FindLocalMaxV2(input)
	checkResult(t, output, expected)
}

func checkResult(t *testing.T, output, expected []int) {
	if !compareSlice(output, expected) {
		t.Errorf("Failed: expected %v and got %v", expected, output)
	}
}

func compareSlice(s1, s2 []int) bool {
	n1, n2 := len(s1), len(s2)
	if n1 != n2 {
		return false
	}

	for i := 0; i < n1; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}

	return true
}

/* Benchmarks */

func BenchmarkV1(b *testing.B) {
	input := []int{0, 7, 7, 7, 2, 3, 3, 3, 0, 4, 5, 7, 8, 2, 0, 1, 1, 1, 1, 4, -2, 4, 0}
	for i := 0; i < b.N; i++ {
		FindLocalMaxV1(input)
	}
}

func BenchmarkV2(b *testing.B) {
	input := []int{0, 7, 7, 7, 2, 3, 3, 3, 0, 4, 5, 7, 8, 2, 0, 1, 1, 1, 1, 4, -2, 4, 0}
	for i := 0; i < b.N; i++ {
		FindLocalMaxV2(input)
	}
}
