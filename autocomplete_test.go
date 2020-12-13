package input_autocomplete

import (
	"errors"
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
		want   string
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
			want: "./home",
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
			want: "/home/",
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
			want: "./auto",
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
			want: "/aut",
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
			want: "/binary",
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
			want: "",
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
			want: "./file.txt",
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
			want: "/bi",
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

func Test_cleanFilePathUnix(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "/ dir file clean",
			args: args{
				path: "/bi",
			},
			want: "/bi",
		},
		{
			name: "nested dir file clean",
			args: args{
				path: "/bi/as/d",
			},
			want: "/bi/as/d",
		},
		{
			name: "simple file clean",
			args: args{
				path: "f",
			},
			want: "./f",
		},
		{
			name: "simple file clean2",
			args: args{
				path: "/f",
			},
			want: "/f",
		},
		{
			name: "empty file clean",
			args: args{
				path: "",
			},
			want: "./",
		},
		{
			name: "/ dir clean",
			args: args{
				path: "/",
			},
			want: "/",
		},
		{
			name: "./ dir clean",
			args: args{
				path: "./",
			},
			want: "./",
		},
		{
			name: "./ dir file clean",
			args: args{
				path: "./bi",
			},
			want: "./bi",
		},
		{
			name: "nested dir file clean",
			args: args{
				path: "./bi/as/d",
			},
			want: "./bi/as/d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanFilePathUnix(tt.args.path); got != tt.want {
				t.Errorf("cleanFilePathUnix() = %v, want %v", got, tt.want)
			}
		})
	}
}
