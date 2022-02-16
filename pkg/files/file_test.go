package files

import (
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadFile(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteFile(t *testing.T) {
	type args struct {
		filePath   string
		data       []byte
		permission int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteFile(tt.args.filePath, tt.args.data, tt.args.permission); (err != nil) != tt.wantErr {
				t.Errorf("WriteFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckFileExistence(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckFileExistence(tt.args.filePath); got != tt.want {
				t.Errorf("CheckFileExistence() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateFileBackup(t *testing.T) {
	type args struct {
		filePath   string
		permission int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateFileBackup(tt.args.filePath, tt.args.permission); (err != nil) != tt.wantErr {
				t.Errorf("CreateFileBackup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDeleteExistingFile(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteExistingFile(tt.args.filePath); (err != nil) != tt.wantErr {
				t.Errorf("DeleteExistingFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
