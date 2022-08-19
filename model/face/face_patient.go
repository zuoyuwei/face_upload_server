package face

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type FacePatient struct {
	global.GVA_MODEL
	PatientName       string              `json:"PatientName" form:"PatientName" gorm:"comment:客户名"`   // 客户名
	Age               int                 `json:"Age" form:"Age" gorm:"comment:年龄"`                    // 年龄
	Gender            string              `json:"Gender" form:"Gender" gorm:"comment:性别"`              // 性别
	PhoneNumber       string              `json:"PhoneNumber" form:"PhoneNumber" gorm:"comment:联系方式"`  // 联系方式
	PatientCode       string              `json:"patientCode" form:"patientCode" gorm:"comment:患者编号"`  // 患者编号
	MedicalRecordList []FaceMedicalRecord `json:"medicalRecordList" form:"medicalRecordList" gorm:"-"` // 病历列表
}
