package main

import (
	"strconv"
	"strings"
)

type Node struct {
	ID          int
	Plug        string
	LeftSocket  string
	RightSocket string

	Parent     int
	LeftChild  int
	RightChild int
}

type Rules struct {
	AllowWeak    bool
	AllowReplace bool
}

type BondType int

const (
	NoBond BondType = iota
	StrongBond
	WeakBond
	ReplaceBond
)

func parseInput(lines []string) map[int]*Node {

	tree := map[int]*Node{}

	for _, line := range lines {

		parts := strings.Split(line, ", ")

		fields := map[string]string{}
		for _, p := range parts {
			kv := strings.SplitN(p, "=", 2)
			fields[kv[0]] = kv[1]
		}

		id, _ := strconv.Atoi(fields["id"])

		tree[id] = &Node{
			ID:          id,
			Plug:        fields["plug"],
			LeftSocket:  fields["leftSocket"],
			RightSocket: fields["rightSocket"],
		}
	}

	return tree
}

func checkBond(tree map[int]*Node, thisNode *Node, socketNode *Node, side string, rules Rules) BondType {

	plug := thisNode.Plug
	socket := socketNode.LeftSocket
	child := socketNode.LeftChild

	if side == "right" {
		socket = socketNode.RightSocket
		child = socketNode.RightChild
	}

	plugSplit := strings.Split(plug, " ")
	socketSplit := strings.Split(socket, " ")

	if child == 0 {

		if plug == socket {
			return StrongBond
		}

		if rules.AllowWeak &&
			(socketSplit[0] == plugSplit[0] || socketSplit[1] == plugSplit[1]) {
			return WeakBond
		}

	} else {

		if rules.AllowReplace &&
			socket != tree[child].Plug &&
			plug == socket {
			return ReplaceBond
		}

	}

	return NoBond
}

func switchNodes(tree map[int]*Node, currNode *Node, side string, thisNode **Node) {

	var oldNodeID int

	if side == "left" {
		oldNodeID = currNode.LeftChild
		currNode.LeftChild = (*thisNode).ID
	} else {
		oldNodeID = currNode.RightChild
		currNode.RightChild = (*thisNode).ID
	}

	(*thisNode).Parent = currNode.ID
	*thisNode = tree[oldNodeID]
}

func findParent(tree map[int]*Node, currNode *Node, rules Rules, thisNode **Node) (*Node, string) {

	lCheck := checkBond(tree, *thisNode, currNode, "left", rules)

	if lCheck != NoBond {

		if lCheck != ReplaceBond {
			return currNode, "left"
		}

		switchNodes(tree, currNode, "left", thisNode)

	} else if currNode.LeftChild != 0 {

		if p, side := findParent(tree, tree[currNode.LeftChild], rules, thisNode); p != nil {
			return p, side
		}
	}

	rCheck := checkBond(tree, *thisNode, currNode, "right", rules)

	if rCheck != NoBond {

		if rCheck != ReplaceBond {
			return currNode, "right"
		}

		switchNodes(tree, currNode, "right", thisNode)

	} else if currNode.RightChild != 0 {

		if p, side := findParent(tree, tree[currNode.RightChild], rules, thisNode); p != nil {
			return p, side
		}
	}

	return nil, ""
}

func contains(arr []int, v int) bool {
	for _, x := range arr {
		if x == v {
			return true
		}
	}
	return false
}

func solveTree(puzzleInput []string, rules Rules) int {

	tree := parseInput(puzzleInput)

	root := 1
	length := len(tree)

	var thisNode *Node

	for i := root + 1; i <= length; i++ {

		thisNode = tree[i]

		var parent *Node
		var side string

		for parent == nil {
			parent, side = findParent(tree, tree[root], rules, &thisNode)
		}

		thisNode.Parent = parent.ID

		if side == "left" {
			parent.LeftChild = thisNode.ID
		} else {
			parent.RightChild = thisNode.ID
		}
	}

	seen := map[int]bool{}
	ans := []int{}
	currID := 1

	for len(ans) < length {

		node := tree[currID]

		if !seen[currID] {

			seen[currID] = true

			if node.LeftChild == 0 {

				ans = append(ans, currID)

				if node.RightChild != 0 {
					currID = node.RightChild
				} else {
					currID = node.Parent
				}

			} else {
				currID = node.LeftChild
			}

			continue
		}

		if !contains(ans, currID) {
			ans = append(ans, currID)
		}

		if node.RightChild != 0 {
			currID = node.RightChild
		} else {
			currID = node.Parent
		}

		for contains(ans, currID) {
			currID = tree[currID].Parent
		}
	}

	checksum := 0

	for i, v := range ans {
		checksum += (i + 1) * v
	}

	return checksum
}
