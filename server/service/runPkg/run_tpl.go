/*
 * @Author: xx
 * @Date: 2023-05-17 15:44:34
 * @LastEditTime: 2023-05-18 16:22:40
 * @Description:
 */
package runPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/runPkg"
	runPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/runPkg/request"
	"gorm.io/gorm"
)

type RunTplService struct {
}

// CreateRunTpl 创建RunTpl记录
// Author [piexlmax](https://github.com/piexlmax)
func (runTplService *RunTplService) CreateRunTpl(runTpl *runPkg.RunTpl) (err error) {
	err = global.GVA_DB.Create(runTpl).Error
	return err
}

// DeleteRunTpl 删除RunTpl记录
// Author [piexlmax](https://github.com/piexlmax)
func (runTplService *RunTplService) DeleteRunTpl(runTpl runPkg.RunTpl) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&runPkg.RunTpl{}).Where("id = ?", runTpl.ID).Update("deleted_by", runTpl.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&runTpl).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteRunTplByIds 批量删除RunTpl记录
// Author [piexlmax](https://github.com/piexlmax)
func (runTplService *RunTplService) DeleteRunTplByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&runPkg.RunTpl{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&runPkg.RunTpl{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateRunTpl 更新RunTpl记录
// Author [piexlmax](https://github.com/piexlmax)
func (runTplService *RunTplService) UpdateRunTpl(runTpl runPkg.RunTpl) (err error) {
	err = global.GVA_DB.Save(&runTpl).Error
	return err
}

// GetRunTpl 根据id获取RunTpl记录
// Author [piexlmax](https://github.com/piexlmax)
func (runTplService *RunTplService) GetRunTpl(id uint) (runTpl runPkg.RunTpl, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&runTpl).Error
	return
}

// GetRunTplByTplId 根据tplId获取RunTpl记录
// Author [piexlmax](https://github.com/piexlmax)
func (runTplService *RunTplService) GetRunTplByTplId(tplId string) (runTpl runPkg.RunTpl, err error) {
	err = global.GVA_DB.Where("tpl_id = ?", tplId).First(&runTpl).Error
	return
}

// GetRunTplInfoList 分页获取RunTpl记录
// Author [piexlmax](https://github.com/piexlmax)
func (runTplService *RunTplService) GetRunTplInfoList(info runPkgReq.RunTplSearch) (list []runPkg.RunTpl, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// db2 := global.GVA_DB.Model(&runPkg.RunTpl{})

	// searchOptions := make(map[string][]string)
	// var tplTypeNames[]string
	// if info.UserId != 1 {
	// 	db2.Where("user_id = ?", info.UserId).Distinct().Pluck("")
	// }

	// 创建db
	db := global.GVA_DB.Model(&runPkg.RunTpl{})
	var runTpls []runPkg.RunTpl
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.UserId != 1 {
		db = db.Where("user_id = ?", info.UserId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Order("created_at desc").Limit(limit).Offset(offset).Find(&runTpls).Error
	return runTpls, total, err
}
