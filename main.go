package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Ange vattenpelare i meter samt brunnens diameter.")
		os.Exit(1)
	}
	pi := float64(3.14159265359)
	vp, _ := strconv.ParseFloat(os.Args[1], 64) // Vattenpelare
	bd, _ := strconv.ParseFloat(os.Args[2], 64) // Brunnens diameter
	br := bd / 2                                // Brunnens radie
	v := vp * pi * br * br * 1000
	fmt.Println("Vattenmängd i brunnen:", int(v), "Liter.")
}
