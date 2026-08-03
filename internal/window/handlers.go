package window

import (
	"context"

	"charm.land/log/v2"
)

func (w *Window) HandleRiverWindowV1Closed(ctx context.Context) {
	log.Info("window closed", "id", w.Object)
	w.Closed = true
}
