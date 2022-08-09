package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/face"
)

type FacePatientResponse struct {
	FacePatient face.FacePatient `json:"FacePatient"`
}
