package window

import "context"

func (w *Window) HandleRiverWindowV1Closed(ctx context.Context) {
	w.Closed = true
}
