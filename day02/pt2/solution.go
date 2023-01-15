package main

import (
	"log"
	"os"
	"strings"
)

type TypeSign uint8

const (
	R TypeSign = iota
	P
	S
)

type Type struct {
	sign     TypeSign
	beats    TypeSign
	isBeaten TypeSign
	value    int
}

func main() {

	content, err := os.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	rock := Type{
		sign:     R,
		beats:    S,
		isBeaten: P,
		value:    1,
	}

	paper := Type{
		sign:     P,
		beats:    R,
		isBeaten: S,
		value:    2,
	}

	scissor := Type{
		sign:     S,
		beats:    P,
		isBeaten: R,
		value:    3,
	}

	var signMap = make(map[string]Type)
	signMap["A"] = rock
	signMap["B"] = paper
	signMap["C"] = scissor

	signMap["X"] = rock
	signMap["Y"] = paper
	signMap["Z"] = scissor

	var typeMap = make(map[TypeSign]Type)
	typeMap[R] = rock
	typeMap[P] = paper
	typeMap[S] = scissor

	points := 0

	for _, line := range strings.Split(string(content), "\n") {

		parts := strings.Split(line, " ")
		opponentType := signMap[strings.TrimSpace(parts[0])]
		var yourType Type
		winIndicator := strings.TrimSpace(parts[1])

		if winIndicator == "X" {
			yourType = typeMap[opponentType.beats]
		} else if winIndicator == "Y" {
			yourType = typeMap[opponentType.sign]
		} else {
			yourType = typeMap[opponentType.isBeaten]
		}

		points += yourType.value

		if yourType.beats == opponentType.sign {
			points += 6
		} else if yourType.sign == opponentType.sign {
			points += 3
		}
	}

	log.Println("Points:", points)
}
