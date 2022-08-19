package example

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	exampleRes "github.com/flipped-aurora/gin-vue-admin/server/model/example/response"
	modelFace "github.com/flipped-aurora/gin-vue-admin/server/model/face"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"mime/multipart"
)

type FileUploadAndDownloadApi struct{}

type FilePatientCode struct {
	PatientCode string `json:"patientCode" form:"patientCode"`       // 患者编号
	Type        string `json:"type" form:"type" gorm:"comment:文件类型"` // 文件类型(image, video)
}

// @Tags ExaFileUploadAndDownload
// @Summary 上传文件示例
// @Security ApiKeyAuth
// @accept multipart/form-data
// @Produce  application/json
// @Param file formData file true "上传文件示例"
// @Param data query FilePatientCode true "患者编号,文件类型"
// @Success 200 {object} response.Response{data=exampleRes.ExaFileResponse,msg=string} "上传文件示例,返回包括文件详情"
// @Router /fileUploadAndDownload/upload [post]
func (b *FileUploadAndDownloadApi) UploadFile(c *gin.Context) {
	var file example.ExaFileUploadAndDownload
	noSave := c.DefaultQuery("noSave", "0")
	_, header, err := c.Request.FormFile("file")
	patientcode := c.Request.FormValue("patientCode")
	tp := c.Request.FormValue("type")
	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	//file, err = fileUploadAndDownloadService.UploadFile(header, noSave, patientcode, tp) // 文件上传后拿到文件路径
	file, err = UploadFile_Son(header, noSave, patientcode, tp)
	if err != nil {
		global.GVA_LOG.Error("修改数据库链接失败!", zap.Error(err))
		response.FailWithMessage("修改数据库链接失败", c)
		return
	}
	response.OkWithDetailed(exampleRes.ExaFileResponse{File: file}, "上传成功", c)
}

func UploadFile_Son(header *multipart.FileHeader, noSave string, patientcode string, t string) (file example.ExaFileUploadAndDownload, err error) {
	// 根据患者编号查询是否有此患者
	patientId := 0
	medicalRecordId := 0
	face_med_rec, err := faceMedicalRecordService.FindFaceMedicalRecord(patientcode)
	medicalRecordId = int(face_med_rec.ID)
	// 如果有，继续下一步，如果没有，就新建患者
	if medicalRecordId == 0 {
		var FacePatient modelFace.FacePatient
		FacePatient.PatientCode = patientcode
		patientId, err = facePatientService.CreateFacePatient(FacePatient)

		var FaceMedicalRecord modelFace.FaceMedicalRecord
		FaceMedicalRecord.PatientCode = patientcode
		FaceMedicalRecord.PatientId = patientId
		medicalRecordId, err = faceMedicalRecordService.CreateFaceMedicalRecord(FaceMedicalRecord)
	}
	// 建立联系
	file, err = fileUploadAndDownloadService.UploadFile(header, noSave, medicalRecordId, patientcode, t) // 文件上传后拿到文件路径
	return file, err
}

// EditFileName 编辑文件名或者备注
func (b *FileUploadAndDownloadApi) EditFileName(c *gin.Context) {
	var file example.ExaFileUploadAndDownload
	_ = c.ShouldBindJSON(&file)
	if err := fileUploadAndDownloadService.EditFileName(file); err != nil {
		global.GVA_LOG.Error("编辑失败!", zap.Error(err))
		response.FailWithMessage("编辑失败", c)
		return
	}
	response.OkWithMessage("编辑成功", c)
}

// @Tags ExaFileUploadAndDownload
// @Summary 删除文件
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body example.ExaFileUploadAndDownload true "传入文件里面id即可"
// @Success 200 {object} response.Response{msg=string} "删除文件"
// @Router /fileUploadAndDownload/deleteFile [post]
func (b *FileUploadAndDownloadApi) DeleteFile(c *gin.Context) {
	var file example.ExaFileUploadAndDownload
	_ = c.ShouldBindJSON(&file)
	if err := fileUploadAndDownloadService.DeleteFile(file); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// @Tags ExaFileUploadAndDownload
// @Summary 分页文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "页码, 每页大小"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "分页文件列表,返回包括列表,总数,页码,每页数量"
// @Router /fileUploadAndDownload/getFileList [post]
func (b *FileUploadAndDownloadApi) GetFileList(c *gin.Context) {
	var pageInfo request.PageInfo
	_ = c.ShouldBindJSON(&pageInfo)
	list, total, err := fileUploadAndDownloadService.GetFileRecordInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
