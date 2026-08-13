package bsp

import (
	"charm.land/log/v2"
	"codeberg.org/everesh/labe/internal/window"
)

type Node struct {
	Window *window.Window

	Left  *Node
	Right *Node
}

func New(win *window.Window, root *Node) *Node {
	node := &Node{
		Window: win,
	}

	if root == nil {
		return node
	}

	old := *root

	*root = Node{
		Window: nil,

		Left:  &old,
		Right: node,
	}

	return node
}

// Find returns the node containing win, or nil if it isn't in the tree.
func (n *Node) Find(win *window.Window) *Node {
	if n == nil {
		return nil
	}

	if n.Window == win {
		return n
	}

	if found := n.Left.Find(win); found != nil {
		return found
	}

	return n.Right.Find(win)
}

func (n *Node) Remove(win *window.Window) bool {
	if n.Window == win {
		return true
	}

	if n.Left == nil || n.Right == nil {
		return false
	}

	if n.Left.Remove(win) {
		*n = *n.Right
		return false
	}

	if n.Right.Remove(win) {
		*n = *n.Left
		return false
	}

	return false
}

func (n *Node) Align(x, y, width, height int32) {
	if n.Window == nil {
		if n.Left == nil || n.Right == nil {
			log.Error("align call on bsp node failed, window and at least one of the leafs are nil", "x", x, "y", y, "width", width, "height", height)
			return
		}

		if width < height {
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
