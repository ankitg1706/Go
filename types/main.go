package main

import(
	"fmt"
	"strconv"
)

func main() {

	fmt.Printf("Data types and type conversion")

	var a int
	a = 12345
	b := 5
	var c int16
	c = int16(a)
	d := strconv.Itoa(a)
	x , e := strconv.Atoi(d)
	fmt.Println("Type casting of integers", a, b, c)
	fmt.Println("Type casting of integers= ", d, x, e)
}
