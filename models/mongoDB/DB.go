package mongoDB

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"log"
)
var Dbsession *mgo.Session
func init() {
   url:=beego.AppConfig.String("mongodb::dburl")
   session,err:=mgo.Dial(url)
   if err!=nil{
   	log.Fatal("数据库连接失败")
   	return
   }
   Dbsession=session
}
