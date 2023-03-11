package main

import "fmt"

func getHello() (string, string) {
	return "Engkos", "Kostaman"

}

func main() {
	name, _ := getHello()
	fmt.Println(name)
}
