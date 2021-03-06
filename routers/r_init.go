package routers

import (
	"github.com/astaxie/beego"
	"testapi/controllers"
)

func init() {
	// api路由,一级
	ns := beego.NewNamespace("/test",
		// api路由,二级
		beego.NSNamespace("/api",
			//api路由,三级
			beego.NSNamespace(
				"/v1",
				//测试接口
				beego.NSRouter("/", &controllers.API{}, "get:Hello"),
				// 查询用户, 分页
				beego.NSRouter("/users", &controllers.API{}, "get:SelectUsers"),
				// 查询用户
				beego.NSRouter("/user", &controllers.API{}, "get:SelectUser"),
				// 新增用户
				beego.NSRouter("/user", &controllers.API{}, "put:AddUser"),
				// 更新用户
				beego.NSRouter("/user", &controllers.API{}, "post:UpdateUser"),
				// 删除用户
				beego.NSRouter("/user", &controllers.API{}, "delete:DelUser"),
			),
		),
	)
	// 注册自定义namespace
	beego.AddNamespace(ns)

	// 执行过滤器
	RouterFilter()
}
