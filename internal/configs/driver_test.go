package configs

import (
	"os"
	"testing"
)

func TestInitializeTerradepConfig(t *testing.T) {
	type args struct {
		path  string
		force bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "InitializeTerradepConfig",
			args: args{
				path:  "",
				force: false,
			},
			wantErr: false,
		},
		{
			name: "InitializeTerradepConfigWithFileAlreadyPresent",
			args: args{
				path:  "",
				force: false,
			},
			wantErr: true,
		},
		{
			name: "InitializeTerradepConfigWithFileAlreadyPresentForce",
			args: args{
				path:  "",
				force: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitializeTerradepConfig(tt.args.path, tt.args.force); (err != nil) != tt.wantErr {
				t.Errorf("InitializeTerradepConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	err := os.RemoveAll(".terradep")
	if err != nil {
		t.Error("Failed to clean test resources: ", err)
	}
}

func TestInitializeDcfConfig(t *testing.T) {
	os.Mkdir("dcftest", 0755)
	type args struct {
		path  string
		force bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "InitializeDcfConfig",
			args: args{
				path:  "dcftest",
				force: false,
			},
			wantErr: false,
		},
		{
			name: "InitializeDcfConfigWithFileAlreadyPresent",
			args: args{
				path:  "dcftest",
				force: false,
			},
			wantErr: true,
		},
		{
			name: "InitializeDcfConfigWithFileAlreadyPresentForce",
			args: args{
				path:  "dcftest",
				force: true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitializeDcfConfig(tt.args.path, tt.args.force); (err != nil) != tt.wantErr {
				t.Errorf("InitializeDcfConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	err := os.RemoveAll("dcftest")
	if err != nil {
		t.Error("Failed to clean test resources: ", err)
	}
}
