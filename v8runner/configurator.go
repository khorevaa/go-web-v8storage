package v8runner

import (
	"fmt"
	//"os"
	//"io/ioutil"
	//"github.com/mash/go-tempfile-suffix"
	"os/exec"
	"strings"
	log "github.com/sirupsen/logrus"
	//"github.com/constabulary/gb/testdata/src/f"
)

type Конфигуратор struct {
	Контекст              *Контекст
	ФайлИнформации        string
	ОчищатьФайлИнформации bool
	путьКПлатформе1С      string
	ЭтоWindows            bool
	версияПлатформы       string
}

// new func

func НовыйКонфигуратор() (*Конфигуратор, error) {

	ctx := Контекст{}

	conf := &Конфигуратор{}
	conf.Контекст = &ctx

	var err error

	return conf, err
}

func НовыйКонфигураторСКонтекстом(ctx *Контекст) (*Конфигуратор, error) {

	conf, err := НовыйКонфигуратор()
	conf.Контекст = ctx

	return conf, err
}

func НовыйКонфигураторСОпциями(opts ...func(*Конфигуратор)) (*Конфигуратор, error) {

	conf, err := НовыйКонфигуратор()

	for _, opt := range opts {
		opt(conf)
	}

	return conf, err

}

//export func

func (conf *Конфигуратор) ПолучитьКонтекст() *Контекст {
	return conf.Контекст
}

func (conf *Конфигуратор) КлючСоединенияСБазой() string {

	if ЗначениеЗаполнено(conf.Контекст.КлючСоединенияСБазой) {
		conf.ИнициализироватьВременныйКонтекст()
	}
	return conf.Контекст.КлючСоединенияСБазой
}

func (conf *Конфигуратор) ИнициализироватьВременныйКонтекст() {

	tmpDir := ВременныйКаталог()

	var TempDB = ВременнаяБаза{
		ПутьКБазе:            tmpDir,
		КлючСоединенияСБазой: fmt.Sprintf("/F%s", tmpDir),
	}

	TempDB.ИнициализироватьВременнуюБазу()

	conf.Контекст.КлючСоединенияСБазой = TempDB.КлючСоединенияСБазой
}

func (conf *Конфигуратор) СтандартныеПараметрыЗапускаКонфигуратора() []string {

	var p []string
	var мОчищатьФайлИнформации bool
	var ctx *Контекст = conf.Контекст

	мОчищатьФайлИнформации = true

	p = append(p, "DESIGNER")
	p = append(p, "КлючСоединенияСБазой()")

	p = append(p, "/Out", НовыйФайлИнформации())

	if мОчищатьФайлИнформации {
		p = append(p, " -NoTruncate")
	}

	if ЗначениеЗаполнено(ctx.Пользователь) {
		p = append(p, fmt.Sprintf("/N %s", ctx.Пользователь))
	}
	if ЗначениеЗаполнено(ctx.Пароль) {
		p = append(p, fmt.Sprintf("/P %s", ctx.Пароль))
	}
	p = append(p, "/DisableStartupMessages")
	p = append(p, "/DisableStartupDialogs")

	return p
}

// private run func

func (conf *Конфигуратор) выполнить(args []string) error {

	binary, lookErr := exec.LookPath("echo")
	if lookErr != nil {
		panic(lookErr)
	}
	fmt.Errorf(binary)
	//env := os.Environ()
	//fmt.Errorf()
	//strings.Fields(args)
	out, execErr := exec.Command(binary, strings.Join(args, " ")).Output()
	if execErr != nil {
		panic(execErr)
	}
	log.Debugf("%s", out)
	return execErr

}
