package url

import (
	"testing"
)

func Test_cleanBaseUrl(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "google.com",
			args: args{url: "google.com"},
			want: "https://google.com/",
		},
		{
			name: "https://google.com/?key=value",
			args: args{url: "https://google.com/?key=value"},
			want: "https://google.com/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CleanBase(tt.args.url); got != tt.want {
				t.Errorf("CleanBase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_cleanUrl(t *testing.T) {
	type args struct {
		url    string
		domain string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "http://domain.com/",
			args: args{
				url:    "",
				domain: "http://domain.com/",
			},
			want: "http://domain.com/",
		},
		{
			name: "https://domain.com/page",
			args: args{
				url:    "/page",
				domain: "https://domain.com/",
			},
			want: "https://domain.com/page/",
		},
		{
			name: "https://domain.com/page/",
			args: args{
				url:    "/page/",
				domain: "https://domain.com/",
			},
			want: "https://domain.com/page/",
		},
		{
			name: "https://domain.com/page/ without trailing slash",
			args: args{
				url:    "https://domain.com/page",
				domain: "https://domain.com",
			},
			want: "https://domain.com/page/",
		},
		{
			name: "https://domain.com/page/ with trailing slash",
			args: args{
				url:    "https://domain.com/page/",
				domain: "https://domain.com",
			},
			want: "https://domain.com/page/",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clean(tt.args.url, tt.args.domain); got != tt.want {
				t.Errorf("cleanUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inTargetDomain(t *testing.T) {
	type args struct {
		url    string
		domain string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "in domain",
			args: args{
				url:    "google.com/about",
				domain: "google.com",
			},
			want: true,
		},
		{
			name: "subdomain",
			args: args{
				url:    "subdomain.google.com/about",
				domain: "google.com",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InTargetDomain(tt.args.url, tt.args.domain); got != tt.want {
				t.Errorf("inTargetDomain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeHash(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "with hash",
			args: args{url: "google.com#about"},
			want: "google.com",
		},
		{
			name: "without hash string",
			args: args{url: "google.com"},
			want: "google.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeHash(tt.args.url); got != tt.want {
				t.Errorf("removeHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeQueryString(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "with query string",
			args: args{url: "google.com?key=value"},
			want: "google.com",
		},
		{
			name: "without query string",
			args: args{url: "google.com"},
			want: "google.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeQueryString(tt.args.url); got != tt.want {
				t.Errorf("removeQueryString() = %v, want %v", got, tt.want)
			}
		})
	}
}
