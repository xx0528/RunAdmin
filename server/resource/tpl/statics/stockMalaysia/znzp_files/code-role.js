// 规则
window.getGuojinMarketType = function (seccode) {
    if (seccode == null) {
        return null;
    }
    // 20200804 xujb 针对创业板注册制，300判断未创业板股票不合理修复
    if(seccode && seccode.length == 10) {
        return seccode;
    }

    var codeFront = seccode.substr(0, 3);
    if(codeFront == "688") { //科创板
        return '0107' + seccode;
    }else if (codeFront == "600" || codeFront == "601" || codeFront == "603") {//沪市A股
        return "0101" + seccode;
    } else if (codeFront == "000" || codeFront == "001") {//深市A股
        return "0001" + seccode;
    } else if (codeFront == "002" || codeFront == "003") {//中小板
        return "0001" + seccode;
    } else if (codeFront == "200") {//b股
        return "0001" + seccode;
    } else if (codeFront == "300" || codeFront.substr(0, 2) == "30") {//创业板
        return "0001" + seccode;
    } else {
        return "0101" + seccode;//其他
    }
};
