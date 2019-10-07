package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type World struct {
	Cells []int
	Width int
}

func NeweWorld(width int, density float64) *World {
	cells := make([]int, width)

	for i, _ := range cells {
		if rand.Float64() < density {
			cells[i] = rand.Intn(19) - 9
		}
	}

	return &World{
		Cells: cells,
		Width: width,
	}
}

func (w *World) Reproduce() {
	var sc = bufio.NewScanner(os.Stdin)
	newcells := make([]int, w.Width)
	for i, old := range w.Cells {
		for j := old; j != 0; {
			sc.Scan()
			k := mod(i+j, w.Width)
			fmt.Println("i,old,j,k:", i, old, j, k)
			fmt.Println(w.Cells)
			fmt.Println(newcells)

			if newcells[k] > 0 {
				// mutation
				break
			} else {
				newcells[k] = old
				j = w.Cells[k]
			}
		}
	}
	w.Cells = newcells
}

func mod(d, m int) int {
	return ((d % m) + m) % m
}

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	rand.Seed(time.Now().UnixNano())
	// world := NeweWorld(30, 0.2)
	world := &World{
		Width: 26,
		Cells: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, -3, 1, -3, 0, -3, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	for i := 0; i < 30000; i++ {
		sc.Scan()
		for _, v := range world.Cells {
			fmt.Printf("%3d", v)
		}

		world.Reproduce()
	}
}
