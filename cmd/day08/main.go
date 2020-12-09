package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/asmolik/aoc2020/internal/utils"
)

type Instruction interface {
	execute(*Program)
	value() int
}

type nop struct {
	_value int
}

func (i nop) execute(p *Program) {
	p.pointer++
}

func (i nop) value() int {
	return i._value
}

type acc struct {
	_value int
}

func (i acc) execute(p *Program) {
	p.accumulator += i._value
	p.pointer++
}

func (i acc) value() int {
	return i._value
}

type jmp struct {
	_value int
}

func (i jmp) execute(p *Program) {
	p.pointer += i._value
}

func (i jmp) value() int {
	return i._value
}

type Program struct {
	accumulator  int
	pointer      int
	instructions []Instruction
	terminated   bool
}

func (program *Program) run() int {
	program.pointer = 0
	program.accumulator = 0
	runInstructions := map[int]bool{}
	for {
		_, exists := runInstructions[program.pointer]
		if exists {
			break
		}
		runInstructions[program.pointer] = true
		i := program.instructions[program.pointer]
		i.execute(program)
		if program.pointer >= len(program.instructions) {
			program.terminated = true
			break
		}
	}
	return program.accumulator
}

func parseInput(lines []string) []Instruction {
	ret := []Instruction{}
	for _, line := range lines {
		a := strings.Split(line, " ")
		instruction_name := a[0]
		v, _ := strconv.Atoi(a[1])
		switch instruction_name {
		case "nop":
			ret = append(ret, nop{v})
		case "acc":
			ret = append(ret, acc{v})
		case "jmp":
			ret = append(ret, jmp{v})
		}
	}
	return ret
}

func main() {
	data_str := utils.ReadLines("test_input")
	run_for_data(data_str)
	data_str = utils.ReadLines("input")
	run_for_data(data_str)
}

func run_for_data(data []string) {
	instructions := parseInput(data)
	fmt.Println(part1(instructions))
	fmt.Println(part2(instructions))
}

func part1(instructions []Instruction) int {
	program := Program{0, 0, instructions, false}
	return program.run()
}

func part2(instructions []Instruction) int {
	program := Program{0, 0, instructions, false}
	for pointer, instruction := range program.instructions {
		program.instructions[pointer] = exchangeInstruction(instruction)
		program.run()
		if program.terminated {
			break
		}
		program.instructions[pointer] = instruction
	}
	return program.accumulator
}

func exchangeInstruction(i interface{}) Instruction {
	switch t := i.(type) {
	case jmp:
		return nop{t.value()}
	case nop:
		return jmp{t.value()}
	case acc:
		return t
	}
	return nop{}
}
