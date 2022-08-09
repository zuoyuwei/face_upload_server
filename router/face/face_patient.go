package face

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type FacePatientRouter struct{}

func (e *FacePatientRouter) InitFacePatientRouter(Router *gin.RouterGroup) {
	//facePatientRouter := Router.Group("facePatient").Use(middleware.OperationRecord())
	facePatientRouterWithoutRecord := Router.Group("facePatient")
	FacePatientApi := v1.ApiGroupApp.FaceApiGroup.FacePatientApi
	{
		facePatientRouterWithoutRecord.POST("facePatient", FacePatientApi.CreateFacePatient)   // 创建患者
		facePatientRouterWithoutRecord.PUT("facePatient", FacePatientApi.UpdateFacePatient)    // 更新患者
		facePatientRouterWithoutRecord.DELETE("facePatient", FacePatientApi.DeleteFacePatient) // 删除患者
	}
	{
		facePatientRouterWithoutRecord.GET("facePatient", FacePatientApi.GetFacePatient)         // 获取单一患者信息
		facePatientRouterWithoutRecord.GET("facePatientList", FacePatientApi.GetFacePatientList) // 获取患者列表
	}
}
