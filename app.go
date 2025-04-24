package main

import (
	"fmt"

	"example.hu/app/bank"
	"example.hu/app/deferusage"
	"example.hu/app/investment"
	"example.hu/app/vertex"
)

func main() {
	fmt.Println("Welcome to the practicing app!")
	var mainChoice int

	for {
		fmt.Println("What would you like to do?")
		printMainOptions()
		fmt.Scan(&mainChoice)

		switch mainChoice {
		case 1:
			bank.Banking()
		case 2:
			investment.Investment()
		case 3:
			deferusage.DeferPrinting()
		case 4:
			tryVertex()
		case 5:
			outputNullValues()
		case 6:
			fmt.Println("Exiting the app.")
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		if mainChoice == 6 {
			break
		}
	}
	fmt.Println("Good bye!")
}

func printMainOptions() {
	fmt.Println("----------------------------")
	fmt.Println("1. Banking app")
	fmt.Println("2. Investment app")
	fmt.Println("3. Defer usage")
	fmt.Println("4. Vertex app")
	fmt.Println("5. Output null values")
	fmt.Println("6. Exit")
	fmt.Println("----------------------------")
}

func outputNullValues() {
	var a int
	var b float64
	var c string
	var d bool
	var e complex64

	fmt.Print("Int: ", a, "\n")
	fmt.Print("Float: ", b, "\n")
	fmt.Print("String: ", c, "\n")
	fmt.Print("Boolean: ", d, "\n")
	fmt.Print("Complex: ", e, "\n")
}

func tryVertex() {
	v1 := vertex.NewVertex(1, 2, 3)
	v2 := vertex.NewVertex(4, 5, 6)
	v3 := vertex.NewVertex(7, 8, 9)

	fmt.Println("v1:", v1)
	fmt.Println("v2:", v2)
	fmt.Println("v3:", v3)

	v4 := v1.Add(v2).Sub(v3)
	fmt.Println("v4:", v4)

	v5 := vertex.NewScaledVertex(1, 2, 3, 2)
	fmt.Println("v5:", v5)

	v6 := vertex.CopyVertex(v1)
	fmt.Println("v6:", v6)

	v7 := vertex.NewVertex(10, 11, 12)
	distance := v1.Distance(v7)
	fmt.Printf("Distance between v1 and v7: %.2f\n", distance)

	v8 := vertex.NewVertex(13, 14, 15)
	absValue := v8.Abs()
	fmt.Printf("Abs value of v8: %.2f\n", absValue)

	v9 := vertex.NewVertex(16, 17, 18)
	v9.Scale(2)
	fmt.Println("v9 after scaling:", v9)
}
