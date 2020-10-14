package input_autocomplete

import (
	"errors"
	"testing"
)

type DirListerCustomMock struct {
	listContentMock func(path string) ([]string, error)
}

func (d DirListerCustomMock) ListContent(path string) ([]string, error) {
	return d.listContentMock(path)
}

func Test_autocomplete_unixAutocomplete(t *testing.T) {
	type fields struct {
		cmd DirLister
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "success to find some dir or file to autocomplete",
			fields: fields{
				cmd: DirListerCustomMock{
					listContentMock: func(path string) ([]string, error) {
						return []string{".", "..", "home", "binary", "etc"}, nil
					},
				},
			},
			args:   args{
				path: "ho",
			},
			want:   "./home",
		},
		{
			name:   "failed to find some dir or file to autocomplete",
			fields: fields{
				cmd: DirListerCustomMock{
					listContentMock: func(path string) ([]string, error) {
						return []string{".", "..", "home", "binary", "etc"}, nil
					},
				},
			},
			args:   args{
				path: "auto",
			},
			want:   "./auto",
		},
		{
			name:   "failed to find some dir or file to autocomplete",
			fields: fields{
				cmd: DirListerCustomMock{
					listContentMock: func(path string) ([]string, error) {
						return []string{".", "..", "home", "binary", "etc"}, nil
					},
				},
			},
			args:   args{
				path: "/aut",
			},
			want:   "/aut",
		},
		{
			name:   "success to find some dir or file to autocomplete",
			fields: fields{
				cmd: DirListerCustomMock{
					listContentMock: func(path string) ([]string, error) {
						return []string{".", "..", "home", "binary", "etc"}, nil
					},
				},
			},
			args:   args{
				path: "/bi",
			},
			want:   "/binary",
		},
		{
			name:   "success with empty path",
			fields: fields{
				cmd: DirListerCustomMock{
					listContentMock: func(path string) ([]string, error) {
						return []string{".", "..", "home", "binary", "etc"}, nil
					},
				},
			},
			args:   args{
				path: "",
			},
			want:   "",
		},
		{
			name:   "success with already completed path",
			fields: fields{
				cmd: DirListerCustomMock{
					listContentMock: func(path string) ([]string, error) {
						return []string{".", "..", "home", "binary", "etc", "file.txt "}, nil
					},
				},
			},
			args:   args{
				path: "./file.txt ",
			},
			want:   "./file.txt ",
		},
		{
			name:   "failed to list content",
			fields: fields{
				cmd: DirListerCustomMock{
					listContentMock: func(path string) ([]string, error) {
						return nil, errors.New("some error")
					},
				},
			},
			args:   args{
				path: "/bi",
			},
			want:   "/bi",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := autocomplete{
				cmd: tt.fields.cmd,
			}
			if got := a.unixAutocomplete(tt.args.path); got != tt.want {
				t.Errorf("unixAutocomplete() = %v, want %v", got, tt.want)
			}
		})
	}
}