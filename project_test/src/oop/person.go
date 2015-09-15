package oop
import "fmt"

type Person struct{
	Name string
	Age int
}

func (p Person) Say (content string){
	fmt.Printf("%s said:%s\n",p.Name,content);
}


type Integer int

func (a Integer) Less(b Integer) bool{
	return a<b
}

func (a *Integer) Add(b Integer){
	*a += b
}

type Number interface{
	Less(b Integer) bool
	Add(b Integer)
}