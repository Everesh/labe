package seat

import "codeberg.org/everesh/labe/internal/proto"

type Seat struct {
	proto.RiverSeatV1Stub

	Object proto.RiverSeatV1

	New     bool
	Removed bool
}

func (s *Seat) Manage() {
	if s.New {
		s.New = false
	}
}

func (s *Seat) MaybeDestroy() bool {
	if !s.Removed {
		return false
	}

	s.Object.Destroy()
	return true
}

func NewSeat(object proto.RiverSeatV1) *Seat {
	seat := &Seat{
		Object: object,
		New:    true,
	}

	seat.Object.SetUserData(seat)
	return seat
}
