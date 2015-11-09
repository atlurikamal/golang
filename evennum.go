package main

import (

	"fmt"
)

func main() {

for a:=0; a <=100 ; a++ {
 if a%3 == 0 {
 fmt.Print("Fizz ")
}else if a%5 ==0 {
 fmt.Print("Buzz ")
} else if(a%3==0 && a%5==0){
 fmt.Print("FizzBizz ")
}else {
  fmt.Print("  ",a,"  ")
}
}
}