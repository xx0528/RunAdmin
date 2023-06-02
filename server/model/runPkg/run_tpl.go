/*
 * @Author: xx
 * @Date: 2023-05-17 15:44:34
 * @LastEditTime: 2023-05-18 20:03:25
 * @Description:
 */
// 自动生成模板RunTpl
package runPkg

import (
	"math/rand"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// RunTpl 结构体
type RunTpl struct {
	global.GVA_MODEL
	TplId      string `json:"tplId" form:"tplId" gorm:"column:tpl_id;comment:页面ID;"`
	UserId     uint   `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;"`
	TplName    string `json:"tplName" form:"tplName" gorm:"column:tpl_name;comment:页面名;"`
	TplType    int    `json:"tplType" form:"tplType" gorm:"column:tpl_type;comment:模板类型;"`
	ClickUrl   string `json:"clickUrl" form:"clickUrl" gorm:"column:click_url;comment:点击链接;"`
	PageUrl    string `json:"pageUrl" form:"pageUrl" gorm:"column:page_url;comment:展示链接;"`
	ClickDesc  string `json:"clickDesc" form:"clickDesc" gorm:"column:click_desc;comment:按钮描述;"`
	DialogDesc string `json:"dialogDesc" form:"dialogDesc" gorm:"column:dialog_desc;comment:弹窗描述;"`
	Text1      string `json:"text1" form:"text1" gorm:"column:text1;type:text;comment:描述1;"`
	Text2      string `json:"text2" form:"text2" gorm:"column:text2;type:text;comment:描述2;"`
	Text3      string `json:"text3" form:"text3" gorm:"column:text3;type:text;comment:描述3;"`
	Text4      string `json:"text4" form:"text4" gorm:"column:text4;type:text;comment:描述4;"`
	Text5      string `json:"text5" form:"text5" gorm:"column:text5;type:text;comment:描述5;"`
	Text6      string `json:"text6" form:"text6" gorm:"column:text6;type:text;comment:描述6;"`
	Text7      string `json:"text7" form:"text7" gorm:"column:text7;type:text;comment:描述7;"`
	Text8      string `json:"text8" form:"text8" gorm:"column:text8;type:text;comment:描述8;"`
	Text9      string `json:"text9" form:"text9" gorm:"column:text9;type:text;comment:描述9;"`
	Text10     string `json:"text10" form:"text10" gorm:"column:text10;type:text;comment:描述10;"`
	Text11     string `json:"text11" form:"text11" gorm:"column:text11;type:text;comment:描述11;"`
	Text12     string `json:"text12" form:"text12" gorm:"column:text12;type:text;comment:描述12;"`
	Text13     string `json:"text13" form:"text13" gorm:"column:text13;type:text;comment:描述13;"`
	Text14     string `json:"text14" form:"text14" gorm:"column:text14;type:text;comment:描述14;"`
	Text15     string `json:"text15" form:"text15" gorm:"column:text15;type:text;comment:描述15;"`
	Text16     string `json:"text16" form:"text16" gorm:"column:text16;type:text;comment:描述16;"`
	Text17     string `json:"text17" form:"text17" gorm:"column:text17;type:text;comment:描述17;"`
	Text18     string `json:"text18" form:"text18" gorm:"column:text18;type:text;comment:描述18;"`
	Text19     string `json:"text19" form:"text19" gorm:"column:text19;type:text;comment:描述19;"`
	Text20     string `json:"text20" form:"text20" gorm:"column:text20;type:text;comment:描述20;"`
	PicName1   string `json:"picName1" form:"picName1" gorm:"column:pic_name1;comment:图片1;"`
	PicName2   string `json:"picName2" form:"picName2" gorm:"column:pic_name2;comment:图片2;"`
	PicName3   string `json:"picName3" form:"picName3" gorm:"column:pic_name3;comment:图片3;"`
	PicName4   string `json:"picName4" form:"picName4" gorm:"column:pic_name4;comment:图片4;"`
	PicName5   string `json:"picName5" form:"picName5" gorm:"column:pic_name5;comment:图片5;"`
	PicName6   string `json:"picName6" form:"picName6" gorm:"column:pic_name6;comment:图片6;"`
	PicName7   string `json:"picName7" form:"picName7" gorm:"column:pic_name7;comment:图片7;"`
	PicName8   string `json:"picName8" form:"picName8" gorm:"column:pic_name8;comment:图片8;"`
	PicName9   string `json:"picName9" form:"picName9" gorm:"column:pic_name9;comment:图片9;"`
	PicName10  string `json:"picName10" form:"picName10" gorm:"column:pic_name10;comment:图片10;"`
	CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName RunTpl 表名
func (RunTpl) TableName() string {
	return "run_tpl"
}

// 根据域名生成落地页url
func (e *RunTpl) GenerateUrl() (string, string) {
	rand.Seed(time.Now().UnixNano())
	rndStr := RandStringRunes(8)
	// 获取main_page
	var sysUser system.SysUser
	if err := global.GVA_DB.Where("id = ?", e.UserId).First(&sysUser).Error; err != nil {
		return "", ""
	}

	return rndStr, sysUser.MainPage + "tpl/get?tplId=" + rndStr
}

const (
	TplType_Normal         = 1 //通用模板
	TplType_Stock_Thai     = 2 //泰国股票
	TplType_Stock_Malaysia = 3 //马来西亚股票
)
