package face

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/face"
	faceRes "github.com/flipped-aurora/gin-vue-admin/server/model/face/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FacePhotoApi struct{}

// @Tags FacePhoto
// @Summary 创建图片
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body face.FacePhoto true "图片用户名, 图片手机号码"
// @Success 200 {object} response.Response{msg=string} "创建图片"
// @Router /FacePhoto/FacePhoto [post]
func (e *FacePhotoApi) CreateFacePhoto(c *gin.Context) {
	var FacePhoto face.FacePhoto
	_ = c.ShouldBindJSON(&FacePhoto)
	//if err := utils.Verify(FacePhoto, utils.FacePhotoVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	//FacePhoto.SysUserID = utils.GetUserID(c)
	//FacePhoto.SysUserAuthorityID = utils.GetUserAuthorityId(c)
	if err := facePhotoService.CreateFacePhoto(FacePhoto); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags FacePhoto
// @Summary 删除图片
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body face.FacePhoto true "图片ID"
// @Success 200 {object} response.Response{msg=string} "删除图片"
// @Router /FacePhoto/FacePhoto [delete]
func (e *FacePhotoApi) DeleteFacePhoto(c *gin.Context) {
	var FacePhoto face.FacePhoto
	_ = c.ShouldBindJSON(&FacePhoto)
	//if err := utils.Verify(FacePhoto.GVA_MODEL, utils.IdVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	if err := facePhotoService.DeleteFacePhoto(FacePhoto); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags FacePhoto
// @Summary 更新图片信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body face.FacePhoto true "图片ID, 图片信息"
// @Success 200 {object} response.Response{msg=string} "更新图片信息"
// @Router /FacePhoto/FacePhoto [put]
func (e *FacePhotoApi) UpdateFacePhoto(c *gin.Context) {
	var FacePhoto face.FacePhoto
	_ = c.ShouldBindJSON(&FacePhoto)
	//if err := utils.Verify(FacePhoto.GVA_MODEL, utils.IdVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	//if err := utils.Verify(FacePhoto, utils.FacePhotoVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	if err := facePhotoService.UpdateFacePhoto(&FacePhoto); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags FacePhoto
// @Summary 获取单一图片信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query face.FacePhoto true "图片ID"
// @Success 200 {object} response.Response{data=faceRes.FacePhotoResponse,msg=string} "获取单一图片信息,返回包括图片详情"
// @Router /FacePhoto/FacePhoto [get]
func (e *FacePhotoApi) GetFacePhoto(c *gin.Context) {
	var FacePhoto face.FacePhoto
	_ = c.ShouldBindQuery(&FacePhoto)
	//if err := utils.Verify(FacePhoto.GVA_MODEL, utils.IdVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	data, err := facePhotoService.GetFacePhoto(FacePhoto.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(faceRes.FacePhotoResponse{FacePhoto: data}, "获取成功", c)
	}
}

// @Tags FacePhoto
// @Summary 分页获取权限图片列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取权限图片列表,返回包括列表,总数,页码,每页数量"
// @Router /FacePhoto/FacePhotoList [get]
func (e *FacePhotoApi) GetFacePhotoList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	//if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	FacePhotoList, total, err := facePhotoService.GetFacePhotoInfoList(utils.GetUserAuthorityId(c), pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     FacePhotoList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
