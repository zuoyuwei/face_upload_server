package face

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
)

type FaceMedicalRecord struct {
	global.GVA_MODEL
	PatientName         string                             `json:"patientName" form:"patientName" gorm:"comment:患者姓名"`                     // 患者姓名
	PatientCode         string                             `json:"patientCode" form:"patientCode" gorm:"comment:患者编号"`                     // 患者编号
	PatientId           int                                `json:"patientId" form:"patientId" gorm:"comment:患者ID"`                         // 患者ID
	IsAdenoidBodyLarge  string                             `json:"isAdenoidBodyLarge" form:"isAdenoidBodyLarge" gorm:"comment:有无腺样体肥大"`    // 有无腺样体肥大
	TonsilDegree        string                             `json:"tonsilDegree" form:"tonsilDegree" gorm:"comment:扁桃体级别"`                  // 扁桃体级别
	IsOpenMouthBreathe  string                             `json:"isOpenMouthBreathe" form:"isOpenMouthBreathe" gorm:"comment:是否张口呼吸"`     // 是否张口呼吸
	IsTeethWhiteDeposit string                             `json:"isTeethWhiteDeposit" form:"isTeethWhiteDeposit" gorm:"comment:牙齿有无白色沉积"` // 牙齿有无白色沉积
	FileList            []example.ExaFileUploadAndDownload `json:"fileList" form:"fileList" gorm:"-"`                                      // 文件列表
}
