# Go-Modulus

A fizz buzz is a great way to showcase that I understand the basic functionality of a language using functions and control flow. I have understood basic arithmatic and how they work I have wondered for a while how modulus works in a language setting. Upon searching for a modulus function source code I could not find any refferences to code that would satisfy my curiosity. After 45 minutes staring at my screen I finally was able to implement a 64bit modulus function. Instead of just putting it out there I decided to release a fizz buzz program utilizing the modulus function.

TODO
- [x] Upload v 1.0
- [ ] Edit Readme
- [ ] Provide More Examples



## How To Use
### Import Modulus Function

go get github.com/kalebosborne/go-modulus

    import "github.com/kalebosborne/go-modulus"

### Use Modulus Function

    // num = the number you want to divide
    // div = the number you want to divide by
    x := mod.Modulus(num,div)
    fmt.Println(x) //
    
    

### Try It Out In FizzBuzz
cd $GOPATH/src/github.com/kalebosborne/go-modulus/fizzbuzz
#### Option 1
go run main.go

#### Option 2

go build -o fizzbuzz

fizzbuzz.exe (windows) ./fizz-buzz (linux)

or execute fizzbuzz.exe||fizzbuzz in file explorer


### Edit FizzBuzz

    //fizz(a,b) a = starting number b = ending number
    //eg...
    fizz(0,10)
    
 output:
 **1
2
fizz
4
buzz
fizz
7
8
fizz
buzz**
