package main

import "fmt"

type Position struct {
	x float32
	y float32
}

type badGuy struct {
	name   string
	health int
	pos    Position
}

func whereIsBadGuy(b badGuy) {
	x := b.pos.x
	y := b.pos.y
	fmt.Println("(", x, ",", y, ")")
}

func main() {

	p := Position{4, 2}
	fmt.Println(p.x)

	b := badGuy{"Jabba The Hut", 100, p}
	whereIsBadGuy(b)
}
