package input_autocomplete

import (
	"errors"
	"reflect"
	"runtime"
	"testing"
)

type DirListCheckerCustomMock struct {
	listContentMock func(path string) ([]string, error)
	isDirMock       func(path string) (bool, error)
}

func (d DirListCheckerCustomMock) ListContent(path string) ([]string, error) {
	return d.listContentMock(path)
}

func (d DirListCheckerCustomMock) IsDir(path string) (bool, error) {
	return d.isDirMock(path)
}

func Test_autocomplete_unixAutocomplete(t *testing.T) {
	if runtime.GOOS != "linux" && runtime.GOOS != "darwin" {
		t.Skipf("Skip test because OS is %v", runtime.GOOS)
	}
	type fields struct {
		cmd DirListChecker
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		{
			name: "success to find some dir or file to autocomplete",
			fields: fields{
				cmd: DirListCheckerCustomMock{
					listContentMock: func(path string) ([]string, error) {
						return []string{".", "..", "home", "binary", "etc"}, nil
					},
					isDirMock: func(path string) (bool, error) {
						return false, nil
					},
				},
			},
			args: args{
				path: "ho",
			},
			want: []string{"./home"},
		},
		{
			name: "success to find some dir to autocomplete with absolute path",
			fields: fields{
				cmd: DirListCheckerCustomMock{
					listContentMock: func(path string) ([]string, error) {
						return []string{".", "..", "home", "binary", "etc"}, nil
					},
					isDirMock: func(path string) (bool, error) {
						return true, nil
					},
				},
			},
			args: args{
				path: "/ho",
			},
			want: []string{"/home/"},
		},
		{
			name: "failed to find some dir or file to autocomplete",
			fields: fields{
				cmd: DirListCheckerCustomMock{
					listContentMock: func(path string) ([]string, error) {
						return []string{".", "..", "home", "binary", "etc"}, nil
					},
					isDirMock: func(path string) (bool, error) {
						return false, nil
					},
				},
			},
			args: args{
				path: "auto",
			},
			want: []string{"./auto"},
		},
		{
			name: "failed to find some dir or file to autocomplete",
			fields: fields{
				cmd: DirListCheckerCustomMock{
					listContentMock: func(path string) ([]string, error) {
						return []string{".", "..", "home", "binary", "etc"}, nil
					},
					isDirMock: func(path string) (bool, error) {
						return false, nil
					},
				},
			},
			args: args{
				path: "/aut",
			},
			want: []string{"/aut"},
		},
		{
			name: "success to find some dir or file to autocomplete",
			fields: fields{
				cmd: DirListCheckerCustomMock{
					listContentMock: func(path string) ([]string, error) {
						return []string{".", "..", "home", "binary", "etc"}, nil
					},
					isDirMock: func(path string) (bool, error) {
						return false, nil
					},
				},
			},
			args: args{
				path: "/bi",
			},
			want: []string{"/binary"},
		},
		{
			name: "success to find multiple dirs or files to autocomplete",
			fields: fields{
				cmd: DirListCheckerCustomMock{
					listContentMock: func(path string) ([]string, error) {
						return []string{".", "..", "home", "bin", "binary", "etc"}, nil
					},
					isDirMock: func(path string) (bool, error) {
						return false, nil
					},
				},
			},
			args: args{
				path: "/bi",
			},
			want: []string{"/bin", "/binary"},
		},
		{
			name: "success with empty path",
			fields: fields{
				cmd: DirListCheckerCustomMock{
					listContentMock: func(path string) ([]string, error) {
						return []string{".", "..", "home", "binary", "etc"}, nil
					},
					isDirMock: func(path string) (bool, error) {
						return true, nil
					},
				},
			},
			args: args{
				path: "",
			},
			want: []string{""},
		},
		{
			name: "success with already completed path",
			fields: fields{
				cmd: DirListCheckerCustomMock{
					listContentMock: func(path string) ([]string, error) {
						return []string{".", "..", "home", "binary", "etc", "file.txt"}, nil
					},
					isDirMock: func(path string) (bool, error) {
						return false, nil
					},
				},
			},
			args: args{
				path: "./file.txt",
			},
			want: []string{"./file.txt"},
		},
		{
			name: "failed to list content",
			fields: fields{
				cmd: DirListCheckerCustomMock{
					listContentMock: func(path string) ([]string, error) {
						return nil, errors.New("some error")
					},
					isDirMock: func(path string) (bool, error) {
						return false, nil
					},
				},
			},
			args: args{
				path: "/bi",
			},
			want: []string{"/bi"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := autocomplete{
				cmd: tt.fields.cmd,
			}
			if got := a.unixAutocomplete(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unixAutocomplete() = %v, want %v", got, tt.want)
			}
		})
	}
}
