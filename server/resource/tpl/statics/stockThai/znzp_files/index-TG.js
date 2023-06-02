function checkData(data, pro, code) {
    var reg = new RegExp("^\d{6}$");
    var flag = false;
    for (var i = 0; i < data.length; i++) {
        var temp = data[i];
        if (temp[pro] == code) {
            flag = true;
            break;
        }
    }
    return flag;
}

function queryStock(queryStock) {
    console.log(queryStock);
    var isCatch = false;
	var infoDefault = {c: 'AJ', n: 'เอ.เจ.พลาสท์'};
	var infoArr = {};
    if (queryStock) {
        if (!/^\d{6}$/.test(queryStock)) {

            // console.log(s.length);
            for (var i = 0, len = s.length; i < len; i++) {
                var item = s[i];
                // alert(item);
                // alert(item.c);
                // alert(item.n);
                if ((item.c || "")
                    .toLowerCase()
                    .indexOf(queryStock.toLowerCase()) != -1) {
                    isCatch = true;
					infoArr.c = item.c
					infoArr.n = item.n
                    // queryStock = item.c;
					// console.log(item);
                    break;
                }
            }
        } else {
            isCatch = true;
        }
        if (isCatch) {
            return infoArr;
        } else {
				
            return infoDefault;
        }
    } else {
        return infoDefault;
    }
}


layer.config({
    maxWidth: '560'
});

function GetQueryString(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)");
    var r = window.location.search.substr(1).match(reg);
    if (r != null) return decodeURIComponent(r[2]);
    return null;
};

var cNum = GetQueryString("keyword") ? GetQueryString("keyword") : "AJ";
infoArr = queryStock(cNum);
$('#stock-input').val(infoArr.c);
  if (cNum != '') getCodeInfo(infoArr.c,infoArr.n);

function getCodeInfo(code,name) {
	// console.log(name);
  
	if(code!=undefined){
		 $('#stock-input').val(code);
		 $('.gCode').text("("+code+")");
		 $('.gpName,.cName,.banner_code,.dimensionName').text(code);
		 $('.gName').text(code);
		 $('.cnname').text(name);
	}
};

function showLocale(objD) {
    var str, colorhead, colorfoot;
    var yy = objD.getYear();
    if (yy < 1900) yy = yy + 1900;
    var MM = objD.getMonth() + 1;
    if (MM < 10) MM = '0' + MM;
    var dd = objD.getDate();
    if (dd < 10) dd = '0' + dd;
    var hh = objD.getHours();
    if (hh < 10) hh = '0' + hh;
    var mm = objD.getMinutes();
    if (mm < 10) mm = '0' + mm;
    var ss = objD.getSeconds();
    if (ss < 10) ss = '0' + ss;
    var ww = objD.getDay();
    if (ww == 0) colorhead = "<font>";
    if (ww > 0 && ww < 7) colorhead = "<font >";
    str = colorhead + yy + "/" + MM + "/" + dd + " &nbsp&nbsp" + hh + ":" + mm + ":" + ss + " ";
    return (str);
}



//图片延迟加载
ImgDelay();
function ImgDelay(){
    var lazyImgs = $('[data-src]');
    if(lazyImgs.length > 0){
        $("body").one("click",function(event){
            [].slice.call(lazyImgs).forEach(function(element){
                var imgSrc = element.getAttribute('data-src');
                element.style.backgroundImage = 'url('+imgSrc+')';
                element.removeAttribute('data-src');
            });
        });
    }
}

//confirm拒绝
$('.confirm_refuse').on('click', function() {
    $('.confirm_stay,.zhegai').hide();
})

//confirm接受
$('.confirm_accept').on('click', function() {
    try{
        setBrowseInfo('Wait-wanliuye','clk');
    }catch (e){
    }
    $('.confirm_stay').hide();
    $("[data-v='" + $(this).attr('data-v') + "']").show();
})

//带确认弹窗的cancel事件
function cancelComfirm(t) {
    try{
        setBrowseInfo('Wait-wanliuye','shw');
    }catch (e){
    }
    $(t).parent().css("display", "none");
    $(t).parent().attr('data-v', Math.random());
    $('.confirm_stay').css("display", "block");
    $('.confirm_accept').attr('data-v', $(t).parent().attr('data-v'));
}

//最新消息按钮，显示指定文案
function  newsWords() {
    var time = new Date();
    var h = time.getHours();
    var m = time.getMinutes();
    var num = "";
    var newTxt = ["最新消息已于今日9:00整合完成","最新消息已于今日13:00整合完成","最新消息已于今日15:30整合完成","最新消息已于昨日15:30整合完成"];
    if(h >= 9 && h < 13){
        num = 0;
    }else if(h >= 13 && h< 15){
        num = 1;
    }else if(h >= 15 && h > 0){
        if(h == 15 && m <= 30){
            num = 1;
        }else{
            num = 2;
        }
    }else if(h > 0 && h < 9){
        num = 3;
    }
    return newTxt[num];
}

//获取页面跟踪信息
scrollTable2();
scrollTable();
discText();


$('.circle2').click(function () {
    if($(this).parent().css('display')=='block'){
        cancelComfirm(this);
    }
});

if( sessionStorage.getItem("detain") == undefined){
    sessionStorage.setItem("detain", 1);
}
$('.circle5,.circle4,.circle3,.circle6').click(function(){
    if($(this).parent().css('display')=='block'){

        cancelComfirm(this);
    }
});

$('.circle7').click(function () {
    $('.detain,.zhegai').hide();
});



$('.plList li').one('click',function(){
    var n = parseInt($(this).find('span').eq(1).text());
    $(this).find('span').eq(1).text( n+1 );
});


//维度申请弹窗
$('.dimension_bg3 p').each(function (index) {
    $(this).click(function () {
        var dimension_txt = $('.dimension_txt');
        var dimension_txt1 = $('.dimension_txt1');
        var dimension_btn  = $('.dimension_btn');
        var dimension_input  = $('.dimension_input');
        // var codeName = $(".cnname").eq(0).text();
        var codeName = $(".cnname").eq(0).text();
        // console.log('cnname' + codeName);
        changeSpeedTxt(1,codeName);
        switch(index){
            case 0:
                note = "会不会涨";
                report_txt = "涨跌报告";
                dimension_txt.html("涨跌报告已调取完成");
                dimension_txt1.html("何时涨？涨多少？");
                dimension_input.attr("placeholder","输入手机号，涨跌预测私发您");
                dimension_btn.attr("id","dimensionBtn1");
                aloneLoad($(".dimension,.zhegai"));
                break;
            case 1:
                note = "最新消息";
                report_txt = "最新消息";
                var txts = newsWords();
                dimension_txt.html(txts);
                dimension_txt1.html("最新消息！ 直接影响股价");
                dimension_input.attr("placeholder","输入手机号，最新消息私发您");
                dimension_btn.attr("id","dimensionBtn2");
                aloneLoad($(".dimension,.zhegai"));
                break;
            case 2:
                note = $('.p3_txt').text();
                report_txt = note;
                if(note == "涨跌原因"){
                    dimension_txt.html("涨跌预测已调取完成");
                    dimension_input.attr("placeholder","输入手机号，涨跌原因私发您");
                    dimension_txt1.html("搞清原因的散户 都能提前知涨跌");
                }else if(note == "停牌原因"){
                    dimension_txt.html("停牌原因已调取完成");
                    dimension_input.attr("placeholder","输入手机号，停牌原因私发您");
                    dimension_txt1.html("搞清原因的散户 都能提前知涨跌");
                }else if(note == "涨停原因"){
                    $('.dimension_txt').html("涨停原因已调取完成");
                    $('.dimension_input').attr("placeholder","输入手机号，涨停原因私发您");
                    dimension_txt1.html("搞清原因的散户 都能提前知涨跌");
                }else if(note == "跌停原因"){
                    $('.dimension_txt').html("跌停原因已调取完成");
                    $('.dimension_input').attr("placeholder","输入手机号，跌停原因私发您");
                    dimension_txt1.html("搞清原因的散户 都能提前知涨跌");
                }else if(note == "超跌反弹时间"){
                    $('.dimension_txt').html("超跌反弹时间已调取成功");
                    $('.dimension_input').attr("placeholder","输入手机号，超跌反弹时间私发您");
                    dimension_txt1.html("抓住机会的散户 就能先下手为强");
                }
                dimension_btn.attr("id","dimensionBtn3");
                aloneLoad($(".dimension,.zhegai"));
                break;
            case 3:
                note = yearTime[0] + "年目标价";
                report_txt = yearTime[0] + "年目标价";
                $('.dialog5_txt').text( yearTime[0] + "年目标价已调取完成");
                $(".target").show();
                $(".buying").hide();
                $('.dialog5_input').attr("placeholder","输入手机号，目标价私发您");
                $('.dialog5_prompt').removeClass("prompts5");
                aloneLoad($(".dialog5,.zhegai"));
                break;
            case 4:
                note = $('.p5_txt').text();
                report_txt = note;
                if(note == "最佳买卖点"){
                    $('.dialog5_txt').text("最佳买卖点预测已调取完成");
                    $('.dialog5_input').attr("placeholder","输入手机号，买卖点私发您");
                    $(".target").hide();
                    $(".buying").show();
                    $('.dialog5_prompt').addClass("prompts5").css("margin","0");
                }else{
                    $('.dialog5_txt').text("复牌时间已调取完成");
                    $('.dialog5_input').attr("placeholder","输入手机号，复牌时间私发您");
                    $(".target,.buying").hide();
                    $('.dialog5_prompt').addClass("prompts5").css("margin","0.4rem 0");
                }
                aloneLoad($(".dialog5,.zhegai"));
                break;
            default:
                break;
        }

    });
});

function magic_number(value,name) {
    var num = $("."+name);
    num.animate({count: value}, {
        duration: 500,
        step: function() {
            num.text(String(parseInt(this.count)));
        }
    });
}


//诊股表格滚动
function scrollTable2() {
    var i = 1;
    var len = $('.sharesList tr').length;
    $('.sharesList').append($('.sharesList tr').clone());
    var _table = $('.sharesList').eq(0);
    setInterval(function() {
        var num = (-0.66 * i).toFixed(2);
        _table.css('marginTop', num + 'rem');
        i++;
        if (i == len + 1) {
            setTimeout(function() {
                _table.css('transition', 'none');
                _table.css('marginTop', 2);
                i = 1;
                setTimeout(function() {
                    _table.css('transition', 'all .7s')
                }, 700);
            }, 1000)
        }
    }, 2500);
}


function animate() {
    $('.tan_content').show();
    $(".dialog").hide();
    $('html, body').animate({
      scrollTop: 0
    }, 100);
    $(".charts").animate({
      width: "25%"
    }, 500, "", function () {
      $(".discuss").html("正在通过最小二乘法OLS确定必要报 酬率...");
    });
    $(".charts").animate({
      width: "50%"
    }, 600, "", function () {
      $(".discuss").html("正在通过VAR系统确认风险值...");
    });
    $(".charts").animate({
      width: "75%"
    }, 700, "", function () {
      $(".discuss").html("正在通过量价交易模型...");
    });
    $(".charts").animate({
      width: "100%"
    }, 800, "", function () {
      $(".discuss").html("正在通过量价交易模型...");
      $('.charts').css('width', '0');
      $('.tan_content').hide();
      $(".dialog").show();
      $('.phonec').focus();
    });
  }

  function saveMobile() {
    zpcj = cNum;
    if (zpcj == '' || !checkData(s, "c", zpcj)) {
      layer.msg('请输入正确的股票代码');
      return false;
    } else {
      $(".discuss").html("正在通过事件驱动策略模型...");
      animate();
      $("#SonContent0").show();
    }
  }

//进度条动画
var spotNum = '';

//进度条调用 Popup为Jq的DOM对象
function loadLittle(Popup) {
    var n = '.';
    $("#loadBox").css("display", "block");
    spotNum = setInterval(function() {
        if (n.length < 4) {
            $('.load_discuss .spot').text(n);
            n = n + '.';
        } else {
            n = '.';
            $('.load_discuss .spot').text(n);
        }
    }, 400);
    $(".load_charts").animate({
        width: "100%"
    }, 2000, "", function() {
        hideLittle();
        Popup.show();
    });
}

//结束后隐藏进度条
function hideLittle() {
    clearInterval(spotNum);
    $("#loadBox").hide();
    $(".load_charts").stop(true);
    $(".load_charts").width("0px");
}


//时间变换
function showLocale(objD) {
    var str, colorhead, colorfoot;
    var yy = objD.getYear();
    if (yy < 1900) yy = yy + 1900;
    var MM = objD.getMonth() + 1;
    if (MM < 10) MM = '0' + MM;
    var dd = objD.getDate();
    if (dd < 10) dd = '0' + dd;
    var hh = objD.getHours();
    if (hh < 10) hh = '0' + hh;
    var mm = objD.getMinutes();
    if (mm < 10) mm = '0' + mm;
    var ss = objD.getSeconds();
    if (ss < 10) ss = '0' + ss;
    var ww = objD.getDay();
    if (ww == 0) colorhead = "<font>";
    if (ww > 0 && ww < 7) colorhead = "<font >";
    str = colorhead + yy + "-" + MM + "-" + dd +" &nbsp&nbsp"+ hh + ":" + mm + ":" + ss + " " ;
    return (str);
}



addTime();
function addTime(){
    var day=86400000;
    $('.plTime').each(function(i){
        var date=new Date().getTime();
        var n=Math.floor($('.plTime').eq(i).parents('p').find('.red').text()/300);
        var newDate=date-(day*n);
        var mm=new Date(newDate).getMonth()+1;
        var dd=new Date(newDate).getDate();
        $('.plTime').eq(i).text(mm+'-'+dd);
    })
}

//诊断表格滚动

function scrollTable() {
    var i = 1;
    var len = $('.table2bg tr').length;
    $('.table2bg table').append($('.table2bg tr').clone());
    var _table = $('.table2bg table').eq(0);
    setInterval(function() {
        var num = (-0.7 * i).toFixed(2);
        _table.css('marginTop',num +'rem');
        i++;
        if (i == len + 1) {
            _table.css('transition', 'none').css('marginTop', 0);
            i = 1;
            setTimeout(function() {
                _table.css('transition', 'all  2.5s linear');
            }, 800);
        }
    }, 2500);
}


$('.zhegai').on('click',function(){
    $('.circle6,.circle5,.circle4,.circle3,.circle2').click();
});

//获取当前月份和日期
var btnTime = "";
var yearTime = "";
setYearTime();
function setYearTime(){
    var time = new Date();
    var year = time.getFullYear();
    var month = time.getMonth()+1;
    var date = time.getDate();
    var hour = time.getHours();
    var minute = time.getMinutes();
    var second = time.getSeconds();
    var day = time.getDay();
    return yearTime = [year,month,date,hour,minute,second,day];
}

//获取按钮的时间（考虑到停牌时没有时间）
function getThisTime(){
    if($('.btnBg_time').length == 0){
        addMonthDate();
    }
    return btnTime;
}

//设置时间(年)
$(".yearTime").text(yearTime[0]);

function changeSpeedTxt(num,measureName) {
    // alert(measureName);
    if(measureName){
        $('.load_discuss').css("padding","0.56rem 0 0.3rem");
    }else{
        $('.load_discuss').css("padding","0.56rem 0");
    }
    if(num == 1){
        $('.loadName').html(measureName);
        $('.loadCode').html($('.gCode').html().substring(0,7));
    }else{
        $('.load_discuss').html('正在从<span style="color:#3d8de3"></span>数据库中调取<span style="color:#3d8de3">'+measureName+'</span>结果<span class="spot">.</span>');
    }
}
//位查看多少位
addpeople();
function addpeople() {
    var num = $('#number');
    setInterval(function() {
        var penum = parseInt(num.text());
        var pe = Math.floor(Math.random() * 5 + 1);
        num.text(penum + pe);
    }, 3000);
}
var day = new Date();
var hour = day.getHours();
var minu = day.getMinutes();
var no_number = hour*252+minu*5;
var num2 =  $('#no_number');
var num3 =  $('#no_number1');
var loop = "";
num2.text(no_number);
num3.text(no_number);
loop = setInterval(function() {
    var penum2 = parseInt(num3.text());
    // alert('penum2' + penum2)
    // var pe2 = Math.floor(Math.random() * 5 + 1);
    var  pe2 = minu;
    // alert('pe2' + pe2)
    num2.text(penum2 + pe2);
    num3.text(penum2 + pe2);
}, 15000);

$(function () {
    var num = 3414317;
    var _i = 0;
    var loop = setInterval(function () {
      _i = _i + 1000000;
      if (_i >= num) {
        _i = num;
        clearInterval(loop);
  
        magic_number(num);
        setTimeout(function () {
          magic_number(num);
        }, 100);
        return;
      }
      magic_number(_i);
    }, 500);
  })

function magic_number(value) {
    var num = $("#number");
    num.animate({
      count: value
    }, {
      duration: 500,
      step: function () {
        num.text(String(parseInt(this.count)));
      }
    });
  }


//根据涨跌幅显示指定文案
function riseFallTxt(now,rise,fall,stop) {
    if(stop){
        $('.p3_txt').text("停牌原因");
        $('.p5_txt').text("复牌时间");
        $('.p3_txt').prev().show();
        $('.trend_txt').text('复牌走势预测');
        $('.dialog3_txt').text('复牌走势预测已调取完成');
    }else{
        if(now == rise){
            $('.p3_txt').text("涨停原因");
            $('.p3_txt').prev().show();
            $('.p5_txt').text("最佳买卖点");
        }else if(now == fall){
            $('.p3_txt').html("超跌<br/>反弹时间");
            $('.p3_txt').prev().hide();
            $('.p5_txt').text("最佳买卖点");
        }else{
            $('.p3_txt').html("超跌<br/>反弹时间");
            $('.p3_txt').prev().hide();
            $('.p5_txt').text("最佳买卖点");
        }
        $('.dialog3_txt').html("<span class='btnBg_time'></span><span class=''>预测报告已调取完成</span>");
        $('.trend_txt').html("<span class='btnBg_time'></span>预测报告");
        addMonthDate();
    }
}

//翻倍潜力股
if( sessionStorage.getItem("detain_btn") == undefined){
    sessionStorage.setItem("detain_btn",1);
}
$('.detain_btn').click(function () {
    kind = 1;
    try{
        if (sessionStorage.getItem("detain_btn") != undefined && sessionStorage.getItem("detain_btn") == 1) {
            sessionStorage.setItem("detain_btn", 0);
            setBrowseInfo('Wait-wanliuye','clk');
        }
    }catch (e){
    }
    doAjax_jiangu('detain_input');
});

//单独的加载动画
function aloneLoad(Popup){
showSpeed();
var load_discuss2 = $(".load_discuss2");
$(".loadBox2").show();
setTimeout(function(){
    load_discuss2.css("height","0.72rem");
},400);
setTimeout(function(){
    load_discuss2.css("height","1.08rem");
},800);
setTimeout(function(){
    load_discuss2.css("height","1.44rem");
},1200);
setTimeout(function(){
    load_discuss2.css("height","1.8rem");
},1600);
setTimeout(function(){
    load_discuss2.css("height","2.16rem");
},2000);
setTimeout(function(){
    load_discuss2.css("height","0.36rem");
    $(".loadBox2").hide();
    Popup.show();
},2700);
}

//边框动画效果
function showSpeed (){
$(".ico3").stop();
$(".ico3").animate({
    left: "3rem"
},700,"linear",function(){
    $(".ico3").css({"left":"auto","width":"0.07rem","height":"2.8rem","top":"-2.8rem","right":"0.01rem","background":"-webkit-linear-gradient(left,#fe790d,#ffc158)","background":"linear-gradient(to right,#fe790d,#ffc158)"});
});
$(".ico3").animate({
    top: "3.2rem"
},700,"linear",function(){
    $(".ico3").css({"top":"auto","width":"3rem","height":"0.07rem","right":"-4rem","bottom":"0.01rem","background":"-webkit-linear-gradient(left,#ffc158,#fe790d)","background":"linear-gradient(to right,##ffc158,#fe790d)"});
});
$(".ico3").animate({
    right: "5.3rem"
},700,"linear",function(){
    $(".ico3").css({"right":"auto","width":"0.07rem","height":"2rem","top":"3.25rem","left":"0","background":"-webkit-linear-gradient(left,#fe790d,#ffc158)","background":"linear-gradient(to right,#fe790d,#ffc158)"});
});
$(".ico3").animate({
    top: "-2rem"
},700,"linear",function(){
    $(".ico3").css({"width":"3rem","height":"0.07rem","top":"0","left":"-3rem","background":"-webkit-linear-gradient(left,#ffc158,#fe790d)","background":"linear-gradient(to right,#ffc158,#fe790d)"});
});
}

// 收盘，开盘判断

function discText () {
var newDate = new Date();
var geDays = newDate.getDay();
var dtext = '收盘';
// 判断日期是否处于周末
if (geDays > 0  && geDays < 6 ) {
    var dates = newDate.toLocaleDateString();
    var nowtime = newDate.getTime();
    var time930 = new Date(dates).getTime() + 9 * 60 * 60 * 1000 + 30 * 60 * 1000;
    var time1130 = new Date(dates).getTime() + 11 * 60 * 60 * 1000 + 30 * 60 * 1000;
    var time1300 = new Date(dates).getTime() + 13 * 60 * 60 * 1000;
    var time1500 = new Date(dates).getTime() + 15 * 60 * 60 * 1000;
    //console.log(newDate) 
    if ((nowtime >= time930 && nowtime <= time1130) || (nowtime >= time1300 && nowtime <= time1500)) {
        dtext = '交易中';
    }
}
$('#labelState').text(dtext);
setTimeout(function () {
discText();
}, 1000)
};