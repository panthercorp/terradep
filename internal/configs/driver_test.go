package configs

import "testing"

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
		// TODO fix tests
		// {
		// 	name: "InitializeTerradepConfig",
		// 	args: args{
		// 		path:  "",
		// 		force: false,
		// 	},
		// 	wantErr: false,
		// },
		// {
		// 	name: "InitializeTerradepConfig",
		// 	args: args{
		// 		path:  "",
		// 		force: false,
		// 	},
		// 	wantErr: true,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InitializeTerradepConfig(tt.args.path, tt.args.force); (err != nil) != tt.wantErr {
				t.Errorf("InitializeTerradepConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInitializeDcfConfig(t *testing.T) {
	type args struct {
		path  string
		force bool
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
			if err := InitializeDcfConfig(tt.args.path, tt.args.force); (err != nil) != tt.wantErr {
				t.Errorf("InitializeDcfConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
