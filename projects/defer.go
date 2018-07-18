package main

import "fmt"
import "strconv"

func t(i *int) *int {
	*i+=1
	fmt.Println("world : "+ strconv.Itoa(*i))
	return i
}

func main() {
	a := 2
	defer t(&a)
	
	a = 6
	fmt.Println("hello : " + strconv.Itoa(a))
}
