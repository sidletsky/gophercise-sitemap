package internal

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

func Test_parse(t *testing.T) {
	type args struct {
		r io.Reader
	}
	var arguments [4]args
	for i := 0; i < 4; i++ {
		path := fmt.Sprintf("client_test/ex%d.html", i+1)
		file, _ := os.Open(path)
		arg := args{r: file}
		arguments[i] = arg
	}
	tests := []struct {
		name    string
		args    args
		want    []Url
		wantErr bool
	}{
		{
			name:    "client_test/ex1.html",
			args:    arguments[0],
			want:    []Url{{Loc: "/other-page"}},
			wantErr: false,
		},
		{
			name: "client_test/ex2.html",
			args: arguments[1],
			want: []Url{
				{Loc: "https://www.twitter.com/joncalhoun"},
				{Loc: "https://github.com/gophercises"}},
			wantErr: false,
		},
		{
			name: "client_test/ex3.html",
			args: arguments[2],
			want: []Url{
				{Loc: "#"},
				{Loc: "/lost"},
				{Loc: "https://twitter.com/marcusolsson"}},
			wantErr: false,
		},
		{
			name:    "client_test/ex4.html",
			args:    arguments[3],
			want:    []Url{{Loc: "/dog-cat"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
