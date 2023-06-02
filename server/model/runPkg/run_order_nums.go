/*
 * @Author: xx
 * @Date: 2023-04-27 16:03:57
 * @LastEditTime: 2023-05-26 15:49:42
 * @Description:
 */
package runPkg

type RunSaveOrderNum struct {
	Num       string `json:"num"`
	PageId    string `json:"pageId"`
	State     int    `json:"state"`
	UserNum   int    `json:"userNum"`
	UserLimit int    `json:"userLimit"`
	NumType   int    `json:"numType"`
}

type RunSaveOrderInfo struct {
	UserID       uint              `json:"userID"`
	OrderUrl     string            `json:"orderUrl"`
	OrderName    string            `json:"orderName"`
	UserAllLimit int               `json:"userAllLimit"`
	NumList      []RunSaveOrderNum `json:"nums"`
}

// 监控工单的状态
type RunSaveOrderState struct {
	Warning70    bool              `json:"Warning70"`
	Warning90    bool              `json:"Warning90"`
	Warning100   bool              `json:"Warning100"`
	IntoAllFuns  float64           `json:"intoAllFuns"`
	LostTimes    int               `json:"lostTimes"`
	LostIsNotice bool              `json:"lostIsNotice"`
	NumList      []RunSaveOrderNum `json:"nums"`
}

type DDMsgCfg struct {
	Msg       string
	AtMobiles []string
	IsAtAll   bool
}

// 工单链接类型
const (
	OrderType_Share = "share"
	OrderType_URL   = "url"
	OrderType_KF007 = "kf007"
	OrderType_GooSu = "goosu"
	OrderType_XYZ   = "xyz"
)

// 号码类型
const (
	WhatsApp      = 1
	LINE          = 2
	Telegram      = 3
	Zalo          = 4
	WhatsAppGroup = 5
)

const (
	NumType_OffLine = 0 //离线
	NumType_OnLine  = 1 //在线
	NumType_Lock    = 2 //封号
	NumType_Freeze  = 3 //冻结
	NumType_Lost    = 4 //丢失
)
