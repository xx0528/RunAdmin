/*
 * @Author: xx
 * @Date: 2023-04-24 17:49:18
 * @LastEditTime: 2023-06-02 08:45:37
 * @Description:
 */
package runPkg

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/runPkg"
	"github.com/gin-gonic/gin"
	"github.com/oschwald/geoip2-golang"
	"go.uber.org/zap"
)

type RunOrderNumsApi struct {
}

var rePageIdx = make(map[string]int)
var ipCityDB *geoip2.Reader

func init() {
	var err error
	ipCityDB, err = geoip2.Open("resource/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
}
func (runOrderNumsApi *RunOrderNumsApi) RedirectURLOld(c *gin.Context) {
	subUuid := c.Query("subUuid")

	runPage, err := runPageService.GetRunPageByRemark(subUuid)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	runNums, err := runNumService.GetRunNumsByPageId(runPage.PageId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	redirect(c, runPage.PageId, runNums)
}

// 重定向URL
func (runOrderNumsApi *RunOrderNumsApi) RedirectURL(c *gin.Context) {
	pageId := c.Query("pageId")
	//找到落地页绑定的所有可用号码，可能不止一个工单
	runNums, err := runNumService.GetRunNumsByPageId(pageId)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	redirect(c, pageId, runNums)

}

func redirect(c *gin.Context, pageId string, runNums []runPkg.RunNum) {

	ip := c.ClientIP()
	country := "国家未知"
	city := "城市未知"
	if ip == "127.0.0.1" {
		country = "本地"
		city = "localhost"
	} else {
		ipParse := net.ParseIP(ip)
		record, err := ipCityDB.City(ipParse)
		if err == nil {
			country = record.Country.Names["zh-CN"]
			city = record.City.Names["en"]
			if len(city) > 0 {
				city = fmt.Sprintf("-%s", city)
			}
		}
	}

	runPage, errRunPage := runPageService.GetRunPageByPageId(pageId)

	if len(runNums) <= 0 {
		// response.FailWithMessage("nothing", c)
		pName := runPage.PageName
		pCountry := runPage.Country
		if errRunPage != nil {
			pName = fmt.Sprintf("无此落地页 pageId=%s", pageId)
			pCountry = fmt.Sprintf("未知工单国家 pageId=%s", pageId)
		}
		// if err == nil {
		// 	warningMsg := runPkg.DDMsgCfg{
		// 		IsAtAll: false,
		// 		Msg:     fmt.Sprintf("重要提示:\n落地页:%s\n链接:%s\n无可用号码，仍然在进粉", runPage.PageName, runPage.Url),
		// 	}
		// 	reqUser, err := systemService.UserServiceApp.FindUserById(int(runPage.UserId))
		// 	if err != nil {
		// 		warningMsg.AtMobiles = append(warningMsg.AtMobiles, reqUser.Phone)
		// 	}
		// 	errSend := notifyService.ServiceGroupApp.SendTextMessage(warningMsg.Msg, warningMsg.AtMobiles, warningMsg.IsAtAll)
		// 	if errSend != nil {
		// 		fmt.Println("发送错误--", errSend.Error())
		// 	}
		// }

		//记录进粉
		var runUserRecord = runPkg.RunUserRecord{
			IpAddr:      ip,
			Country:     fmt.Sprintf("%s%s", country, city),
			PageId:      pageId,
			PageName:    pName,
			PageCountry: pCountry,
			OrderName:   "未知工单",
			UserId:      runPage.UserId,
			Num:         "无号码",
		}
		runUserRecordService.CreateRunUserRecord(&runUserRecord)
		return
	}
	//获取到跳转到第几个
	if _, ok := rePageIdx[pageId]; ok {
		rePageIdx[pageId]++
		if rePageIdx[pageId] >= len(runNums) {
			rePageIdx[pageId] = 0
		}
	} else {
		rePageIdx[pageId] = 0
	}
	numInfo := runNums[rePageIdx[pageId]]
	redirectURL := ""
	// 如果是whatsapp 群 那号码填的就是地址 直接转向这个地址
	if numInfo.NumType == runPkg.WhatsAppGroup {
		redirectURL = numInfo.Num
	} else {
		newUrl := getUrlByNumType(numInfo.NumType, numInfo.SayHi)
		if len(newUrl) == 0 {
			err := errors.New("未找到号码类型")
			global.GVA_LOG.Error("未找到号码类型", zap.Error(err))
			response.FailWithMessage(err.Error(), c)

			//记录进粉
			var runUserRecord = runPkg.RunUserRecord{
				IpAddr:      ip,
				Country:     fmt.Sprintf("%s%s", country, city),
				PageId:      pageId,
				PageName:    numInfo.PageName,
				PageCountry: runPage.Country,
				OrderName:   numInfo.OrderName,
				UserId:      numInfo.UserId,
				Num:         numInfo.Num + "-无号码类型",
			}
			runUserRecordService.CreateRunUserRecord(&runUserRecord)
			return
		}

		// 根据访问的域名和pageId的值，构造要跳转的链接
		if len(numInfo.SayHi) > 0 {
			redirectURL = fmt.Sprintf(newUrl, numInfo.Num, numInfo.SayHi)
		} else {
			redirectURL = fmt.Sprintf(newUrl, numInfo.Num)
		}
	}

	c.Redirect(http.StatusFound, redirectURL)

	//记录进粉
	var runUserRecord = runPkg.RunUserRecord{
		IpAddr:      ip,
		Country:     fmt.Sprintf("%s%s", country, city),
		PageId:      pageId,
		PageName:    numInfo.PageName,
		PageCountry: runPage.Country,
		OrderName:   numInfo.OrderName,
		UserId:      numInfo.UserId,
		Num:         numInfo.Num,
	}
	runUserRecordService.CreateRunUserRecord(&runUserRecord)
}

func getUrlByNumType(numType int, sayHi string) string {
	if len(sayHi) > 0 {
		switch numType {
		case runPkg.WhatsApp:
			return "https://wa.me/%s?text=%s"
		case runPkg.LINE:
			return "https://line.me/ti/p/~%s?text=%s"
		case runPkg.Telegram:
			return "https://t.me/%s?text=%s"
		case runPkg.Zalo:
			return "https://zalo.me/%s?text=%s"
		default:
			return ""
		}
	} else {
		switch numType {
		case runPkg.WhatsApp:
			return "https://wa.me/%s"
		case runPkg.LINE:
			return "https://line.me/ti/p/~%s"
		case runPkg.Telegram:
			return "https://t.me/%s"
		case runPkg.Zalo:
			return "https://zalo.me/%s"
		default:
			return ""
		}
	}

}
