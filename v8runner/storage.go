package v8runner


import (
	//log "github.com/sirupsen/logrus"
	"fmt"
)

const (
	// Типы протоколов подключения

	ТипПротоколПодключенияHTTP = ТипПротоколПодключения("http")
	ТипПротоколПодключенияTCP = ТипПротоколПодключения("tcp")
	ТипПротоколПодключенияFILE = ТипПротоколПодключения("file")

	// Права пользоателей

	правоТолькоЧтение = ЗначениеПравоПользователяХранилища("ReadOnly")
	правоЗахватаОбъектов = ЗначениеПравоПользователяХранилища("LockObjects")
	правоИзмененияВерсий = ЗначениеПравоПользователяХранилища("ManageConfigurationVersions")
	правоАдминистрирования = ЗначениеПравоПользователяХранилища("Administration")
)

type ПравоПользователяХранилища interface {
	ToString() string
}

type ТипПротоколПодключения string
type ЗначениеПравоПользователяХранилища string

var ПраваПользователяХранилища  = []ЗначениеПравоПользователяХранилища{правоТолькоЧтение, правоЗахватаОбъектов, правоИзмененияВерсий, правоАдминистрирования}
var ТипыПротоколаПодключения =  []ТипПротоколПодключения{ТипПротоколПодключенияFILE, ТипПротоколПодключенияHTTP, ТипПротоколПодключенияTCP}


type ХранилищеКонфигурации struct {
	
	Конфигуратор *Конфигуратор
	СтрокаПодключения string
	Пользователь string
	Пароль string
	ТипПротоколПодключения ТипПротоколПодключения

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

func (s *ЗначениеПравоПользователяХранилища)ToString() (string) {

	stringValue := fmt.Sprintf("%s",s)

	return stringValue
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

// ConfigurationRepositoryCreate
func (storage *ХранилищеКонфигурации) ПодключитьсяКХранилищюКонфигурации (ИгнорироватьНаличиеПодключеннойБД bool, ЗаменитьКонфигурациюБД bool ) (err error) {


	var c = storage.Конфигуратор.СтандартныеПараметрыЗапускаКонфигуратора()

	c = append(c, storage.ПараметрыПодключения()...)
	c = append(c, "/ConfigurationRepositoryBindCfg", "-NoBind", )

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
func (storage *ХранилищеКонфигурации) ДобавитьПользователяВХранилище (Пользователь string, Пароль string, ПраваПользователя ПравоПользователяХранилища, ВосстановитьУдаленного bool) (err error) {


	var c = storage.Конфигуратор.СтандартныеПараметрыЗапускаКонфигуратора()

	c = append(c, storage.ПараметрыПодключения()...)
	c = append(c, "/ConfigurationRepositoryAddUser", "-NoBind", )
	c = append (c, "-User", Пользователь)

	if ЗначениеЗаполнено(Пароль) {
		c = append (c, "-Pwd", Пароль)
	}

	c = append(c, "-Rights", ПраваПользователя.ToString())

	if ВосстановитьУдаленного {
		c = append (c, "-RestoreDeletedUser")
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