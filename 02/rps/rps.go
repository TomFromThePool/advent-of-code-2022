package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Rock = 1
//Paper = 2
//Scissors = 3

const file = "input.txt"

func Readln(r *bufio.Reader) (string, error) {
	var (
		isPrefix bool  = true
		err      error = nil
		line, ln []byte
	)
	for isPrefix && err == nil {
		line, isPrefix, err = r.ReadLine()
		ln = append(ln, line...)
	}
	return string(ln), err
}

func Split(line string) []int {
	s := strings.Split(line, " ")
	vals := make([]int, len(s))
	for i, shape := range s {
		vals[i] = Normalise(shape)
	}
	return vals
}

const ASCII_A = 65
const ASCII_C = 67
const ASCII_X = 88

// ASCII FUN - char value to shape score
func Normalise(shape string) int {
	x := int(shape[0])
	var val int
	if x <= ASCII_C {
		val = 1 + x - ASCII_A //Rock: A, Paper: B, Scissors: C
	} else {
		val = 1 + x - ASCII_X //Rock: X, Paper: Y, Scissors: Z
	}
	return val
}

const win = 6
const lose = 0
const draw = 3

func MatchInstruction(opponent int, instruction int) int {
	var i int
	if instruction == 1 {
		i = opponent - 1
		//lose
	} else if instruction == 2 {
		//draw
		i = opponent
	} else {
		//win
		i = opponent + 1
	}

	if i < 1 {
		i = 3
	} else if i > 3 {
		i = 1
	}
	return i
}

func Score(round []int, instruct bool) int {
	o := round[0]
	var m int
	if instruct {
		m = MatchInstruction(o, round[1])
	} else {
		m = round[1]
	}
	var result string
	var score int
	var bonus int
	var rps = (m - o) % 3
	flop := false

	if rps < 0 {
		rps = -rps
		flop = true
	}

	if rps == 0 {
		result = "Draw"
		bonus = draw
	} else if (!flop && rps == 1) || (flop && rps > 1) {
		result = "Win"
		bonus = win
	} else {
		result = "Lose"
		bonus = lose
	}

	score = bonus + m

	fmt.Printf("Instruction ? %v: %v (%v) VS %v (%v): %v - Bonus: %v | Score %v\n", InstructionName(round[1], instruct), Name(o), o, Name(m), Value(Name(m)), result, bonus, score)
	return score
}

func InstructionName(x int, instruct bool) string {
	if instruct {
		if x == 1 {
			return "Lose"
		} else if x == 2 {
			return "Draw"
		} else {
			return "Win"
		}
	} else {
		return "Compete"
	}
}

func Name(x int) string {
	if x == 1 {
		return "Rock"
	} else if x == 2 {
		return "Paper"
	} else {
		return "Scissors"
	}
}

func Value(name string) int {
	var i int
	switch name {
	case "Rock":
		i = 1
	case "Paper":
		i = 2
	case "Scissors":
		i = 3
	}
	return i
}

func Play(round string, instruct bool) int {
	return Score(Split(round), instruct)
}

func main() {
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Could not read %v\n", err)
	}

	r := bufio.NewReader(f)
	s, e := Readln(r)
	part1Total := 0
	part2Total := 0
	for e == nil {
		part1Total += Play(s, false)
		part2Total += Play(s, true)
		s, e = Readln(r)
	}

	fmt.Printf("Part 1 Total Score: %v\n", part1Total)
	fmt.Printf("Part 2 Total Score: %v\n", part2Total)

}
