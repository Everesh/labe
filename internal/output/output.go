package output

import "codeberg.org/everesh/labe/internal/proto"

type Output struct {
	proto.RiverOutputV1Stub

	Object proto.RiverOutputV1

	Removed bool
}

func (o *Output) MaybeDestroy() bool {
	if !o.Removed {
		return false
	}

	o.Object.Destroy()
	return true
}

func NewOutput(object proto.RiverOutputV1) *Output {
	output := &Output{
		Object: object,
	}

	object.SetUserData(output)
	return output
}
