// 自动生成模板UploadFile
package model

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 如果含有time.Time 请自行import time包
type FtpFile struct {
	global.GVA_MODEL
	Shangchuanrenid int    `json:"shangchuanrenid" form:"shangchuanrenid" gorm:"column:shangchuanrenid;comment:上传人id;type:int;"`
	Wenjianid       int    `json:"wenjianid" form:"wenjianid" gorm:"column:wenjianid;comment:文件id;type:int;"`
	Wenjianming     string `json:"wenjianming" form:"wenjianming" gorm:"column:wenjianming;comment:文件名;"`
	Lujing          string `json:"lujing" form:"lujing" gorm:"column:lujing;comment:路径;type:varchar(255);size:255;"`
	//Pingtai         string `json:"pingtai" form:"pingtai" gorm:"column:pingtai;comment:平台;"`
}

func (FtpFile) TableName() string {
	return "personal_ftp_file"
}

// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type UploadFileWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	UploadFile   `json:"business"`
// }

// func (UploadFile) TableName() string {
// 	return "crm_upload_file"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["uploadFile"] = func() model.GVA_Workflow {
//   return new(model.UploadFileWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["uploadFile"] = func() interface{} {
// 	return new(model.UploadFile)
// }
