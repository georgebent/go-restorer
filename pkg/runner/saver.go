package runner

import (
	"fmt"
	"github.com/georgebent/go-restorer/pkg/core"
	"github.com/georgebent/go-restorer/pkg/file_manager"
	"github.com/georgebent/go-restorer/pkg/io_manager"
)

func QuickSave() error {
	return saveByName("quick_save")
}

func Save() error {
	name := io_manager.Read("Enter backup name: \n")

	return saveByName(name)
}

func saveByName(name string) error {
	backups := core.GetEnv("BACKUP_DIR")
	folders, err := file_manager.ListFolders(backups)
	if err != nil {
		return err
	}

	length := len(folders) + 1
	name = fmt.Sprintf("%d.%s", length, name)

	source := core.GetEnv("ORIGIN_DIR")
	backupPath := fmt.Sprintf("%s/%s", backups, name)

	err = file_manager.Copy(source, backupPath)
	if err != nil {
		return err
	}

	io_manager.Write(fmt.Sprintf("Folder %s saved to %s", source, backupPath))

	return nil
}
