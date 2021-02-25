package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Adrian Mucha, Miękka klucha :)")

	if stars := os.Getenv("FUCK_PEACE"); stars != "" {
		num, err := strconv.Atoi(stars)
		switch {
		case err != nil:
			panic("maciej kabala hehe ;~~~D")
		case num == 8:
			fmt.Println("***** ***")
		default:
			for i := 0; i < num; i++ {
				for j := 0; j < (num - i - 1); j++ {
					fmt.Print(" ")
				}
				for j := 0; j < (i*2)+1; j++ {
					fmt.Print("*")
				}
				fmt.Println()
			}
		}
	}
}
