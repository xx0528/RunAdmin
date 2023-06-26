import http.client
import re
import json
import requests
import cloudscraper
from flask import Flask, request
from urllib.parse import urlparse, parse_qs

# cloudscraper 要最新的 要依赖nodejs  记得在服务端安装nodejs
# python -m pip install requests
# pip install requests-toolbelt
# pip install js2py
# pip install v8eval

app = Flask(__name__)

class OrderType:
    Share = "share"
    URL = "url"
    KF007 = "kf007"
    GooSu = "goosu"
    XYZ = "xyz"

class NumStateType:
    NumType_OffLine = 0 #离线
    NumType_OnLine  = 1 #在线
    NumType_Lock    = 2 #封号
    NumType_Freeze  = 3 #冻结
    NumType_Lost    = 4 #丢失
    
@app.route('/getData', methods=['POST'])
def get_data():
    url = request.form["orderUrl"]
    step = request.form["step"]
    psw = request.form["orderPsw"]
    orderType = request.form["orderType"]
    print("开始提取工单 orderType = " + orderType + " psw=" + psw + " url = " + url + '\n')
    if step == "firstStep":
        # 云控工单
        if orderType == OrderType.Share:
            data = get_Share(url)
        # 007工单
        elif orderType == OrderType.KF007:
            data = get_KF007(url, psw)
        # goo.su工单
        elif orderType == OrderType.GooSu:
            data = get_GooSu(url, psw)
        # ok0.xyz工单
        elif orderType == OrderType.XYZ:
            data = get_XYZ(url, psw)
        # url工单
        elif orderType == OrderType.URL:
            data = get_URL(url, psw)

        data["orderType"] = orderType
        print("提取工单结束 第一步 返回---" + json.dumps(data) + '\n\n')
        return data
    if step == "secondStep":
        # 云控工单
        if orderType == OrderType.Share:
            data = get_Share(url)
        # 007工单
        elif orderType == OrderType.KF007:
            data = get_KF0072(url)
        # goo.su工单
        elif orderType == OrderType.GooSu:
            data = get_GooSu2(url)
        # ok0.xyz工单
        elif orderType == OrderType.XYZ:
            data = get_XYZ2(url)
        # url工单
        elif orderType == OrderType.URL:
            data = get_URL2(url)
        
        data["orderType"] = orderType
        print("提取工单结束 第二步 返回---" + json.dumps(data) + '\n\n')
        return data
def get_Share(url):
    index_html_index = url.find("index.html")
    base_url = url[:index_html_index]
    token = url.split("=")[1]
    newUrl = base_url + "/api_yinliu_count.html?token=" + token + "&page=1&limit=200&is_repet=1"
    scraper = cloudscraper.create_scraper(disableCloudflareV1=True)
    response = scraper.get(newUrl, allow_redirects=True)
    print("share 工单 status_code = " + str(response.status_code) + '\n')
    if response.status_code != 200:
        new_data = {
            "code": response.status_code,
            "data": {},
            "msg": "访问失败 " + url,
        }
        return new_data
    data = json.loads(response.text)
    print("share 工单 data code = " + str(data["code"]) + '\n')
    new_data = {}
    new_data["code"] = data["code"]
    new_data["msg"] = "成功"

    if data["code"] != 0:
        new_data["code"] = data["code"]
        new_data["msg"] = "请求地址出错 url = " + url + " status = " + str(data["code"])
        return new_data
    
    data = json.loads(response.text)
    #因为工单不同 code 正确 返回值不同 有的0 有的200 这里统一200
    new_data["code"] = 200
    new_data["data"] = {}
    new_data["data"]["intoAllFuns"] = int(data["totalRow"]["day_sum"])
    new_data["data"]["list"] = []
    for item in data["data"]:
        state = item["online"]
        # print(item["user"] + "  -- " + str(item["online"]))
        if state == 3:  #封号
            state = NumStateType.NumType_Lock
            # print("封号状态-------------------" + item["user"] + " num lock -- " + str(NumStateType.NumType_Lock))
        elif state == 2:  #离线
            state = NumStateType.NumType_OffLine
            # print("离线状态=====-------------------")
        new_item = {
            "numId": item["user"],
            "state": state,
            "intoFans": int(item["day_sum"]),
            "repeatFans": 0
        }
        new_data["data"]["list"].append(new_item)
    new_data["orderUrl"] = url

    return new_data

def get_KF007(url, psw):

    scraper = cloudscraper.create_scraper(disableCloudflareV1=True)
    response = scraper.get(url, allow_redirects=True)
    location = response.url
    print("KF007 工单1  status_code = " + str(response.status_code) + " location = " + location + '\n')

    order_id = ""
    matchOrderId = re.search(r"shared-order/(.*?)\?u=", location)
    if matchOrderId:
        order_id = matchOrderId.group(1)

    order_u = ""
    matchOrderU = re.search(r"\?u=(.*?)&code=", location)
    if matchOrderU:
        order_u = matchOrderU.group(1)

    code = ""
    matchCode = re.search(r"&code=(.*)", location)
    if matchCode:
        code = matchCode.group(1)

    print("KF007 工单1  order_id = " + order_id + " order_u = " + order_u + " code = " + code + '\n')
    #https://kf.007.tools/shared-order/efe4b9b2f249fc117444b8e2e9f1cf67868f10b0a4a4e769d5d1d8cdebf2a71a?u=0cb692081b8e43818c849ad960e4d71c&code=LXJimlFK
    #https://kf.007.tools/counter-api/detail/user-detail/get_share_list?order_id=efe4b9b2f249fc117444b8e2e9f1cf67868f10b0a4a4e769d5d1d8cdebf2a71a&uuid=0cb692081b8e43818c849ad960e4d71c&code=LXJimlFK&share_pwd=qq123&page=1&perpage=50
    newUrl = "https://kf.007.tools/counter-api/detail/user-detail/get_share_list?order_id=" + order_id + "&uuid=" + order_u + "&code=" + code + "&share_pwd=" + psw + "&page=1&perpage=200"
    return get_KF0072(newUrl)

def get_KF0072(url):
    scraper = cloudscraper.create_scraper(disableCloudflareV1=True)
    response = scraper.get(url)
    print("KF007 工单2 访问 url = " + url + '\n')
    print('KF007 工单2 status_code = ' + str(response.status_code) + '\n')
    if response.status_code != 200:
        redata = {
            "code": response.status_code,
            "data": {},
            "msg": "访问失败 " + url,
        }
        return redata
    text = response.text
    # print("text -- \n" + text)
    new_data = {}

    if response.status_code == 403:
        new_data["code"] = 7
        new_data["msg"] = "工单地址拒绝访问"
        return new_data

    data = json.loads(text)
    print("KF007 工单2 data code == " + str(data["code"]) + '\n')

    if not data or not data["code"] or data["code"] != 200:
        new_data["code"] = data["code"]
        new_data["msg"] = data["msg"]
        return new_data
    
    #因为工单不同 code 正确 返回值不同 有的0 有的200 这里统一200
    new_data["code"] = 200
    new_data["data"] = {}
    new_data["data"]["intoAllFuns"] = int(data["data"]["total"])
    new_data["data"]["list"] = []
    
    for item in data["data"]["list"]:
        if item["line_id"] and len(item["line_id"]) > 0:
            numId = item["line_id"]
        else:
            numId = item["line_account"]
        new_item = {
            "numId": numId,
            "state": item["status"],
            "intoFans": int(item["day_target"]),
            "repeatFans": 0
        }
        new_data["data"]["list"].append(new_item)
    new_data["orderUrl"] = url
    return new_data

def get_GooSu(url, psw):
    # TODO: 此处实现获取 goo.su 工单的代码
    scraper = cloudscraper.create_scraper(disableCloudflareV1=True)
    response = scraper.get(url, allow_redirects=True)
    location = response.url
    #https://007.mn/staff-share-list?uuid=659707fc333d493ca33fe8c4139dcc0f
    print('GooSu 工单1 status_code = ' + str(response.status_code) + " location = " + location + '\n')

    newUrl = getNewUrlByLocation(location, psw)
    if newUrl == "":
        redata = {
            "code": response.status_code,
            "data": {},
            "msg": "访问失败-没有对应网址类型：" + url + " \n location--" + location,
        }
        return redata
    
    return get_URL2(newUrl)

def get_GooSu2(url):
    print('GooSu 工单2 url = ' + url)
    return get_URL2(url)

def get_XYZ(url, psw):
    scraper = cloudscraper.create_scraper(disableCloudflareV1=True)
    response = scraper.get(url, allow_redirects=True)
    location = response.url
    #https://007.mn/staff-share-list?uuid=659707fc333d493ca33fe8c4139dcc0f
    print('XYZ 工单1 status_code = ' + str(response.status_code) + " location = " + location + '\n')

    newUrl = getNewUrlByLocation(location, psw)
    if newUrl == "":
        redata = {
            "code": response.status_code,
            "data": {},
            "msg": "访问失败-没有对应网址类型：" + url + " \n location--" + location,
        }
        return redata
    
    return get_URL2(newUrl)

def get_XYZ2(url):
    print('XYZ 工单2 url = ' + url + '\n')
    return get_URL2(url)

def get_URL(url, psw):
    # 发送第一个请求
    scraper = cloudscraper.create_scraper(disableCloudflareV1=True)
    response = scraper.get(url, allow_redirects = True)
    location = response.url
    print('URL 工单1 status_code = ' + str(response.status_code) + " location = " + location + '\n')
    newUrl = getNewUrlByLocation(location, psw)
    if newUrl == "":
        redata = {
            "code": response.status_code,
            "data": {},
            "msg": "访问失败-没有对应网址类型：" + url + " \n location--" + location,
        }
        return redata
    
    return get_URL2(newUrl)

def get_URL2(url):
    print("URL 工单2 url = " + url + '\n')
    scraper = cloudscraper.create_scraper(disableCloudflareV1=True, delay=1000)
    response = scraper.get(url)

    print("URL 工单2 status_code = " + str(response.status_code) + '\n')
    new_data = {}
    if response.status_code == 403:
        new_data["code"] = 7
        new_data["msg"] = "工单地址拒绝访问"
        return new_data
    
    data = json.loads(response.text)
    print("URL 工单2 data code = " + str(data["code"]) + '\n')
    if not data or not data["code"] or data["code"] != 200:
        new_data["code"] = data["code"]
        new_data["msg"] = data["msg"]
        return new_data
    
    #因为工单不同 code 正确 返回值不同 有的0 有的200 这里统一200
    new_data["code"] = 200
    new_data["data"] = {}
    #不同地址 请求来的数据这个总数key值不同 有的还没有 所以这里自己计算吧
    intoAllFans = 0
    new_data["data"]["list"] = []
    for item in data["data"]["list"]:
        state = item["online_status"]
        if state == 2:  #2是离线
            state =  NumStateType.NumType_OffLine
        elif state == 3:  #3是封号
            state =  NumStateType.NumType_Lock
        #可能被踢掉了
        if item["single_into_fans_num"] == None:
            continue
        intoAllFans += int(item["single_into_fans_num"])
        new_item = {
            "numId": item["username"],
            "state": state,
            "intoFans": int(item["single_into_fans_num"]),
            "repeatFans": int(item["single_repeat_fans_num"])
        }
        new_data["data"]["list"].append(new_item)
    new_data["orderUrl"] = url
    new_data["data"]["intoAllFuns"] = intoAllFans
    return new_data

def getNewUrlByLocation(location, psw):
    # print("location ---- " + location)
    uuid_regex = re.compile(r"uuid=([^&]+)")
    sid_regex = re.compile(r"sid=([^&]+)")
    # 根据重定向链接获取分享链接
    #这里有两种域名 007.mn 007.ma
    domain = location.split("//")[1].split("/")[0].split(":")[0]
    newUrl = ''
    # domain = "007.ma"

    #req url -- https://url04.top/b/NcFvCfYN4502/5
    #location -- https://007.mn/ws-add-fans-detail-share-list?uuid=16aa1d3d8437476b80e83580fd61640f&sid=d0b6bbdca9d5451c9cc68784e31dab7c
    #new url -- https://007.mn/java-api/shunt/fans_details_list?page=1&per_page=10&link_uuid=16aa1d3d8437476b80e83580fd61640f&sid=d0b6bbdca9d5451c9cc68784e31dab7c&password=1234
    if "shunt" in location or "ws-add-fans" in location:
        uuid = uuid_regex.search(location).group(1)
        sidSearch = sid_regex.search(location)
        if sidSearch:
            sid = sidSearch.group(1)
            newUrl = "https://" + domain + "/java-api/shunt/fans_details_list?page=1&per_page=150&link_uuid=" + uuid+"&sid=" + sid + "&password=" + psw
        else :
            newUrl = "https://" + domain + "/java-api/shunt/fans_details_list?page=1&per_page=150&link_uuid=" + uuid + "&password=" + psw        
    #req url --     https://url05.top/b/CvrhlmoX2175/4
    #location ---   https://007.mn/staff-share-list?uuid=d698e7cd4987443cb2cc8342abb95c35
    #new url --     https://007.mn/java-api/staff/get_share_list?page=1&per_page=10&uuid=d698e7cd4987443cb2cc8342abb95c35&password=12345
    elif "staff-share-list" in location:
        uuid = uuid_regex.search(location).group(1)
        sidSearch = sid_regex.search(location)
        if sidSearch:
            sid = sidSearch.group(1)
            newUrl = "https://" + domain + "/java-api/staff/get_share_list?page=1&per_page=150&uuid=" + uuid + "&sid=" + sid + "&password=" + psw
        else :
            newUrl = "https://" + domain + "/java-api/staff/get_share_list?page=1&per_page=150&uuid=" + uuid + "&password=" + psw
    #req url --     https://url08.top/b/ndSOEwjN9588/3
    #location ---   https://007.mn/shated-staff-list?uuid=ca8bfbc2a8b44afca0b37e6d157002bc 
    #new url --     https://007.mn/java-api/staff/staff-share/get_account?page=1&per_page=10&uuid=ca8bfbc2a8b44afca0b37e6d157002bc&password=a1234
    elif "shated-staff-list" in location:
        uuid = uuid_regex.search(location).group(1)
        sidSearch = sid_regex.search(location)
        if sidSearch:
            sid = sidSearch.group(1)
            newUrl = "https://" + domain + "/java-api/staff/staff-share/get_account?page=1&per_page=150&uuid=" + uuid + "&sid=" + sid + "&password=" + psw
        else :
            newUrl = "https://" + domain + "/java-api/staff/staff-share/get_account?page=1&per_page=150&uuid=" + uuid + "&password=" + psw
    else:
        newUrl = ""
    
    return newUrl
    
if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)