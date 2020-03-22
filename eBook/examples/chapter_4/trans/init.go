package trans

import "math"
import "fmt"

var Pi float64

func init() {
	Pi = 4 * math.Atan(1) // init() function computes Pi
	fmt.Printf("pi值：", Pi)
}
