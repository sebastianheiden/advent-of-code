package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func solution_pt1() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	max := 0
	tmp := 0
	i := 0

	for _, line := range strings.Split(string(content), "\n") {

		trimmedLine := strings.TrimSpace(line)

		if trimmedLine == "" {
			i++
			if tmp > max {
				max = tmp
				tmp = 0
				log.Printf("Elf #%d with new max %d", i, max)
			} else {
				log.Printf("Elf #%d with lower than max %d", i, tmp)
				tmp = 0
			}
			continue
		}

		caloriesInLine, err := strconv.Atoi(trimmedLine)
		if err != nil {
			log.Fatal(err)
		}
		tmp += caloriesInLine
	}

	log.Println("Max calories:", max)
}
