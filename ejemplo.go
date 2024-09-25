package main

import "fmt"

var var1 = "Este es el nivel 1"

func scope_go() {
	var var2 = "Este es el nivel 2"
	{
		var var3 = "Este es el nivel 3"
		fmt.Println(var3)
	}
	fmt.Println(var1)
	fmt.Println(var2)
}
func main() {
	scope_go()
}
