package files

import (
	"io/fs"
	"io/ioutil"
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
