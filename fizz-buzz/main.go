/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.

*/


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
