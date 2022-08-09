package face

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/face"
)

type FacePhotoService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateFacePhoto
//@description: 创建图片
//@param: e model.FacePhoto
//@return: err error

func (exa *FacePhotoService) CreateFacePhoto(e face.FacePhoto) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除图片
//@param: e model.FacePhoto
//@return: err error

func (exa *FacePhotoService) DeleteFacePhoto(e face.FacePhoto) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateFacePhoto
//@description: 更新图片
//@param: e *model.FacePhoto
//@return: err error

func (exa *FacePhotoService) UpdateFacePhoto(e *face.FacePhoto) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFacePhoto
//@description: 获取图片信息
//@param: id uint
//@return: FacePhoto model.FacePhoto, err error

func (exa *FacePhotoService) GetFacePhoto(id uint) (FacePhoto face.FacePhoto, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&FacePhoto).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFacePhotoInfoList
//@description: 分页获取图片列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *FacePhotoService) GetFacePhotoInfoList(sysUserAuthorityID uint, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&face.FacePhoto{})
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
	var FacePhotoList []face.FacePhoto
	err = db.Count(&total).Error
	if err != nil {
		return FacePhotoList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&FacePhotoList).Error
	}
	return FacePhotoList, total, err
}
