package io_manager

import (
	"fmt"
	"os"
	"strings"
)

func Write(description string) {
	fmt.Println(description)
}

func Ask(question string, options map[int]string) int {
	if len(options) == 0 {
		os.Exit(0)
	}
	fmt.Println(question)
	for i, variant := range options {
		fmt.Printf("\x1b[34m%d. %s\n\x1b[0m", i, variant)
	}
	fmt.Println("\x1b[90m\nВведіть номер відповідного варіанту або 'exit' для виходу:\u001B[0m")

	var input string
	_, err := fmt.Scanf("%v", &input)
	if err != nil {
		fmt.Println("Помилка зчитування вводу:", err)
		return Ask("Помилка зчитування вводу. Будь ласка, спробуйте ще раз.", options)
	}

	if strings.ToLower(input) == "exit" {
		fmt.Println("Вихід з програми...")
		os.Exit(0)
	}

	choice := strings.TrimSpace(input)
	var choiceNum int
	_, err = fmt.Sscanf(choice, "%d", &choiceNum)
	if err != nil || choiceNum < 1 || choiceNum > len(options) {
		return Ask("Вибір не вірний. Будь ласка, спробуйте ще раз.", options)
	}

	return choiceNum
}
