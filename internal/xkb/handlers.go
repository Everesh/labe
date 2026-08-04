package xkb

import (
	"context"
)

func (p *Binding) HandleRiverXkbBindingV1Pressed(ctx context.Context) {
	p.Seat.SetPendingAction(p.ActionFunc)
}

func (p *PointerBinding) HandleRiverPointerBindingV1Pressed(ctx context.Context) {
	p.Seat.SetPendingAction(p.ActionFunc)
}
