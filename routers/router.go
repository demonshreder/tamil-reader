package routers

import (
	"github.com/astaxie/beego"
	"github.com/demonshreder/tamil-reader/controllers"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	beego.Router("/new", &controllers.NewController{})
	beego.Router("/new/book", &controllers.NewBook{})
}
