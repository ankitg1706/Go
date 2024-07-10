package main

import "fmt"

type student struct {
	Name       string
	Class      string
	RollNumber int
	Deptt      string
}

func main() {
	var Ankit student
	var Vishal student
	var add_of_name *string
	add_of_name = &Ankit.Deptt
	Ankit.Name = "Ankit Kumar Gupta"
	Ankit.Class = "12"
	Vishal.Class = "12"
	Vishal.Name = "Vishal Singh"
	Ankit.Deptt = "Science"
	Vishal.Deptt = "Science"
	fmt.Println("NAmwe ", Ankit, Vishal)
	fmt.Println("Address of name for non pointer = ", add_of_name)
	fmt.Println(`hello worldgo run `)
}
