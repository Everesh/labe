package bsp

import (
	"fmt"

	"codeberg.org/everesh/labe/internal/layout"
	"codeberg.org/everesh/labe/internal/window"
)

var _ layout.Layout = (*Tree)(nil)

type Tree struct {
	Root *Node
}

func New() *Tree {
	return &Tree{}
}

func (t *Tree) Add(w *window.Window) error {
	if t.Root.Find(w) != nil {
		return fmt.Errorf("bsp: window is already part of the tree")
	}

	if t.Root == nil {
		t.Root = newNode(w, nil)
		return nil
	}

	target := t.Root
	if st := w.SplitTarget; st != nil {
		if found := t.Root.Find(st); found != nil {
			target = found
		}
	}
	w.SplitTarget = nil

	newNode(w, target)
	w.HintTiled(0b1111)
	return nil
}

func (t *Tree) Remove(w *window.Window) error {
	if t.Root == nil {
		return nil
	}

	if t.Root.Remove(w) {
		t.Root = nil
	}

	w.HintTiled(0b0000)
	return nil
}

func (t *Tree) Align(x, y, width, height int32) error {
	if t.Root == nil {
		return nil
	}

	return t.Root.Align(x, y, width, height)
}

type Node struct {
	Window *window.Window

	Left  *Node
	Right *Node
}

func newNode(win *window.Window, root *Node) *Node {
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

func (n *Node) Align(x, y, width, height int32) error {
	if n.Window == nil {
		if n.Left == nil || n.Right == nil {
			return fmt.Errorf("bsp: align called on a node with a nil window and at least one nil leaf (x=%d y=%d width=%d height=%d)", x, y, width, height)
		}

		if width < height {
			if err := n.Left.Align(x, y, width, height/2); err != nil {
				return err
			}
			return n.Right.Align(x, y+height/2, width, height/2)
		}

		if err := n.Left.Align(x, y, width/2, height); err != nil {
			return err
		}
		return n.Right.Align(x+width/2, y, width/2, height)
	}

	n.Window.SetPosition(x, y)
	n.Window.ProposeDimensions(width, height, false)
	return nil
}
