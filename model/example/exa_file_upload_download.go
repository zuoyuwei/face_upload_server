package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

type ExaFileUploadAndDownload struct {
	global.GVA_MODEL
	PatientCode string `json:"patientCode" form:"patientCode" gorm:"comment:患者编号"` // 患者编号
	Type        string `json:"type" form:"type" gorm:"comment:文件类型"`               // 文件类型(image, video)
	Name        string `json:"name" gorm:"comment:文件名"`                            // 文件名
	Url         string `json:"url" gorm:"comment:文件地址"`                            // 文件地址
	Tag         string `json:"tag" gorm:"comment:文件标签"`                            // 文件标签
	Key         string `json:"key" gorm:"comment:编号"`                              // 编号
}

func (ExaFileUploadAndDownload) TableName() string {
	return "exa_file_upload_and_downloads"
}
