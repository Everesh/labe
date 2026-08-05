package output

import "codeberg.org/everesh/labe/internal/proto"

// required

type WindowManager interface {
	GetLayerShellOutput(proto.RiverOutputV1) proto.RiverLayerShellOutputV1
}

// advertised

func (o *Output) GetX() int32 {
	return o.X
}

func (o *Output) GetY() int32 {
	return o.Y
}

func (o *Output) GetWidth() int32 {
	return o.Width
}

func (o *Output) GetHeight() int32 {
	return o.Height
}
