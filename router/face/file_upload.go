package face

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type FileUploadRouter struct{}

func (e *FileUploadRouter) InitFileUploadRouter(Router *gin.RouterGroup) {
	//fileUploadRouter := Router.Group("fileUpload").Use(middleware.OperationRecord())
	fileUploadRouterWithoutRecord := Router.Group("fileUpload")
	FileUploadApi := v1.ApiGroupApp.FaceApiGroup.PersonalUploadFileApi
	{
		fileUploadRouterWithoutRecord.POST("fileUpload", FileUploadApi.UploadGo) // 上传文件
		fileUploadRouterWithoutRecord.POST("fileDown", FileUploadApi.DownGo)     // 下载文件
	}
	{
		//fileUploadRouterWithoutRecord.GET("fileUpload", FileUploadApi.GetFileUpload)         // 获取单一患者信息
		//fileUploadRouterWithoutRecord.GET("fileUploadList", FileUploadApi.GetFileUploadList) // 获取患者列表
	}
}
