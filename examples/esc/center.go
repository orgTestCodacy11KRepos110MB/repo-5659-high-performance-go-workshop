package main

import "fmt"

// START OMIT

type Point struct {
	X, Y int
}

const (
	Width  = 640
	Height = 480
)

func Center(p *Point) {
	p.X += Width / 2
	p.Y += Height / 2
}

func NewPoint() {
	p := new(Point)
	Center(p)
	fmt.Println(p.X, p.Y)
}

// END OMIT