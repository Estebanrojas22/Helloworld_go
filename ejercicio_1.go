package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// import "fmt"

// func imprimir() {
// 	fmt.Println("Texto desde la funcion imprimir")

// }
// // func main() {
// 	fmt.Println("Texto desde la funcion main")
// 	imprimir()
// }

// func Hola_string(s string) string {
// 	return "Hola" + s
// }

// func main() {
// 	fmt.Println(Hola_string(" Esteban"))
// }

// func sumar(a int, b int) int {
// 	return a + b
// }

// func main() {
// 	fmt.Println(sumar(3, 5))
// }

// func comparar_bool() {
// 	var a bool
// 	var b bool
// 	a = true
// 	b = false

// 	if a == b {
// 		fmt.Println("A y B son iguales")
// 	} else {
// 		fmt.Println("A y B no son iguales")
// 	}
// }

// func main() {
// 	comparar_bool()

// }

// func arreglo() {
// 	var aprendices [5]string
// 	aprendices[0] = "Leonadardo"
// 	aprendices[1] = "Juan Sebastian"
// 	aprendices[2] = "Frank"
// 	aprendices[3] = "Juan Jose"
// 	aprendices[4] = "Jaider"

// 	fmt.Println(aprendices[1])
// }

// func main() {
// 	arreglo()
// }

// func tipos_datos() {
// 	var texto string = "Esteban"
// 	var numero int = 1000
// 	var decimal float64 = 0.0001

// 	fmt.Println(reflect.TypeOf(texto))
// 	fmt.Println(reflect.TypeOf(numero))
// 	fmt.Println(reflect.TypeOf(decimal))
// }

// func main() {
// 	// tipos_datos()
// }

func convertir_string_boolean() {
	var palabra string = "false"
	boolean, err := strconv.ParseBool(palabra)
	if err != nil {
		fmt.Println("Este es el error ", err)
	} else {
		fmt.Println(boolean, reflect.TypeOf(boolean))
	}
}

func main() {
	convertir_string_boolean()
}

// func covertir_boolean_string() {
// 	var boolean bool = true
// 	strbool := strconv.FormatBool(boolean)
// 	fmt.Println(strbool, reflect.TypeOf(strbool))
// }

// func main() {
// 	covertir_boolean_string()
// }
