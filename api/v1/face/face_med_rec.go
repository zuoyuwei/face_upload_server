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

type FaceMedicalRecordApi struct{}

// @Tags FaceMedicalRecord
// @Summary 创建病历
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body face.FaceMedicalRecord true "患者姓名, 患者编号, 有无腺样体肥大, 扁桃体级别, 是否张口呼吸, 牙齿有无白色沉积"
// @Success 200 {object} response.Response{msg=string} "创建病历"
// @Router /faceMedicalRecord/faceMedicalRecord [post]
func (e *FaceMedicalRecordApi) CreateFaceMedicalRecord(c *gin.Context) {
	//fmt.Println("test")
	var FaceMedicalRecord face.FaceMedicalRecord
	_ = c.ShouldBindJSON(&FaceMedicalRecord)
	//if err := utils.Verify(FaceMedicalRecord, utils.FaceMedicalRecordVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	//FaceMedicalRecord.SysUserID = utils.GetUserID(c)
	//FaceMedicalRecord.SysUserAuthorityID = utils.GetUserAuthorityId(c)
	if _, err := faceMedicalRecordService.CreateFaceMedicalRecord(FaceMedicalRecord); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags FaceMedicalRecord
// @Summary 删除病历
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body face.FaceMedicalRecord true "病历ID"
// @Success 200 {object} response.Response{msg=string} "删除病历"
// @Router /faceMedicalRecord/faceMedicalRecord [delete]
func (e *FaceMedicalRecordApi) DeleteFaceMedicalRecord(c *gin.Context) {
	var FaceMedicalRecord face.FaceMedicalRecord
	_ = c.ShouldBindJSON(&FaceMedicalRecord)
	//if err := utils.Verify(FaceMedicalRecord.GVA_MODEL, utils.IdVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	if err := faceMedicalRecordService.DeleteFaceMedicalRecord(FaceMedicalRecord); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags FaceMedicalRecord
// @Summary 更新病历信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body face.FaceMedicalRecord true "病历ID, 病历信息"
// @Success 200 {object} response.Response{msg=string} "更新病历信息"
// @Router /faceMedicalRecord/faceMedicalRecord [put]
func (e *FaceMedicalRecordApi) UpdateFaceMedicalRecord(c *gin.Context) {
	var FaceMedicalRecord face.FaceMedicalRecord
	_ = c.ShouldBindJSON(&FaceMedicalRecord)
	//if err := utils.Verify(FaceMedicalRecord.GVA_MODEL, utils.IdVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	//if err := utils.Verify(FaceMedicalRecord, utils.FaceMedicalRecordVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	if err := faceMedicalRecordService.UpdateFaceMedicalRecord(&FaceMedicalRecord); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags FaceMedicalRecord
// @Summary 获取单一病历信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query face.FaceMedicalRecord true "病历ID"
// @Success 200 {object} response.Response{data=faceRes.FaceMedicalRecordResponse,msg=string} "获取单一病历信息,返回包括病历详情"
// @Router /faceMedicalRecord/faceMedicalRecord [get]
func (e *FaceMedicalRecordApi) GetFaceMedicalRecord(c *gin.Context) {
	var FaceMedicalRecord face.FaceMedicalRecord
	_ = c.ShouldBindQuery(&FaceMedicalRecord)
	//if err := utils.Verify(FaceMedicalRecord.GVA_MODEL, utils.IdVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	//data, err := faceMedicalRecordService.GetFaceMedicalRecord(FaceMedicalRecord.ID)
	data, err := GetFaceMedicalRecord_Son(FaceMedicalRecord.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(faceRes.FaceMedicalRecordResponse{FaceMedicalRecord: data}, "获取成功", c)
	}
}

func GetFaceMedicalRecord_Son(medicRecId uint) (medic_rec face.FaceMedicalRecord, err error) {
	medic_rec, _ = faceMedicalRecordService.GetFaceMedicalRecord(medicRecId)
	var pageInfo request.PageInfo
	pageInfo.GuanLianId = int(medic_rec.ID)
	fileList, _, err := exaFileUploadFileService.GetFileRecordInfoList_A(pageInfo)
	medic_rec.FileList = fileList
	return
}

// @Tags FaceMedicalRecord
// @Summary 分页获取权限病历列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取权限病历列表,返回包括列表,总数,页码,每页数量"
// @Router /faceMedicalRecord/faceMedicalRecordList [get]
func (e *FaceMedicalRecordApi) GetFaceMedicalRecordList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	//if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	FaceMedicalRecordList, total, err := faceMedicalRecordService.GetFaceMedicalRecordInfoList(utils.GetUserAuthorityId(c), pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     FaceMedicalRecordList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
