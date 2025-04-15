package main

import (
	"fmt"
)

// Red-Black Tree Node
type Node struct {
	data   int
	color  bool // true for Red, false for Black
	left   *Node
	right  *Node
	parent *Node
}

// Red-Black Tree structure
type RedBlackTree struct {
	root *Node
}

// Constants for Red and Black colors
const (
	RED   = true
	BLACK = false
)

// Rotate Left
func (t *RedBlackTree) rotateLeft(x *Node) {
	y := x.right            // Step 1: `y` को `x` के Right Child से Point करो
	x.right = y.left        // Step 2: `y` का Left Subtree, `x` का Right बन जाता है
	if y.left != nil {  // if exists currently 
		y.left.parent = x   // Step 3: अगर `y.left` exist करता है, तो उसका Parent `x` होगा
	}
	y.parent = x.parent     // Step 4: `y` का नया Parent `x` का पुराना Parent होगा
	if x.parent == nil {
		t.root = y          // Step 5: अगर `x` Root था, तो अब `y` नया Root बनेगा
	} else if x == x.parent.left {
		x.parent.left = y   // Step 6: अगर `x`, Left Child था, तो `y` भी Left Child बनेगा
	} else {
		x.parent.right = y  // Step 7: अगर `x`, Right Child था, तो `y` भी Right Child बनेगा
	}
	y.left = x             // Step 8: `x`, `y` का Left Child बनेगा
	x.parent = y           // Step 9: `x` का नया Parent `y` होगा
}


// Rotate Right
func (t *RedBlackTree) rotateRight(y *Node) {
	x := y.left
	y.left = x.right
	if x.right != nil {
		x.right.parent = y
	}
	x.parent = y.parent
	if y.parent == nil {
		t.root = x
	} else if y == y.parent.right {
		y.parent.right = x
	} else {
		y.parent.left = x
	}
	x.right = y
	y.parent = x
}

// Insert a node
func (t *RedBlackTree) insert(data int) {
	newNode := &Node{data: data, color: RED}
	t.root = bstInsert(t.root, newNode)
	t.fixInsert(newNode)
}

// BST Insert (Helper function)
func bstInsert(root, newNode *Node) *Node {
	if root == nil {
		return newNode
	}
	if newNode.data < root.data {
		root.left = bstInsert(root.left, newNode)
		root.left.parent = root
	} else {
		root.right = bstInsert(root.right, newNode)
		root.right.parent = root
	}
	return root
}





// create secnd helper functions



// Fix Red-Black Tree after Insert
func (t *RedBlackTree) fixInsert(node *Node) {
	for node != t.root && node.parent.color == RED {
		if node.parent == node.parent.parent.left {
			uncle := node.parent.parent.right
			if uncle != nil && uncle.color == RED { // Case 1: Uncle is Red
				node.parent.color = BLACK
				uncle.color = BLACK
				node.parent.parent.color = RED
				node = node.parent.parent
			} else {
				if node == node.parent.right { // Case 2: Node is Right Child
					node = node.parent
					t.rotateLeft(node)
				}
				node.parent.color = BLACK // Case 3: Uncle is Black
				node.parent.parent.color = RED
				t.rotateRight(node.parent.parent)
			}
		} else { // Symmetric Cases for Right Subtree
			uncle := node.parent.parent.left
			if uncle != nil && uncle.color == RED {
				node.parent.color = BLACK
				uncle.color = BLACK
				node.parent.parent.color = RED
				node = node.parent.parent
			} else {
				if node == node.parent.left {
					node = node.parent
					t.rotateRight(node)
				}
				node.parent.color = BLACK
				node.parent.parent.color = RED
				t.rotateLeft(node.parent.parent)
			}
		}
	}
	t.root.color = BLACK
}

// Inorder Traversal (For Testing)
func (t *RedBlackTree) inorderTraversal(node *Node) {
	if node != nil {
		t.inorderTraversal(node.left)
		fmt.Printf("%d (%s) ", node.data, colorToString(node.color))
		t.inorderTraversal(node.right)
	}
}

// Helper function to convert color to string
func colorToString(color bool) string {
	if color == RED {
		return "RED"
	}
	return "BLACK"
}


// Main function
func main() {
	rbt := &RedBlackTree{}

	values := []int{20, 15, 25, 10, 18, 30, 5, 12, 17, 19}
	for _, v := range values {
		rbt.insert(v)
	}

	fmt.Println("Inorder Traversal of Red-Black Tree:")
	rbt.inorderTraversal(rbt.root)
	fmt.Println()

	fmt.Println("Last the botted")
}
