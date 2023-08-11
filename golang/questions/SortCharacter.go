package questions

import (
	"fmt"
	"sort"
	"strings"
)

func SortCharacter(word string) {
	var (
		vowel     string
		consonant string
	)

	word = strings.ToLower(word)
	for i := 0; i < len(word); i++ {
		if strings.ContainsAny(string(word[i]), "aeiou") {
			vowel += string(word[i])
		} else if word[i] >= 'a' && word[i] <= 'z' {
			consonant += string(word[i])
		}
	}

	vowelsSorted := sortString(vowel)
	consonantsSorted := sortString(consonant)

	fmt.Println("Vowel Characters : ", vowelsSorted)
	fmt.Println("Consonant Characters : ", consonantsSorted)
}

func sortString(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
