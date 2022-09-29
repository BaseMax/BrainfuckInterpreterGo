/*
 * Name: Brainfuck Interpreter in Go
 * Repository: https://github.com/BaseMax/BrainfuckInterpreterGo
 * Author: Max Base
 * Date: 2022/09/23
 */

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	MvRight    = '>'
	MvLeft     = '<'
	IncMem     = '+'
	DecMem     = '-'
	Output     = '.'
	Input      = ','
	BraceOpen  = '['
	BraceClose = ']'
)

var stdinReader = bufio.NewReader(os.Stdin)

func readFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func interpret(bf []byte) []byte {
	var (
		emptySlice = make([]byte, 10)
		memory     = make([]byte, 10)
		pointer    = 0
		builder    = strings.Builder{}
	)

	for i := 0; i < len(bf); i++ {
		// increase memory size if i is higher or equal to length of memory slice
		if pointer >= len(memory) {
			memory = append(memory, emptySlice...)
		}
		switch bf[i] {
		case MvRight:
			pointer++
		case MvLeft:
			pointer--
		case IncMem:
			memory[pointer]++
		case DecMem:
			memory[pointer]--
		case Output:
			builder.WriteByte(memory[pointer])

		case Input:
			in, err := stdinReader.ReadByte()
			if err != nil {
				panic(err)
			}
			memory[pointer] = in
		case BraceOpen:
			if memory[pointer] == 0 {
				count := 1
				for count > 0 {
					i++
					if bf[i] == BraceOpen {
						count++
					} else if bf[i] == BraceClose {
						count--
					}
				}
			}
		case BraceClose:
			if memory[pointer] != 0 {
				count := 1
				for count > 0 {
					i--
					if bf[i] == BraceClose {
						count++
					}
					if bf[i] == BraceOpen {
						count--
					}
				}
			}
		}
	}

	return []byte(builder.String())
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <file>", os.Args[0])
		return
	}

	content, err := readFile(os.Args[1])
	if err != nil {
		log.Fatalf("err occured: %s\n", err.Error())
	}
	res := interpret(content)
	fmt.Printf("%s\n", res)
}
