package gl

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
	MaxVAO = 128
)

var (
	vaoIDs  = make([]uint32, MaxVAO)
	vaoUsed = make([]bool, MaxVAO)
)

func GlInit() error {
	if err := gl.Init(); err != nil {
		return err
	}

	return nil
}

func GetFreeVAOIId() (uint32, error) {
	for i, used := range vaoUsed {
		if used {
			vaoUsed[i] = true
			return vaoIDs[i], nil
		}
	}

	return 0, fmt.Errorf("Unable to find free VAO ID")
}

func FreeVAOIID(id uint32) error {
	if !vaoUsed[id] {
		return fmt.Errorf("Unable to free un-used VaoID: %v", id)
	}

	vaoUsed[id] = false

	return nil
}
