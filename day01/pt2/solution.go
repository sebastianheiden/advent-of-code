package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	tmp := 0
	i := 0
	sumedCalories := []int{}

	for _, line := range strings.Split(string(content), "\n") {

		trimmedLine := strings.TrimSpace(line)

		if trimmedLine == "" {
			sumedCalories = append(sumedCalories, tmp)
			log.Printf("Elf #%d with carries %d calories", i, tmp)
			i++
			tmp = 0
			continue
		}

		caloriesInLine, err := strconv.Atoi(trimmedLine)
		if err != nil {
			log.Fatal(err)
		}
		tmp += caloriesInLine
	}

	sort.Sort(sort.Reverse(sort.IntSlice(sumedCalories)))
	maxThree := sumedCalories[0] + sumedCalories[1] + sumedCalories[2]

	log.Println("Max three:", maxThree)
}
