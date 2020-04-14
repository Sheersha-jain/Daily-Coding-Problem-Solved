// Merge two trees by adding nodes at same postion

package main

import "fmt"

type node struct {
	data int
	left *node
	right *node
}


func mergeTree(newTreenode *node, treenode *node) {
	newTreenode.data += treenode.data
	if treenode.left != nil {
		if newTreenode.left != nil {
			mergeTree(newTreenode.left, treenode.left)
		} else {
			newNode := &node{data: 0, left: nil, right: nil}
			newTreenode.left = newNode
			mergeTree(newNode, treenode.left)
		}
	}

	if treenode.right != nil {
		if newTreenode.right != nil {
			mergeTree(newTreenode.right, treenode.right)
		} else {
			newNode := &node{data: 0, left: nil, right: nil}
			newTreenode.right = newNode
			mergeTree(newNode, treenode.right)
		}
	}
}

func constructTree() *node {
	nodeone := &node{data: -20, left:nil, right: nil}
	nodetwo := &node{data: 21, left:nil, right:nil}
	nodethree :=  &node{data: 23, left:nodeone, right:nodetwo}
	nodeseven := &node{data: 8, left: nil, right:nil}
	nodefour := &node{data: 4 , left:nodeseven, right: nil}
	nodefive := &node{data: -264, left:nil, right:nil}
	nodesix :=  &node{data: 23, left:nodefour, right:nodefive}
	rootnode := &node{data: -121, left: nodethree, right: nodesix}
	return rootnode
}

func displayTree(root *node) {
	fmt.Println(root.data)
	fmt.Println(root.left.data)
	fmt.Println(root.right.data)
	fmt.Println(root.left.left.data)
	fmt.Println(root.left.right.data)
	fmt.Println(root.right.left.data)
	fmt.Println(root.right.right.data)
	fmt.Println(root.right.left.left.data)
}

func main() {
	rootfirsttree := constructTree()
	rootsecondtree := constructTree()
	rootnewtree := &node{data: 0, left: nil, right: nil}
	mergeTree(rootnewtree, rootfirsttree)
	mergeTree(rootnewtree, rootsecondtree)
	displayTree(rootnewtree)
}