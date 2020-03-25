package main

import "fmt"

func main () {
	for i:=0;i<25;i++ {
		for j:=0;j<i+1;j++{
			fmt.Printf("G")
		}
		fmt.Printf("\n")
	}
}
