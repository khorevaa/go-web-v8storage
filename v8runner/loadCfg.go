package v8runner

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"./dumpMode"
)

func (conf *Конфигуратор) ЗагрузитьКонфигурациюИзФайлов(КаталогЗагрузки string) (e error) {
	return conf.loadConfigFromFiles(КаталогЗагрузки, "", "", false)
}

func (conf *Конфигуратор) ЗагрузитьКонфигурациюИзФайла(ПутьКФайлуКонфигуации string) (e error) {
	return conf.loadCfg(ПутьКФайлуКонфигуации)
}

// private func

func (conf *Конфигуратор) loadCfg(cfg string) (e error) {

	var c = conf.СтандартныеПараметрыЗапускаКонфигуратора()

	c = append(c, fmt.Sprintf("/LoadCfg %s", cfg))

	err := conf.выполнить(c)

	return err
}

func (conf *Конфигуратор) loadConfigFromFiles(dir string, pListFile string, format string, updDumpInfo bool) (e error) {

	var c = conf.СтандартныеПараметрыЗапускаКонфигуратора()

	c = append(c, fmt.Sprintf("/LoadConfigFromFiles %s", dir))

	if ЗначениеЗаполнено(pListFile) {

		if ok, _ := РежимВыгрузкиКонфигурации.РежимДоступен(format); ok {
			c = append(c, fmt.Sprintf("-format %s", format))
		} else {
			return
		}
		c = append(c, fmt.Sprintf("-listFile %s", pListFile))

		if updDumpInfo {
			//Если ПроверитьВозможностьОбновленияФайловВыгрузки(КаталогВыгрузки, ПутьКФайлуВерсийДляСравнения, ФорматВыгрузки) Тогда
			c = append(c, "-updateConfigDumpInfo", "-force")
		}

	}

	log.Debugf("Параметры запуска: %s", c)

	err := conf.выполнить(c)

	return err
}
