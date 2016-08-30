package localmax

func FindLocalMaxV1(input []int) []int {
	n := len(input)
	output := []int{}

	for i := 1; i < n; i++ {
		current_value := input[i]

		if input[i-1] > current_value {
			continue
		}

		for j := i + 1; j < n; j++ {
			if input[j] == current_value {
				i = j
			} else if current_value > input[j] {
				output = append(output, current_value)
				break
			} else {
				break
			}
		}
	}

	return output
}

func FindLocalMaxV2(input []int) []int {
	n := len(input)
	if n == 0 {
		return []int{}
	}

	output := []int{}
	prev := input[0]

	// we assume there's no max at pos 0 or n-1
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
