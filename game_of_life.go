package main

import (

)

func (c *cell) checkState(cells [][]*cell) {
	c.alive =  c.aliveNext
	c.aliveNext = c.alive
	liveCount := c.liveNeighbors(cells)

	if c.alive {
		// 1. Any live cell with fewer than two live neighbors dies
		if liveCount < 2 {
			c.aliveNext = false
		}

		// 2. Any live cell with two or three live neighbors live
		if liveCount == 2 || liveCount == 3 {
			c.aliveNext = true
		}

		// 3. Any live cell with more than three neighbors dies
		if liveCount > 3 {
			c.aliveNext = false
		}
	} else {
		// 4. Any dead cell with exactly three live neighbors become live
		if liveCount == 3 {
			c.aliveNext = true
		}
	}
		
	
}

// liveNeighbors returns the number of live neighbors for a cell.
func (c *cell) liveNeighbors(cells [][]*cell) int {
    var liveCount int
    add := func(x, y int) {
        // If we're at an edge, check the other side of the board.
        if x == len(cells) {
            x = 0
        } else if x == -1 {
            x = len(cells) - 1
        }
        if y == len(cells[x]) {
            y = 0
        } else if y == -1 {
            y = len(cells[x]) - 1
        }
        
        if cells[x][y].alive {
            liveCount++
        }
    }
    
    add(c.x-1, c.y)   // To the left
    add(c.x+1, c.y)   // To the right
    add(c.x, c.y+1)   // up
    add(c.x, c.y-1)   // down
    add(c.x-1, c.y+1) // top-left
    add(c.x+1, c.y+1) // top-right
    add(c.x-1, c.y-1) // bottom-left
    add(c.x+1, c.y-1) // bottom-right
    
    return liveCount
}