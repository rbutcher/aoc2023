/*
Copyright © 2023 Ryan Butcher <ryanbutcher06@gmail.com>
*/

package matrix

type direction struct {
	X int
	Y int
}

var (
	directions = []direction{
		{-1, -1}, {0, -1}, {1, -1},
		{-1, 0}, {1, 0},
		{-1, 1}, {0, 1}, {1, 1},
	}
)
