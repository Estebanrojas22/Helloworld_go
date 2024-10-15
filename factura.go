package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jung-kurt/gofpdf"
	// Se importan varios paquetes necesarios para la funcionalidad del programa
	// bufio: Para leer la entrada del usuario de forma eficiente.
	// fmt: Para la impresión y formateo de cadenas.
	// log: Para manejar errores.
	// os: Para operaciones del sistema, como leer y escribir archivos.
	// strconv: Para convertir cadenas a tipos numéricos.
	// strings: Para manipular cadenas.
	// gofpdf: Para crear documentos PDF.
)

// Se define una estructura llamada Item que almacena la descripción del producto, la cantidad y el precio unitario.
type Item struct {
	Descripcion    string
	Cantidad       int
	PrecioUnitario float64
}

var cliente string

func main() {
	var items []Item
	fmt.Print("Nombre del cliente: ")
	// Se solicita al usuario que ingrese su nombre.

	fmt.Scanln(&cliente)
	// Entrada del usuario para agregar productos
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("¿Desea añadir un producto a la factura? (s/n)")
		confirm, _ := reader.ReadString('\n')
		confirm = strings.TrimSpace(confirm)
		if confirm == "n" {
			break
		}
		// Leer descripción del producto
		fmt.Print("Descripción del producto: ")
		descripcion, _ := reader.ReadString('\n')
		descripcion = strings.TrimSpace(descripcion)
		// Leer cantidad
		fmt.Print("Cantidad: ")
		cantidadStr, _ := reader.ReadString('\n')
		cantidadStr = strings.TrimSpace(cantidadStr)
		cantidad, _ := strconv.Atoi(cantidadStr)

		// Se usa un bucle for para preguntar al usuario si desea añadir un producto a la factura. Si la respuesta es "n", el bucle se detiene.
		// Lectura de detalles del producto:

		// Dentro del bucle, se solicita al usuario que ingrese la descripción del producto, la cantidad y el precio unitario.
		// Cada entrada se limpia de espacios innecesarios y se convierte a los tipos adecuados (usando strconv.Atoi para enteros y strconv.ParseFloat para números decimales).

		// Leer precio unitario
		fmt.Print("Precio unitario: ")
		precioStr, _ := reader.ReadString('\n')
		precioStr = strings.TrimSpace(precioStr)
		precioUnitario, _ := strconv.ParseFloat(precioStr, 64)
		// Se crea un nuevo Item con la información ingresada y se agrega a la lista de ítems.
		// Agregar el producto a la lista de items
		item := Item{
			Descripcion:    descripcion,
			Cantidad:       cantidad,
			PrecioUnitario: precioUnitario,
		}
		items = append(items, item)

	}

	generarFacturaPDF(items)
	// Después de que el usuario decide no agregar más productos, se llama a la función generarFacturaPDF para crear el archivo PDF con la factura.
}

func generarFacturaPDF(items []Item) {
	// Crear un nuevo PDF
	pdf := gofpdf.New("P", "mm", "A4", "") // Inicializa el PDF
	pdf.AddPage()

	// Establecer fuente para el título
	pdf.SetFont("Arial", "B", 20)
	pdf.Cell(0, 10, "Factura de Compra")
	pdf.Ln(10)

	// Establecer fuente para los detalles
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, "Numero de Factura: 12345")
	pdf.Ln(5)

	// Obtener fecha y hora actuales
	now := time.Now()
	fecha := now.Format("02/01/2006")
	hora := now.Format("15:04:05")
	pdf.Cell(0, 10, fmt.Sprintf("Fecha: %s", fecha))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("Hora: %s", hora))
	pdf.Ln(5)
	pdf.Cell(0, 10, fmt.Sprintf("Cliente: %s", cliente))
	pdf.Ln(10)

	pdf.Cell(0, 10, "-----------------------------------------------------")
	pdf.Ln(5)

	// Encabezados de la tabla
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(80, 10, "Descripcion")
	pdf.Cell(30, 10, "Cantidad")
	pdf.Cell(40, 10, "Precio Unitario")
	pdf.Cell(40, 10, "Total")
	pdf.Ln(5)

	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, "-----------------------------------------------------")
	pdf.Ln(5)

	// Agregar los ítems y calcular total
	var totalFinal float64
	for _, item := range items {
		total := float64(item.Cantidad) * item.PrecioUnitario
		totalFinal += total
		pdf.Cell(80, 10, item.Descripcion)
		pdf.Cell(30, 10, strconv.Itoa(item.Cantidad))
		pdf.Cell(40, 10, fmt.Sprintf("%.2f", item.PrecioUnitario))
		pdf.Cell(40, 10, fmt.Sprintf("%.2f", total))
		pdf.Ln(5)
	}

	// Total final de la factura
	pdf.Cell(0, 10, "-----------------------------------------------------")
	pdf.Ln(5)
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(80, 10, "Total Final")
	pdf.Cell(30, 10, "")
	pdf.Cell(40, 10, "")
	pdf.Cell(40, 10, fmt.Sprintf("%.2f", totalFinal))
	pdf.Ln(10)
	pdf.Cell(0, 10, "-----------------------------------------------------")

	// Guardar el PDF
	err := pdf.OutputFileAndClose("factura.pdf")
	if err != nil {
		log.Fatal(err)
	}
}
