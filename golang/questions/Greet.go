package questions

func Greet(name string) string {
	if name == "Johnny" {
		name = "my love"
	}
	return "Hello, " + name + "!"
}
