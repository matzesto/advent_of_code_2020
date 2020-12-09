package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// Instruction type
type Instruction struct {
	cmd      string
	arg      int
	executed bool
}

// Ctx Execution
type Ctx struct {
	program []Instruction
	ip      int
	acc     int
}

// Return a fresh Execution context
func makeCtx(program []Instruction) Ctx {
	ctx := Ctx{
		acc:     0,
		ip:      0,
		program: program,
	}
	for i := range ctx.program {
		ctx.program[i].executed = false
	}
	return ctx
}

func parse(input string) Ctx {
	var program []Instruction

	for _, line := range strings.Split(input, "\n") {
		var instr Instruction
		tok := strings.Split(line, " ")
		if len(tok) != 2 {
			continue
		}
		instr.cmd = tok[0]
		instr.arg, _ = strconv.Atoi(tok[1])
		instr.executed = false
		program = append(program, instr)
	}

	return makeCtx(program)
}

// Given a context, execute one step
func exec(ctx Ctx) (Ctx, error) {
	ctx.program[ctx.ip].executed = true
	switch instr := ctx.program[ctx.ip]; instr.cmd {
	case "nop":
		ctx.ip++
	case "acc":
		ctx.acc += instr.arg
		ctx.ip++
	case "jmp":
		ctx.ip += instr.arg // Check wether this jump is legal yadda yadda
		if ctx.ip >= len(ctx.program) {
			return ctx, errors.New("Illegal Jump")
		}
	}

	return ctx, nil

}

// execute until termination or loop detected
func execUntil(ctx Ctx) Ctx {
	for {
		if ctx.ip >= len(ctx.program) {
			return ctx
		}
		instr := ctx.program[ctx.ip]
		if instr.executed {
			return ctx
		}
		ctx, _ = exec(ctx)
	}
}

// Exec until termination
func execUntilTerminated(ctx Ctx) (Ctx, bool) {
	ctx = execUntil(ctx)
	if ctx.ip == len(ctx.program) {
		return ctx, true
	}
	return ctx, false
}

// Test wether the program termiantes or loops
func check(program []Instruction) bool {
	ctx := makeCtx(program)
	ctx, terminated := execUntilTerminated(ctx)
	if terminated {
		return true
	}
	return false

}

// Compute a terminating program with one instruction changed
func fix(program []Instruction) []Instruction {
	for i, instr := range program {
		var fixedProgram = make([]Instruction, len(program))
		copy(fixedProgram, program)
		switch instr.cmd {

		// Replace NOPs with jumps of at most len(program) arg
		case "nop":
			fixedProgram[i].cmd = "jmp"
			for j := range program {
				fixedProgram[i].arg = j
				if check(fixedProgram) {
					return fixedProgram
				}
			}
		// Replace JMP with a nop, ignoring all args
		case "jmp":
			fixedProgram[i].cmd = "nop"
			if check(fixedProgram) {
				return fixedProgram
			}
		default:
			continue
		}
	}
	return nil
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	r := bufio.NewReader(file)

	text, err := ioutil.ReadAll(r)

	ctx := parse(string(text))
	fmt.Println("Taks 1: ", execUntil(ctx).acc)

	ctx = parse(string(text))
	fixedProgram := fix(ctx.program)

	ctx = makeCtx(fixedProgram)
	fmt.Println("Task 2: ", execUntil(ctx).acc)
}
