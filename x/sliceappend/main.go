package main

func main() {
	var a []string
	a = append(a, "x")
	a = append(a, "y")

	b := []string{"X", "Y"}
	a = append(a, b...)
}
