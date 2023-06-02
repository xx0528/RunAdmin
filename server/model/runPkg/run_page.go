/*
 * @Author: xx
 * @Date: 2023-04-24 17:49:18
 * @LastEditTime: 2023-05-18 14:56:27
 * @Description:
 */
// 自动生成模板RunPage
package runPkg

import (
	"math/rand"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// RunPage 结构体
type RunPage struct {
	global.GVA_MODEL
	PageName  string `json:"pageName" form:"pageName" gorm:"column:page_name;comment:落地页名字;"`
	Url       string `json:"url" form:"url" gorm:"column:url;comment:链接;"`
	Remark    string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`
	Country   string `json:"country" form:"country" gorm:"column:country;comment:国家;"`
	UserNum   int    `json:"userNum" form:"user_num" gorm:"column:user_num;comment:用户数;"`
	UserId    uint   `json:"userId" form:"user_id" gorm:"column:user_id;comment:用户id;"`
	State     int    `json:"state" form:"state" gorm:"column:state;comment:状态;"`
	PageId    string `json:"pageId" form:"page_id" gorm:"column:page_id;comment:落地页id;"`
	CreatedBy uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName RunPage 表名
func (RunPage) TableName() string {
	return "run_page"
}

// 根据域名生成落地页url
func (e *RunPage) GenerateUrl() (string, string) {
	rand.Seed(time.Now().UnixNano())
	rndStr := RandStringRunes(8)
	// 获取main_page
	var sysUser system.SysUser
	if err := global.GVA_DB.Where("id = ?", e.UserId).First(&sysUser).Error; err != nil {
		return "", ""
	}

	return rndStr, sysUser.MainPage + "land/get?pageId=" + rndStr
}

func RandStringRunes(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
