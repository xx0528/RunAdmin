/*
 * @Author: xx
 * @Date: 2023-04-24 18:50:51
 * @LastEditTime: 2023-05-15 18:58:23
 * @Description:
 */
// 自动生成模板RunNum
package runPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// RunNum 结构体
type RunNum struct {
	global.GVA_MODEL
	Num          string `json:"num" form:"num" gorm:"column:num;comment:号码;"`
	State        int    `json:"state" form:"state" gorm:"column:state;comment:状态;"`
	NumType      int    `json:"numType" form:"num_type" gorm:"column:num_type;comment:号码类型;"`
	UserId       uint   `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;"`
	UserNum      uint   `json:"userNum" form:"user_num" gorm:"column:user_num;comment:用户数;"`
	PageId       string `json:"pageId" form:"page_id" gorm:"column:page_id;comment:落地页id;"`
	PageName     string `json:"pageName" form:"page_name" gorm:"column:page_name;comment:落地页名;"`
	SayHi        string `json:"sayHi" form:"sayHi" gorm:"column:say_hi;comment:打招呼;"`
	OrderName    string `json:"orderName" form:"order_name" gorm:"column:order_name;comment:工单名字;"`
	EachEnterNum int    `json:"eachEnterNum" form:"each_enter_num" gorm:"column:each_enter_num;comment:平均进粉数;"`
	CreatedBy    uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName RunNum 表名
func (RunNum) TableName() string {
	return "run_num"
}
