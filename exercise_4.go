/*
*Declare a named type called counter with a base type of int. 
Declare and initialize a variable of this named 
type to its zero value. Display the value of this variable and the variables type.
*/

package main
import (
	"fmt"
)
type Counter int;
func main() {

	var i Counter;

	fmt.Printf("i = %d\n",i);
	fmt.Printf("Type of i is  %T\n",i)
}

