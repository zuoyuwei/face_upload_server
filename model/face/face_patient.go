package face

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type FacePatient struct {
	global.GVA_MODEL
	PatientName string `json:"PatientName" form:"PatientName" gorm:"comment:客户名"` // 客户名
}
