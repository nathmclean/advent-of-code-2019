package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main()  {
	file, err := ioutil.ReadFile("day-1/input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")
	var total int

	for _, line := range lines {
		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			panic(err)
		}

		by3 := value / 3
		moduleWeight := int(by3 -2)

		total += moduleWeight + calculateFuelWeight(moduleWeight)
	}
	fmt.Println("Total:", total)
}

func calculateFuelWeight(weight int) int {
	fuelWeight := (weight / 3) - 2
	if fuelWeight > 0 {
		return fuelWeight + calculateFuelWeight(fuelWeight)
	}
	return 0
}