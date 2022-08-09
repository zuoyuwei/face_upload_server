package face

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/face"
)

type FaceVideoService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateFaceVideo
//@description: 创建视频
//@param: e model.FaceVideo
//@return: err error

func (exa *FaceVideoService) CreateFaceVideo(e face.FaceVideo) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除视频
//@param: e model.FaceVideo
//@return: err error

func (exa *FaceVideoService) DeleteFaceVideo(e face.FaceVideo) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateFaceVideo
//@description: 更新视频
//@param: e *model.FaceVideo
//@return: err error

func (exa *FaceVideoService) UpdateFaceVideo(e *face.FaceVideo) (err error) {
	err = global.GVA_DB.Updates(e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFaceVideo
//@description: 获取视频信息
//@param: id uint
//@return: FaceVideo model.FaceVideo, err error

func (exa *FaceVideoService) GetFaceVideo(id uint) (FaceVideo face.FaceVideo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&FaceVideo).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFaceVideoInfoList
//@description: 分页获取视频列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *FaceVideoService) GetFaceVideoInfoList(sysUserAuthorityID uint, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&face.FaceVideo{})
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
	var FaceVideoList []face.FaceVideo
	err = db.Count(&total).Error
	if err != nil {
		return FaceVideoList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&FaceVideoList).Error
	}
	return FaceVideoList, total, err
}
