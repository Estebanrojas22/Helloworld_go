package main

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

// func convertir_string_boolean() {
// 	var palabra string = "false"
// 	boolean, err := strconv.ParseBool(palabra)
// 	if err != nil {
// 		fmt.Println("Este es el error ", err)
// 	} else {
// 		fmt.Println(boolean, reflect.TypeOf(boolean))
// 	}
// }

// func main() {
// 	convertir_string_boolean()
// }

// func covertir_boolean_string() {
// 	var boolean bool = true
// 	strbool := strconv.FormatBool(boolean)
// 	fmt.Println(strbool, reflect.TypeOf(strbool))
// }

// func main() {
// 	covertir_boolean_string()
// }

// Ejercicio 15-------------------------------------------------------------------

// func variables() {
// 	var (
// 		nombre     string = "Esteban"
// 		edad       int    = 25
// 		pensionado bool   = false
// 	)
// 	fmt.Println("Nombre:", nombre)
// 	fmt.Println("Edad:", edad)
// 	fmt.Println("Pensionado:", pensionado)
// }

// func main() {
// 	variables()
// }

// Ejercicio 16-------------------------------------------------------------------

// func valor_cero() {
// 	var nombre string
// 	var edad int
// 	var peso float64
// 	var estudiante bool

// 	fmt.Println("Nombre: ", nombre)
// 	fmt.Println("Edad: ", edad)
// 	fmt.Println("Peso: ", peso)
// 	fmt.Println("Estudiante: ", estudiante)
// }

// func main() {
// 	valor_cero()
// }

// Ejercicio 17-------------------------------------------------------------------

// func variables_cortas() {
// 	nombre := "Benjamin Button"
// 	edad := 99
// 	peso := 80
// 	estudiante := false
// 	fmt.Println("Nombre: ", nombre)
// 	fmt.Println("Edad: ", edad)
// 	fmt.Println("Peso: ", peso)
// 	fmt.Println("Estudiante: ", estudiante)
// }

//	func main() {
//		variables_cortas()
//	}
//
// Ejercicio 17.1-------------------------------------------------------------------
// var profesion = "Deportista"

// func variables_go() {
// 	sueldo := 1000000
// 	fmt.Println("Profesion: ", profesion)
// 	fmt.Println("Sueldo: ", sueldo)
// }

// func main() {
// 	variables_go()
// }

// Ejercicio 18-------------------------------------------------------------------

// var var1 = "Este es el nivel 1"

// func scope_go() {
// 	var var2 = "Este es el nivel 2"
// 	{
// 		var var3 = "Este es el nivel 3"
// 		fmt.Println(var3)
// 	}
// 	fmt.Println(var1)
// 	fmt.Println(var2)
// }

// func main() {
// 	scope_go()
// }

// Ejercicio 19-------------------------------------------------------------------

// func uso_puntero() {
// 	color := "rojo"
// 	fmt.Println(color, &color)
// }

//	func main() {
//		uso_puntero()
//	}
//
// // Ejercicio 19.1-------------------------------------------------------------------
// func puntero_asterisco() {

// 	vehiculo1 := "rojo"
// 	fmt.Println("El vehiculo1 es", vehiculo1)

// 	vehiculo2 := vehiculo1
// 	fmt.Println("El vehiculo2 es", vehiculo2)

// 	vehiculo3 := &vehiculo1
// 	fmt.Println("El vehiculo3 es", *vehiculo3)

// 	vehiculo1 = "gris"

// 	fmt.Println("El vehiculo1 es", vehiculo1, &vehiculo1)
// 	fmt.Println("El vehiculo2 es", vehiculo2, &vehiculo2)
// 	fmt.Println("El vehiculo3 es", *vehiculo3, vehiculo3)
// }

// func main() {
// 	puntero_asterisco()
// }

// Ejercicio 20-------------------------------------------------------------------

// func equivalenciaEnPies(altura float32) float32 {
// 	altura = altura * 3.28
// 	return altura
// }
// func main() {
// 	var altura float32 = 1.70

// 	fmt.Println("La altura es:", altura, "mts")
// 	fmt.Println("La altura es:", equivalenciaEnPies(altura), " pies")
// 	fmt.Println("La nueva altura es:", altura, "mts")
// }
// Ejercicio 20.1-------------------------------------------------------------------

// func conversion_Pies(altura *float32) float32 {
// 	*altura = *altura - 0.10
// 	return *altura
// }

// var altura float32 = 1.70

// func main() {
// 	fmt.Println("La altura es:", altura, "mts")
// 	fmt.Println("Al envejecer:", conversion_Pies(&altura), "mts")
// 	fmt.Println("Despues de envejecer:", altura, "mts")
// }

// Ejercicio 22----------------------------------------------------------------------

// const Pi = 3.1416

// func area(radio float64) float64 {
// 	return Pi * radio * radio
// }

// func main() {
// 	fmt.Println("El area de un circulo cuyo radio es 3 es: ", area(3))
// }
