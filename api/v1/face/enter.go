package face

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	FacePatientApi
	FacePhotoApi
	FaceVideoApi
	FaceMedicalRecordApi
	PersonalUploadFileApi
}

var (
	facePatientService       = service.ServiceGroupApp.FaceServiceGroup.FacePatientService
	facePhotoService         = service.ServiceGroupApp.FaceServiceGroup.FacePhotoService
	faceVideoService         = service.ServiceGroupApp.FaceServiceGroup.FaceVideoService
	faceMedicalRecordService = service.ServiceGroupApp.FaceServiceGroup.FaceMedicalRecordService
	exaFileUploadFileService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
)
