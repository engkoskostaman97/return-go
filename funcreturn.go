package main

import "fmt"

func getHello(name string) string {

	// return "Hello" + name
	if name == "" {
		return "Hello Bro"
	} else {
		return "Hello" + name
	}
}

// func main() {
// 	result := getHello("Engkos")
// 	fmt.Println(result)
// }

func main() {
	result := getHello("Engkos")
	fmt.Println(result)
	fmt.Println(getHello(""))
}
