package face

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
)

type FaceMedicalRecord struct {
	global.GVA_MODEL
	PatientName         string                             `json:"PatientName" form:"PatientName" gorm:"comment:客户名"`                      // 客户名
	PatientCode         string                             `json:"patientCode" form:"patientCode" gorm:"comment:患者编号"`                     // 患者编号
	PatientId           int                                `json:"patientId" form:"patientId" gorm:"comment:患者ID"`                         // 患者ID
	IsAdenoidBodyLarge  string                             `json:"isadenoidbodylarge" form:"isadenoidbodylarge" gorm:"comment:有无腺样体肥大"`    // 有无腺样体肥大
	TonsilDegree        string                             `json:"tonsildegree" form:"tonsildegree" gorm:"comment:扁桃体级别"`                  // 扁桃体级别
	IsOpenMouthBreathe  string                             `json:"isopenmouthbreathe" form:"isopenmouthbreathe" gorm:"comment:是否张口呼吸"`     // 是否张口呼吸
	IsTeethWhiteDeposit string                             `json:"isteethwhitedeposit" form:"isteethwhitedeposit" gorm:"comment:牙齿有无白色沉积"` // 牙齿有无白色沉积
	FileList            []example.ExaFileUploadAndDownload `json:"FileList" form:"FileList" gorm:"-"`                                      // 文件列表
}
