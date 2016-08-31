package localmax

func FindLocalMax(input []int) []int {
	n := len(input)
	if n == 0 {
		return []int{}
	}

	output := []int{}
	prev := input[0]

	// we dont' consider the first and last positions
	for i := 1; i < n-1; i++ {
		current := input[i]
		next := input[i+1]

		// skip plateau
		if current == next {
			continue
		}

		// check local max
		if prev < current && current > next {
			output = append(output, current)
			prev = next
			i += 1
		}
	}

	return output
}
