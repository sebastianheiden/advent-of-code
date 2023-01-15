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

	var i int
	var sum int

	for _, line := range strings.Split(string(content), "\n") {

		line = strings.TrimSpace(line)

		first := line[:len(line)/2]
		second := line[len(line)/2:]

		for _, char := range first {
			in := string(char)
			priority := getPriority(char)
			if strings.Contains(second, in) {
				log.Printf("Duplicate char is %s with priority %d", in, priority)
				i++
				sum += priority
				break
			}
		}
	}

	log.Printf("Found %d duplicate chars", i)
	log.Printf("Sum priority: %d", sum)
}

func getPriority(char rune) int {
	if char > 96 {
		return int(char - 96)
	} else {
		return int(char - 38)
	}
}
