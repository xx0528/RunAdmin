/*
 * @Author: xx
 * @Date: 2023-05-09 10:23:48
 * @LastEditTime: 2023-05-12 13:54:28
 * @Description:
 */
// 自动生成模板RunOrder
package runPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// RunOrder 结构体
type RunOrder struct {
	global.GVA_MODEL
	OrderName    string `json:"orderName" form:"orderName" gorm:"column:order_name;comment:工单名;"`
	PageId       string `json:"pageId" form:"pageId" gorm:"column:page_id;comment:落地页ID;"`
	UserId       uint   `json:"userId" form:"userId" gorm:"column:user_id;comment:用户ID;"`
	PageName     string `json:"pageName" form:"pageName" gorm:"column:page_name;comment:落地页名;"`
	MaxEnterNum  int    `json:"maxEnterNum" form:"maxEnterNum" gorm:"column:max_enter_num;comment:进粉限制;"`
	EachEnterNum int    `json:"eachEnterNum" form:"eachEnterNum" gorm:"column:each_enter_num;comment:平均进粉;"`
	UserNum      int    `json:"userNum" form:"userNum" gorm:"column:user_num;comment:绑定号码数;"`
	CreatedBy    uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName RunOrder 表名
func (RunOrder) TableName() string {
	return "run_order"
}
