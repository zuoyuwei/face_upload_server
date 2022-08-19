package face

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type FaceMedicalRecordRouter struct{}

func (e *FaceMedicalRecordRouter) InitFaceMedicalRecordRouter(Router *gin.RouterGroup) {
	//faceMedicalRecordRouter := Router.Group("faceMedicalRecord").Use(middleware.OperationRecord())
	faceMedicalRecordRouterWithoutRecord := Router.Group("faceMedicalRecord")
	FaceMedicalRecordApi := v1.ApiGroupApp.FaceApiGroup.FaceMedicalRecordApi
	{
		faceMedicalRecordRouterWithoutRecord.POST("faceMedicalRecord", FaceMedicalRecordApi.CreateFaceMedicalRecord)   // 创建患者
		faceMedicalRecordRouterWithoutRecord.PUT("faceMedicalRecord", FaceMedicalRecordApi.UpdateFaceMedicalRecord)    // 更新患者
		faceMedicalRecordRouterWithoutRecord.DELETE("faceMedicalRecord", FaceMedicalRecordApi.DeleteFaceMedicalRecord) // 删除患者
	}
	{
		faceMedicalRecordRouterWithoutRecord.GET("faceMedicalRecord", FaceMedicalRecordApi.GetFaceMedicalRecord)         // 获取单一患者信息
		faceMedicalRecordRouterWithoutRecord.GET("faceMedicalRecordList", FaceMedicalRecordApi.GetFaceMedicalRecordList) // 获取患者列表
	}
}
