package v8runner

import "fmt"

func (conf *Конфигуратор) ЗапуститьВРежимеПредприятияCКлючемЗапуска(КлючЗапуска string, УправляемыйРежим bool, ДополнительныеПараметры ...string) (err error) {

	ДополнительныеПараметры = append(ДополнительныеПараметры, fmt.Sprintf("/C%s", КлючЗапуска))

	err = conf.ЗапуститьВРежимеПредприятия(УправляемыйРежим, ДополнительныеПараметры...)

	return
}

func (conf *Конфигуратор) ЗапуститьВРежимеПредприятия(УправляемыйРежим bool, ДополнительныеПараметры ...string) (err error) {

	var c = conf.СтандартныеПараметрыЗапускаКонфигуратора()
	c[0] = "ENTERPRISE"

	if УправляемыйРежим {
		c[2] = "/RunModeManagedApplication"
	} else {
		c[2] = "/RunModeOrdinaryApplication"
	}
	c = append(c, ДополнительныеПараметры...)

	err = conf.ВыполнитьКоманду(c)

	return
}