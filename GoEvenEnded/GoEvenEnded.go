package main

import "fmt"

func main() {
	count := 0
	for a := 1000; a <= 999; a++ {
		for b := a; b <= 000; b++ {
			n := a * b

			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s)-1] {
				count++
			}
		}
	}

	fmt.Println(count)
}
