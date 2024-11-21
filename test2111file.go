package main

import "fmt"

func main() {

	s := 5
	n := 10
	nearwhite1 := s
	nearwhite2 := s
	red := 1

	for day := 1; day < n+1; day = day + 1 {
		if nearwhite1-1 > 0 {
			nearwhite1 = nearwhite1 - 1
			red = red + 1
		}
		if nearwhite2+1 < n+1 {
			nearwhite2 = nearwhite2 + 1
			red = red + 1
		}

		fmt.Println("Red tomato count: ", red, "Day count: ", day)

		if red == n {
			break
		}
	}
}