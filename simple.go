package main

import (
	"fmt"
	"bufio"
	"os"
//	"bytes"
	"regexp"
	"unicode"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a string: ")
	input, _ := reader.ReadString('\n')
	// Declare regular expressions 'handlers'
	id_regex, _ := regexp.Compile("^[a-z][a-z]*")
	num_regex, _ := regexp.Compile("^[0-9][0-9]*")
	ws_regex, _ := regexp.Compile("^\\s")

	// Program logic
	flag := true
	var current_position int
	var current_char rune
	for flag {
		if input[0] == '+' {
			fmt.Println("PLUS")
			input = input[1:]
		} else if input[0] == '*' {
			fmt.Println("TIMES")
			input = input[1:]
		} else if ws_regex.MatchString(input) {
			input = input[1:]
		} else if id_regex.MatchString(input){
			current_position = 0
			current_char = rune(input[current_position])
			for unicode.IsLetter(current_char) {
				current_position = current_position + 1
				if current_position >= len(input){
					break
				}
				current_char = rune(input[current_position])
			}
			lexeme := input[:current_position]
			input = input[current_position:]
			fmt.Println("ID", lexeme)

		} else if num_regex.MatchString(input){
			current_position = 0
			current_char = rune(input[current_position])
			for unicode.IsNumber(current_char) {
				current_position = current_position + 1
				if current_position >= len(input){
					break
				}
				current_char = rune(input[current_position])
			}
			lexeme := input[:current_position]
			input = input[current_position:]
			fmt.Println("NUMBER", lexeme)

		} else {
			fmt.Println("Lexical Error")
			flag = false
		}
		if len(input) == 0 {
			flag = false
		}
	} 


}
