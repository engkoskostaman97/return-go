package main

import (
	"fmt"
)
func main(){

	person  := map[string]string{
		"name":"engkos",
		"addres":"tangerang",
	}

	person["title"] = "programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["addres"])

	var book map[string]string = make(map[string]string)
	book["title"]="belajar golang"
	book["author"]="engkos"
	book["ups"]="salah"

	fmt.Println(book)
	fmt.Println(len(book))


	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))



}