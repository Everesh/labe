package output

import "codeberg.org/everesh/labe/internal/proto"

// required

type WindowManager interface {
	GetLayerShellOutput(proto.RiverOutputV1) proto.RiverLayerShellOutputV1
}

// advertised
