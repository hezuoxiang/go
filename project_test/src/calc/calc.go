package main

import "fmt"
import "simplemath"
import "oop"

func main(){
	he := &oop.Person{"hezuoxiang",10}
	he.Say("fuck")
	
	var a oop.Integer = 1
	var b oop.Number = &a
	fmt.Println("Result: ", b.Less(2))
	fmt.Println("Result: ", simplemath.Add(3,4))
}