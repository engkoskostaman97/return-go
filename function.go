package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World")
}

// func main() {
// 	sayHello()
// 	sayHello()
// 	sayHello()
// }

func main() {
	for i := 0; i < 10; i++ {
		sayHello()
	}
}
