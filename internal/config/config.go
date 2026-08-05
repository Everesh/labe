package config

import (
	"charm.land/log/v2"
	"codeberg.org/everesh/labe/internal/proto"
	"codeberg.org/everesh/labe/internal/xkb"
)

// TODO - hook up this to a conf file parser so the keybinds are not hardcoded

func ConfigureBindings(s Seat) ([]*xkb.Binding, []*xkb.PointerBinding) {

	// 1 - L-ALT, 4 - Super
	const mainMod = proto.RiverSeatV1ModifiersMod4

	bindings := []*xkb.Binding{
		xkb.NewBinding(s, xkb.KEY_Return, mainMod, func() {
			Spawn("alacritty")
		}),
		xkb.NewBinding(s, xkb.KEY_Return, mainMod+proto.RiverSeatV1ModifiersShift, func() {
			Spawn("ghostty")
		}),
		xkb.NewBinding(s, xkb.KEY_Return, mainMod+proto.RiverSeatV1ModifiersShift+proto.RiverSeatV1ModifiersCtrl, func() {
			Spawn("foot")
		}),
		xkb.NewBinding(s, xkb.KEY_c, mainMod, func() {
			if w := s.GetFocused(); w != nil {
				w.Object.Close()
			}
		}),
		xkb.NewBinding(s, xkb.KEY_z, mainMod, func() {
			Spawn("wofi", "--show", "run")
		}),
		xkb.NewBinding(s, xkb.KEY_z, mainMod+proto.RiverSeatV1ModifiersShift, func() {
			Spawn("zeditor")
		}),
		xkb.NewBinding(s, xkb.KEY_f, mainMod+proto.RiverSeatV1ModifiersShift, func() {
			Spawn("nautilus")
		}),
		xkb.NewBinding(s, xkb.KEY_w, mainMod+proto.RiverSeatV1ModifiersShift, func() {
			Spawn("firefox")
		}),
		xkb.NewBinding(s, xkb.KEY_n, mainMod, func() {
			windows := s.GetWindows()
			if len(windows) > 0 {
				s.Focus(windows[0])
			}
		}),
		xkb.NewBinding(s, xkb.KEY_c, mainMod, func() {
			s.ExitSession()
		}),
	}

	pointerBindings := []*xkb.PointerBinding{
		xkb.NewPointerBinding(s, xkb.KEY_LeftPointer, mainMod, func() {
			w := s.GetHovered()
			if w == nil {
				return
			}

			log.Debug("operation requested, window move, via keybind", "window", w.Object)
			s.PointerMove(w)
		}),
		xkb.NewPointerBinding(s, xkb.KEY_RightPointer, mainMod, func() {
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
