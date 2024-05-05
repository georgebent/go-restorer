package main

import (
	"fmt"
	"github.com/georgebent/go-restorer/pkg/runner"
)

func main() {
	fmt.Println("Run restore program")

	err := runner.Restore("1.quick_save")
	if err != nil {
		fmt.Println("Got error: ", err)
	}
}
