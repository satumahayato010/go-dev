package main

import "fmt"

func main() {
	var emptiy interface{}

	emptiy = "Hello World"
	emptiy = 100
	emptiy = true

	fmt.Println(emptiy)

}
