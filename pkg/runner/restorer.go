package runner

import (
	"fmt"
	"github.com/georgebent/go-restorer/pkg/core"
	"github.com/georgebent/go-restorer/pkg/file_manager"
	"github.com/georgebent/go-restorer/pkg/io_manager"
)

func Restore() error {
	backups := core.GetEnv("BACKUP_DIR")
	source := core.GetEnv("ORIGIN_DIR")

	folders, err := file_manager.ListFolders(backups)
	if err != nil {
		return err
	}

	var options []string
	for _, folder := range folders {
		options = append(options, folder)
	}

	chosen := io_manager.Ask("Choose restore file", options)

	backupsTmpPath := fmt.Sprintf("%s/tmp", backups)
	backupPath := fmt.Sprintf("%s/%s", backups, chosen)

	err = file_manager.ForceCopy(source, backupsTmpPath)
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
