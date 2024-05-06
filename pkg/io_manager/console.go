package io_manager

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Write(description string) {
	fmt.Println(description)
}

func Read(question string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(question)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	fmt.Println("Ви ввели:", text)

	return text
}

func Ask(question string, options map[string]string) string {
	if len(options) == 0 {
		os.Exit(0)
	}

	menu := NewMenu(question)
	for i, variant := range options {
		menu.AddItem(variant, i)
	}

	choice := menu.Display()

	fmt.Printf("Choice: %s\n", options[choice])

	if strings.ToLower(choice) == "exit" {
		fmt.Println("Вихід з програми...")
		os.Exit(0)
	}

	_, ok := options[choice]
	if !ok {
		return Ask("Вибір не вірний. Будь ласка, спробуйте ще раз.", options)
	}

	return choice
}
