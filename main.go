package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type Transaccion struct {
	Id                int      `form:"Id" json:"Id"`
	CodigoTransaccion string   `form:"CodigoTransaccion" json:"CodigoTransaccion"`
	Moneda            []string `form:"Moneda" json:"Moneda"`
	Monto             float64  `form:"Monto" json:"Monto"`
	Emisor            string   `form:"Emisor" json:"Emisor"`
	Receptor          string   `form:"Receptor" json:"Receptor`
	FechaTransaccion  string   `form:"FechaTransaccion" json:"FechaTransaccion"`
}

func handleTransaccion(ctx *gin.Context) {
	var trans1 []Transaccion
	readFile, _ := os.ReadFile("./transaccion.json")

	if err := json.Unmarshal(readFile, &trans1); err != nil {
		fmt.Sprintf("%v", err)
	}
	ctx.JSON(200, trans1)
}

func getAllWithFilter(ctx *gin.Context) {
	var tr1 Transaccion
	var tr2 Transaccion
	var tr3 Transaccion
	arregloTrans := make([]Transaccion, 0)

	readFile, _ := os.ReadFile("./trans.json")
	readFile2, _ := os.ReadFile("./trans2.json")
	readFile3, _ := os.ReadFile("./trans3.json")

	if err := json.Unmarshal(readFile, &tr1); err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(readFile2, &tr2); err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(readFile3, &tr3); err != nil {
		log.Fatal(err)
	}
	arregloTrans = append(arregloTrans, tr1, tr2, tr3)

	filtroId := ctx.Query("Id")
	filtroEmisor := ctx.Query("Emisor")
	filtroReceptor := ctx.Query("Receptor")
	filtroFecha := ctx.Query("FechaTransaccion")
	fmt.Println("··········· Filtros de búsqueda ············")
	fmt.Printf("Id:%v\nEmisor: %v\nReceptor: %v\nFecha: %v\n", filtroId, filtroEmisor, filtroReceptor, filtroFecha)
	fmt.Println("·············································")

	arrayFiltrado := make([]Transaccion, 0)

	for _, trans := range arregloTrans {
		if filtroId == fmt.Sprint(trans.Id) {
			fmt.Println("Encontrado")
			fmt.Printf("%v\n\n", trans)
			arrayFiltrado = append(arrayFiltrado, trans)
			continue
		}
		if filtroFecha == trans.FechaTransaccion {
			fmt.Println("Encontrado")
			fmt.Printf("%v\n\n", trans)
			arrayFiltrado = append(arrayFiltrado, trans)
			continue
		}
		if filtroReceptor == trans.Receptor {
			fmt.Println("Encontrado")
			fmt.Printf("%v\n\n", trans)
			arrayFiltrado = append(arrayFiltrado, trans)
			continue
		}
		if filtroEmisor == trans.Emisor {
			fmt.Println("Encontrado")
			fmt.Printf("%v\n\n", trans)
			arrayFiltrado = append(arrayFiltrado, trans)
		}
	}
	ctx.JSON(200, arrayFiltrado)
}

func GetOneById(ctx gin.Context) {

}

func main() {
	router := gin.Default()
	// Turno mañana - Ejercicio 1
	router.GET("/index/:nombre", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": c.Param("nombre"),
		})
	})

	// Nos devuelvo las transacciones
	router.GET("/transacciones", handleTransaccion)

	// Turno tarde - Ejercicio 1
	// Le pasamos los filtros Id, Emisor, Receptor o FechaTransaccion
	router.GET("/getAllWithFilter", getAllWithFilter)

	// Turno tarde - Ejercicio 2
	router.GET("/getOne/:id", func(ctx *gin.Context) {
		// Todo este viaje es para leer el Arreglo
		var tr1 Transaccion
		var tr2 Transaccion
		var tr3 Transaccion
		arregloTrans := make([]Transaccion, 0)
		readFile, _ := os.ReadFile("./trans.json")
		readFile2, _ := os.ReadFile("./trans2.json")
		readFile3, _ := os.ReadFile("./trans3.json")
		if err := json.Unmarshal(readFile, &tr1); err != nil {
			log.Fatal(err)
		}
		if err := json.Unmarshal(readFile2, &tr2); err != nil {
			log.Fatal(err)
		}
		if err := json.Unmarshal(readFile3, &tr3); err != nil {
			log.Fatal(err)
		}
		arregloTrans = append(arregloTrans, tr1, tr2, tr3)

		//Acá comienza la lógica
		idAux := ctx.Param("id")
		for _, trans := range arregloTrans {
			if idAux == fmt.Sprint(trans.Id) {
				ctx.JSON(200, trans)
				break
			}
		}
		ctx.JSON(404, fmt.Sprint("Not Found (Acá tengo que devolver un NotFound"))
	})
	router.Run()
}
