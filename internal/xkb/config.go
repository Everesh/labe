package xkb

import (
	"charm.land/log/v2"
	"codeberg.org/everesh/labe/internal/proto"
)

func ConfigureBindings(s Seat) ([]*Binding, []*PointerBinding) {
	const (
		// See xkbcommon-keysyms.h
		escSym   = 0xff1b
		enterSym = 0xff0d
		spaceSym = 0x0020

		nSym = 0x006e
		qSym = 0x0071
		cSym = 0x0063
		zSym = 0x007a
		fSym = 0x0066
		wSym = 0x0077

		leftButton  = 0x110
		rightButton = 0x111
	)

	// 1 - L-ALT, 4 - Super
	const mainMod = proto.RiverSeatV1ModifiersMod1

	bindings := []*Binding{
		NewBinding(s, enterSym, mainMod, func() {
			Spawn("alacritty")
		}),
		NewBinding(s, enterSym, mainMod+proto.RiverSeatV1ModifiersShift, func() {
			Spawn("ghostty")
		}),
		NewBinding(s, enterSym, mainMod+proto.RiverSeatV1ModifiersShift+proto.RiverSeatV1ModifiersCtrl, func() {
			Spawn("foot")
		}),
		NewBinding(s, cSym, mainMod, func() {
			if w := s.GetFocused(); w != nil {
				w.Object.Close()
			}
		}),
		NewBinding(s, zSym, mainMod, func() {
			Spawn("wofi", "--show", "run")
		}),
		NewBinding(s, fSym, mainMod+proto.RiverSeatV1ModifiersShift, func() {
			Spawn("nautilus")
		}),
		NewBinding(s, wSym, mainMod+proto.RiverSeatV1ModifiersShift, func() {
			Spawn("firefox")
		}),
		NewBinding(s, nSym, mainMod, func() {
			windows := s.GetWindows()
			if len(windows) > 0 {
				s.Focus(windows[0])
			}
		}),
		NewBinding(s, escSym, mainMod, func() {
			s.ExitSession()
		}),
	}

	pointerBindings := []*PointerBinding{
		NewPointerBinding(s, leftButton, mainMod, func() {
			w := s.GetHovered()
			if w == nil {
				return
			}

			log.Debug("operation requested, window move, via keybind", "window", w.Object)
			s.PointerMove(w)
		}),
		NewPointerBinding(s, rightButton, mainMod, func() {
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
