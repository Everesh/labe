package seat

import (
	"context"

	"charm.land/log/v2"
	"codeberg.org/everesh/labe/internal/proto"
	"codeberg.org/everesh/labe/internal/window"
)

func (s *Seat) HandleRiverSeatV1Removed(ctx context.Context) {
	log.Info("seat removed", "seat", s.Object)
	s.Removed = true
}

func (s *Seat) HandleRiverSeatV1PointerEnter(ctx context.Context, w proto.RiverWindowV1) {
	log.Debug("pointer entered", "seat", s.Object, "window", w)
	s.Hovered = w.UserData().(*window.Window)
}

func (s *Seat) HandleRiverSeatV1PointerLeave(ctx context.Context) {
	log.Debug("pointer left", "seat", s.Object, "window", s.Hovered.Object)
	s.Hovered = nil
}

func (s *Seat) HandleRiverSeatV1WindowInteraction(ctx context.Context, w proto.RiverWindowV1) {
	log.Debug("pointer interacted", "seat", s.Object, "window", w)
	s.Interacted = w.UserData().(*window.Window)
	s.WM.MarkSeatLastActive(s)
}

func (s *Seat) HandleRiverSeatV1PointerPosition(ctx context.Context, x int32, y int32) {
	s.X = x
	s.Y = y
}

func (s *Seat) HandleRiverSeatV1OpDelta(ctx context.Context, dx int32, dy int32) {
	s.OpState.Dx, s.OpState.Dy = dx, dy
}

func (s *Seat) HandleRiverSeatV1OpRelease(ctx context.Context) {
	log.Debug("operation released", "seat", s.Object)
	s.OpReleased = true
}
