package routers

import (
	"github.com/demonshreder/tamil-reader/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
