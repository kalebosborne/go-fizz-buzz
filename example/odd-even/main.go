package main

import "fmt"
import "github.com/kalebosborne/go-modulus"


func main() {
	isOdd(5)
	
}


func isOdd(x float64)  {
	var y = mod.Modulus(x,2)
	if y == 0 {
		fmt.Println(x, "is even")
	} else {
		fmt.Println(x, "is odd")
	}
}