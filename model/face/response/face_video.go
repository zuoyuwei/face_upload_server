package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/face"
)

type FaceVideoResponse struct {
	FaceVideo face.FaceVideo `json:"faceVideo"`
}
