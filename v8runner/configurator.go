package v8runner

import (
	"fmt"
	"os/exec"
	log "github.com/sirupsen/logrus"
	"syscall"
	"github.com/pkg/errors"
	//"github.com/cweill/gotests/testdata"
	//"github.com/stretchr/testify/assert"
	//"path"
)

type Конфигуратор struct {
	Контекст              *Контекст
	ФайлИнформации        string
	ОчищатьФайлИнформации bool
	ЭтоWindows            bool
	ВерсияПлатформы       *ВерсияПлатформы
	выводКоманды          string

}

// new func

func НовыйКонфигуратор() (conf *Конфигуратор) {

	conf = &Конфигуратор{}
	conf.Контекст = НовыйКонтекст()
	conf.ФайлИнформации = НовыйФайлИнформации()
	conf.ВерсияПлатформы = ПолучитьВерсиюПоУмолчанию()

	return conf
}

func НовыйКонфигураторСКонтекстом(ctx *Контекст) (conf *Конфигуратор) {

	conf = НовыйКонфигуратор()
	conf.Контекст = ctx

	return conf
}

func НовыйКонфигураторСОпциями(opts ...func(*Конфигуратор)) (conf *Конфигуратор) {

	conf = НовыйКонфигуратор()

	for _, opt := range opts {
		opt(conf)
	}

	return conf

}

//export func

func (conf *Конфигуратор) ПолучитьКонтекст() *Контекст {
	return conf.Контекст
}

func (conf *Конфигуратор) КлючСоединенияСБазой() string {

	log.Debugf("Получен ключа соединения с базой: %s", conf.Контекст.КлючСоединенияСБазой)

	if ПустаяСтрока(conf.Контекст.КлючСоединенияСБазой) {
		return conf.КлючВременногоСоединенияСБазой()
	}

	return conf.Контекст.КлючСоединенияСБазой
}

func (conf *Конфигуратор) КлючВременногоСоединенияСБазой() string {

	//conf.Контекст.ВременнаяБаза = НоваяВременнаяБаза(tempDBPath)

	log.Debugf("Получение временного ключа соединения с базой: %s", conf.Контекст.ВременнаяБаза.КлючСоединенияСБазой )

	return conf.Контекст.ВременнаяБаза.КлючСоединенияСБазой
}


func (conf *Конфигуратор) СтандартныеПараметрыЗапускаКонфигуратора() (p []string) {

	var мОчищатьФайлИнформации bool
	var ctx *Контекст = conf.Контекст

	мОчищатьФайлИнформации = true

	p = append(p, "DESIGNER")
	p = append(p, conf.КлючСоединенияСБазой())

	p = append(p, "/Out", conf.ФайлИнформации)

	if !мОчищатьФайлИнформации {
		p = append(p, "-NoTruncate")
	}

	if ЗначениеЗаполнено(ctx.Пользователь) {
		p = append(p, fmt.Sprintf("/N %s", ctx.Пользователь))
	}
	if ЗначениеЗаполнено(ctx.Пароль) {
		p = append(p, fmt.Sprintf("/P %s", ctx.Пароль))
	}
	p = append(p, "/DisableStartupMessages")
	p = append(p, "/DisableStartupDialogs")

	return
}

// private run func
const defaultFailedCode = 1

func (conf *Конфигуратор) Выполнить (args []string) (e error) {

	var exitCode int

	procName := conf.ВерсияПлатформы.V8
	cmd := exec.Command(procName, args...) // strings.Join(args, " "))

	log.Debugf("Строка запуска: %s", cmd.Args)

	out, e := cmd.Output()

	if e != nil {
		// try to get the exit code
		if exitError, ok := e.(*exec.ExitError); ok {
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		} else {
			// This will happen (in OSX) if `name` is not available in $PATH,
			// in this situation, exit code could not be get, and stderr will be
			// empty string very likely, so we use the default fail code, and format err
			// to string and set to stderr
			log.Debugf("Could not get exit code for failed program: %v, %v", procName, args)
			exitCode = defaultFailedCode
		}
	} else {
		// success, exitCode should be 0 if go is ok
		ws := cmd.ProcessState.Sys().(syscall.WaitStatus)
		exitCode = ws.ExitStatus()
	}

	conf.установитьВыводКоманды(conf.прочитатьФайлИнформации())

	if exitCode == 1 {
		e = errors.New(conf.выводКоманды)
	}

	log.Debugf("КодЗавершения команды: %v", exitCode)
	log.Debugf("Результат выполнения команды: %s, out: %s", conf.выводКоманды, out)
	return e

}


func (conf *Конфигуратор) ВыполнитьКоманду(args []string) (e error) {

	if ok, err := conf.ПроверитьВозможностьВыполнения(); !ok {
		e = err
		return
	}

	e = conf.Выполнить(args)

	return
}

func (c *Конфигуратор) ПроверитьВозможностьВыполнения() (ok bool, err error) {

	ok = true

	if c.ВерсияПлатформы == nil {
		err = errors.Wrap(err, "Не найдена доступная версия платформы")
		ok = false
	}


	log.Debugf("КлючСоединенияСБазой: %s", c.КлючСоединенияСБазой())
	log.Debugf("ВременныйКлючСоединенияСБазой: %s", c.КлючВременногоСоединенияСБазой())
	log.Debugf("ВременнаяБазаСуществует: %v", c.Контекст.ВременнаяБаза.существует)

	if c.КлючСоединенияСБазой() == c.КлючВременногоСоединенияСБазой() && !c.Контекст.ВременнаяБаза.существует{
		c.Контекст.ВременнаяБаза.ИнициализироватьВременнуюБазу()
	}


	return

}

func (c *Конфигуратор) установитьВыводКоманды(s string) {
	c.выводКоманды = s
	log.Debugf("Установлен вывод команды 1С: %s", s)
}

func (c *Конфигуратор) прочитатьФайлИнформации() (str string) {

	log.Debugf("Читаю файла информации 1С: %s", c.ФайлИнформации)

	b, err := ReadFileUTF16(c.ФайлИнформации) // just pass the file name
	if err != nil {
		log.Debugf("Обшибка чтения файла информации 1С %s: %v", c.ФайлИнформации, err)
		str = ""
		return
		//fmt.Print(err)
	}

	str = string(b) // convert content to a 'string'

	return
}