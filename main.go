package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/demonshreder/tamil-reader/models"
	_ "github.com/demonshreder/tamil-reader/routers"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)

	orm.RegisterDataBase("default", "postgres", "user=tamil dbname=tamil_reader sslmode=disable")

	orm.DefaultTimeLoc = time.Local
	orm.RunCommand()
}
func main() {

	beego.Run()
	orm.RunCommand()
	o := orm.NewOrm()
	o.Using("default")
<<<<<<< HEAD
	err := orm.RunSyncdb("default", true, true)
	if err != nil {
		fmt.Println(err)
	}
=======
>>>>>>> cb86ddb3ee3187f8690a5ebef792c882c2310ad4
}
