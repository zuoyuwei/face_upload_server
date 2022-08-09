package face

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type FacePhotoRouter struct{}

func (e *FacePhotoRouter) InitFacePhotoRouter(Router *gin.RouterGroup) {
	//facePhotoRouter := Router.Group("facePhoto").Use(middleware.OperationRecord())
	facePhotoRouterWithoutRecord := Router.Group("facePhoto")
	FacePhotoApi := v1.ApiGroupApp.FaceApiGroup.FacePhotoApi
	{
		facePhotoRouterWithoutRecord.POST("facePhoto", FacePhotoApi.CreateFacePhoto)   // 创建图片
		facePhotoRouterWithoutRecord.PUT("facePhoto", FacePhotoApi.UpdateFacePhoto)    // 更新图片
		facePhotoRouterWithoutRecord.DELETE("facePhoto", FacePhotoApi.DeleteFacePhoto) // 删除图片
	}
	{
		facePhotoRouterWithoutRecord.GET("facePhoto", FacePhotoApi.GetFacePhoto)         // 获取单一图片信息
		facePhotoRouterWithoutRecord.GET("facePhotoList", FacePhotoApi.GetFacePhotoList) // 获取图片列表
	}
}
