package face

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type FaceVideoRouter struct{}

func (e *FaceVideoRouter) InitFaceVideoRouter(Router *gin.RouterGroup) {
	//faceVideoRouter := Router.Group("faceVideo").Use(middleware.OperationRecord())
	faceVideoRouterWithoutRecord := Router.Group("faceVideo")
	FaceVideoApi := v1.ApiGroupApp.FaceApiGroup.FaceVideoApi
	{
		faceVideoRouterWithoutRecord.POST("faceVideo", FaceVideoApi.CreateFaceVideo)   // 创建视频
		faceVideoRouterWithoutRecord.PUT("faceVideo", FaceVideoApi.UpdateFaceVideo)    // 更新视频
		faceVideoRouterWithoutRecord.DELETE("faceVideo", FaceVideoApi.DeleteFaceVideo) // 删除视频
	}
	{
		faceVideoRouterWithoutRecord.GET("faceVideo", FaceVideoApi.GetFaceVideo)         // 获取单一视频信息
		faceVideoRouterWithoutRecord.GET("faceVideoList", FaceVideoApi.GetFaceVideoList) // 获取视频列表
	}
}
