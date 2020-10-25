package main

import (
	"time"
	"runtime"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
    width  = 500
	height = 500
	
	rows = 100
	columns = 100
	threshold = 0.15
	fps = 10
)

var (
	// triangle = [] float32{
	// 	0, 0.5, 0, // top
	// 	-0.5, -0.5, 0, // left
	// 	0.5, -0.5, 0, // right
	// }
	triangle = []float32{
		-0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
	}

	square = []float32{
		-0.5, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
	
		-0.5, 0.5, 0,
		0.5, 0.5, 0,
		0.5, -0.5, 0,
	}
)

func main() {
    runtime.LockOSThread()

    window := initGlfw()
    defer glfw.Terminate()

	program := initOpenGL()
	
	// vao := makeVao(square)
	cells := makeCells()

    // for !window.ShouldClose() {
	// 	// draw(vao, window, program)
	// 	for x := range cells {
	// 		for _, c := range cells[x] {
	// 			c.checkState(cells)
	// 		}
	// 	}
	// 	draw(cells, window, program)
		
	// }
	
	for !window.ShouldClose() {
		t := time.Now()

		for x := range cells {
			for _, c := range cells[x] {
				c.checkState(cells)
			}
		}

		draw(cells, window, program)

		time.Sleep(time.Second/time.Duration(fps) - time.Since(t))
	}
}

