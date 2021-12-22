package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

type Transaccion struct {
	Id                int
	CodigoTransaccion string
	Moneda            []string
	Monto             float64
	Emisor            string
	Receptor          string
	FechaTransaccion  string
}

func handleTransaccion(ctx *gin.Context) {
	var trans1 []Transaccion
	readFile, _ := os.ReadFile("./transaccion.json")

	if err := json.Unmarshal(readFile, &trans1); err != nil {
		fmt.Sprintf("%v", err)
	}
	ctx.JSON(200, trans1)
}

func main() {
	router := gin.Default()
	router.GET("/index/:nombre", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": c.Param("nombre"),
		})
	})
	router.GET("/transacciones", handleTransaccion)
	router.Run()
}
