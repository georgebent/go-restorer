package main

import (
	"fmt"
	"github.com/georgebent/go-restorer/pkg/runner"
)

func main() {
	err := runner.Save("quick_save")
	if err != nil {
		fmt.Println("Got error: ", err)
	}
}
