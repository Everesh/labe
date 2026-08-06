package window

import (
	"context"
	"time"

	"charm.land/log/v2"
	"codeberg.org/everesh/labe/internal/proto"
)

func (w *Window) HandleRiverWindowV1Closed(ctx context.Context) {
	log.Info("window closed", "window", w.Object)
	w.Closed = true
}

func (w *Window) HandleRiverWindowV1Dimensions(ctx context.Context, width int32, height int32) {
	w.Width, w.Height = width, height

	// guard against spwaning unnecessary gorutines
	if log.GetLevel() != log.DebugLevel {
		return
	}

	if w.debounceTimer != nil {
		w.debounceTimer.Stop()
	}
	w.debounceTimer = time.AfterFunc(150*time.Millisecond, func() {
		log.Debug("window dimensions changed", "window", w.Object, "width", width, "height", height)
	})
}

func (w *Window) HandleRiverWindowV1PointerMoveRequested(ctx context.Context, s proto.RiverSeatV1) {
	log.Debug("operation requested, window move", "window", w.Object, "seat", s)
	w.PointerMoveRequested = s.UserData().(Seat)
}

func (w *Window) HandleRiverWindowV1PointerResizeRequested(ctx context.Context, s proto.RiverSeatV1, edges uint32) {
	log.Debug("operation requested, window resize", "window", w.Object, "seat", s)
	w.PointerResizeRequested = s.UserData().(Seat)
	w.PointerResizeRequestedEdges = edges
}

func (w *Window) HandleRiverWindowV1DecorationHint(ctx context.Context, hint uint32) {
	log.Debug("decorations hinted", "window", w.Object, "hint", hint)
	w.DecorationHinted = true
	w.DecorationHint = hint
}

func (w *Window) HandleRiverWindowV1DimensionsHint(ctx context.Context, minWidth int32, minHeight int32, maxWidth int32, maxHeight int32) {
	log.Debug("dimensions hinted", "window", w.Object, "minWidth", minWidth, "minHeight", minHeight, "maxWidth", maxWidth, "maxHeight", maxHeight)
	w.MinWidth = minWidth
	w.MinHeight = minHeight
	w.MaxWidth = maxWidth
	w.MaxHeight = maxHeight
}
