package face

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type FacePhoto struct {
	global.GVA_MODEL
	PatientId uint   `json:"patientId" form:"patientId" gorm:"comment:患者ID"` // 患者ID
	Url       string `json:"url" form:"url" gorm:"comment:图片地址"`             // 图片地址
}
