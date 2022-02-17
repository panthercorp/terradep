package files

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestReadFile(t *testing.T) {
	data := []byte("Hello World")
	err := ioutil.WriteFile("./test_hello", data, 0755)
	if err != nil {
		t.Error("failed to initialize test")
	}
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "read file test if exists",
			args: args{
				filePath: "./test_hello",
			},
			want:    data,
			wantErr: false,
		},
		{
			name: "throw error if file does not exist",
			args: args{
				filePath: "./test_hello_not_exist",
			},
			want:    nil,
			wantErr: true,
		},
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
	err = os.Remove("./test_hello")
	if err != nil {
		t.Error("failed to cleanup test resources")
	}
}

func TestWriteFile(t *testing.T) {
	data := []byte("Hello World")
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
		{
			name: "write file test",
			args: args{
				filePath:   "./test_hello",
				data:       data,
				permission: 0755,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := WriteFile(tt.args.filePath, tt.args.data, tt.args.permission); (err != nil) != tt.wantErr {
				t.Errorf("WriteFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			_, err := ioutil.ReadFile(tt.args.filePath)
			if err != nil {
				t.Error("failed to read file:", err)
			}
		})
	}
	err := os.Remove("./test_hello")
	if err != nil {
		t.Error("failed to cleanup test resources")
	}
}

func TestCheckFileExistence(t *testing.T) {
	data := []byte("Hello World")
	err := ioutil.WriteFile("./test_hello", data, 0755)
	if err != nil {
		t.Error("failed to initialize test")
	}
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "check file existence test",
			args: args{
				filePath: "./test_hello",
			},
			want: true,
		},
		{
			name: "check file existence test",
			args: args{
				filePath: "./test_hello_not_exist",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckFileExistence(tt.args.filePath); got != tt.want {
				t.Errorf("CheckFileExistence() = %v, want %v", got, tt.want)
			}
		})
	}
	err = os.Remove("./test_hello")
	if err != nil {
		t.Error("failed to cleanup test resources")
	}
}

func TestCreateFileBackup(t *testing.T) {
	data := []byte("Hello World")
	err := ioutil.WriteFile("./test_hello", data, 0755)
	if err != nil {
		t.Error("failed to initialize test")
	}
	type args struct {
		filePath   string
		permission int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "create file backup test",
			args: args{
				filePath:   "./test_hello",
				permission: 0755,
			},
			wantErr: false,
		},
		{
			name: "create file backup of non existing files test",
			args: args{
				filePath:   "./test_hello_not_exist",
				permission: 0755,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateFileBackup(tt.args.filePath, tt.args.permission); (err != nil) != tt.wantErr {
				t.Errorf("CreateFileBackup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	files, err := filepath.Glob("./test_hello*")
	if err != nil {
		t.Error("failed to get test generated files")
	}
	for _, file := range files {
		err = os.Remove(file)
		if err != nil {
			t.Error("failed to cleanup test resources")
		}
	}
}

func TestDeleteExistingFile(t *testing.T) {
	data := []byte("Hello World")
	err := ioutil.WriteFile("./test_hello", data, 0755)
	if err != nil {
		t.Error("failed to initialize test")
	}
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "delete existing file test",
			args: args{
				filePath: "./test_hello",
			},
			wantErr: false,
		},
		{
			name: "delete non existing file test",
			args: args{
				filePath: "./test_hello_not_exist",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteExistingFile(tt.args.filePath); (err != nil) != tt.wantErr {
				t.Errorf("DeleteExistingFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	_, err = ioutil.ReadFile("./test_hello")
	if err == nil {
		err = os.Remove("./test_hello")
		if err != nil {
			t.Error("failed to cleanup test resources")
		}
	}

}
