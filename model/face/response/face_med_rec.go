package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/face"
)

type FaceMedicalRecordResponse struct {
	FaceMedicalRecord face.FaceMedicalRecord `json:"FaceMedicalRecord"`
}
