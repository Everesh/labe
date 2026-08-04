package window

import (
	"context"

	"charm.land/log/v2"
	"codeberg.org/everesh/labe/internal/proto"
)

func (w *Window) HandleRiverWindowV1Closed(ctx context.Context) {
	log.Info("window closed", "window", w.Object)
	w.Closed = true
}

func (w *Window) HandleRiverWindowV1Dimensions(ctx context.Context, width int32, height int32) {
	w.Width, w.Height = width, height
}

func (w *Window) HandleRiverWindowV1PointerMoveRequested(ctx context.Context, s proto.RiverSeatV1) {
	log.Debug("operation requested, window move", "window", w.Object, "seat", s)
	w.PointerMoveRequested = s.UserData().(PointerRequester)
}

func (w *Window) HandleRiverWindowV1PointerResizeRequested(ctx context.Context, s proto.RiverSeatV1, edges uint32) {
	log.Debug("operation requested, window resize", "window", w.Object, "seat", s)
	w.PointerResizeRequested = s.UserData().(PointerRequester)
	w.PointerResizeRequestedEdges = edges
}
