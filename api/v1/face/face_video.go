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

type FaceVideoApi struct{}

// @Tags FaceVideo
// @Summary 创建视频
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body face.FaceVideo true "视频用户名, 视频手机号码"
// @Success 200 {object} response.Response{msg=string} "创建视频"
// @Router /FaceVideo/FaceVideo [post]
func (e *FaceVideoApi) CreateFaceVideo(c *gin.Context) {
	var FaceVideo face.FaceVideo
	_ = c.ShouldBindJSON(&FaceVideo)
	//if err := utils.Verify(FaceVideo, utils.FaceVideoVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	//FaceVideo.SysUserID = utils.GetUserID(c)
	//FaceVideo.SysUserAuthorityID = utils.GetUserAuthorityId(c)
	if err := faceVideoService.CreateFaceVideo(FaceVideo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags FaceVideo
// @Summary 删除视频
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body face.FaceVideo true "视频ID"
// @Success 200 {object} response.Response{msg=string} "删除视频"
// @Router /FaceVideo/FaceVideo [delete]
func (e *FaceVideoApi) DeleteFaceVideo(c *gin.Context) {
	var FaceVideo face.FaceVideo
	_ = c.ShouldBindJSON(&FaceVideo)
	//if err := utils.Verify(FaceVideo.GVA_MODEL, utils.IdVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	if err := faceVideoService.DeleteFaceVideo(FaceVideo); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags FaceVideo
// @Summary 更新视频信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body face.FaceVideo true "视频ID, 视频信息"
// @Success 200 {object} response.Response{msg=string} "更新视频信息"
// @Router /FaceVideo/FaceVideo [put]
func (e *FaceVideoApi) UpdateFaceVideo(c *gin.Context) {
	var FaceVideo face.FaceVideo
	_ = c.ShouldBindJSON(&FaceVideo)
	//if err := utils.Verify(FaceVideo.GVA_MODEL, utils.IdVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	//if err := utils.Verify(FaceVideo, utils.FaceVideoVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	if err := faceVideoService.UpdateFaceVideo(&FaceVideo); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags FaceVideo
// @Summary 获取单一视频信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query face.FaceVideo true "视频ID"
// @Success 200 {object} response.Response{data=faceRes.FaceVideoResponse,msg=string} "获取单一视频信息,返回包括视频详情"
// @Router /FaceVideo/FaceVideo [get]
func (e *FaceVideoApi) GetFaceVideo(c *gin.Context) {
	var FaceVideo face.FaceVideo
	_ = c.ShouldBindQuery(&FaceVideo)
	//if err := utils.Verify(FaceVideo.GVA_MODEL, utils.IdVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	data, err := faceVideoService.GetFaceVideo(FaceVideo.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(faceRes.FaceVideoResponse{FaceVideo: data}, "获取成功", c)
	}
}

// @Tags FaceVideo
// @Summary 分页获取权限视频列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取权限视频列表,返回包括列表,总数,页码,每页数量"
// @Router /FaceVideo/FaceVideoList [get]
func (e *FaceVideoApi) GetFaceVideoList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	//if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	FaceVideoList, total, err := faceVideoService.GetFaceVideoInfoList(utils.GetUserAuthorityId(c), pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     FaceVideoList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
