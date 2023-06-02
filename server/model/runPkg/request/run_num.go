/*
 * @Author: xx
 * @Date: 2023-04-24 18:50:51
 * @LastEditTime: 2023-05-17 14:43:17
 * @Description:
 */
package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/runPkg"
)

type RunNumSearch struct {
	runPkg.RunNum
	SearchNum      string     `json:"searchNum" form:"searchNum"`
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type RunCreateInfo struct {
	runPkg.RunNum
	Nums []string `json:"nums" form:"nums"`
}

type RunNumUpdateByIds struct {
	runPkg.RunNum
	Ids []int `json:"ids" form:"ids"`
}
