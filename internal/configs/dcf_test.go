package configs

import (
	"reflect"
	"testing"
)

func TestUnmarshallDcf(t *testing.T) {
	type args struct {
		file []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Dcf
		wantErr bool
	}{
		// TODO: Add test cases.
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
