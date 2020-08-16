package main

import (
	"fmt"
)

type Node struct {
    tag      string
    text     string
    children []*Node
}

bf := Node{
	tag:  "bf",
	text: "Hello, World!",
}


em := Node{
	tag:      "em",
	children: []*Node{&bf},
}

ae := Node{
	tag:      "ae",
	children: []*Node{&em},
}



// main will build and walk the tree
func main() {
	fmt.Println("Building Tree")
	root := buildTree()
	fmt.Println("Walking Tree")
	root.walk()

}

