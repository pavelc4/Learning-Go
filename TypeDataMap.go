package main

import "fmt"

func main() {
	var person map[string]string = map[string]string{}
	person["name"] = "Dimas"
	person["age"] = "19"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["age"])

	game := map[string]string{
		"name":  "Delta force",
		"price": "free",
	}

	fmt.Println(game)
	fmt.Println(game["name"])
	fmt.Println(game["price"])

}
