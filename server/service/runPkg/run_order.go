package runPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/runPkg"
	runPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/runPkg/request"
	"gorm.io/gorm"
)

type RunOrderService struct {
}

// CreateRunOrder 创建RunOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (runOrderService *RunOrderService) CreateRunOrder(runOrder *runPkg.RunOrder) (err error) {
	err = global.GVA_DB.Create(runOrder).Error
	return err
}

// DeleteRunOrder 删除RunOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (runOrderService *RunOrderService) DeleteRunOrder(runOrder runPkg.RunOrder) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&runPkg.RunOrder{}).Where("id = ?", runOrder.ID).Update("deleted_by", runOrder.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&runOrder).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteRunOrderByIds 批量删除RunOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (runOrderService *RunOrderService) DeleteRunOrderByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&runPkg.RunOrder{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&runPkg.RunOrder{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (runOrderService *RunOrderService) GetRunOrderByIds(ids request.IdsReq) (results []runPkg.RunOrder, err error) {
	db := global.GVA_DB.Model(&runPkg.RunOrder{})
	err = db.Where("id in (?)", ids.Ids).Find(&results).Error
	return
}

// UpdateRunOrder 更新RunOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (runOrderService *RunOrderService) UpdateRunOrder(runOrder runPkg.RunOrder) (err error) {
	err = global.GVA_DB.Save(&runOrder).Error
	return err
}

// UpdateRunPage 更新RunPage记录
// Author [piexlmax](https://github.com/piexlmax)
func (runOrderService *RunOrderService) UpdateRunOrderUsers(runOrder runPkg.RunOrder) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&runPkg.RunOrder{}).Where("order_name = ? AND user_id = ?", runOrder.OrderName, runOrder.UserId).Update("user_num", runOrder.UserNum).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// GetRunOrder 根据id获取RunOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (runOrderService *RunOrderService) GetRunOrder(id uint) (runOrder runPkg.RunOrder, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&runOrder).Error
	return
}

// GetRunOrder 根据userId获取RunOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (runOrderService *RunOrderService) GetRunOrdersByUserId(id uint) (orderInfos []string, err error) {
	db := global.GVA_DB.Model(&runPkg.RunOrder{})
	//获取搜索选项
	if id != 1 {
		err = db.Where("user_id = ?", id).Distinct().Select("order_name").Scan(&orderInfos).Error
	} else {
		err = db.Distinct().Select("order_name").Scan(&orderInfos).Error
	}

	return
}

// GetRunOrderByName 根据name获取RunOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (runOrderService *RunOrderService) GetRunOrderByName(name string, userId uint) (runOrder runPkg.RunOrder, err error) {
	err = global.GVA_DB.Where("order_name = ? AND user_id = ? ", name, userId).First(&runOrder).Error
	return
}

// GetRunOrderInfoList 分页获取RunOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (runOrderService *RunOrderService) GetRunOrderInfoList(info runPkgReq.RunOrderSearch) (list []runPkg.RunOrder, total int64, options map[string][]string, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db2 := global.GVA_DB.Model(&runPkg.RunOrder{})
	//获取搜索选项
	searchOptions := make(map[string][]string)
	var pageNames []string
	if info.UserId != 1 {
		db2.Where("user_id = ?", info.UserId).Distinct().Pluck("page_name", &pageNames)
	} else {
		db2.Distinct().Pluck("page_name", &pageNames)
	}
	db2 = global.GVA_DB.Model(&runPkg.RunOrder{})
	var orderNames []string
	if info.UserId != 1 {
		db2.Where("user_id = ?", info.UserId).Distinct().Pluck("order_name", &orderNames)
	} else {
		db2.Distinct().Pluck("order_name", &orderNames)
	}
	searchOptions["pageName"] = pageNames
	searchOptions["orderName"] = orderNames

	// 创建db
	db := global.GVA_DB.Model(&runPkg.RunOrder{})
	var runOrders []runPkg.RunOrder
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

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Order("created_at desc").Limit(limit).Offset(offset).Find(&runOrders).Error
	return runOrders, total, searchOptions, err
}
