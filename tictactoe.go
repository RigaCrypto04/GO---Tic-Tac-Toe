package main

import (
	"fmt"
)

func main() {
	r1 := [3]string{"_", "_", "_"}
	r2 := [3]string{"_", "_", "_"}
	r3 := [3]string{"_", "_", "_"}
	var pos int
	var row int
	var sign string
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println()
			fmt.Println("Player X turn")
			sign = "X"
		} else {
			fmt.Println()
			fmt.Println("Player O turn")
			sign = "O"
		}
		fmt.Println(r1[0], "", r1[1], "", r1[2])
		fmt.Println(r2[0], "", r2[1], "", r2[2])
		fmt.Println(r3[0], "", r3[1], "", r3[2])
		fmt.Println("Choose row")
		_, _ = fmt.Scanln(&row)
		fmt.Println("Choose a position")
		_, _ = fmt.Scanln(&pos)
		switch {
		case row == 1:
			switch {
			case pos ==1:
				r1[0] = sign
			case pos == 2:
				r1[1] = sign
			case pos == 3:
				r1[2] = sign
			}
		case row == 2:
			switch {
			case pos == 1:
				r2[0] = sign
			case pos == 2:
				r2[1] = sign
			case pos == 3:
				r3[2] = sign
			}
		case row == 3:
			switch {
			case pos == 1:
				r3[0] = sign
			case pos == 2:
				r3[1] = sign
			case pos == 3:
				r3[2] = sign
			}
		}
		fmt.Println(r1[0], "", r1[1], "", r1[2])
		fmt.Println(r2[0], "", r2[1], "", r2[2])
		fmt.Println(r3[0], "", r3[1], "", r3[2])

	}
}
