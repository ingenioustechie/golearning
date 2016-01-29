/*Write a program that uses the continue keyword in a for loop to 
* print only even numbers up to 100.
*/

package main

import "fmt"

func main() {
	for i:=0; i<=100; i++ {
		if i%2 != 0 {
			continue;
		}
		fmt.Println(i)			
	}
}