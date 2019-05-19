package sitemap

import (
	"reflect"
	"testing"
)

func TestNode_contains(t *testing.T) {
	type fields struct {
		url      string
		children []*Node
		parent   *Node
	}
	type args struct {
		url string
	}
	want1 := &Node{"Hello", []*Node{{url: "World"}}, nil}
	want1.children[0].parent = want1
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "Single child",
			fields: fields{want1.url, want1.children, want1.parent},
			args:   args{want1.children[0].url},
			want:   true,
		},
		{
			name:   "Single child doesn't not contain",
			fields: fields{want1.url, want1.children, want1.parent},
			args:   args{"Unknown string"},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node{
				url:      tt.fields.url,
				children: tt.fields.children,
				parent:   tt.fields.parent,
			}
			if got := node.contains(tt.args.url); got != tt.want {
				t.Errorf("Node.contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNode_addChild(t *testing.T) {
	type fields struct {
		url      string
		children []*Node
		parent   *Node
	}
	type args struct {
		url string
	}
	want1 := &Node{"Hello", []*Node{{url: "World"}}, nil}
	want1.children[0].parent = want1
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Node
	}{
		{
			name:   "Single child add",
			fields: fields{url: "Hello"},
			args:   args{"World"},
			want:   want1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			node := &Node{
				url:      tt.fields.url,
				children: tt.fields.children,
				parent:   tt.fields.parent,
			}
			if got := node.addChild(tt.args.url); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Node.addChild() = %v, want %v", got, tt.want)
			}
		})
	}
}
