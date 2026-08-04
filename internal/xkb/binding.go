package xkb

import (
	"codeberg.org/everesh/labe/internal/proto"
)

type Binding struct {
	proto.RiverXkbBindingV1Stub

	Object proto.RiverXkbBindingV1

	Seat       Seat
	ActionFunc func()
}

func NewBinding(seat Seat, keysym uint32, mods uint32, fn func()) *Binding {
	b := &Binding{
		Object:     seat.GetXkbBinding(keysym, mods),
		Seat:       seat,
		ActionFunc: fn,
	}

	b.Object.SetUserData(b)
	return b
}

type PointerBinding struct {
	proto.RiverPointerBindingV1Stub

	Object proto.RiverPointerBindingV1

	Seat       Seat
	ActionFunc func()
}

func NewPointerBinding(seat Seat, button uint32, mods uint32, fn func()) *PointerBinding {
	b := &PointerBinding{
		Object:     seat.GetPointerBinding(button, mods),
		Seat:       seat,
		ActionFunc: fn,
	}

	b.Object.SetUserData(b)
	return b
}
