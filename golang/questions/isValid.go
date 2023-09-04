package questions

import "fmt"

func IsValid(s string) bool {
	// Membuat peta (map) yang memetakan tanda kurung tutup ke tanda kurung buka yang sesuai. Ini disimpan dalam variabel closeToOpen.
	closeToOpen := map[byte]byte{')': '(', ']': '[', '}': '{'}

	// Membuat tumpukan kosong (stack) yang digunakan untuk melacak tanda kurung buka yang telah ditemui.
	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		// Iterasi melalui setiap karakter dalam string s.
		c := s[i]

		// Jika karakter saat ini adalah tanda kurung tutup (seperti ')', ']', atau '}'), maka kode memeriksa tumpukan.
		if _, ok := closeToOpen[c]; ok {

			// Jika tumpukan tidak kosong dan tanda kurung buka yang sesuai dengan tanda kurung tutup saat ini berada di bagian paling atas tumpukan, maka tanda kurung buka tersebut dihapus dari tumpukan.
			if len(stack) > 0 && stack[len(stack)-1] == closeToOpen[c] {
				stack = stack[:len(stack)-1]

				// Jika tidak sesuai, maka fungsi mengembalikan false, karena ini menunjukkan bahwa ekspresi tidak valid.
			} else {
				return false
			}

			// Jika karakter saat ini adalah tanda kurung buka (seperti '(', '[', atau '{'), maka karakter tersebut dimasukkan ke dalam tumpukan.
		} else {
			stack = append(stack, c)
		}
	}

	fmt.Println(len(stack))
	return len(stack) == 0
}
