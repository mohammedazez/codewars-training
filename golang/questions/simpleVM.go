package questions

import (
	"strconv"
	"strings"
)

// ["PUSH 10", "PUSH 20"] -> [20, 10]
// "PUSH 10", "PUSH 20"
// PUSH
// [20 10]

// ["PUSH 10", "PUSH 20", "ADD"] -> [30]
// "PUSH 10", "PUSH 20"
// PUSH ADD
// 20 + 10
// [30]

// ["PUSH 10", "PUSH 20", "SWAP"] -> [10, 20]
// "PUSH 10", "PUSH 20"
// PUSH SWAP
// 20 10
// [10 20]

// ["PUSH 10", "PUSH 20", "SWAP", "ADD"] -> [30]
// "PUSH 10", "PUSH 20"
// PUSH SWAP ADD
// 20 10
// 10 20
// 10 + 20
// [30]

// ["PUSH 10", "PUSH 20", "MUL", "PUSH 7", "ADD"]
// PUSH SWAP ADD
// (20*10 + 7)
// [207]

// ["PUSH 20", "SWAP", "ADD"] ADD, MUL, and SWAP are element less than 2 -> []

func SimpleVM(a []string) []int {
	stack := []int{}

	for _, v := range a {
		parts := strings.Split(v, " ")
		if len(parts) > 2 {
			continue
		}
		command := parts[0]

		switch command {
		case "PUSH":
			num, err := strconv.Atoi(parts[1])
			if err != nil {
				return []int{}
			}
			stack = append(stack, num)
		case "MUL":
			if len(stack) < 2 {
				return []int{}
			}
			result := stack[len(stack)-2] * stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		case "ADD":
			if len(stack) < 2 {
				return []int{}
			}
			result := stack[len(stack)-2] + stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			stack = append(stack, result)
		case "SWAP":
			if len(stack) < 2 {
				return []int{}
			}
			stack[len(stack)-1], stack[len(stack)-2] = stack[len(stack)-2], stack[len(stack)-1]
		default:
			return []int{}
		}
	}

	if len(stack) == 0 {
		return []int{}
	}

	return stack
}
