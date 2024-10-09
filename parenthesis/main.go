package parenthesis

func GenParentheses(n int) []string {
	return solution1(n)
	// result := []string{}
	// result = exec2(n, "", result, 0, 0)
	// return result
}

func solution1(n int) []string {
	if n == 1 {
		return []string{"()"}
	}

	before := solution1(n - 1)
	after := []string{}
	for _, val := range before {
		after = append(after, "("+val+")")

		if ("()" + val) == (val + "()") {
			after = append(after, "()"+val)
		} else {
			after = append(after, "()"+val)
			after = append(after, val+"()")
		}
	}

	return after
}

func solution2(n int, concencate string, result []string, opened, closed int) []string {
	if opened == closed && opened == n {
		return append(result, concencate)
	}

	switch {
	case opened == closed:
		return solution2(n, concencate+"(", result, opened+1, closed)
	case opened == n:
		return solution2(n, concencate+")", result, opened, closed+1)
	case opened > closed:
		result = solution2(n, concencate+")", result, opened, closed+1)
		result = solution2(n, concencate+"(", result, opened+1, closed)
	}

	return result
}
