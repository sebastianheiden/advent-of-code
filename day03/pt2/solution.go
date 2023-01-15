package main

import (
	"log"
	"os"
	"strings"
)

func main() {

	content, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var sum int

	lines := strings.Split(string(content), "\n")

	for i := 0; i < len(lines); i += 3 {
		l1 := lines[i]
		l2 := lines[i+1]
		l3 := lines[i+2]

		for _, char := range l1 {
			in := string(char)
			priority := getPriority(char)
			if strings.Contains(l2, in) && strings.Contains(l3, in) {
				log.Printf("Duplicate char is %s with priority %d", in, priority)
				sum += priority
				break
			}
		}

	}

	log.Printf("Sum priority: %d", sum)
}

func getPriority(char rune) int {
	if char > 96 {
		return int(char - 96)
	} else {
		return int(char - 38)
	}
}
