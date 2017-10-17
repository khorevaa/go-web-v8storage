package v8runner


import (
	//log "github.com/sirupsen/logrus"
	//"fmt"
	"time"
)

const (
	// Типы протоколов подключения

	ТипПротоколПодключенияHTTP = ТипПротоколПодключения("http")
	ТипПротоколПодключенияTCP = ТипПротоколПодключения("tcp")
	ТипПротоколПодключенияFILE = ТипПротоколПодключения("file")

	// Права пользоателей

	правоТолькоЧтение = правоПользователяХранилища("ReadOnly")
	правоЗахватаОбъектов = правоПользователяХранилища("LockObjects")
	правоИзмененияВерсий = правоПользователяХранилища("ManageConfigurationVersions")
	правоАдминистрирования = правоПользователяХранилища("Administration")
)


type ТипПротоколПодключения string
type правоПользователяХранилища string

var ПраваПользователяХранилища  = []правоПользователяХранилища{правоТолькоЧтение, правоЗахватаОбъектов, правоИзмененияВерсий, правоАдминистрирования}
var ТипыПротоколаПодключения =  []ТипПротоколПодключения{ТипПротоколПодключенияFILE, ТипПротоколПодключенияHTTP, ТипПротоколПодключенияTCP}


type ХранилищеКонфигурации struct {
	
	*Конфигуратор
	СтрокаПодключения string
	Пользователь string
	Пароль string
	ТипПротоколПодключения ТипПротоколПодключения

}

type ИсторияХранилищаКонфигурации struct {

	*ХранилищеКонфигурации
	СписокАвторов []string
	ТаблицаВерсийИстории map[int]ВерсияИсторииХранилища

}

type ВерсияИсторииХранилища struct {
	Номер int
	Дата DateTime
	Время time.Time
	Автор string
	Комментарий string
}


func НовоеХранилищеКонфигурацииПоСтрокеПодключения(СтрокаПодключения string, Пользователь string, Пароль string) *ХранилищеКонфигурации{

	return НовоеХранилищеКонфигурации(НовыйКонфигуратор(), СтрокаПодключения, Пользователь, Пароль)

}
func НовоеХранилищеКонфигурации(Конфигуратор *Конфигуратор,СтрокаПодключения string, Пользователь string, Пароль string) *ХранилищеКонфигурации{

	return &ХранилищеКонфигурации{
		Конфигуратор,
		СтрокаПодключения,
		Пользователь,
		Пароль,
		ОпределитьПротоколПодключенияПоСтрокеПодключения(СтрокаПодключения),

	}

}

func ОпределитьПротоколПодключенияПоСтрокеПодключения(строкаПодключения string) (тип ТипПротоколПодключения) {

	тип = ТипПротоколПодключенияFILE

	return
}

func (storage *ХранилищеКонфигурации) ПараметрыПодключения() (s []string) {

	s = append(s, "/ConfigurationRepositoryF", storage.СтрокаПодключения)
	s = append(s, "/ConfigurationRepositoryS", storage.Пользователь)

	if ЗначениеЗаполнено(storage.Пароль) {
		s = append(s, "/ConfigurationRepositoryP", storage.Пароль)
	}

	return
}

func (storage *ХранилищеКонфигурации) ПараметрыПодключенияДляКопирования() (s []string) {

	s = append(s, "-Path", storage.СтрокаПодключения)
	s = append(s, "-User", storage.Пользователь)

	if ЗначениеЗаполнено(storage.Пароль) {
		s = append(s, "-Pwd", storage.Пароль)
	}


	return
}

// ConfigurationRepositoryCopyUsers
func (storage *ХранилищеКонфигурации) КопироватьПользователейИзХранилища (ХранилищеИсходное *ХранилищеКонфигурации, ВосставновитьУдаленных bool) (err error){


	var c = storage.Конфигуратор.СтандартныеПараметрыЗапускаКонфигуратора()

	c = append(c, "/ConfigurationRepositoryCopyUsers")

	c = append(c, ХранилищеИсходное.ПараметрыПодключенияДляКопирования()...)

	c = append(c, storage.ПараметрыПодключения()...)

	if ВосставновитьУдаленных {
		c = append (c, "-RestoreDeletedUser")
	}

	err = storage.Конфигуратор.ВыполнитьКоманду(c)
	return
}


// ConfigurationRepositoryCreate
func (storage *ХранилищеКонфигурации) СоздатьХранилищеКонфигурации (ДополнительныеПараметры ...string) (err error) {


	var c = storage.Конфигуратор.СтандартныеПараметрыЗапускаКонфигуратора()

	c = append(c, storage.ПараметрыПодключения()...)
	c = append(c, "/ConfigurationRepositoryCreate", "-NoBind", )

	c = append(c, ДополнительныеПараметры...)

	err = storage.Конфигуратор.ВыполнитьКоманду(c)

	return
}

// ConfigurationRepositoryBindCfg
func (storage *ХранилищеКонфигурации) ПодключитьсяКХранилищюКонфигурации (ИгнорироватьНаличиеПодключеннойБД bool, ЗаменитьКонфигурациюБД bool ) (err error) {


	var c = storage.Конфигуратор.СтандартныеПараметрыЗапускаКонфигуратора()

	c = append(c, storage.ПараметрыПодключения()...)
	c = append(c, "/ConfigurationRepositoryBindCfg")

	if ИгнорироватьНаличиеПодключеннойБД {
		c = append (c, "-forceBindAlreadyBindedUser")
	}

	if ЗаменитьКонфигурациюБД {
		c = append (c, "-forceReplaceCfg")
	}

	err = storage.Конфигуратор.ВыполнитьКоманду(c)

	return
}

// ConfigurationRepositoryCreate
func (storage *ХранилищеКонфигурации) ДобавитьПользователяВХранилище (Пользователь string, Пароль string, ПраваПользователя правоПользователяХранилища, ВосстановитьУдаленного bool) (err error) {


	var c = storage.Конфигуратор.СтандартныеПараметрыЗапускаКонфигуратора()

	c = append(c, storage.ПараметрыПодключения()...)
	c = append(c, "/ConfigurationRepositoryAddUser", "-NoBind", )
	c = append (c, "-User", Пользователь)

	if ЗначениеЗаполнено(Пароль) {
		c = append (c, "-Pwd", Пароль)
	}

	c = append(c, "-Rights", string(ПраваПользователя))

	if ВосстановитьУдаленного {
		c = append (c, "-RestoreDeletedUser")
	}

	err = storage.Конфигуратор.ВыполнитьКоманду(c)

	return
}

// ConfigurationRepositoryDumpCfg
func (storage *ХранилищеКонфигурации) СохранитьВерсиюКонфигурацииВФайл(НомерВерсии int, ИмяФайлаКонфигурации string) (err error) {


	var c = storage.Конфигуратор.СтандартныеПараметрыЗапускаКонфигуратора()

	c = append(c, storage.ПараметрыПодключения()...)
	c = append(c, "/ConfigurationRepositoryDumpCfg", ИмяФайлаКонфигурации)

	if НомерВерсии == 0 {
		c = append (c, "-v", string(НомерВерсии))
	}

	err = storage.Конфигуратор.ВыполнитьКоманду(c)

	return
}

// ConfigurationRepositoryDumpCfg
func (storage *ХранилищеКонфигурации) ПоследняяВерсияКонфигурацииВФайл(ИмяФайлаКонфигурации string) (err error) {

	err = storage.СохранитьВерсиюКонфигурацииВФайл(0, ИмяФайлаКонфигурации)

	return
}

// ConfigurationRepositoryOptimizeData
func (storage *ХранилищеКонфигурации) ОптимизироватьБазуХранилища() (err error) {


	var c = storage.Конфигуратор.СтандартныеПараметрыЗапускаКонфигуратора()

	c = append(c, storage.ПараметрыПодключения()...)

	c = append(c, "/ConfigurationRepositoryOptimizeData")

	err = storage.Конфигуратор.ВыполнитьКоманду(c)

	return
}

// ConfigurationRepositoryReport
func (storage *ХранилищеКонфигурации) ПолучитьОтчетПоВерсиям(ПутьКФайлуОтчета string, НомерНачальнойВерсии int, НомерКонечнойВерсии int) (err error) {


	var c = storage.Конфигуратор.СтандартныеПараметрыЗапускаКонфигуратора()

	c = append(c, storage.ПараметрыПодключения()...)

	c = append(c, "/ConfigurationRepositoryReport", ПутьКФайлуОтчета)

	if ПустаяСтрока(string(НомерНачальнойВерсии))  {

		НомерНачальнойВерсии = 1
	}

	c = append(c, "-NBegin", string(НомерНачальнойВерсии))

	if НомерКонечнойВерсии == 0 {
		c = append (c, "-NEnd", string(НомерКонечнойВерсии))
	}

	err = storage.Конфигуратор.ВыполнитьКоманду(c)

	return
}

// ConfigurationRepositoryCreate
func (conf *Конфигуратор) СоздатьХранилищеКонфигурации(ПутьКХранилищю string, Пользователь string, Пароль string, ДополнительныеПараметры ...string) (storage *ХранилищеКонфигурации, err error) {

	storage = НовоеХранилищеКонфигурации(conf, ПутьКХранилищю, Пользователь, Пароль)

	err = storage.СоздатьХранилищеКонфигурации(ДополнительныеПараметры...)

	return
}

// ConfigurationRepositoryBindCfg
func (conf *Конфигуратор) ПодключитьсяКХранилищюКонфигурации(ПутьКХранилищю string, Пользователь string, Пароль string, ИгнорироватьНаличиеПодключеннойБД bool, ЗаменитьКонфигурациюБД bool ) (storage *ХранилищеКонфигурации, err error) {

	storage = НовоеХранилищеКонфигурации(conf, ПутьКХранилищю, Пользователь, Пароль)

	err = storage.ПодключитьсяКХранилищюКонфигурации(ИгнорироватьНаличиеПодключеннойБД, ЗаменитьКонфигурациюБД)

	return
}