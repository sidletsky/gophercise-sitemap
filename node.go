package sitemap

import (
	"fmt"
)

type Node struct {
	Url      string
	children []*Node
	parent   *Node
}

func (node *Node) addChild(url string) *Node {
	if !node.root().contains(url) {
		newNode := Node{Url: url, parent: node}
		node.children = append(node.children, &newNode)
	}
	return node
}

func (node *Node) root() *Node {
	parent := node
	for {
		if parent.parent == nil {
			return parent
		}
		parent = parent.parent
	}
}

func (node *Node) contains(url string) bool {
	if node.Url == url {
		return true
	}
	for _, child := range node.children {
		if child.contains(url) {
			return true
		}
	}
	return false
}

func (node *Node) Print(separator string) {
	fmt.Println(separator, node.Url)
	for _, child := range node.children {
		child.Print(separator + " ")
	}
}

func (node Node) Flat() []Node {
	var nodes []Node
	nodes = append(nodes, node)
	for _, child := range node.children {
		nodes = append(nodes, child.Flat()...)
	}
	return nodes
}
