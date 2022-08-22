package face

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/face"
	faceRes "github.com/flipped-aurora/gin-vue-admin/server/model/face/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type FacePatientApi struct{}

// @Tags FacePatient
// @Summary 创建患者
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body face.FacePatient true "患者用户名, 患者手机号码"
// @Success 200 {object} response.Response{msg=string} "创建患者"
// @Router /facePatient/facePatient [post]
func (e *FacePatientApi) CreateFacePatient(c *gin.Context) {
	//fmt.Println("test")
	var FacePatient face.FacePatient
	//fmt.Println(c.Request)
	_ = c.ShouldBindJSON(&FacePatient)
	//if err := utils.Verify(FacePatient, utils.FacePatientVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	//FacePatient.SysUserID = utils.GetUserID(c)
	//FacePatient.SysUserAuthorityID = utils.GetUserAuthorityId(c)
	if _, err := facePatientService.CreateFacePatient(FacePatient); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags FacePatient
// @Summary 删除患者
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body face.FacePatient true "患者ID"
// @Success 200 {object} response.Response{msg=string} "删除患者"
// @Router /facePatient/facePatient [delete]
func (e *FacePatientApi) DeleteFacePatient(c *gin.Context) {
	var FacePatient face.FacePatient
	_ = c.ShouldBindJSON(&FacePatient)
	//if err := utils.Verify(FacePatient.GVA_MODEL, utils.IdVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	if err := facePatientService.DeleteFacePatient(FacePatient); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags FacePatient
// @Summary 更新患者信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body face.FacePatient true "患者ID, 患者信息"
// @Success 200 {object} response.Response{msg=string} "更新患者信息"
// @Router /facePatient/facePatient [put]
func (e *FacePatientApi) UpdateFacePatient(c *gin.Context) {
	var FacePatient face.FacePatient
	_ = c.ShouldBindJSON(&FacePatient)
	//if err := utils.Verify(FacePatient.GVA_MODEL, utils.IdVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	//if err := utils.Verify(FacePatient, utils.FacePatientVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	if err := facePatientService.UpdateFacePatient(&FacePatient); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags FacePatient
// @Summary 获取单一患者信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query face.FacePatient true "患者ID"
// @Success 200 {object} response.Response{data=faceRes.FacePatientResponse,msg=string} "获取单一患者信息,返回包括患者详情"
// @Router /facePatient/facePatient [get]
func (e *FacePatientApi) GetFacePatient(c *gin.Context) {
	var FacePatient face.FacePatient
	_ = c.ShouldBindQuery(&FacePatient)
	//if err := utils.Verify(FacePatient.GVA_MODEL, utils.IdVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	//data, err := facePatientService.GetFacePatient(FacePatient.ID)
	fmt.Println("ID:", FacePatient.ID)
	data, err := GetFacePatient_Son(FacePatient.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(faceRes.FacePatientResponse{FacePatient: data}, "获取成功", c)
	}
}

func GetFacePatient_Son(patientId uint) (patient face.FacePatient, err error) {
	patient, err = facePatientService.GetFacePatient(patientId)
	var pageInfo request.PageInfo
	pageInfo.GuanLianId = int(patient.ID)
	FaceMedicalRecordList, _, err := faceMedicalRecordService.GetFaceMedicalRecordInfoList_A(0, pageInfo)
	patient.MedicalRecordList = FaceMedicalRecordList
	return
}

// @Tags FacePatient
// @Summary 分页获取权限患者列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页获取权限患者列表,返回包括列表,总数,页码,每页数量"
// @Router /facePatient/facePatientList [get]
func (e *FacePatientApi) GetFacePatientList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindQuery(&pageInfo)
	//if err := utils.Verify(pageInfo, utils.PageInfoVerify); err != nil {
	//	response.FailWithMessage(err.Error(), c)
	//	return
	//}
	FacePatientList, total, err := facePatientService.GetFacePatientInfoList(utils.GetUserAuthorityId(c), pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     FacePatientList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
