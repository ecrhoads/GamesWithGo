//Denotes this package will be an exe.
package main

//Use a built-in package (libraray).
import "fmt"

//Call a main() function. This is a reserved word, and where the program starts.
func main() {
	/*
		a := 5 //Or var a int = 5
		var a : int8 // -127, 127
		var b : uint8 // 0 255
		var c : int16
		var d : uint16
		var e : int // 32bit on 32bit machines, 64 on 64

		//3.14
		var f1 float32
		var f2 float64

		a := 5
		b := 3.14

		fmt.Println(a + int(b))

		torf := true //bool

		for i := 0; i < 10; i++ {
			fmt.Println("Hello World", i)
		}

		//flow control

	*/

	x := 5
	if x > 5 {
		fmt.Println("Hello, World!")
	} else if x < 5 {
		fmt.Println("Goodbye, World!")
	} else {
		fmt.Println("IT IS FIVE!")
	}
}
