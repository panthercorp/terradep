package configs

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/panthercorp/terradep/internal/statics"
	"github.com/panthercorp/terradep/pkg/files"
)

// Path should be till folder only, not file
// TODO fix relative path for copying files
func InitializeTerradepConfig(path string, force bool) error {
	destinationPath := filepath.Join(path, ".terradep", "terradep.yml")
	os.Mkdir(filepath.Join(path, ".terradep"), 0755)
	if files.CheckFileExistence(destinationPath) && !force {
		return fmt.Errorf("terradep config already exists at %v. use --force to override", destinationPath)
	} else if files.CheckFileExistence(destinationPath) && force {
		if err := files.CreateFileBackup(destinationPath, 0755); err != nil {
			return err
		}
	}
	terradepConfigData := []byte(statics.TerradepConfig)
	if err := files.WriteFile(destinationPath, terradepConfigData, 0755); err != nil {
		return err
	}
	return nil
}

// Path should be till folder only, not file
// TODO fix relative path for copying files
func InitializeDcfConfig(path string, force bool) error {
	destinationPath := filepath.Join(path, "dcf.yml")
	if files.CheckFileExistence(destinationPath) && !force {
		return fmt.Errorf("dcf config already exists at %v. use --force to override", destinationPath)
	} else if files.CheckFileExistence(destinationPath) && force {
		if err := files.CreateFileBackup(destinationPath, 0755); err != nil {
			return err
		}
	}
	dcfConfigData := []byte(statics.DcfConfig)
	if err := files.WriteFile(destinationPath, dcfConfigData, 0755); err != nil {
		return err
	}
	return nil
}
