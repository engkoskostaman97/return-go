package main

import "fmt"

func getHello() (string, string) {
	return "Engkos", "Kostaman"

}

func main() {
	name, lastName := getHello()
	fmt.Println(name, lastName)
}
