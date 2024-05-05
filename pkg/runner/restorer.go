package runner

import (
	"fmt"
	"github.com/georgebent/go-restorer/pkg/core"
	"github.com/georgebent/go-restorer/pkg/file_manager"
	"github.com/georgebent/go-restorer/pkg/io_manager"
)

func Restore(name string) error {
	// Todo: add options from backups in folder
	options := map[int]string{
		1: "yes",
		2: "no",
	}

	chosen := io_manager.Ask("Are you sure ?", options)
	if chosen != 1 {
		return nil
	}

	source := core.GetEnv("ORIGIN_DIR")
	backups := core.GetEnv("BACKUP_DIR")
	backupsTmpPath := fmt.Sprintf("%s/tmp", backups)
	backupPath := fmt.Sprintf("%s/%s", backups, name)

	err := file_manager.ForceCopy(source, backupsTmpPath)
	if err != nil {
		return err
	}

	err = file_manager.ForceCopy(backupPath, source)
	if err != nil {
		return err
	}

	io_manager.Write(fmt.Sprintf("Folder %s restored from %s", source, backupPath))

	return nil
}
