package output

import "codeberg.org/everesh/labe/internal/proto"

// required

type WindowManager interface {
	GetLayerShellOutput(proto.RiverOutputV1) proto.RiverLayerShellOutputV1
}

// advertised

func (o *Output) GetX() int32 {
	return o.NonExclusiveX
}

func (o *Output) GetY() int32 {
	return o.NonExclusiveY
}

func (o *Output) GetWidth() int32 {
	return o.NonExclusiveWidth
}

func (o *Output) GetHeight() int32 {
	return o.NonExclusiveHeight
}
