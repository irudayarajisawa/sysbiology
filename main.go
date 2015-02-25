package main

import (
	_ "github.com/irudayarajisawa/sysbiology/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
		// Preload values used by Beego on startup
	_ "github.com/irudayarajisawa/sysbiology/models"
	_ "github.com/irudayarajisawa/sysbiology/routers"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	mysqlReg := beego.AppConfig.String("mysqluser") + ":" +
		beego.AppConfig.String("mysqlpass") + "@tcp(127.0.0.1:3306)/" +
		beego.AppConfig.String("mysqldb")
	orm.RegisterDataBase("default", "mysql", mysqlReg)
}

func main() {
	beego.Run()
}

