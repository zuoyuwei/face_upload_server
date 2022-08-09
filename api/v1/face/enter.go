package face

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	FacePatientApi
	FacePhotoApi
	FaceVideoApi
	PersonalUploadFileApi
}

var (
	facePatientService        = service.ServiceGroupApp.FaceServiceGroup.FacePatientService
	facePhotoService          = service.ServiceGroupApp.FaceServiceGroup.FacePhotoService
	faceVideoService          = service.ServiceGroupApp.FaceServiceGroup.FaceVideoService
	personalUploadFileService = service.ServiceGroupApp.FaceServiceGroup.FaceVideoService
)
