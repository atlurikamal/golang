package main


import "fmt"



const (

	A = iota
	B = iota*2
        C = iota*2
)


 func main() {

	
 fmt.Println(A)
 fmt.Println(B)
 fmt.Println(C)
}
