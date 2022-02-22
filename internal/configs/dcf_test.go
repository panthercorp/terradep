package configs

import (
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestUnmarshallDcf(t *testing.T) {
	dcfData, err := ioutil.ReadFile(filepath.Join("..", "res", "dcf_test.yml"))
	if err != nil {
		t.Errorf("Failed to initialize test data: %s", err)
	}
	var dcf Dcf
	if err = yaml.Unmarshal(dcfData, &dcf); err != nil {
		t.Errorf("Failed to unmarshal test data: %s", err)
	}
	type args struct {
		file []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Dcf
		wantErr bool
	}{
		{
			name: "UnmarshalDcf",
			args: args{
				file: dcfData,
			},
			want:    dcf,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshallDcf(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshallDcf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnmarshallDcf() = %v, want %v", got, tt.want)
			}
		})
	}
}
