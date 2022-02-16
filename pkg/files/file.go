package files

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"time"
)

func ReadFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func WriteFile(filePath string, data []byte, permission int) error {
	err := ioutil.WriteFile(filePath, data, fs.FileMode(permission))
	if err != nil {
		return err
	}
	return nil
}

func CheckFileExistence(filePath string) bool {
	_, err := ioutil.ReadFile(filePath)
	return err == nil
}

func CreateFileBackup(filePath string, permission int) error {
	if CheckFileExistence(filePath) {
		data, err := ReadFile(filePath)
		if err != nil {
			return err
		}
		return WriteFile(filePath+time.Now().String()+".bak", data, permission)
	} else {
		return fmt.Errorf("file %v does not exist", filePath)
	}
}

func DeleteExistingFile(filePath string) error {
	if CheckFileExistence(filePath) {
		return os.Remove(filePath)
	} else {
		return fmt.Errorf("file %v does not exist", filePath)
	}
}
