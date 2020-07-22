package main

// import fmt to display fizzbuzz to console
import "fmt"

//import modulus function
import "github.com/kalebosborne/go-modulus"


// if modulus is dividable to 3 print fizz
// if modulus is dividable by 5 print buzz
// else print number


func main() {
	fizz(0,1000) //specify start and end
}


func fizz(x,y float64){
	for x := x; x <= y; x++ {
		if mod.Modulus(x,3) == 0 {
			fmt.Println("fizz")
		} else if mod.Modulus(x,5) == 0 {
			fmt.Println("buzz")
		} else {fmt.Println(x)}
	}
}