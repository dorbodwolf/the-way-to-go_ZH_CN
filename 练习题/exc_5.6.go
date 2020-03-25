// 使用按位补码从 0 到 10，使用位表达式 %b 来格式化输出。
package main
import "fmt"

func main () {
	for i:=0;i<10;i++ {
		fmt.Printf("%b的补码是：%b\n", i, ^i)
	}
}
