package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	twoD := make([][]uint8, dy)
	for i := 0; i < dy; i++ {

		twoD[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			twoD[i][j] = uint8((2 * i * j) / 1)
		}
	}
	return twoD
}

func main() {
	pic.Show(Pic)
}
