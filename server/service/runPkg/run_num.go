/*
 * @Author: xx
 * @Date: 2023-04-24 18:50:51
 * @LastEditTime: 2023-06-26 16:36:32
 * @Description:
 */
/*
 * @Author: xx
 * @Date: 2023-04-24 18:50:51
 * @LastEditTime: 2023-04-28 18:28:54
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

type RunNumService struct {
}

// CreateRunNum 创建RunNum记录
// Author [piexlmax](https://github.com/piexlmax)
func (runNumService *RunNumService) CreateRunNum(runNum *runPkg.RunNum) (err error) {
	err = global.GVA_DB.Create(runNum).Error
	return err
}

// DeleteRunNum 删除RunNum记录
// Author [piexlmax](https://github.com/piexlmax)
func (runNumService *RunNumService) DeleteRunNum(runNum runPkg.RunNum) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&runPkg.RunNum{}).Where("id = ?", runNum.ID).Update("deleted_by", runNum.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&runNum).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteRunNum 删除RunNum记录
// Author [piexlmax](https://github.com/piexlmax)
func (runNumService *RunNumService) DeleteRunNumByOrder(runOrder runPkg.RunOrder) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&runPkg.RunNum{}).Where("order_name = ? AND user_id = ?", runOrder.OrderName, runOrder.UserId).Update("deleted_by", runOrder.DeletedBy).Error; err != nil {
			return err
		}

		if err = tx.Where("order_name = ? AND user_id = ?", runOrder.OrderName, runOrder.UserId).Delete(&runPkg.RunNum{}).Error; err != nil {
			return err
		}

		return nil
	})
	return err
}

// DeleteRunNumByIds 批量删除RunNum记录
// Author [piexlmax](https://github.com/piexlmax)
func (runNumService *RunNumService) DeleteRunNumByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&runPkg.RunNum{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&runPkg.RunNum{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateRunNum 更新RunNum记录
// Author [piexlmax](https://github.com/piexlmax)
func (runNumService *RunNumService) UpdateRunNum(runNum runPkg.RunNum) (err error) {
	err = global.GVA_DB.Save(&runNum).Error
	return err
}

// UpdateRunNum 更新RunNum记录
// Author [piexlmax](https://github.com/piexlmax)
func (runNumService *RunNumService) UpdateRunNumState(runOrder runPkg.RunOrder) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&runPkg.RunNum{}).Where("page_id = ? AND order_name = ? AND user_id = ?", runOrder.PageId, runOrder.OrderName, runOrder.UserId).Update("state", runOrder.State).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// GetRunNum 根据id获取RunNum记录
// Author [piexlmax](https://github.com/piexlmax)
func (runNumService *RunNumService) GetRunNum(id uint) (runNum runPkg.RunNum, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&runNum).Error
	return
}

func (runNumService *RunNumService) GetRunNumByIds(ids request.IdsReq) (results []runPkg.RunNum, err error) {
	db := global.GVA_DB.Model(&runPkg.RunNum{})
	err = db.Where("id in (?)", ids.Ids).Find(&results).Error
	return
}

// GetRunNumInfoList 分页获取RunNum记录
// Author [piexlmax](https://github.com/piexlmax)
func (runNumService *RunNumService) GetRunNumInfoList(info runPkgReq.RunNumSearch) (list []runPkg.RunNum, total int64, options map[string][]string, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db2 := global.GVA_DB.Model(&runPkg.RunNum{})
	//获取搜索选项
	searchOptions := make(map[string][]string)
	var orderNames []string
	if info.UserId != 1 {
		db2.Where("user_id = ?", info.UserId).Distinct().Pluck("order_name", &orderNames)
	} else {
		db2.Distinct().Pluck("order_name", &orderNames)
	}

	db2 = global.GVA_DB.Model(&runPkg.RunNum{})
	var pageNames []string
	if info.UserId != 1 {
		db2.Where("user_id = ?", info.UserId).Distinct().Pluck("page_name", &pageNames)
	} else {
		db2.Distinct().Pluck("page_name", &pageNames)
	}
	searchOptions["orderName"] = orderNames
	searchOptions["pageName"] = pageNames

	// 创建db
	db := global.GVA_DB.Model(&runPkg.RunNum{})
	var runNums []runPkg.RunNum
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ? AND user_id = ?", info.StartCreatedAt, info.EndCreatedAt)
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
	if info.State != -1 {
		db = db.Where("state = ?", info.State)
	}
	if len(info.SearchNum) > 0 {
		db = db.Where("num = ?", info.SearchNum)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Order("created_at desc").Limit(limit).Offset(offset).Find(&runNums).Error
	return runNums, total, searchOptions, err
}

func (runNumService *RunNumService) GetRunNumsByPageId(pageId string) (list []runPkg.RunNum, err error) {
	db := global.GVA_DB.Model(&runPkg.RunNum{})
	var runNums []runPkg.RunNum
	err = db.Where("page_id = ? AND state = 1", pageId).Find(&runNums).Error
	return runNums, err
}

func (runNumService *RunNumService) GetAllRunNumsByPageId(pageId string) (list []runPkg.RunNum, err error) {
	db := global.GVA_DB.Model(&runPkg.RunNum{})
	var runNums []runPkg.RunNum
	err = db.Where("page_id = ?", pageId).Find(&runNums).Error
	return runNums, err
}

func (runNumService *RunNumService) GetAllRunNumsByOrderName(orderName string, userId uint) (list []runPkg.RunNum, err error) {
	db := global.GVA_DB.Model(&runPkg.RunNum{})
	var runNums []runPkg.RunNum
	err = db.Where("order_name = ? AND user_id = ?", orderName, userId).Find(&runNums).Error
	return runNums, err
}

func (runNumService *RunNumService) CreateRunOrderNums(createNumsInfo runPkgReq.RunCreateNumsInfo) (err error) {
	runNums := make([]*runPkg.RunNum, len(createNumsInfo.Nums))
	index := 0
	for _, num := range createNumsInfo.Nums {
		if len(num) == 0 {
			global.GVA_LOG.Error("有空号码!")
			continue
		}
		runNum := &runPkg.RunNum{
			Num:          num,
			NumType:      createNumsInfo.NumType,
			UserId:       createNumsInfo.UserId,
			PageId:       createNumsInfo.PageId,
			PageName:     createNumsInfo.PageName,
			OrderName:    createNumsInfo.OrderName,
			SayHi:        createNumsInfo.SayHi,
			UserNum:      0,
			State:        1,
			EachEnterNum: createNumsInfo.EachEnterNum,
		}
		runNums[index] = runNum
		index++
	}

	err = global.GVA_DB.Create(&runNums).Error
	return err
}
