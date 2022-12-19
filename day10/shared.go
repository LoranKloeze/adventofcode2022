package day10

import (
	"bufio"
	"io"
	"log"
	"strconv"
	"strings"
)

func parseInstructions(r io.Reader) []Instruction {
	instructions := []Instruction{}

	s := bufio.NewScanner(r)
	for s.Scan() {
		spl := strings.Split(s.Text(), " ")
		instruction := Instruction{Operation: spl[0]}

		if len(spl) == 2 {
			val, err := strconv.Atoi(spl[1])
			if err != nil {
				log.Fatalf("Expected a number: %v", spl[0])
			}
			instruction.Value = val
		}

		instructions = append(instructions, instruction)
	}

	return instructions
}
