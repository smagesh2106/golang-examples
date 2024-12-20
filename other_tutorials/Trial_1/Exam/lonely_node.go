package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res = []int{}

func getLonelyNodes(root *TreeNode) []int {
	// Write your code here
	if root == nil {
		return res
	} else if (root.Left != nil) && (root.Right != nil) {
		getLonelyNodes(root.Left)
		getLonelyNodes(root.Right)
	} else if root.Left == nil && root.Right != nil {
		res = append(res, root.Val)
		getLonelyNodes(root.Right)
	} else { //} ( root.Right == nil && root.Left != nil){
		res = append(res, root.Val)
		getLonelyNodes(root.Left)
	}

	return res
}

/**
R E A D M E
DO NOT CHANGE the code below, we use it to grade your submission. If changed your submission will be failed automatically.
**/

func printResults(results []int) {
	reslen := len(results)
	fmt.Print("[")
	for index, element := range results {

		fmt.Printf("%v", element)

		if reslen != index+1 {
			fmt.Print(", ")
		}
	}
	fmt.Print("]")
}

func getNode(nodes []*TreeNode) (*TreeNode, []*TreeNode) {
	if len(nodes) > 0 {
		return nodes[len(nodes)-1], nodes[:len(nodes)-1]
	}
	return nil, nil
}

func BuildTree(nodes []*TreeNode, root *TreeNode) *TreeNode {

	var new_nodes []*TreeNode

	for i := len(nodes) - 1; i >= 0; i-- {
		new_nodes = append(new_nodes, nodes[i])
	}

	root = new_nodes[len(new_nodes)-1]
	new_nodes = new_nodes[:len(new_nodes)-1]

	for _, node := range nodes {

		var n *TreeNode
		if node != nil {
			n, new_nodes = getNode(new_nodes)
			if n != nil {
				node.Left = n
			}
			n, new_nodes = getNode(new_nodes)
			if n != nil {
				node.Right = n
			}
		}
	}
	return root

}

func main() {

	reader := bufio.NewReader(os.Stdin)
	rawArray, _ := reader.ReadString('\n')
	rawArray = strings.Replace(rawArray, "\n", " ", -1)
	rawArray = strings.Replace(rawArray, ",", " ", -1)
	rawArray = strings.Replace(rawArray, "\r", " ", -1)
	elements := strings.Split(rawArray, " ")
	var input []*TreeNode

	for index := range elements {
		if elements[index] == "null" {
			input = append(input, nil)
		} else {
			num, err := strconv.Atoi(elements[index])
			if err == nil {
				input = append(input, &TreeNode{num, nil, nil})
			}
		}
	}
	var root *TreeNode
	root = BuildTree(input, root)
	printResults(getLonelyNodes(root))

}
