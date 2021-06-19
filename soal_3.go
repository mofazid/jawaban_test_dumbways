package main

import "fmt"

func LoopingDumbwaysID() {
	for i := 1; i <= 7; i++ {

		arr := []string{"D", "U", "M", "B", "W", "A", "Y", "S", "I", "D"}

		signOne := "*"
		signTwo := "="

		for j := 0; j <= 9; j++ {

			if i == 4 {
				fmt.Print(arr[j])
				continue
			}

			if i == 3 || i == 5 {
				signOne = "="
				signTwo = "*"
			}

			if j%2 == 0 {
				fmt.Print(signOne)
			} else {
				fmt.Print(signTwo)
			}
		}

		fmt.Println()
	}
}