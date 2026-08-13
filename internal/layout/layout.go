package layout

import "codeberg.org/everesh/labe/internal/window"

type Layout interface {
	Add(w *window.Window) error

	Remove(w *window.Window) error

	Align(x, y, width, height int32) error
}
