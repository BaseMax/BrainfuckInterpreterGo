/*
 * Name: Brainfuck Interpreter in Go
 * Repository: https://github.com/BaseMax/BrainfuckInterpreterGo
 * Author: Max Base
 * Date: 2022/09/23
 */

package main

import (
	"os"
	"fmt"
	"log"
	"bufio"
)

const (
	MvRight    byte = '>'
	MvLeft     byte = '<'
	IncMem     byte = '+'
	DecMem     byte = '-'
	Output     byte = '.'
	Input      byte = ','
	BraceOpen  byte = '['
	BraceClose byte = ']'
	MemSize         = 30000
)

var stdinReader = bufio.NewReader(os.Stdin)

func readFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func interpret(bf []byte) [MemSize]byte {
	var (
		output        [MemSize]byte
		memory        [MemSize]byte
		pointer       = 0
		outputPointer = 0
	)

	i := 0
	for i < len(bf) {
		if bf[i] == MvRight {
			pointer++
		}

		if bf[i] == MvLeft {
			pointer--
		}

		if bf[i] == IncMem {
			memory[pointer]++
		}

		if bf[i] == DecMem {
			memory[pointer]--
		}

		if bf[i] == Output {
			output[outputPointer] = memory[pointer]
			outputPointer++
		}

		if bf[i] == Input {
			in, err := stdinReader.ReadByte()
			if err != nil {
				panic(err)
			}
			memory[pointer] = in
		}

		if bf[i] == BraceOpen {
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
		}

		if bf[i] == BraceClose {
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
		i++
	}

	return output
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
