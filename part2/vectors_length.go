package main

import (
	"fmt"
	"math"
)

type Vector struct {
	x, y, z float64
}

func createVector(x float64, y float64, z float64) Vector {
	return Vector{x, y, z}
}

func (vector *Vector) length() int {
	return int(math.Sqrt(vector.x*vector.x + vector.y*vector.y + vector.z*vector.z))
}

func main() {
	a := createVector(6, 3, 2)
	b := createVector(1, 2, 4)
	fmt.Print(a.length(), " ", b.length())
}
