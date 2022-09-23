/*
 * Name: Brainfuck Interpreter in Go
 * Repository: https://github.com/BaseMax/BrainfuckInterpreterGo
 * Author: Max Base
 * Date: 2022/09/23
 */

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

char* read_file(char* file_path)
{
    char* data = NULL;
    FILE* file = fopen(file_path, "r");
    if (file == NULL) {
        return NULL;
    }

    fseek(file, 0, SEEK_END);
    long file_size = ftell(file);
    rewind(file);

    data = malloc(file_size + 1);
    fread(data, file_size, 1, file);
    data[file_size] = '\0';

    fclose(file);
    return data;
}

char* interpret(char* bf)
{
    // >	Move the pointer to the right
    // <	Move the pointer to the left
    // +	Increment the memory cell at the pointer
    // -	Decrement the memory cell at the pointer
    // .	Output the character signified by the cell at the pointer
    // ,	Input a character and store it in the cell at the pointer
    // [	Jump past the matching ] if the cell at the pointer is 0
    // ]	Jump back to the matching [ if the cell at the pointer is nonzero
    // All characters other than ><+-.,[] should be considered comments and ignored. But, see extensions below.

    char* output = malloc(30000);
    // initialize memory
    int memory[30000];
    // initialize pointer
    int pointer = 0;
    // initialize output
    int output_pointer = 0;

    // loop through input
    int i = 0;
    while (i < strlen(bf)) {
        // >	Move the pointer to the right
        if (bf[i] == '>') {
            pointer += 1;
        }
        // <	Move the pointer to the left
        else if (bf[i] == '<') {
            pointer -= 1;
        }
        // +	Increment the memory cell at the pointer
        else if (bf[i] == '+') {
            memory[pointer] += 1;
        }
        // -	Decrement the memory cell at the pointer
        else if (bf[i] == '-') {
            memory[pointer] -= 1;
        }
        // .	Output the character signified by the cell at the pointer
        else if (bf[i] == '.') {
            output[output_pointer] = memory[pointer];
            output_pointer += 1;
        }
        // ,	Input a character and store it in the cell at the pointer
        else if (bf[i] == ',') {
            memory[pointer] = getchar();
        }
        // [	Jump past the matching ] if the cell at the pointer is 0
        else if (bf[i] == '[') {
            if (memory[pointer] == 0) {
                // find matching ]
                int count = 1;
                while (count > 0) {
                    i += 1;
                    if (bf[i] == '[') {
                        count += 1;
                    }
                    else if (bf[i] == ']') {
                        count -= 1;
                    }
                }
            }
        }
        // ]	Jump back to the matching [ if the cell at the pointer is nonzero
        else if (bf[i] == ']') {
            if (memory[pointer] != 0) {
                // find matching [
                int count = 1;
                while (count > 0) {
                    i -= 1;
                    if (bf[i] == ']') {
                        count += 1;
                    }
                    else if (bf[i] == '[') {
                        count -= 1;
                    }
                }
            }
        }
        else {
            // pass
        }
        i += 1;
    }

    // add \0 to output
    output[output_pointer] = '\0';

    return output;
}

int main(int argc, char** argv)
{
    if (argc < 2) {
        printf("Usage: %s <file>", argv[0]);
        return 1;
    }
    else {
        char* code = read_file(argv[1]);
        if (code == NULL) {
            printf("File not found");
            return 1;
        }
        else {
            char* result = interpret(code);
            printf("%s", result);

            free(code);
            free(result);
        }
    }
    return 0;
}
