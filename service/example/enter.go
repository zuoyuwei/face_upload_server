package example

import "github.com/flipped-aurora/gin-vue-admin/server/service/face"

type ServiceGroup struct {
	face.FacePatientService
	ExcelService
	CustomerService
	FileUploadAndDownloadService
}
