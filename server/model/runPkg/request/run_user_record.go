/*
 * @Author: xx
 * @Date: 2023-04-24 18:55:47
 * @LastEditTime: 2023-05-08 14:11:41
 * @Description:
 */
package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/runPkg"
)

type RunUserRecordSearch struct {
	runPkg.RunUserRecord
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}
