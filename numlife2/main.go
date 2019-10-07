package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type World struct {
	ThisGeneration []int
	NextGeneration []int
	Width          int
}

func NeweWorld(width int, density float64) *World {
	this := make([]int, width)
	next := make([]int, width)

	for i, _ := range this {
		if rand.Float64() < density {
			// [-9,9]
			this[i] = rand.Intn(19) - 9
		}
	}

	return &World{
		ThisGeneration: this,
		NextGeneration: next,
		Width:          width,
	}
}

func (w *World) Reproduce() {
	var sc = bufio.NewScanner(os.Stdin)
	for i, n := range w.ThisGeneration {
		for j := n; j != 0; {
			sc.Scan()
			k := mod(i+j, w.Width)
			fmt.Println("i,n,j,k:", i, n, j, k)
			fmt.Println("this:", w.ThisGeneration)
			fmt.Println("next:", w.NextGeneration)

			if w.NextGeneration[k] > 0 {
				// mutation
				fmt.Println("mutation!")
				break
			} else {
				w.NextGeneration[k] = n
				j = w.ThisGeneration[k]
			}
		}
	}
	copy(w.ThisGeneration, w.NextGeneration)
}

func mod(d, m int) int {
	return ((d % m) + m) % m
}

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	rand.Seed(time.Now().UnixNano())
	// world := NeweWorld(30, 0.2)
	world := &World{
		Width:          26,
		ThisGeneration: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 5, -3, 1, -3, 0, -3, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		NextGeneration: make([]int, 26),
		// ThisGeneration: []int{0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 0, 0, 3, 0, 0, 0, -2, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	for i := 0; i < 30000; i++ {
		sc.Scan()
		for _, v := range world.ThisGeneration {
			fmt.Printf("%3d", v)
		}
		world.Reproduce()
	}
}
