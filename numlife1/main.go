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
	newcells := make([]int, w.Width)
	cc := make([]int, w.Width)
	for i, old := range w.Cells {
		if old != 0 {
			// i = mod(i+old, w.Width)
			i = i + old
			if 0 <= i && i < w.Width {
				cc[i] += 1
				newcells[i] += old
			}
		}
	}
	for i, old := range w.Cells {
		if cc[i] > 0 {
			w.Cells[i] = newcells[i] - old*(cc[i]-1)
		}
	}
}

func mod(d, m int) int {
	return ((d % m) + m) % m
}

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	rand.Seed(time.Now().UnixNano())
	world := NeweWorld(30, 0.2)
	for i := 0; i < 30000; i++ {
		sc.Scan()
		for _, v := range world.Cells {
			fmt.Printf("%3d", v)
		}

		world.Reproduce()
	}
}
