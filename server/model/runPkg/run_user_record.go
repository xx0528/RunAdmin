/*
 * @Author: xx
 * @Date: 2023-04-24 18:55:47
 * @LastEditTime: 2023-06-01 15:31:25
 * @Description:
 */
// 自动生成模板RunUserRecord
package runPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// RunUserRecord 结构体
type RunUserRecord struct {
	global.GVA_MODEL
	IpAddr      string `json:"ipAddr" form:"ipAddr" gorm:"column:ip_addr;comment:ip地址;"`
	Country     string `json:"country" form:"country" gorm:"column:country;comment:所属国家;"`
	PageId      string `json:"pageId" form:"pageId" gorm:"column:page_id;comment:落地页ID;"`
	PageName    string `json:"pageName" form:"pageName" gorm:"column:page_name;comment:落地页名;"`
	OrderName   string `json:"orderName" form:"orderName" gorm:"column:order_name;comment:工单名;"`
	PageCountry string `json:"pageCountry" form:"pageCountry" gorm:"column:page_country;comment:国家名;"`
	UserId      uint   `json:"userId" form:"user_id" gorm:"column:user_id;comment:用户id;"`
	Num         string `json:"num" form:"num" gorm:"column:num;comment:对应号码;"`
}

// TableName RunUserRecord 表名
func (RunUserRecord) TableName() string {
	return "run_user_record"
}
