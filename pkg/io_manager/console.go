package io_manager

import (
	"bufio"
	"fmt"
	"github.com/manifoldco/promptui"
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

func Ask(question string, options []string) string {
	if len(options) == 0 {
		os.Exit(0)
	}

	prompt := promptui.Select{
		Label: question,
		Items: options,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)

		return ""
	}

	fmt.Printf("You choose %q\n", result)

	return result
}
