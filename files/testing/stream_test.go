package testing_test

import (
	"testing"

	tst "github.com/aqyuki/goutil/files/testing"
)

func TestFileStreamChecker_Write(t *testing.T) {
	type args struct {
		p []byte
	}
	tests := []struct {
		name    string
		s       *tst.FileStreamChecker
		args    args
		wantN   int
		wantErr bool
	}{
		{
			name: "Write normal",
			s: &tst.FileStreamChecker{
				StringContent: "",
			},
			args: args{
				p: []byte("sample"),
			},
			wantN:   0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotN, err := tt.s.Write(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileStreamChecker.Write() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("FileStreamChecker.Write() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}

func TestFileStreamChecker_EqualResult(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name    string
		s       *tst.FileStreamChecker
		content string
		args    args
		want    bool
	}{
		{
			name:    "Equal string",
			s:       &tst.FileStreamChecker{},
			content: "sample",
			args: args{
				content: "sample",
			},
			want: true,
		},
		{
			name:    "Equal failure",
			s:       &tst.FileStreamChecker{},
			content: "aaa",
			args: args{
				content: "string",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Write([]byte(tt.content))
			if got := tt.s.EqualResult(tt.args.content); got != tt.want {
				t.Errorf("FileStreamChecker.EqualResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
