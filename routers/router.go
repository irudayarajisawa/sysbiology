package routers

import (
	"github.com/irudayarajisawa/sysbiology/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
