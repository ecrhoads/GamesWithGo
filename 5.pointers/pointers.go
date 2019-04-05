//pointers are there no matter what, in any language... some hide them.
//Go lets you use pointers, syntax similar to C

package main

import "fmt"

func addOne(x *int) {
	*x = *x + 1
}

type position struct {
	x float32
	y float32
}

type badGuy struct {
	name   string
	health int
	pos    position
}

//Here I am saying set the input var type as a pointer of badGuy.
//Go will de-reference the types automatically, meaning it will get/use
//the values of where Badguy is stored in memory.
func whereIsBadGuy(b *badGuy) {
	x := b.pos.x
	y := b.pos.y
	fmt.Println("(", x, ",", y, ")")
}

func main() {

	x := 5
	fmt.Println(x)

	//Pointer is the actual address in the computer(virtual memory where # lives)
	//imagine computer has numbers living in memory address spaces:
	//1 -- 5
	//2 -- 4
	//3 -- 2
	//4 -- 1

	xPtr := &x
	fmt.Println(xPtr)

	//why use?
	addOne(xPtr)   //remember this is making a copy to function.
	fmt.Println(x) //this now prints the modified value, not the original value of x, thanks to using the pointer.

	p := position{4, 2}
	b := badGuy{"Jabba The Hut", 100, p}
	whereIsBadGuy(&b) //I am saying, pass in the actual b variable in the memory space to the function
}