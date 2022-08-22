package face

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type FacePatient struct {
	global.GVA_MODEL
	PatientName       string              `json:"patientName" form:"patientName" gorm:"comment:客户名"`   // 客户名
	Age               int                 `json:"age" form:"age" gorm:"comment:年龄"`                    // 年龄
	Gender            string              `json:"gender" form:"gender" gorm:"comment:性别"`              // 性别
	PhoneNumber       string              `json:"phoneNumber" form:"phoneNumber" gorm:"comment:联系方式"`  // 联系方式
	PatientCode       string              `json:"patientCode" form:"patientCode" gorm:"comment:患者编号"`  // 患者编号
	MedicalRecordList []FaceMedicalRecord `json:"medicalRecordList" form:"medicalRecordList" gorm:"-"` // 病历列表
}
