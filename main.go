package main


import (
	"flag"
	"./v8runner"
	log "github.com/sirupsen/logrus"
	//"os"
)


func main() {
	flag.Parse()

	log.SetLevel(log.DebugLevel)

	var c, _ = v8runner.НовыйКонфигуратор()

	//ctx := c.ПолучитьКонтекст()

	p := "/home/khorevaa/1c/tempDB"
	//os.RemoveAll(p)




	c.СоздатьФайловуюБазуПоУмолчанию(p)

	v8runner.ОчиститьВременныйКаталог()
	//fmt.Println("Версия платформы:", c.ВерсияПлатформы.V8)
}

