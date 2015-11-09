package main


import "fmt"


func main() {

	
var a int
var b *int
	a = 22
        fmt.Println("a value is = ",a)
	fmt.Println("a's memory address is =",&a)
	b = &a
	fmt.Println("Value of B: ",b)
	*b = 30
	fmt.Println("The value of a is: ",a)
}
