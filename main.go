package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Ange vattenpelare i cm samt brunnens diameter.")
		os.Exit(1)
	}
	vp, _ := strconv.ParseFloat(os.Args[1], 64) // Vattebpelare
	bd, _ := strconv.ParseFloat(os.Args[2], 64) // Brunnens diameter
	br := bd / 100 / 2                          // Brunnens radie
	v := (vp / 100) * float64(3.14) * br * br * 1000
	fmt.Println("Vattenmängd i brunnen:", int(v), "Liter")
}
