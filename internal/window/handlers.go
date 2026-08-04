package window

import (
	"context"

	"charm.land/log/v2"
)

func (w *Window) HandleRiverWindowV1Closed(ctx context.Context) {
	log.Info("window closed", "window", w.Object)
	w.Closed = true
}
