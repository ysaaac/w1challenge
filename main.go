package main

import (
	"fmt"
	"os"

	"trucode.com/w1challenge/calculator"
	"trucode.com/w1challenge/evenodd"
	"trucode.com/w1challenge/loops"
	"trucode.com/w1challenge/mario"
	"trucode.com/w1challenge/physics"
)

func main() {
	args := os.Args
	functionName := args[1]

	switch functionName {
	case "evenodd":
		var num int
		fmt.Println("Qué número quieres evaluar?")
		fmt.Scan(&num)
		fmt.Println(evenodd.Evenodd(num))
	case "ohm":
		var v float32
		var r float32
		var i float32
		fmt.Println("Dame los valores para v, r , e i")
		fmt.Scan(&v, &r, &i)
		fmt.Println(physics.Ohm(v, r, i))
	case "foobar":
		var limit int
		fmt.Println("Cuál es el límite de Foobar?")
		fmt.Scan(&limit)
		fmt.Println(loops.Foobar(limit))
	case "bmi":
		var weight float32
		var height float32
		fmt.Println("How much do you weigh? (don't lie)")
		fmt.Scan(&weight)
		fmt.Println("How tall are you? (barefoot)")
		fmt.Scan(&height)
		fmt.Println(calculator.CalculateBMI(weight, height))
	case "mario":
		for {
			fmt.Print("Pyramid height: ")
			var height int
			fmt.Scan(&height)
	
			if height >= 1 && height <= 8 {
				fmt.Println(mario.PrintPyramid(height))
				break
			} 
		}
	}
}
