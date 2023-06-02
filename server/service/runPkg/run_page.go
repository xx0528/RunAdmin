/*
 * @Author: xx
 * @Date: 2023-04-24 17:49:18
 * @LastEditTime: 2023-05-18 20:25:00
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

type RunPageService struct {
}

// CreateRunPage 创建RunPage记录
// Author [piexlmax](https://github.com/piexlmax)
func (runPageService *RunPageService) CreateRunPage(runPage *runPkg.RunPage) (err error) {
	err = global.GVA_DB.Create(runPage).Error
	return err
}

// DeleteRunPage 删除RunPage记录
// Author [piexlmax](https://github.com/piexlmax)
func (runPageService *RunPageService) DeleteRunPage(runPage runPkg.RunPage) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&runPkg.RunPage{}).Where("id = ?", runPage.ID).Update("deleted_by", runPage.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&runPage).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteRunPageByIds 批量删除RunPage记录
// Author [piexlmax](https://github.com/piexlmax)
func (runPageService *RunPageService) DeleteRunPageByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&runPkg.RunPage{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&runPkg.RunPage{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateRunPage 更新RunPage记录
// Author [piexlmax](https://github.com/piexlmax)
func (runPageService *RunPageService) UpdateRunPage(runPage runPkg.RunPage) (err error) {
	err = global.GVA_DB.Save(&runPage).Error
	return err
}

// UpdateRunPage 更新RunPage记录
// Author [piexlmax](https://github.com/piexlmax)
func (runPageService *RunPageService) UpdateRunPageUsers(runPage runPkg.RunPage) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&runPkg.RunPage{}).Where("page_id = ?", runPage.PageId).Update("user_num", runPage.UserNum).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// GetRunPage 根据id获取RunPage记录
// Author [piexlmax](https://github.com/piexlmax)
func (runPageService *RunPageService) GetRunPage(id uint) (runPage runPkg.RunPage, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&runPage).Error
	return
}

// GetRunPage 根据userId获取RunPage记录
// Author [piexlmax](https://github.com/piexlmax)
func (runPageService *RunPageService) GetRunPageNamesByUserId(id uint) (pageInfos []runPkgReq.RunPageNameInfos, err error) {
	db := global.GVA_DB.Model(&runPkg.RunPage{})
	//获取搜索选项
	if id != 1 {
		// err = db.Where("user_id = ?", id).Distinct().Pluck("page_name", &pageNames).Error
		err = db.Where("user_id = ?", id).Distinct().Select("page_id, page_name").Scan(&pageInfos).Error
	} else {
		// err = db.Distinct().Pluck("page_name", &pageNames).Error
		err = db.Distinct().Select("page_id, page_name").Scan(&pageInfos).Error
	}

	return
}

// GetRunPage 根据pageId获取RunPage记录
// Author [piexlmax](https://github.com/piexlmax)
func (runPageService *RunPageService) GetRunPageByPageId(pageId string) (runPage runPkg.RunPage, err error) {
	err = global.GVA_DB.Where("page_id = ?", pageId).First(&runPage).Error
	return
}

// GetRunPage 根据id获取RunPage记录
// Author [piexlmax](https://github.com/piexlmax)
func (runPageService *RunPageService) GetRunPageByRemark(remark string) (runPage runPkg.RunPage, err error) {
	err = global.GVA_DB.Where("remark = ?", remark).First(&runPage).Error
	return
}

// 这个获取到的不正确 同名只能获取到一个 弃用
// GetRunPage 根据id获取RunPage记录
// Author [piexlmax](https://github.com/piexlmax)
func (runPageService *RunPageService) GetRunPageByName(runPageInfo runPkg.RunPage) (runPage runPkg.RunPage, err error) {
	err = global.GVA_DB.Where("page_name = ? AND user_id = ?", runPageInfo.PageName, runPageInfo.UserId).First(&runPage).Error
	return
}

// GetRunPageInfoList 分页获取RunPage记录
// Author [piexlmax](https://github.com/piexlmax)
func (runPageService *RunPageService) GetRunPageInfoList(info runPkgReq.RunPageSearch) (list []runPkg.RunPage, total int64, options map[string][]string, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db2 := global.GVA_DB.Model(&runPkg.RunPage{})
	//获取搜索选项
	searchOptions := make(map[string][]string)
	var pageNames []string
	if info.UserId != 1 {
		db2.Where("user_id = ?", info.UserId).Distinct().Pluck("page_name", &pageNames)
	} else {
		db2.Distinct().Pluck("page_name", &pageNames)
	}
	db2 = global.GVA_DB.Model(&runPkg.RunPage{})
	var countryNames []string
	if info.UserId != 1 {
		db2.Where("user_id = ?", info.UserId).Distinct().Pluck("country", &countryNames)
	} else {
		db2.Distinct().Pluck("country", &countryNames)
	}
	searchOptions["pageName"] = pageNames
	searchOptions["countryName"] = countryNames

	// 创建db
	db := global.GVA_DB.Model(&runPkg.RunPage{})
	var runPages []runPkg.RunPage
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	//超级管理员就不用过滤了
	if info.UserId != 1 {
		db = db.Where("user_id = ?", info.UserId)
	}
	if info.PageName != "" {
		db = db.Where("page_name LIKE ?", "%"+info.PageName+"%")
	}
	if info.State != -1 {
		db = db.Where("state = ?", info.State)
	}
	if info.Country != "" {
		db = db.Where("country LIKE ?", "%"+info.Country+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("created_at desc").Limit(limit).Offset(offset).Find(&runPages).Error
	return runPages, total, searchOptions, err
}
