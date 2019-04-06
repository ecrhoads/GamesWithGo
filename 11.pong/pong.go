//In this we create an 800x600 window, keep it open for 2 seconds, setup pixels, and render color within the window.

package main

// Experiment! draw some crazy stuff!

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const winWidth int = 800
const winHeight int = 600

type color struct {
	r, g, b byte
}

type pos struct {
	x, y float32 //smaller, better for game dev
}

//in OOL that have inheritance -- you may have class entity {x,y} then class ball : entity (inherits) {radius}
//get all things from thing you are inheriting from. Go does not have inheritance. Go can achieve similar thing...
//called composition. just using type for ball and paddle.

type ball struct {
	pos
	radius int
	xv     float32
	yv     float32
	color  color
}

func (ball *ball) draw(pixels []byte) {
	for y := -ball.radius; y < ball.radius; y++ {
		for x := -ball.radius; x < ball.radius; x++ {
			if x*x+y*y < ball.radius*ball.radius {
				setPixel(int(ball.x)+x, int(ball.y)+y, ball.color, pixels)
			}
		}
	}
}

type paddle struct {
	pos
	w     int
	h     int
	color color
}

//game loop -- goes forever until game ends:
//get all input (keyboard/mouse etc)
//update all your things (physics, ai)
//draw all the things
//usually grouped into Update function and draw function

func (paddle *paddle) draw(pixels []byte) {
	startX := int(paddle.x) - paddle.w/2 //starting in the middle of the paddle, going half the width to get to the edge
	startY := int(paddle.y) - paddle.h/2 //starting from x going half height to go to the top left corner to begin drawing

	//reason to start with y -- going through memory in order.
	for y := 0; y < paddle.h; y++ {
		for x := 0; x < paddle.w; x++ {
			setPixel(startX+x, startY+y, paddle.color, pixels)
		}
	}
}

func setPixel(x, y int, c color, pixels []byte) {
	index := (y*winWidth + x) * 4

	if index < len(pixels)-4 && index >= 0 {
		pixels[index] = c.r
		pixels[index+1] = c.g
		pixels[index+2] = c.b
	}
}

func main() {

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("Testing SDL2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer renderer.Destroy()

	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ABGR8888, sdl.TEXTUREACCESS_STREAMING, int32(winWidth), int32(winHeight))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer tex.Destroy()

	pixels := make([]byte, winWidth*winHeight*4)

	for y := 0; y < winHeight; y++ {
		for x := 0; x < winWidth; x++ {
			setPixel(x, y, color{byte(x % 255), byte(y % 255), 0}, pixels)
		}
	}

	tex.Update(nil, pixels, winWidth*4)
	renderer.Copy(tex, nil, nil)
	renderer.Present()

	player1 := paddle{pos{100, 100}, 20, 100, color{255, 255, 255}}
	ball := ball{pos{300, 300}, 20, 0, 0, color{255, 255, 255}}
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
	}
	player1.draw(pixels)
	ball.draw(pixels)
	sdl.Delay(16)

}
