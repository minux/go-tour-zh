package main

import "tour/pic"

func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for i := range p {
		p[i] = make([]uint8, dx)
	}

	for y := range p {
		for x, row := range p[y] {
			row[x] = uint8(x * y)
		}
	}

	return p
}

func main() {
	pic.Show(Pic)
}
