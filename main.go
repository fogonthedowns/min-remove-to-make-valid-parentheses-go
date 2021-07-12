package main

func minRemoveToMakeValid(s string) string {
	stack := []int{}

	// loop through string length
	for i := 0; i < len(s); i++ {
		// add open parentheses to stack
		if s[i] == '(' {
			stack = append(stack, i)
			// look for closing parentheses
		} else if s[i] == ')' {
			// if the stack is full
			// pop off the last element
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			} else {
				// otherwise delete the ')' character
				// and adjust the index iterator
				s = rmClosing(s, i) // delete ')'
				i--
			}
		}
	}

	// get the last index from the stack
	// pop off the index from the stack
	// delete the character from the string, using the index from get
	for len(stack) != 0 {
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		s = s[:pop] + s[pop+1:]
	}

	// return the string
	return s
}

func rmClosing(s string, i int) string {
	return s[:i] + s[i+1:]
}
