package main

import (

	"fmt"
)

func main() {

var a int = 0

for i:=0 ; i<1000 ; i++{
if i%3 == 0{
 a = a+i;
}else if(i%5==0){
 a=a+i;
}

}
 fmt.Println("The Sum of all the numbers is: ",a)
}