/*
Write a program that uses the Scanf function from the fmt package 
to read user input. This program should take in a number entered 
by the user and double it. Display the result
*/

package main

import "fmt"

func main() {
	var num int;

	input:
		fmt.Println("Enter a number")
		a, err := fmt.Scanf("%d", &num);

	if err != nil{
		fmt.Println(a, err)
		goto input
	}
	fmt.Println(num*2)
	
}