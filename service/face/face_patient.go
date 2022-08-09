package face

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/face"
)

type FacePatientService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateFacePatient
//@description: 创建患者
//@param: e model.FacePatient
//@return: err error

func (exa *FacePatientService) CreateFacePatient(e face.FacePatient) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除患者
//@param: e model.FacePatient
//@return: err error

func (exa *FacePatientService) DeleteFacePatient(e face.FacePatient) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateFacePatient
//@description: 更新患者
//@param: e *model.FacePatient
//@return: err error

func (exa *FacePatientService) UpdateFacePatient(e *face.FacePatient) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFacePatient
//@description: 获取患者信息
//@param: id uint
//@return: FacePatient model.FacePatient, err error

func (exa *FacePatientService) GetFacePatient(id uint) (FacePatient face.FacePatient, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&FacePatient).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFacePatientInfoList
//@description: 分页获取患者列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *FacePatientService) GetFacePatientInfoList(sysUserAuthorityID uint, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&face.FacePatient{})
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
	var FacePatientList []face.FacePatient
	err = db.Count(&total).Error
	if err != nil {
		return FacePatientList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&FacePatientList).Error
	}
	return FacePatientList, total, err
}
