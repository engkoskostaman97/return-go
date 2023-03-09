package main

import "fmt"

func main() {

	name := "Engkos"

	switch name {
	case "Engkos":
		fmt.Println("Hello Engkos")
	case "Joko":
		fmt.Println("Hello Joko")
	default:
		fmt.Println("Hai")

	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("NAMA")
	}

}
