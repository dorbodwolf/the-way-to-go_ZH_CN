package main

func main() {
	var a int
	var b int32
	a = 15
	b = a + 5 // compiler error
	b = b + 5 // ok: 5 is a constant
}
