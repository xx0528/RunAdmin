/*
 * @Author: xx
 * @Date: 2023-04-24 17:49:18
 * @LastEditTime: 2023-05-15 20:24:44
 * @Description:
 */
package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/runPkg"
)

type RunPageSearch struct {
	runPkg.RunPage
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type RunOrderUrlInfo struct {
	Url string `json:"url"` // 工单地址
	Psw string `json:"psw"` //工单密码
}

type RunCreateNumsInfo struct {
	ID           uint     `json:"id"`
	UserId       uint     `json:"userID"`
	NumType      int      `json:"numType"`
	OrderName    string   `json:"orderName"`
	OrderUrl     string   `json:"orderUrl"`
	OrderUrlType string   `json:"orderUrlType"`
	PageName     string   `json:"pageName"`
	PageId       string   `json:"pageId"`
	SayHi        string   `json:"SayHi"`
	MaxEnterNum  int      `json:"maxEnterNum"`
	EachEnterNum int      `json:"eachEnterNum"`
	Nums         []string `json:"nums"`
}

type RunPageNameInfos struct {
	PageID   string `json:"pageId"`
	PageName string `json:"pageName"`
}
