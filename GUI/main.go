package main

import (
	"strconv"
	"fmt"
	g "github.com/AllenDang/giu"
)

var (
	sumando1 string
	sumando2 string
	suma string
)

func formulario() {
	g.SingleWindow().Layout(
		g.PrepareMsgbox(),
		g.Row(
			g.Label("Ingrese un numero: "),
			g.InputText(&sumando1),
		),
		g.Row(
			g.Label("Ingrese un numero: "),
			g.InputText(&sumando2),
		),
		g.Button(" Realizar la suma de dos numeros y mostrarlo ").OnClick(sumar),
		g.Row(
			g.Label("El resultado es: "),
			g.InputText(&suma),
		),	
	)
}

func sumar() {
	n1, err := strconv.ParseFloat(sumando1, 64)
	if err != nil {
		g.Msgbox("Error", "Ingrese un numero")
		return
	}
	n2, err := strconv.ParseFloat(sumando2, 64)
	if err != nil {
		g.Msgbox("Error", "Ingrese un numero")
		return
	}
	total := n1 + n2
	suma = fmt.Sprint(total)
}

func main() {
	ventana := g.NewMasterWindow("suma de dos numeros", 335, 125, 0)
	ventana.Run(formulario)
}
