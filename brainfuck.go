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
	"bytes"
)

const (
	Input      = ','
	IncMem     = '+'
	DecMem     = '-'
	Output     = '.'
	MvLeft     = '<'
	MvRight    = '>'
	BraceOpen  = '['
	BraceClose = ']'
)

var stdinReader = bufio.NewReader(os.Stdin)

func readFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

func interpret(bf []byte) []byte {
	var (
		output        []byte = make([]byte, 30000)
		memory        []byte = make([]byte, 30000)
		pointer       = 0
		outputPointer = 0
	)

	for i := 0; i < len(bf); i++ {
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
			output[outputPointer] = memory[pointer]
			outputPointer++
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
	
	res = bytes.Trim(res, "\x00")
	fmt.Printf("%#v\n", res)
}
