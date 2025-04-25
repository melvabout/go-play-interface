package main

import (
	"fmt"
)

func main() {

	var p Person = Human{Name: "Dave"}
	fmt.Println(p.Jump())
	fmt.Println(p.Walk())
	fmt.Println(p.Sat())

}
