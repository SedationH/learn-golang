package main

type Node struct {
	Value       int
	Left, Right *Node
}

func (node *Node) Print() {
	print(node.Value, " ")
}

func (node *Node) PreOrderTraverseWithFn(fn func(*Node)) {
	if node == nil {
		return
	}
	fn(node)
	node.Left.PreOrderTraverseWithFn(fn)
	node.Right.PreOrderTraverseWithFn(fn)
}

func (node *Node) PreOrderTraverseWithChannel() chan *Node {
	out := make(chan *Node)

	go func() {
		node.PreOrderTraverseWithFn(func(n *Node) {
			out <- n
		})
		close(out)
	}()

	return out
}

func main() {
	root := Node{
		Value: 3,
		Left: &Node{
			Value: 4,
			Right: &Node{
				Value: 5,
			},
		},
		Right: &Node{
			Value: 1,
			Left: &Node{
				Value: 2,
			},
		},
	}
	root.PreOrderTraverseWithFn(func(n *Node) {
		n.Print()
	})

	out := root.PreOrderTraverseWithChannel()
	for node := range out {
		node.Print()
	}
}
