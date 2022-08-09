package face

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type FileUploadService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateFileUpload
//@description: 创建图片
//@param: e model.FileUpload
//@return: err error

func (exa *FileUploadService) CreateFileUpload(e model.FtpFile) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除图片
//@param: e model.FileUpload
//@return: err error

func (exa *FileUploadService) DeleteFileUpload(e model.FtpFile) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateFileUpload
//@description: 更新图片
//@param: e *model.FileUpload
//@return: err error

func (exa *FileUploadService) UpdateFileUpload(e *model.FtpFile) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFileUpload
//@description: 获取图片信息
//@param: id uint
//@return: FileUpload model.FileUpload, err error

func (exa *FileUploadService) GetFileUpload(id uint) (FileUpload model.FtpFile, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&FileUpload).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFileUploadInfoList
//@description: 分页获取图片列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *FileUploadService) GetFileUploadInfoList(sysUserAuthorityID uint, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.FtpFile{})
	//var a system.SysAuthority
	//a.AuthorityId = sysUserAuthorityID
	//auth, err := systemService.AuthorityServiceApp.GetAuthorityInfo(a)
	//if err != nil {
	//	return
	//}
	//var dataId []uint
	//for _, v := range auth.DataAuthorityId {
	//	dataId = append(dataId, v.AuthorityId)
	//}
	var FileUploadList []model.FtpFile
	err = db.Count(&total).Error
	if err != nil {
		return FileUploadList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&FileUploadList).Error
	}
	return FileUploadList, total, err
}
