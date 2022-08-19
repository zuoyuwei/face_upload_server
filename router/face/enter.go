package face

type RouterGroup struct {
	FacePatientRouter
	FacePhotoRouter
	FaceVideoRouter
	FaceMedicalRecordRouter
	FileUploadRouter
}
