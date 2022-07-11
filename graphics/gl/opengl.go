package gl

import (
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

func GetFreeVAOIId() (out uint32) {
	gl.GenVertexArrays(1, &out)

	return
	// for i, used := range vaoUsed {
	// 	if !used {
	// 		vaoUsed[i] = true
	// 		return vaoIDs[i], nil
	// 	}
	// }

	// return 0, fmt.Errorf("Unable to find free VAO ID")
}

func FreeVAOIID(id uint32) {
	gl.DeleteVertexArrays(1, &id)
}
