package main

import (
	//"fmt"
	//"net/http"
	//"strings"

	"github.com/gocraft/web"
	"./oscript_runner"
	//"./v8runner"
	//"./v8runner/dumpMode"
	"fmt"
	"strings"
	//"golang.org/x/tools/go/gcimporter15/testdata"
	//"log"
	//log "github.com/sirupsen/logrus"
	//"database/sql/driver"
	//"reflect"
	//"path"
	//"os"
)

type Context struct {
	HelloCount int
}

func (c *Context) SetHelloCount(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HelloCount = 3
	next(rw, req)
}

var r = oscript.Runner{
	"Еуеы",
	"ее",
}

//var Контекст  =  oscript.Контекст{
//	"соеднинениесбазой",
//	"Пользователь",
//	"Пароль",
//}

//oscript.Run()

// SayHello Doc
func (c *Context) SayHello(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, strings.Repeat("Hello ", c.HelloCount), "World!")
}

func main() {

	//var Контекст = v8runner.Контекст{
	//	"соеднинениесбазой",
	//	"Пользователь",
	//	"Пароль",
	//	"",
	//	"",
	//	"",
	//}
	//
	//log.SetLevel(log.DebugLevel)
	////log.SetFormatter(&log.JSONFormatter{})
	//
	//var v8r, err = v8runner.НовыйКонфигураторСОпциями(v8runner.ФайлИнформации(v8runner.НовыйФайлИнформации()), v8runner.УстановитьКонтекст(&Контекст))
	//if err != nil {
	//	log.Error(err)
	//}
	////v8r.ВыгрузитьКонфигурацию("v8r", РежимВыгрузкиКонфигурации.Иерархический)
	//
	//ctx := v8r.ПолучитьКонтекст()
	//ctx.Пользователь = "Проверочный пользователь"
	//
	////var v8r_2, err_2 = v8runner.НовыйКонфигураторСКонтекстом(ctx)
	////if err_2 != nil {
	////	log.Error(err_2)
	////}
	////
	////ctx.Пользователь = "Пользователь 2"
	////var m = make(map[string]interface{})
	////
	////m["ПутьКШаблону"] = "CNHJRF CNHJ"
	//
	//var МассивПутей = []string{}
	//
	////var СуффиксРасположения = path.Join("1C","1CEStart","1CEStart.cfg")
	//
	//var envs = []string{
	//	"ALLUSERSPROFILE",
	//	"APPDATA",
	//}
	//
	//for _, env := range envs {
	//	//if cat, ok := os.LookupEnv(env); ok {
	//
	//	v8runner.ДополнитьМассивРасположенийИзКонфигурационногоФайла("", &МассивПутей)
	//
	//	//}
	//	log.Debugf("env: %s", env)
	//
	//}
	//log.Debug(МассивПутей)
	//
	////return МассивПутей
	//
	////v8r_2.СоздатьФайловуюБазу("v8r_2", map[string] interface{}{
	////	"ПутьКШаблону": "ТестоваяОбция"})
	//
	/////http.ListenAndServe("localhost:3000", router) // Start the server!
}
