package sitemap

import (
	"fmt"
)

type Node struct {
	url      string
	children []*Node
	parent   *Node
}

func (node *Node) addChild(url string) *Node {
	if !node.root().contains(url) {
		newNode := Node{url: url, parent: node}
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
	if node.url == url {
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
	fmt.Println(separator, node.url)
	for _, child := range node.children {
		child.Print(separator + " ")
	}
}

func (node Node) flat() []Node {
	var nodes []Node
	nodes = append(nodes, node)
	for _, child := range node.children {
		nodes = append(nodes, child.flat()...)
	}
	return nodes
}
