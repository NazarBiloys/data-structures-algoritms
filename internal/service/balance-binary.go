package service

import (
	"fmt"
	"os"
	"unsafe"
)

type Node struct {
	value  int
	height int
	left   *Node
	right  *Node
}

type AVLTree struct {
	root *Node
}

func (t *AVLTree) Insert(value int) {
	t.root = t.insert(t.root, value)
}

func (t *AVLTree) insert(node *Node, value int) *Node {
	if node == nil {
		return &Node{value, 1, nil, nil}
	}

	if value < node.value {
		node.left = t.insert(node.left, value)
	} else if value > node.value {
		node.right = t.insert(node.right, value)
	} else {
		return node
	}

	node.height = max(height(node.left), height(node.right)) + 1
	balance := getBalance(node)

	if balance > 1 && value < node.left.value {
		return t.rotateRight(node)
	}

	if balance < -1 && value > node.right.value {
		return t.rotateLeft(node)
	}

	if balance > 1 && value > node.left.value {
		node.left = t.rotateLeft(node.left)
		return t.rotateRight(node)
	}

	if balance < -1 && value < node.right.value {
		node.right = t.rotateRight(node.right)
		return t.rotateLeft(node)
	}

	return node
}

func (t *AVLTree) Delete(value int) {
	t.root = t.delete(t.root, value)
}

func (t *AVLTree) delete(node *Node, value int) *Node {
	if node == nil {
		return node
	}

	if value < node.value {
		node.left = t.delete(node.left, value)
	} else if value > node.value {
		node.right = t.delete(node.right, value)
	} else {
		if node.left == nil || node.right == nil {
			var temp *Node

			if node.left == nil {
				temp = node.right
			} else {
				temp = node.left
			}

			if temp == nil {
				temp = node
				node = nil
			} else {
				node = temp
			}
		} else {
			temp := minValueNode(node.right)
			node.value = temp.value
			node.right = t.delete(node.right, temp.value)
		}
	}

	if node == nil {
		return node
	}

	node.height = max(height(node.left), height(node.right)) + 1
	balance := getBalance(node)

	if balance > 1 && getBalance(node.left) >= 0 {
		return t.rotateRight(node)
	}

	if balance > 1 && getBalance(node.left) < 0 {
		node.left = t.rotateLeft(node.left)
		return t.rotateRight(node)
	}

	if balance < -1 && getBalance(node.right) <= 0 {
		return t.rotateLeft(node)
	}

	if balance < -1 && getBalance(node.right) > 0 {
		node.right = t.rotateRight(node.right)
		return t.rotateLeft(node)
	}

	return node
}
func (t *AVLTree) Search(value int) bool {
	node := t.root

	for node != nil {
		if value < node.value {
			node = node.left
		} else if value > node.value {
			node = node.right
		} else {
			return true
		}
	}

	return false
}

func (t *AVLTree) estimateAVLTreeMemoryUsage() uintptr {
	nodeSize := unsafe.Sizeof(Node{})
	totalNodes := CountNodes(t.root)
	totalMemoryUsage := nodeSize * uintptr(totalNodes)

	return totalMemoryUsage
}

func CountNodes(node *Node) int {
	if node == nil {
		return 0
	}

	return 1 + CountNodes(node.left) + CountNodes(node.right)
}

func (t *AVLTree) InorderTraversal(node *Node) {
	if node != nil {
		t.InorderTraversal(node.left)
		fmt.Printf("%d ", node.value)
		t.InorderTraversal(node.right)
	}
}

func (t *AVLTree) rotateLeft(node *Node) *Node {
	rightNode := node.right
	node.right = rightNode.left
	rightNode.left = node

	node.height = max(height(node.left), height(node.right)) + 1
	rightNode.height = max(height(rightNode.left), height(rightNode.right)) + 1

	return rightNode
}

func (t *AVLTree) rotateRight(node *Node) *Node {
	leftNode := node.left
	node.left = leftNode.right
	leftNode.right = node

	node.height = max(height(node.left), height(node.right)) + 1
	leftNode.height = max(height(leftNode.left), height(leftNode.right)) + 1

	return leftNode
}

func minValueNode(node *Node) *Node {
	current := node

	for current.left != nil {
		current = current.left
	}

	return current
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func height(node *Node) int {
	if node == nil {
		return 0
	}

	return node.height
}

func getBalance(node *Node) int {
	if node == nil {
		return 0
	}

	return height(node.left) - height(node.right)
}

func (node *Node) Output(file *os.File) {
	if node == nil {
		return
	}
	node.left.Output(file)
	WriteIntoFile(fmt.Sprintf("%d ", node.value), file)
	node.right.Output(file)
}

func (tree *AVLTree) Output(file *os.File) {
	if tree.root == nil {
		return
	}
	tree.root.Output(file)
}
