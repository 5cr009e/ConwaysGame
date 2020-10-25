package main

import (
	"math/rand"
	"time"
	"github.com/go-gl/gl/v4.1-core/gl" // OR: github.com/go-gl/gl/v2.1/gl
)

type cell struct {
	drawable uint32

	alive     bool
	aliveNext bool
	
	x int
	y int

}

func (c *cell) draw() {
	if !c.alive{
		return
	}

	gl.BindVertexArray(c.drawable)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(len(square) / 3))
}



func makeCells() [][]*cell {
	rand.Seed(time.Now().UnixNano())

    cells := make([][]*cell, rows, rows)
    for x := 0; x < rows; x++ {
        for y := 0; y < columns; y++ {
            c := newCell(x, y)
			
			c.alive = rand.Float64() < threshold
			c.aliveNext = c.alive
			
			cells[x] = append(cells[x], c)
			
        }
    }
    
    return cells
}

func newCell(x, y int) *cell {
    points := make([]float32, len(square), len(square))
    copy(points, square)

    for i := 0; i < len(points); i++ {
        var position float32
        var size float32
        switch i % 3 {
        	case 0:
                size = 1.0 / float32(columns)
                position = float32(x) * size
        	case 1:
                size = 1.0 / float32(rows)
                position = float32(y) * size
        	default:
                continue
        }

        if points[i] < 0 {
			points[i] = (position * 2) - 1
        } else {
			points[i] = ((position + size) * 2) - 1
        }
    }

    return &cell{
        drawable: makeVao(points),

        x: x,
        y: y,
    }
}
