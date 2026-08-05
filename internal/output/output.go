package output

import "codeberg.org/everesh/labe/internal/proto"

type Output struct {
	proto.RiverOutputV1Stub
	proto.RiverLayerShellOutputV1Stub

	Object           proto.RiverOutputV1
	LayerShellObject proto.RiverLayerShellOutputV1

	X, Y          int32
	Width, Height int32

	NonExclusiveX, NonExclusiveY          int32
	NonExclusiveWidth, NonExclusiveHeight int32

	Removed                     bool
	PendingSetDefaultLayerShell bool
}

func (o *Output) Contains(x, y int32) bool {
	return x >= o.X && x < o.X+o.Width &&
		y >= o.Y && y < o.Y+o.Height
}

func (o *Output) Manage() {
	if o.PendingSetDefaultLayerShell {
		o.LayerShellObject.SetDefault()
		o.PendingSetDefaultLayerShell = false
	}
}

func (o *Output) MaybeDestroy() bool {
	if !o.Removed {
		return false
	}

	o.Object.Destroy()
	o.LayerShellObject.Destroy()
	return true
}

func (o *Output) SetDefaultLayerShell() {
	o.PendingSetDefaultLayerShell = true
}

func NewOutput(object proto.RiverOutputV1, wm WindowManager) *Output {
	layerShellOutput := wm.GetLayerShellOutput(object)
	output := &Output{
		Object:           object,
		LayerShellObject: layerShellOutput,
	}

	object.SetUserData(output)
	layerShellOutput.SetUserData(output)
	return output
}
