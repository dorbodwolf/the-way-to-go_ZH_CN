package main

import "fmt"

func main () {
	s := "G"
	for i:=0;i<25;i++ {
		fmt.Printf(s+"\n")
		s= s + "G"
	}
}
