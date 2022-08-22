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
		faceMedicalRecordRouterWithoutRecord.POST("faceMedicalRecord", FaceMedicalRecordApi.CreateFaceMedicalRecord)   // 创建病历
		faceMedicalRecordRouterWithoutRecord.PUT("faceMedicalRecord", FaceMedicalRecordApi.UpdateFaceMedicalRecord)    // 更新病历
		faceMedicalRecordRouterWithoutRecord.DELETE("faceMedicalRecord", FaceMedicalRecordApi.DeleteFaceMedicalRecord) // 删除病历
	}
	{
		faceMedicalRecordRouterWithoutRecord.GET("faceMedicalRecord", FaceMedicalRecordApi.GetFaceMedicalRecord)         // 获取单一病历信息
		faceMedicalRecordRouterWithoutRecord.GET("faceMedicalRecordList", FaceMedicalRecordApi.GetFaceMedicalRecordList) // 获取病历列表
	}
}
