package files_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/aqyuki/goutil/files"
	"github.com/google/uuid"
)

func TestExistFile(t *testing.T) {
	tmp := t.TempDir()

	type args struct {
		filename string
	}

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "ExistFile-case-1",
			args: args{
				filename: "sample.txt",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "ExistFile-case-2",
			args: args{
				filename: "non.txt",
			},
			want:    false,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		path := filepath.Join(tmp, tt.args.filename)
		if tt.want {
			f, err := os.Create(path)
			if err != nil {
				t.Errorf("failure create file to test\n")
			}
			f.Close()
		}

		got, err := files.ExistFile(path)

		if err != nil && !tt.wantErr {
			t.Errorf("%+v failure : got -> %+v want -> %+v err -> %+v wantError -> %+v\n", tt.name, got, tt.want, err, tt.wantErr)
		}

		if got != tt.want {
			t.Errorf("%+v failure : got -> %+v want -> %+v err -> %+v wantError -> %+v\n", tt.name, got, tt.want, err, tt.wantErr)
		}

	}
}

func TestExistDir(t *testing.T) {
	tmp := t.TempDir()

	u, err := uuid.NewRandom()
	if err != nil {
		t.Errorf("%+v\n", err)
	}

	type args struct {
		path string
	}

	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "ExistDir-case-1",
			args: args{
				path: tmp,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "ExistDir-case-2",
			args: args{
				path: filepath.Join(tmp, u.String()),
			},
			want:    false,
			wantErr: false,
		},
	}

	for _, tt := range tests {

		got, err := files.ExistDir(tt.args.path)

		if err != nil && !tt.wantErr {
			t.Errorf("%+v failure : got -> %+v want -> %+v err -> %+v wantError -> %+v\n", tt.name, got, tt.want, err, tt.wantErr)
		}

		if got != tt.want {
			t.Errorf("%+v failure : got -> %+v want -> %+v err -> %+v wantError -> %+v\n", tt.name, got, tt.want, err, tt.wantErr)
		}
	}
}
