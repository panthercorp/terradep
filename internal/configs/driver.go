package configs

import (
	"fmt"
	"path/filepath"

	"github.com/panthercorp/terradep/pkg/files"
)

// Path should be till folder only, not file
func InitializeTerradepConfig(path string, force bool) error {
	terradepBase := filepath.Join("internal", "res", "terradepBase.yml")
	destinationPath := filepath.Join(path, ".terradep", "terradep.yml")
	if files.CheckFileExistence(destinationPath) && !force {
		return fmt.Errorf("terradep config already exists at %v", destinationPath)
	} else if files.CheckFileExistence(destinationPath) && force {
		if err := files.CreateFileBackup(destinationPath, 0644); err != nil {
			return err
		}
	}
	terradepConfigData, err := files.ReadFile(terradepBase)
	if err != nil {
		return err
	}
	if err := files.WriteFile(destinationPath, terradepConfigData, 0644); err != nil {
		return err
	}
	return nil
}

// Path should be till folder only, not file
func InitializeDcfConfig(path string, force bool) error {
	dcfBase := filepath.Join("internal", "res", "dcfBase.yml")
	destinationPath := filepath.Join(path, "dcf.yml")
	if files.CheckFileExistence(destinationPath) && !force {
		return fmt.Errorf("dcf config already exists at %v. use --force to override", destinationPath)
	} else if files.CheckFileExistence(destinationPath) && force {
		if err := files.CreateFileBackup(destinationPath, 0644); err != nil {
			return err
		}
	}
	dcfConfigData, err := files.ReadFile(dcfBase)
	if err != nil {
		return err
	}
	if err := files.WriteFile(destinationPath, dcfConfigData, 0644); err != nil {
		return err
	}
	return nil
}
