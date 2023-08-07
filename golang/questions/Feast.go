package questions

func Feast(beast string, dish string) bool {
	// your code here
	return string(beast[0])+string(beast[len(beast)-1]) == string(dish[0])+string(dish[len(dish)-1])
}
