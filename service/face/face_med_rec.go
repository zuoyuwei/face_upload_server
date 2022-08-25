package face

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/face"
)

type FaceMedicalRecordService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateFaceMedicalRecord
//@description: 创建病历
//@param: e model.FaceMedicalRecord
//@return: err error

func (exa *FaceMedicalRecordService) CreateFaceMedicalRecord(e face.FaceMedicalRecord) (id int, err error) {
	err = global.GVA_DB.Create(&e).Error
	return int(e.ID), err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除病历
//@param: e model.FaceMedicalRecord
//@return: err error

func (exa *FaceMedicalRecordService) DeleteFaceMedicalRecord(e face.FaceMedicalRecord) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateFaceMedicalRecord
//@description: 更新病历
//@param: e *model.FaceMedicalRecord
//@return: err error

func (exa *FaceMedicalRecordService) UpdateFaceMedicalRecord(e *face.FaceMedicalRecord) (err error) {
	err = global.GVA_DB.Save(e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFaceMedicalRecord
//@description: 获取病历信息
//@param: id uint
//@return: FaceMedicalRecord model.FaceMedicalRecord, err error

func (exa *FaceMedicalRecordService) GetFaceMedicalRecord(id uint) (FaceMedicalRecord face.FaceMedicalRecord, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&FaceMedicalRecord).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: FindFaceMedicalRecord
//@description: 根据病人编号查询病历信息
//@param: patientcode string
//@return: FaceMedicalRecord model.FaceMedicalRecord, err error

func (exa *FaceMedicalRecordService) FindFaceMedicalRecord(medicalRecordId int) (FaceMedicalRecord face.FaceMedicalRecord, err error) {
	err = global.GVA_DB.Where("id = ?", medicalRecordId).First(&FaceMedicalRecord).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: FindFaceMedicalRecord_Son
//@description: 根据患者编号查询病历信息
//@param: patientcode string
//@return: FaceMedicalRecord model.FaceMedicalRecord, err error

func (exa *FaceMedicalRecordService) FindFaceMedicalRecord_Son(patientcode string) (FaceMedicalRecord face.FaceMedicalRecord, err error) {
	err = global.GVA_DB.Where("patient_code = ?", patientcode).First(&FaceMedicalRecord).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetFaceMedicalRecordInfoList
//@description: 分页获取病历列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *FaceMedicalRecordService) GetFaceMedicalRecordInfoList(sysUserAuthorityID uint, info request.PageInfo) (list interface{}, total int64, err error) {
	return GetFaceMedicalRecordInfoList_Son(sysUserAuthorityID, info)
}

func (exa *FaceMedicalRecordService) GetFaceMedicalRecordInfoList_A(sysUserAuthorityID uint, info request.PageInfo) (list []face.FaceMedicalRecord, total int64, err error) {
	return GetFaceMedicalRecordInfoList_Son(sysUserAuthorityID, info)
}

func GetFaceMedicalRecordInfoList_Son(sysUserAuthorityID uint, info request.PageInfo) (list []face.FaceMedicalRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&face.FaceMedicalRecord{})
	if info.GuanLianId > 0 {
		db.Where("patient_id=?", info.GuanLianId)
	}
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
	var FaceMedicalRecordList []face.FaceMedicalRecord
	err = db.Count(&total).Error
	if err != nil {
		return FaceMedicalRecordList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&FaceMedicalRecordList).Error
	}
	return FaceMedicalRecordList, total, err
}
