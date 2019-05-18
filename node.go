package gophercise_sitemap

type Node struct {
	url      string
	children []*Node
}

func (node *Node) addChild(url string) {
	if node.url == url {
		return
	}
	for _, child := range node.children {
		if child.url == url {
			return
		}
	}
	newNode := Node{url: url}
	node.children = append(node.children, &newNode)
}
