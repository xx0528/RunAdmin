/*
 * @Author: xx
 * @Date: 2023-04-27 15:59:41
 * @LastEditTime: 2023-06-26 16:29:43
 * @Description:
 */
package runPkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/runPkg"
	runPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/runPkg/request"
	notifyService "github.com/flipped-aurora/gin-vue-admin/server/plugin/notify/service"
	systemService "github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"go.uber.org/zap"
)

type RunOrderNumsService struct {
}

var orderInfoMap = make(map[string][]runPkg.RunSaveOrderInfo)  //监控每个人的
var orderStateMap = make(map[string]*runPkg.RunSaveOrderState) //监控工单状态的
var orderWarnNums = make(map[string]int)
var orderInfoMapMutex sync.Mutex

// var mutex = &sync.Mutex{}

// 处理数据
func (runOrderNumsService *RunOrderNumsService) processOrderNumsData(content map[string]interface{}, url string) {
	intoAllFuns := content["intoAllFuns"].(float64)
	// fmt.Println("工单所有进粉数-- ", intoAllFuns)
	list := content["list"].([]interface{})
	// allFans := 0

	//工单进粉限制 取所有人设置的最大值
	orderMaxLimit := 0

	allMsg := ""
	var addedList []string
	var deletedList []string
	//这个工单第一次矫正 如果没有先存下数据
	orderState, ok := orderStateMap[url]
	if !ok {
		addOrderStateMap(url, intoAllFuns, list)
	} else {
		//工单里总粉变少了 都要通知
		if intoAllFuns < orderState.IntoAllFuns {
			allMsg = fmt.Sprintf("总进粉数减少 \n减少前: %d\n现在: %d\n减少: %d\n\n", int(orderState.IntoAllFuns), int(intoAllFuns), int(orderState.IntoAllFuns-intoAllFuns))
		}
		orderStateMap[url].IntoAllFuns = intoAllFuns
		addedList, deletedList = compareNumsThanAdjust(orderStateMap[url].NumList, list)
		if len(addedList) > 0 {
			allMsg = fmt.Sprintf("%s增加了号码:\n", allMsg)
			for _, num := range addedList {
				allMsg = fmt.Sprintf("%s%s\n", allMsg, num)
			}
			allMsg = fmt.Sprintf("%s\n", allMsg)
		}
		if len(deletedList) > 0 {
			allMsg = fmt.Sprintf("%s减少了号码:\n", allMsg)
			for _, num := range deletedList {
				//获取到减少的号的进粉数
				for _, numInfo := range orderStateMap[url].NumList {
					if num == numInfo.Num {
						allMsg = fmt.Sprintf("%s%s 进粉数 >= %d\n", allMsg, num, numInfo.UserNum)
					}
				}
			}
			allMsg = fmt.Sprintf("%s\n", allMsg)
		}
		// 号码有变更了 要重新同步下号码
		if len(addedList) > 0 || len(deletedList) > 0 {
			delete(orderStateMap, url)
			addOrderStateMap(url, intoAllFuns, list)
		}
	}
	//个人工单变化
	userNumMsg := runPkg.DDMsgCfg{
		IsAtAll: false,
	}
	//总工单状态变化
	orderNumMsg := runPkg.DDMsgCfg{
		IsAtAll: false,
		Msg:     allMsg,
	}

	//遍历服务端保存的所有加了这个工单的信息 按人来的
	for idx, orderSave := range orderInfoMap[url] {
		userMsg := ""
		orderMsg := ""
		if orderSave.OrderUrl != url {
			fmt.Println("orderSave.OrderUrl 和 orderInfoMap的key值不一致")
			continue
		}
		reqUser, err := systemService.UserServiceApp.FindUserById(int(orderSave.UserID))
		if err != nil {
			global.GVA_LOG.Error("获取失败!", zap.Error(err))
			continue
		}
		//订单发生变化 AT所有添加这个工单号码的人
		if len(allMsg) > 0 {
			orderNumMsg.AtMobiles = append(orderNumMsg.AtMobiles, reqUser.Phone)
		}
		//找到设置的最大限制进粉数
		if orderSave.UserAllLimit > orderMaxLimit {
			orderMaxLimit = orderSave.UserAllLimit
		}
		// (orderInfoMap[url][0]).NumList[0].UserNum = 10
		numList := &orderInfoMap[url][idx].NumList
		bAllException := true
		//遍历每个人加了这个工单的号码
		for i, numInfo := range *numList {
			//遍历工单的所有号码
			for _, orderNumsData := range list {
				numData, ok := orderNumsData.(map[string]interface{})
				if !ok {
					global.GVA_LOG.Error("invalid orderNumsData type")
					continue
				}
				bUpdate := false
				if len(deletedList) > 0 {
					for _, str := range deletedList {
						if strings.Compare(str, numInfo.Num) == 0 {
							bUpdate = true
							(*numList)[i].State = runPkg.NumType_Lost
							break
						}
					}
				}
				if numInfo.Num == numData["numId"].(string) {
					// userNum := numInfo.UserNum
					// numState := numInfo.State

					if numInfo.UserNum != int(numData["intoFans"].(float64)) {
						(*numList)[i].UserNum = int(numData["intoFans"].(float64))
						// fmt.Println("进粉数 之前", userNum, " 之后 ", (*numList)[i].UserNum)
						numInfo.UserNum = (*numList)[i].UserNum
						// runPkg.NotifyClient(&orderSave)
						bUpdate = true
					}
					//冻结状态就不要改了
					if numInfo.State != runPkg.NumType_Freeze && numInfo.State != int(numData["state"].(float64)) {
						// fmt.Println(numData["numId"].(string), " 发生变化 之前", numState, " 之后 ", int(numData["state"].(float64)))
						(*numList)[i].State = int(numData["state"].(float64))
						// userNumMsg.Msg = fmt.Sprintf("%s号码: %s\n从状态: %s\n变为状态: %s\n----------------------\n", userNumMsg.Msg, numInfo.Num, getNameByState(numState), getNameByState((*numList)[i].State))
						numInfo.State = (*numList)[i].State
						bUpdate = true
					}

					if numInfo.State != runPkg.NumType_Freeze && numInfo.UserNum >= numInfo.UserLimit {
						// fmt.Println("超出进粉限制，冻结号码")
						userNumMsg.AtMobiles = append(userNumMsg.AtMobiles, reqUser.Phone)
						//自动冻结的 就不发通知了
						// userNumMsg.Msg = fmt.Sprintf("%s号码: %s \n进粉数: %d \n超出设定限制 %d\n已被冻结\n----------------------\n", userNumMsg.Msg, numInfo.Num, numInfo.UserNum, numInfo.UserLimit)
						userMsg = fmt.Sprintf("%s号码: %s \n进粉数: %d \n超出设定限制 %d\n已被冻结\n----------------------\n", userMsg, numInfo.Num, numInfo.UserNum, numInfo.UserLimit)
						bUpdate = true
						(*numList)[i].State = runPkg.NumType_Freeze
						numInfo.State = (*numList)[i].State
					}

					//有一个号码在线状态 就说明还能进
					if numInfo.State == runPkg.NumType_OnLine {
						bAllException = false
					}

				}
				if bUpdate {
					// global.GVA_DB.Model(&runPkg.RunNum{}).Where("num = ?", numInfo.Num).UpdateColumns(map[string]interface{}{
					global.GVA_DB.Model(&runPkg.RunNum{}).Where("num = ? AND user_id = ?", numInfo.Num, orderSave.UserID).UpdateColumns(map[string]interface{}{
						"state":    numInfo.State,
						"user_num": numInfo.UserNum,
					})
				}
			}
		}
		if bAllException {
			orderWarnNums[orderSave.OrderName]++
			// 无可用号码只会提示一次 如果更改了号码那么就再重置为0
			if orderWarnNums[orderSave.OrderName] < 2 {
				// userNumMsg.Msg += fmt.Sprintf("%s无可用号码\n----------------------\n", userNumMsg.Msg)
				userMsg += fmt.Sprintf("%s无可用号码\n----------------------\n", userMsg)
				userNumMsg.AtMobiles = append(userNumMsg.AtMobiles, reqUser.Phone)

			}
		}

		if len(userMsg) > 0 {
			userNumMsg.Msg = fmt.Sprintf("工单:%s\n%s", orderSave.OrderName, userMsg)
		}

		if len(allMsg) > 0 {
			orderNumMsg.Msg = fmt.Sprintf("工单:%s\n%s", orderSave.OrderName, orderNumMsg.Msg)
		}
		//一个人的这个工单的所有进粉数
		// fmt.Println("管理人--", orderSave.UserID, " 的所有进粉数 -- ", allFans)
		if !orderStateMap[url].Warning100 && intoAllFuns >= float64(orderMaxLimit) {
			orderMsg = fmt.Sprintf("%s进粉总数: %d  进粉限制: %d\n", orderMsg, int(intoAllFuns), orderMaxLimit)
			orderStateMap[url].Warning100 = true
			orderStateMap[url].Warning90 = true
			orderStateMap[url].Warning70 = true
			//添加要AT的人
			orderNumMsg.AtMobiles = append(orderNumMsg.AtMobiles, reqUser.Phone)
		} else if !orderStateMap[url].Warning90 && intoAllFuns >= float64(orderMaxLimit)*0.9 {
			orderMsg = fmt.Sprintf("%s进粉总数: %d  进粉限制: %d\n", orderMsg, int(intoAllFuns), orderMaxLimit)
			orderStateMap[url].Warning90 = true
			orderStateMap[url].Warning70 = true
			//添加要AT的人
			orderNumMsg.AtMobiles = append(orderNumMsg.AtMobiles, reqUser.Phone)
		} else if !orderStateMap[url].Warning70 && intoAllFuns >= float64(orderMaxLimit)*0.7 {
			orderMsg = fmt.Sprintf("%s进粉总数: %d  进粉限制: %d\n", orderMsg, int(intoAllFuns), orderMaxLimit)
			orderStateMap[url].Warning70 = true
			//添加要AT的人
			orderNumMsg.AtMobiles = append(orderNumMsg.AtMobiles, reqUser.Phone)
		}
		if len(orderMsg) > 0 {
			orderNumMsg.Msg += fmt.Sprintf("工单地址:%s\n%s", url, orderMsg)
		}

	}
	if len(userNumMsg.Msg) > 0 {
		errSend := notifyService.ServiceGroupApp.SendTextMessage(userNumMsg.Msg, userNumMsg.AtMobiles, userNumMsg.IsAtAll)
		if errSend != nil {
			fmt.Println("发送错误--", errSend.Error())
		}
	}
	if len(orderNumMsg.Msg) > 0 {
		errSend := notifyService.ServiceGroupApp.SendTextMessage(orderNumMsg.Msg, orderNumMsg.AtMobiles, orderNumMsg.IsAtAll)
		if errSend != nil {
			fmt.Println("发送错误--", errSend.Error())
		}
	}
}

// 根据号码状态获取中文
func getNameByState(state int) (stateName string) {
	switch state {
	case runPkg.NumType_OffLine:
		return "离线"
	case runPkg.NumType_OnLine:
		return "在线"
	case runPkg.NumType_Lock:
		return "封号"
	case runPkg.NumType_Freeze:
		return "冻结"
	case runPkg.NumType_Lost:
		return "丢失"
	default:
		fmt.Println("未知状态 ---- ", state)
		return "未知"
	}
}

// 添加 orderStateMap值
func addOrderStateMap(url string, intoAllFuns float64, list []interface{}) {
	orderState := runPkg.RunSaveOrderState{
		Warning70:    false,
		Warning90:    false,
		Warning100:   false,
		LostIsNotice: false,
		LostTimes:    0,
		IntoAllFuns:  intoAllFuns,
		NumList:      []runPkg.RunSaveOrderNum{},
	}

	numList := orderState.NumList
	for _, numState := range list {
		numData, ok := numState.(map[string]interface{})
		if !ok {
			global.GVA_LOG.Error("invalid orderNumsData type")
			continue
		}

		num := runPkg.RunSaveOrderNum{
			Num:     numData["numId"].(string),
			State:   int(numData["state"].(float64)),
			UserNum: int(numData["intoFans"].(float64)),
		}
		numList = append(numList, num)
	}

	orderState.IntoAllFuns = intoAllFuns
	orderState.NumList = numList
	orderStateMap[url] = &orderState
}

// 比较前后变化
func compareNumsThanAdjust(prevList []runPkg.RunSaveOrderNum, nextList []interface{}) (addedList []string, deletedList []string) {
	prevMap := make(map[string]bool)
	for _, val := range prevList {
		prevMap[val.Num] = true
	}

	for _, val := range nextList {
		// 如果在 prevMap 中不存在，则表示是新增的元素
		numData, ok := val.(map[string]interface{})
		if !ok {
			global.GVA_LOG.Error("invalid orderNumsData type")
			continue
		}
		if _, ok := prevMap[numData["numId"].(string)]; !ok {
			addedList = append(addedList, numData["numId"].(string))
		} else {
			// 如果在 prevMap 中存在，则表示没有变化
			delete(prevMap, numData["numId"].(string))
		}
	}

	// 遍历 prevMap 剩下的元素，表示它们在 prevList 中存在，但在 nextList 中不存在，即为删除的元素
	for key := range prevMap {
		deletedList = append(deletedList, key)
	}

	return addedList, deletedList
}

// UpdateRunNum 更新RunNum记录
func (runOrderNumsService *RunOrderNumsService) UpdateRunNum(prevNum string, runNum runPkg.RunNum) (err error) {
	for mapIdx, orderList := range orderInfoMap {
		for orderIdx, orderInfo := range orderList {
			if orderInfo.UserID == runNum.UserId {
				for numIdx, numInfo := range orderInfo.NumList {
					if numInfo.Num == prevNum {
						orderInfoMap[mapIdx][orderIdx].NumList[numIdx].Num = runNum.Num
						orderInfoMap[mapIdx][orderIdx].NumList[numIdx].UserLimit = runNum.EachEnterNum
						orderInfoMap[mapIdx][orderIdx].NumList[numIdx].State = runNum.State
						orderInfoMap[mapIdx][orderIdx].NumList[numIdx].NumType = runNum.NumType
						orderWarnNums[orderInfo.OrderName] = 0
						return nil
					}
				}
			}
		}
	}
	return nil
}

// 启动定时矫正数据
func (runOrderNumsService *RunOrderNumsService) StartTimer() {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		for url := range orderInfoMap {
			urlInfo := strings.Split(url, "！")
			if len(urlInfo) < 2 || len(urlInfo[0]) <= 0 {
				global.GVA_LOG.Error("启动定时矫正数据时解析 url 失败 删除矫正数据：", zap.String("url", url))
				delete(orderInfoMap, url)
				continue
			}
			data, err := runOrderNumsService.GetOrderData(urlInfo[0], "", urlInfo[1], "secondStep")
			if err != nil {
				// codeString := data["msg"].(string)
				// global.GVA_LOG.Error("启动定时矫正数据 时 解析 url 失败")
				//没有通知过 到达10次之后通知下 后面就不要再通知了
				if orderStateMap[url] != nil && !orderStateMap[url].LostIsNotice {
					orderStateMap[url].LostTimes++
					if orderStateMap[url].LostTimes > 10 {
						orderStateMap[url].LostIsNotice = true
						msg := fmt.Sprintf("工单:%s\n已长时间不能获取到数据，无法进行矫正\n请进入工单及后台自行对比", url)
						atMob := make([]string, 0)

						for _, persion := range orderInfoMap[url] {
							reqUser, err := systemService.UserServiceApp.FindUserById(int(persion.UserID))
							if err != nil {
								global.GVA_LOG.Error("获取失败!", zap.Error(err))
								continue
							}
							atMob = append(atMob, reqUser.Phone)
						}
						errSend := notifyService.ServiceGroupApp.SendTextMessage(msg, atMob, false)
						if errSend != nil {
							fmt.Println("发送错误--", errSend.Error())
						}
					}
				}

				continue
			}
			if data["code"].(float64) == 210 {
				global.GVA_LOG.Error(data["msg"].(string))
				//分享链接失效 要删掉这个数据
				delete(orderInfoMap, url)
				global.GVA_LOG.Error("分享链接失败 删除矫正数据")
				continue
			} else if data["code"].(float64) == 200 {
				runOrderNumsService.processOrderNumsData(data["data"].(map[string]interface{}), url)
				orderStateMap[url].LostTimes = 0
			}
		}
	}
}

// 获取工单号
func (runOrderNumsService *RunOrderNumsService) GetOrderData(orderUrl string, orderPsw string, orderType string, step string) (map[string]interface{}, error) {

	// link := "http://localhost:5000/getData"
	// link := "http://47.90.250.28:5000/getData"
	// link := "http://47.245.100.232:5000/getData"

	link := global.GVA_CONFIG.Notify.GetDataUrl
	// 构造要发送的参数
	params := url.Values{}
	params.Set("orderUrl", orderUrl)
	params.Set("step", step)
	params.Set("orderPsw", orderPsw)
	params.Set("orderType", orderType)

	// 发送请求获取数据
	resp, err := http.PostForm(link, params)
	if err != nil || resp == nil || resp.StatusCode != 200 {
		return nil, errors.New("获取数据失败")
	}
	defer resp.Body.Close()

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, errors.New("解析数据失败")
	}
	return data, nil
}

// 根据创建订单信息创建一个RunSaveOrderInfo对象
func createRunSaveOrderInfo(createNumsInfo runPkgReq.RunCreateNumsInfo) runPkg.RunSaveOrderInfo {
	numList := make([]runPkg.RunSaveOrderNum, len(createNumsInfo.Nums))
	for i, num := range createNumsInfo.Nums {
		numObj := runPkg.RunSaveOrderNum{
			Num:       num,
			NumType:   createNumsInfo.NumType,
			PageId:    createNumsInfo.PageId,
			State:     1,
			UserNum:   0,
			UserLimit: createNumsInfo.EachEnterNum,
		}
		numList[i] = numObj
	}
	orderObj := runPkg.RunSaveOrderInfo{
		UserID:       createNumsInfo.UserId,
		OrderName:    createNumsInfo.OrderName,
		OrderUrl:     createNumsInfo.OrderUrl + "！" + createNumsInfo.OrderUrlType,
		UserAllLimit: createNumsInfo.MaxEnterNum,
		NumList:      numList,
	}
	orderWarnNums[createNumsInfo.OrderName] = 0
	return orderObj
}

// 向给定订单中添加号码
func addNumsToOrder(orderInfo runPkg.RunSaveOrderInfo, numList []runPkg.RunSaveOrderNum, createNumsInfo runPkgReq.RunCreateNumsInfo) []runPkg.RunSaveOrderNum {
	for i := 0; i < len(createNumsInfo.Nums); i++ {
		isNumExist := false
		for j := 0; j < len(orderInfo.NumList); j++ {
			if orderInfo.NumList[j].Num == createNumsInfo.Nums[i] {
				isNumExist = true
				break
			}
		}
		if !isNumExist {
			numObj := runPkg.RunSaveOrderNum{
				Num:       createNumsInfo.Nums[i],
				PageId:    createNumsInfo.PageId,
				NumType:   createNumsInfo.NumType,
				State:     1,
				UserNum:   0,
				UserLimit: createNumsInfo.EachEnterNum,
			}
			numList = append(numList, numObj)
		}
	}
	return numList
}

func (runOrderNumsService *RunOrderNumsService) SaveOrderNums(createNumsInfo runPkgReq.RunCreateNumsInfo) error {
	jsonBytes, err := json.Marshal(createNumsInfo)
	if err != nil {
		fmt.Println("JSON encoding failed:", err)
		return err
	}

	fmt.Println(string(jsonBytes))

	orderInfoMapMutex.Lock()
	defer orderInfoMapMutex.Unlock()

	urlKey := createNumsInfo.OrderUrl + "！" + createNumsInfo.OrderUrlType

	// 如果没有该订单的值，则新建
	if _, ok := orderInfoMap[urlKey]; !ok {
		orderInfoMap[urlKey] = []runPkg.RunSaveOrderInfo{}
	}

	// 是否存在用户号码
	isExist := false
	for i := 0; i < len(orderInfoMap[urlKey]); i++ {
		if orderInfoMap[urlKey][i].UserID == createNumsInfo.UserId {
			isExist = true
			orderInfoMap[urlKey][i].OrderName = createNumsInfo.OrderName
			orderInfoMap[urlKey][i].NumList = addNumsToOrder(orderInfoMap[urlKey][i], orderInfoMap[urlKey][i].NumList, createNumsInfo)
			orderInfoMap[urlKey][i].OrderUrl = createNumsInfo.OrderUrl + "！" + createNumsInfo.OrderUrlType
			orderInfoMap[urlKey][i].UserAllLimit = createNumsInfo.MaxEnterNum
			break
		}
	}

	// 不存在用户号码则创建
	if !isExist {
		orderInfo := createRunSaveOrderInfo(createNumsInfo)
		orderInfoMap[urlKey] = append(orderInfoMap[urlKey], orderInfo)
	}

	jsonString, err := json.Marshal(orderInfoMap)

	if err != nil {
		fmt.Println("JSON encoding failed:", err)
	}

	fmt.Println(string(jsonString))

	return nil
}

// 添加号码
func (runOrderNumsService *RunOrderNumsService) AddRunNum(runNum runPkg.RunNum) error {
	for url, orderInfo := range orderInfoMap {
		for idx, info := range orderInfo {
			if info.UserID == runNum.UserId && info.OrderName == runNum.OrderName {
				saveNum := runPkg.RunSaveOrderNum{
					Num:       runNum.Num,
					PageId:    runNum.PageId,
					State:     1,
					UserNum:   0,
					UserLimit: runNum.EachEnterNum,
					NumType:   runNum.NumType,
				}

				orderInfoMap[url][idx].NumList = append(orderInfoMap[url][idx].NumList, saveNum)
			}
		}
	}
	return nil
}

// 删除号码
func (runOrderNumsService *RunOrderNumsService) DeleteRunNum(runNum runPkg.RunNum) error {
	for url, orderInfo := range orderInfoMap {
		for idx, info := range orderInfo {
			if info.UserID == runNum.UserId && info.OrderName == runNum.OrderName {
				for numIdx, num := range orderInfoMap[url][idx].NumList {
					if num.Num == runNum.Num {
						//剩一个时无法删除 直接orderInfoMap里去掉
						if len(orderInfoMap[url][idx].NumList) <= 1 {
							orderInfoMap[url][idx].NumList = make([]runPkg.RunSaveOrderNum, 0)
						} else {
							orderInfoMap[url][idx].NumList = append(orderInfoMap[url][idx].NumList[:numIdx], orderInfoMap[url][idx].NumList[numIdx+1:]...)
						}
						return nil
					}
				}
			}
		}
	}
	return nil
}

// 删除工单
func (runOrderNumsService *RunOrderNumsService) DeleteRunOrder(runOrder runPkg.RunOrder) error {
	for url, orderInfo := range orderInfoMap {
		for idx, info := range orderInfo {
			if info.UserID == runOrder.DeletedBy && info.OrderName == runOrder.OrderName {
				//省一个时无法删除 直接orderInfoMap里去掉
				if len(orderInfo) <= 1 {
					orderInfoMap[url] = make([]runPkg.RunSaveOrderInfo, 0)
					break
				} else {
					orderInfoMap[url] = append(orderInfo[:idx], orderInfo[idx+1:]...)
				}

				delete(orderWarnNums, info.OrderName)
			}
		}
	}
	return nil
}

// 更新工单
func (runOrderNumsService *RunOrderNumsService) UpdateRunOrder(runOrder runPkg.RunOrder) error {
	for url, orderInfo := range orderInfoMap {
		for idx, info := range orderInfo {
			if info.UserID == runOrder.UpdatedBy && info.OrderName == runOrder.OrderName {
				orderInfoMap[url][idx].UserAllLimit = runOrder.MaxEnterNum
				orderInfoMap[url][idx].OrderName = runOrder.OrderName

				for numIdx, _ := range orderInfoMap[url][idx].NumList {
					orderInfoMap[url][idx].NumList[numIdx].State = runOrder.State
				}
				orderWarnNums[runOrder.OrderName] = 0
			}
		}
	}
	return nil
}
