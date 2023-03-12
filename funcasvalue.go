package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bay" + name
}

func main() {
	sayGoodByes := getGoodBye
	fmt.Println(sayGoodByes("Engkos"))
}
