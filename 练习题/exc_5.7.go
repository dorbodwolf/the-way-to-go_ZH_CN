// fizz-buzz
package main
import "fmt"

func main () {
	for i:=0;i<100;i++ {
		switch {
			case (i%3==0) && (i%5!=0):
				fmt.Printf("fizz\n")
			case (i%3!=0) && (i%5==0):
				fmt.Printf("buzz\n")
			case (i%3==0) && (i%5==0):
				fmt.Printf("fizz-buzz\n")
			default:
				fmt.Printf("%d\n", i)
		}
	}
}
