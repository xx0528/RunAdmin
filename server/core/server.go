/*
 * @Author: xx
 * @Date: 2023-04-24 10:53:04
 * @LastEditTime: 2023-05-18 16:43:14
 * @Description:
 */
package core

import (
	"fmt"
	"html/template"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/initialize"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}

	// 从db加载jwt数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}

	Router := initialize.Routers()
	Router.Static("/form-generator", "./resource/page")
	Router.Static("/temp-statics", "./resource/tpl")

	//网页渲染
	// Router.HTMLRender = ginview.Default()
	Router.HTMLRender = ginview.New(goview.Config{
		Root:         "resource/tpl",
		Extension:    ".html",
		Master:       "layouts/master",
		Funcs:        template.FuncMap{},
		DisableCache: true,
	})

	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf(`
	欢迎使用 run-admin
	当前版本:v2.5.5
`, address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
