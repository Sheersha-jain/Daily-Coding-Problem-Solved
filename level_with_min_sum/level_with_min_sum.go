//Given a binary tree, return the level of the tree with minimum sum

package main

import (
	"fmt"
	"math"
)

type node struct {
	data int
	left *node
	right *node
}

const (
	intmax = math.MaxInt64
)

func levelwithMinSum(level int, node *node, hash map[int][]int) {
	if (node != nil ) {
		_, ok := hash[level]
		if ok {
			hash[level] = append(hash[level], node.data)
		} else {
			hash[level] = make([]int, 0)
			hash[level] = append(hash[level], node.data)
		}
		levelwithMinSum(level+1, node.left, hash)
		levelwithMinSum(level+1, node.right, hash)
	}
}

func constructTree() *node {
	nodeone := &node{data: -20, left:nil, right: nil}
	nodetwo := &node{data: 21, left:nil, right:nil}
	nodethree :=  &node{data: 23, left:nodeone, right:nodetwo}

	nodefour := &node{data: 4 , left:nil, right: nil}
	nodefive := &node{data: -264, left:nil, right:nil}
	nodesix :=  &node{data: 23, left:nodefour, right:nodefive}

	rootnode := &node{data: -121, left: nodethree, right: nodesix}
	return rootnode
}

func getMinSum(hash map[int][]int) (int,int) {
	minsum := intmax
	level := 0
	for l, value := range hash {
		sum := 0
		for _, v := range value {
			sum += v
		}
		if sum < minsum {
			minsum = sum
			level = l
		}
	}

	return level, minsum
}

func main() {
	hash := map[int][]int {}
	root := constructTree()
	levelwithMinSum(0, root, hash)
	level, minsum := getMinSum(hash)
	//Assuming the root node is at level 0
	if minsum != intmax {
		fmt.Printf("Level with minimum sum is %d and the sum is %d\n", level, minsum)
	} else {
		fmt.Println("Tree does not exist")
	}

}
