package configs

import (
	"reflect"
	"testing"
)

func TestUnmarshallTerradep(t *testing.T) {
	type args struct {
		file []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Terradep
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UnmarshallTerradep(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshallTerradep() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnmarshallTerradep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTerradep_GetGitRepository(t *testing.T) {
	type fields struct {
		Repositories Repositories
	}
	type args struct {
		alias string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Git
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Terradep{
				Repositories: tt.fields.Repositories,
			}
			got, err := tr.GetGitRepository(tt.args.alias)
			if (err != nil) != tt.wantErr {
				t.Errorf("Terradep.GetGitRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Terradep.GetGitRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTerradep_GetHTTPRepository(t *testing.T) {
	type fields struct {
		Repositories Repositories
	}
	type args struct {
		alias string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    HTTP
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Terradep{
				Repositories: tt.fields.Repositories,
			}
			got, err := tr.GetHTTPRepository(tt.args.alias)
			if (err != nil) != tt.wantErr {
				t.Errorf("Terradep.GetHTTPRepository() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Terradep.GetHTTPRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
