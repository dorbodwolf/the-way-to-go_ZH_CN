package main

func main() {
	a, b := 10, 0
	b += a
	c := a / b // panic: runtime error: integer divide by zero

	print(c)
}
