package bsp

import (
	"charm.land/log/v2"
	"codeberg.org/everesh/labe/internal/window"
)

type Node struct {
	Window *window.Window

	Left  *Node
	Right *Node

	Vertical bool
}

func New(win *window.Window, root *Node) *Node {
	node := &Node{
		Window: win,
	}

	if root == nil {
		return node
	}

	node.Vertical = !root.Vertical
	old := *root

	*root = Node{
		Window: nil,

		Left:  &old,
		Right: node,
	}

	return node
}

func (n *Node) Align(x, y, width, height int32) {
	if n.Window == nil {
		if n.Left == nil || n.Right == nil {
			log.Error("align call on bsp node failed, window and at least one of the leafs are nil", "x", x, "y", y, "width", width, "height", height)
			return
		}

		if n.Vertical {
			n.Left.Align(x, y, width, height/2)
			n.Right.Align(x, y+height/2, width, height/2)
		} else {
			n.Left.Align(x, y, width/2, height)
			n.Right.Align(x+width/2, y, width/2, height)
		}

		return
	}

	n.Window.SetPosition(x, y)
	n.Window.ProposeDimensions(width, height, false)
}
