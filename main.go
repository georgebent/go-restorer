package main

import (
	"fmt"
	"github.com/georgebent/go-restorer/pkg/core"
	"github.com/georgebent/go-restorer/pkg/file_manager"
	"github.com/georgebent/go-restorer/pkg/io_manager"
	"github.com/georgebent/go-restorer/pkg/runner"
)

func main() {
	runProgram()
}

func runProgram() {
	options := []string{"Save", "Restore", "List"}
	chosen := io_manager.Ask("Choose action", options)

	println(chosen)

	switch chosen {
	case "Save":
		err := runner.Save()
		if err != nil {
			fmt.Println("Got error: ", err)
		}
	case "Restore":
		err := runner.Restore()
		if err != nil {
			fmt.Println("Got error: ", err)
		}
	case "List":
		backups := core.GetEnv("BACKUP_DIR")
		folders, err := file_manager.ListFolders(backups)
		if err != nil {
			fmt.Println("Got error: ", err)
		}

		for _, folder := range folders {
			println(folder)
		}
	}

	println()

	runProgram()
}
