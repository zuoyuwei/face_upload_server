package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/face"
)

type FacePhotoResponse struct {
	FacePhoto face.FacePhoto `json:"facePhoto"`
}
