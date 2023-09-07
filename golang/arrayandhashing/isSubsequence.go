package arrayandhashing

import (
	"strings"
)

func IsSubsequence(s string, t string) bool {
	wordS := strings.Split(s, "")
	wordT := strings.Split(t, "")
	var res bool

	i := 0
	for i < len(wordS) {
		j := i
		for j < len(wordT) {
			if wordS[i] == wordT[j] && len(wordS) == 1 {
				if wordS[i] != wordT[j] {
					wordT = append(wordT[:j], wordT[j+1:]...)
					j -= 1
				}
			}
			if wordS[i] == wordT[j] && len(wordS) > 1 {
				break
			}
			if wordS[i] != wordT[j] {
				wordT = append(wordT[:j], wordT[j+1:]...)
				j -= 1
			}
			j++
		}
		i++
	}
	if strings.Join(wordT, "") == s {
		res = true
	} else if s == "" && t != "" {
		res = true
	}

	return res
}
