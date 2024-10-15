package main

// import (
// 	"fmt"
// )

// Ejercicio 23 --------------------------------------------------------------

// const Pi = 3.1416

// func circulo(radio float64) (area float64, perimetro float64) {
// 	area = Pi * radio * radio
// 	perimetro = 2 * Pi * radio
// 	return area, perimetro
// }

// func main() {
// 	a, p := circulo(8)
// 	fmt.Println("El area del circulo es: ", a)
// 	fmt.Println("El perimetro del circulo es: ", p)
// }

// Ejercicio 24--------------------------------------------------------------

// func sumar(numeros ...int) int {
// 	// el total inicial es 0
// 	total := 0
// 	// recorrer todos los numeros
// 	for _, numero := range numeros {
// 		// en cada iteración sumar al total el valor del numero
// 		total = numero + total
// 	}
// 	// retornar el valor total
// 	return total
// }

// func main() {
// 	fmt.Println(sumar(2))
// 	fmt.Println(sumar(2, 2))
// 	fmt.Println(sumar(5, 4, 3))
// }

// Ejercicio 25--------------------------------------------------------------

// func ecoDeLaMontana(mensaje string, iteraciones uint) {
// 	if iteraciones > 1 {
// 		ecoDeLaMontana(mensaje, iteraciones-1)
// 	}
// 	fmt.Println(mensaje, iteraciones)
// }

// func main() {
// 	ecoDeLaMontana("yodelayheehoo", 5)
// }

// Ejercicio 26--------------------------------------------------------------

// func circulo(radio float64) (area func() float64, perimetro func() float64) {

// 	area = func() float64 {
// 		return 3.1416 * radio * radio
// 	}

// 	perimetro = func() float64 {
// 		return 2 * 3.1416 * radio
// 	}

// 	return
// }

// func main() {

// 	area, perimetro := circulo(10)
// 	fmt.Println("El area del circulo es", area())
// 	fmt.Println("El perimetro del circulo es", perimetro())

// }

// Ejercicio 30--------------------------------------------------------------

// func main() {
//     var juguete string
//     fmt.Println("Elige persona, animal o cosa:")
//     fmt.Scanln(&juguete)
//     if juguete == "persona" {
//         fmt.Println("El objeto es una persona")
//     } else if juguete == "cosa" {
//         fmt.Println("El objeto es una cosa")
//     } else if juguete == "animal" {
//         fmt.Println("El objeto es un animal")
//     } else {
//         fmt.Println("El objeto es otra categoria")
//     }
// }

// Ejercicio 31--------------------------------------------------------------

// func division_x_y() {
// 	var x int
// 	fmt.Println("Elige el primer valor: ")
// 	fmt.Scanln(&x)

// 	var y int
// 	fmt.Println("Elige el primer valor: ")
// 	fmt.Scanln(&y)

// 	fmt.Println("La division de", x, "divido", y, "es: ", x/y)
// 	// fmt.Println("El resto de", x, "entre", y, "es: ", x%y)
// 	// if x%2 == 0 {
// 	// 	fmt.Println(x, "es par")
// 	// } else {
// 	// 	fmt.Println(x, "es impar")
// 	// }
// }

// func main() {
// 	division_x_y()
// }

// Ejercicio 33 ---------------------------------------------------------------------------------

// func main() {
// 	x, y, z := -5, -10, 5

// 	if (x > 0  && y > 0) || z < 0 {
// 		fmt.Println("Almenos uno de los operadores se cumplen")
//     } else {
// 		fmt.Println("Todos los operadores no se cumplen")
// }

// }

// Ejercicio 33.1 -------------------------------------------------------------------------------

// func contraseña() {
// 	var contraseña string

// 	// Preguntamos por la contraseña del usuario.
// 	fmt.Print("Por favor, ingrese su contraseña: ")
// 	fmt.Scanln(&contraseña)

// 	// Utilizamos el operador or para verificar si la contraseña cumple con los requisitos.
// 	if !strings.ContainsAny(contraseña, "abcdefghijklmnopqrstuvwxyz") || !strings.ContainsAny(contraseña, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") || !strings.ContainsAny(contraseña, "0123456789") || !strings.ContainsAny(contraseña, "*+_!#$&@()[]-./€") || len(contraseña) <= 8 {
// 		fmt.Println("La contraseña no cumple con los requisitos.")
// 	} else {
// 		fmt.Println("La contraseña es válida.")
// 	}
// }

// func main() {
// 	contraseña()
// }

// Ejercicio 35 -------------------------------------------------------------------------------

// func switche() {
// 	var juguete string
// 	fmt.Println("Elija que tipo de juguete agregar? persona, animal, balon o cosa")
// 	fmt.Scanln(&juguete)
// 	switch juguete {
// 	case "persona":
// 		fmt.Println("El juguete es una figura de accion")
// 	case "cosa":
// 		fmt.Println("El juguete es una cosa")
// 	case "animal":
// 		fmt.Println("El juguete es una mascota")
//     case "balon":
// 		fmt.Println("El juguete es un balon")
// 	default:
// 		fmt.Println("El juguete es otra categoria")
// 	}
// }

// func main() {
// 	switche()
// }

// Ejercicio 35.1 -------------------------------------------------------------------------------

// func switche() {
// 	var operacion string
// 	var x int
// 	var y int
// 	fmt.Println("Elija el tipo de caso? suma, resta, multiplicacion o division")
// 	fmt.Scanln(&operacion)
// 	switch operacion {
// 	case "suma":
// 		fmt.Println("Por favor ingrese el primer numero: ")
// 		fmt.Scanln(&x)
// 		fmt.Println("Por favor ingrese el segundo numero: ")
// 		fmt.Scanln(&y)
// 		fmt.Println("La suma es: ", x+y)
// 	case "resta":
// 		fmt.Println("Por favor ingrese el primer numero: ")
// 		fmt.Scanln(&x)
// 		fmt.Println("Por favor ingrese el segundo numero: ")
// 		fmt.Scanln(&y)
// 		fmt.Println("La resta es: ", x-y)
// 	case "multiplicacion":
// 		fmt.Println("Por favor ingrese el primer numero: ")
// 		fmt.Scanln(&x)
// 		fmt.Println("Por favor ingrese el segundo numero: ")
// 		fmt.Scanln(&y)
// 		fmt.Println("La multiplicacion es: ", x*y)
// 	case "division":
// 		fmt.Println("Por favor ingrese el primer numero: ")
// 		fmt.Scanln(&x)
// 		fmt.Println("Por favor ingrese el segundo numero: ")
// 		fmt.Scanln(&y)
// 		if y != 0 {
// 			fmt.Println("La division es: ", float64(x)/float64(y))
// 		} else {
// 			fmt.Println("No se puede dividir por cero")
// 		}
// 	default:
// 		fmt.Println("El valor de la operacion no es valido")
// 	}
// }

// func main() {
// 	switche()
// }

// Ejercicio 36 -------------------------------------------------------------------------------

// func desconectar() {
//     fmt.Println("Se ha desconectado de la base de datos")
// }

// func leer() {
//     fmt.Println("Se han leido los registros de la base de datos")
// }

// func actualizar() {
//     fmt.Println("Se han actalizado registros de la base de datos")
// }

// func conectar() {
//     fmt.Println("Se ha conectado a la base de datos")
// }

// func main() {
//     conectar()
//     defer desconectar()
//     leer()
//     actualizar()
// }

// Ejercicio 38 -------------------------------------------------------------------------------

// func main() {
//     var marcasDeCoches = make([]string, 2)
//     marcasDeCoches[0] = "Mazda"
//     marcasDeCoches[1] = "Toyota"
//     fmt.Println(marcasDeCoches) // [Mazda Toyota]
//     nuevoSlice := append(marcasDeCoches, "Nissan")
//     fmt.Println(nuevoSlice) // [Mazda Toyota Nissan]
// }

// Ejercicio 39 -------------------------------------------------------------------------------

// func main() {
//     razasDePerros := []string{"labrador", "poodle", "doberman", "shitzu", "beagle"}
//     fmt.Println(razasDePerros)
//     razasDePerros = append(razasDePerros[:2], razasDePerros[2+1:]...)
//     fmt.Println(razasDePerros)
// }

// Ejercicio 41 -------------------------------------------------------------------------------

// func main() {
//     var diasDeLaSemana = make(map[string]int)
//     diasDeLaSemana["lunes"] = 1
//     diasDeLaSemana["martes"] = 2
//     diasDeLaSemana["miercoles"] = 3
//     diasDeLaSemana["jueves"] = 4
//     diasDeLaSemana["viernes"] = 5
//     diasDeLaSemana["sabado"] = 6
//     diasDeLaSemana["domingo"] = 7
//     fmt.Println(diasDeLaSemana)
// }

// Ejercicio 42 -------------------------------------------------------------------------------

// func main() {
//     diasDeLaSemanaEnIngles := map[string]string{
//         "lunes":     "Monday",
//         "martes":    "Tuesday",
//         "miercoles": "Wednesday",
//         "jueves":    "Thursday",
//         "sabado":    "Saturday",
//         "domingo":   "Sunday",
//     }
//     fmt.Println(diasDeLaSemanaEnIngles)
//     delete(diasDeLaSemanaEnIngles, "domingo")
//     fmt.Println(diasDeLaSemanaEnIngles)
// }

// Ejercicio 43 -------------------------------------------------------------------------------



