package main 

import "fmt"

func main() {
	const tA = "Hey I'm untyped"
	const tB string = "Hey I'm typed"

	fmt.Println("const tA = 'Hey I'm untyped' :", tA)
	fmt.Println("const tB string = 'Hey I'm typed' :", tB)

	const tC = 2
	const tD = 3

	const tE int = tC*tD
	
	fmt.Println("tD*tD = ", tE)
}