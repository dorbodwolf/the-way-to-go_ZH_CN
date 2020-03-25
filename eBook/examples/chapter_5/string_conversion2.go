package main

import (
	"fmt"
	"strconv"
)
func atoi (s string) (n int) {
	n, _ = strconv.Atoi(s)
	return
}

func main() {
	var orig string = "999"
	// var an int
	var newS string
	// var err error

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	// anInt, err = strconv.Atoi(origStr)
	an := atoi(orig)
	if err:=an; err == 0 {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		return
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
