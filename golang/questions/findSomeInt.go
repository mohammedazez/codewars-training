package questions

import "fmt"

func FindSomeInt(number []int) []int {
	temp := []int{}
	var result []int
	for i := 0; i < len(number); i++ {
		count := 0
		for j := 0; j < len(number); j++ {
			if number[i] == number[j] {
				count++
			}
			if count >= 3 {
				temp = append(temp, number[i])
			}
		}
	}

	unique := make(map[int]bool)
	for _, num := range temp {
		fmt.Println(unique[num])
		if !unique[num] {
			unique[num] = true
			result = append(result, num)
		}
	}

	return result
}
