package xkb

import (
	"charm.land/log/v2"
	"codeberg.org/everesh/labe/internal/proto"
)

func ConfigureBindings(s Seat) ([]*Binding, []*PointerBinding) {

	// 1 - L-ALT, 4 - Super
	const mainMod = proto.RiverSeatV1ModifiersMod1

	bindings := []*Binding{
		NewBinding(s, Key.Return, mainMod, func() {
			Spawn("alacritty")
		}),
		NewBinding(s, Key.Return, mainMod+proto.RiverSeatV1ModifiersShift, func() {
			Spawn("ghostty")
		}),
		NewBinding(s, Key.Return, mainMod+proto.RiverSeatV1ModifiersShift+proto.RiverSeatV1ModifiersCtrl, func() {
			Spawn("foot")
		}),
		NewBinding(s, Key.c, mainMod, func() {
			if w := s.GetFocused(); w != nil {
				w.Object.Close()
			}
		}),
		NewBinding(s, Key.z, mainMod, func() {
			Spawn("wofi", "--show", "run")
		}),
		NewBinding(s, Key.f, mainMod+proto.RiverSeatV1ModifiersShift, func() {
			Spawn("nautilus")
		}),
		NewBinding(s, Key.w, mainMod+proto.RiverSeatV1ModifiersShift, func() {
			Spawn("firefox")
		}),
		NewBinding(s, Key.n, mainMod, func() {
			windows := s.GetWindows()
			if len(windows) > 0 {
				s.Focus(windows[0])
			}
		}),
		NewBinding(s, Key.Esc, mainMod, func() {
			s.ExitSession()
		}),
	}

	pointerBindings := []*PointerBinding{
		NewPointerBinding(s, Key.LeftPointer, mainMod, func() {
			w := s.GetHovered()
			if w == nil {
				return
			}

			log.Debug("operation requested, window move, via keybind", "window", w.Object)
			s.PointerMove(w)
		}),
		NewPointerBinding(s, Key.RightPointer, mainMod, func() {
			w := s.GetHovered()
			if w == nil {
				return
			}

			x, y := s.GetPosition()
			var edges uint32 = proto.RiverWindowV1EdgesNone
			if (x - w.X) > w.Width/2 {
				edges |= proto.RiverWindowV1EdgesRight
			} else {
				edges |= proto.RiverWindowV1EdgesLeft
			}

			if (y - w.Y) > w.Height/2 {
				edges |= proto.RiverWindowV1EdgesBottom
			} else {
				edges |= proto.RiverWindowV1EdgesTop
			}

			log.Debug("operation requested, window resize, via keybind", "window", w.Object)
			s.PointerResize(w, edges)
		}),
	}

	return bindings, pointerBindings
}
