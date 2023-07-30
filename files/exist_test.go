package files_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/aqyuki/goutil/files"
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
			name: "Case 1",
			args: args{
				filename: "sample.txt",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Case 2",
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
			t.Errorf("test failure : got -> %+v want -> %+v err -> %+v wantError -> %+v\n", got, tt.want, err, tt.wantErr)
		}

		if got != tt.want {
			t.Errorf("test failure : got -> %+v want -> %+v err -> %+v wantError -> %+v\n", got, tt.want, err, tt.wantErr)
		}

		t.Logf("pass %+v\n", tt.name)
	}
}
