package algorithm

import (
	"math"
)

// ALVNode alv tree node
type ALVNode struct {
	Data   int32
	LNode  *ALVNode
	RNode  *ALVNode
	Height int16
}

// PreorderTraversal preorder traver
func PreorderTraversal(root *ALVNode) []int32 {
	var list []int32
	if root.LNode != nil {
		list = append(list, PreorderTraversal(root.LNode)...)
	}
	list = append(list, root.Data)
	if root.RNode != nil {
		list = append(list, PreorderTraversal(root.RNode)...)
	}
	return list
}

// Insert insert node in alv tree
func Insert(data int32, root *ALVNode) *ALVNode {
	if root == nil {
		return &ALVNode{
			Data:   data,
			LNode:  nil,
			RNode:  nil,
			Height: 0,
		}
	}

	if data < root.Data {
		root.LNode = Insert(data, root.LNode)
		root.setHeight()
		rootBF := root.getNodeBF()
		lNodeBF := root.LNode.getNodeBF()
		if math.Abs(float64(rootBF)) > 1 {
			sumBF := rootBF + lNodeBF
			if sumBF > 2 || sumBF < -2 {
				if rootBF > 0 {
					root = rightRotate(root)
				} else {
					root = leftRotate(root)
				}
			} else {
				if rootBF > 0 {
					root.LNode = leftRotate(root.LNode)
					root = rightRotate(root)
				} else {
					root.LNode = rightRotate(root.LNode)
					root = leftRotate(root)
				}
			}
		}
	}
	if data > root.Data {
		root.RNode = Insert(data, root.RNode)
		rootBF := root.getNodeBF()
		rNodeBF := root.RNode.getNodeBF()
		if math.Abs(float64(rootBF)) > 1 {
			sumBF := rootBF + rNodeBF
			if sumBF > 2 || sumBF < -2 {
				if rootBF > 0 {
					root = rightRotate(root)
				} else {
					root = leftRotate(root)
				}
			} else {
				if rootBF > 0 {
					root.RNode = leftRotate(root.RNode)
					root = rightRotate(root)
				} else {
					root.RNode = rightRotate(root.RNode)
					root = leftRotate(root)
				}
			}
		}
	}

	root.setHeight()
	return root
}

func (n *ALVNode) getNodeBF() int16 {
	var l, r int16
	if n.LNode == nil {
		l = -1
	} else {
		l = n.LNode.Height
	}
	if n.RNode == nil {
		r = -1
	} else {
		r = n.RNode.Height
	}
	return l - r
}

func rightRotate(node *ALVNode) *ALVNode {
	var newNode *ALVNode

	newNode = node.LNode
	node.LNode = newNode.RNode
	newNode.RNode = node

	node.setHeight()
	newNode.setHeight()

	return newNode
}

func leftRotate(node *ALVNode) *ALVNode {
	var newNode *ALVNode

	newNode = node.RNode
	node.RNode = newNode.LNode
	newNode.LNode = node

	node.setHeight()
	newNode.setHeight()

	return newNode
}

func (n *ALVNode) setHeight() {
	if n.LNode == nil && n.RNode == nil {
		n.Height = 0
	} else if n.LNode == nil {
		n.Height = n.RNode.Height + 1
	} else if n.RNode == nil {
		n.Height = n.LNode.Height + 1
	} else {
		n.Height = int16(math.Max(float64(n.LNode.Height), float64(n.RNode.Height))) + 1
	}
}
