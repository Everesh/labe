package seat

import (
	"context"

	"charm.land/log/v2"
)

func (s *Seat) HandleRiverSeatV1Removed(ctx context.Context) {
	log.Info("seat removed", "id", s.Object)
	s.Removed = true
}
