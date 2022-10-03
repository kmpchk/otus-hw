package main

import "fmt"

func main() {
	var input string = "cat and dog, one dog,two cats and one man" //"Нога, нога. нога. , нога. НоГа noga Noga NoGa" //"cat and dog, one dog,two cats and one man MaN Man Man"
	//Top10(input)
	var top10words []string = Top10(input)
	fmt.Println(top10words)
}
