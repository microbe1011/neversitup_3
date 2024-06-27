package main

import (
	"errors"
	"fmt"
	"log"
)

func findOddOccurs(intArr []int) (int, error) {
	counter := make(map[int]int)
	for _, item := range intArr {
		counter[item] += 1
	}

	for k, v := range counter {
		if v%2 == 1 {
			return k, nil
		}
	}

	return -1, errors.New("no odd number found")
}

func main() {
	intArr := []int{90, 90}
	occurs, err := findOddOccurs(intArr)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(occurs, "found odd times in array")
}
