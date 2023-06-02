/*
 * @Author: xx
 * @Date: 2023-04-24 18:55:47
 * @LastEditTime: 2023-06-01 15:51:43
 * @Description:
 */
package runPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/runPkg"
	runPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/runPkg/request"
)

type RunUserRecordService struct {
}

// CreateRunUserRecord 创建RunUserRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (runUserRecordService *RunUserRecordService) CreateRunUserRecord(runUserRecord *runPkg.RunUserRecord) (err error) {
	err = global.GVA_DB.Create(runUserRecord).Error
	return err
}

// DeleteRunUserRecord 删除RunUserRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (runUserRecordService *RunUserRecordService) DeleteRunUserRecord(runUserRecord runPkg.RunUserRecord) (err error) {
	err = global.GVA_DB.Delete(&runUserRecord).Error
	return err
}

// DeleteRunUserRecordByIds 批量删除RunUserRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (runUserRecordService *RunUserRecordService) DeleteRunUserRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]runPkg.RunUserRecord{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateRunUserRecord 更新RunUserRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (runUserRecordService *RunUserRecordService) UpdateRunUserRecord(runUserRecord runPkg.RunUserRecord) (err error) {
	err = global.GVA_DB.Save(&runUserRecord).Error
	return err
}

// GetRunUserRecord 根据id获取RunUserRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (runUserRecordService *RunUserRecordService) GetRunUserRecord(id uint) (runUserRecord runPkg.RunUserRecord, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&runUserRecord).Error
	return
}

// GetRunUserRecordInfoList 分页获取RunUserRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (runUserRecordService *RunUserRecordService) GetRunUserRecordInfoList(info runPkgReq.RunUserRecordSearch) (list []runPkg.RunUserRecord, total int64, options map[string][]string, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	//获取搜索选项
	db2 := global.GVA_DB.Model(&runPkg.RunUserRecord{})
	searchOptions := make(map[string][]string)
	var orderNames []string
	if info.UserId != 1 {
		db2.Where("user_id = ?", info.UserId).Distinct().Pluck("order_name", &orderNames)
	} else {
		db2.Distinct().Pluck("order_name", &orderNames)
	}
	db2 = global.GVA_DB.Model(&runPkg.RunUserRecord{})
	var pageNames []string
	if info.UserId != 1 {
		db2.Where("user_id = ?", info.UserId).Distinct().Pluck("page_name", &pageNames)
	} else {
		db2.Distinct().Pluck("page_name", &pageNames)
	}
	db2 = global.GVA_DB.Model(&runPkg.RunUserRecord{})
	var pageCountrys []string
	if info.UserId != 1 {
		db2.Where("user_id = ?", info.UserId).Distinct().Pluck("page_country", &pageCountrys)
	} else {
		db2.Distinct().Pluck("page_country", &pageCountrys)
	}
	searchOptions["orderName"] = orderNames
	searchOptions["pageName"] = pageNames
	searchOptions["pageCountry"] = pageCountrys

	// 创建db
	db := global.GVA_DB.Model(&runPkg.RunUserRecord{})
	var runUserRecords []runPkg.RunUserRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	//超级管理员就不用过滤了
	if info.UserId != 1 {
		db = db.Where("user_id = ?", info.UserId)
	}
	if info.OrderName != "" {
		db = db.Where("order_name LIKE ?", "%"+info.OrderName+"%")
	}
	if info.PageName != "" {
		db = db.Where("page_name LIKE ?", "%"+info.PageName+"%")
	}
	if info.PageCountry != "" {
		db = db.Where("page_country LIKE ?", "%"+info.PageCountry+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("created_at desc").Limit(limit).Offset(offset).Find(&runUserRecords).Error
	return runUserRecords, total, searchOptions, err
}
