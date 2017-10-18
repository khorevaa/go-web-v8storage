package v8runner

import "github.com/DATA-DOG/godog"

func stepDefinition1() error {
	return godog.ErrPending
}

func stepDefinition2() error {
	return godog.ErrPending
}

func stepDefinition3(arg1 string) error {
	return godog.ErrPending
}

func stepDefinition4(arg1 string) error {
	return godog.ErrPending
}

func stepDefinition5() error {
	return godog.ErrPending
}

func stepDefinition6(arg1, arg2 string) error {
	return godog.ErrPending
}

func stepDefinition9(arg1 string) error {
	return godog.ErrPending
}

func stepDefinition10(arg1 string) error {
	return godog.ErrPending
}

func stepDefinition11(arg1 string) error {
	return godog.ErrPending
}

func stepDefinition12(arg1 string) error {
	return godog.ErrPending
}

func stepDefinition13() error {
	return godog.ErrPending
}

func stepDefinition21() error {
	return godog.ErrPending
}

func stepDefinition25(arg1 string) error {
	return godog.ErrPending
}

func stepDefinition44() error {
	return godog.ErrPending
}

func stepDefinition46() error {
	return godog.ErrPending
}

func stepDefinition47() error {
	return godog.ErrPending
}

func stepDefinition54(arg1 string) error {
	return godog.ErrPending
}

func stepDefinition62(arg1, arg2 string) error {
	return godog.ErrPending
}

func json() error {
	return godog.ErrPending
}

func json() error {
	return godog.ErrPending
}

func stepDefinition77(arg1 string) error {
	return godog.ErrPending
}

func stepDefinition78() error {
	return godog.ErrPending
}

func stepDefinition85() error {
	return godog.ErrPending
}

func stepDefinition88() error {
	return godog.ErrPending
}

func stepDefinition98(arg1, arg2, arg3 string) error {
	return godog.ErrPending
}

func stepDefinition102(arg1 string) error {
	return godog.ErrPending
}

func stepDefinition103(arg1 string) error {
	return godog.ErrPending
}

func stepDefinition104() error {
	return godog.ErrPending
}

func stepDefinition124() error {
	return godog.ErrPending
}

func stepDefinition163(arg1, arg2, arg3 string) error {
	return godog.ErrPending
}

func stepDefinition168(arg1 string) error {
	return godog.ErrPending
}

func stepDefinition181() error {
	return godog.ErrPending
}

func stepDefinition184(arg1 string) error {
	return godog.ErrPending
}

func stepDefinition185(arg1 string) error {
	return godog.ErrPending
}

func stepDefinition188() error {
	return godog.ErrPending
}

func stepDefinition189() error {
	return godog.ErrPending
}

func stepDefinition230(arg1 string) error {
	return godog.ErrPending
}

func stepDefinition231(arg1 string) error {
	return godog.ErrPending
}

func stepDefinition232(arg1, arg2 string) error {
	return godog.ErrPending
}

func stepDefinition233(arg1 string) error {
	return godog.ErrPending
}

func stepDefinition234(arg1, arg2 string) error {
	return godog.ErrPending
}

func stepDefinition235(arg1, arg2 string) error {
	return godog.ErrPending
}

func stepDefinition236(arg1 string, arg2 int) error {
	return godog.ErrPending
}

func stepDefinition243() error {
	return godog.ErrPending
}

func stepDefinition244() error {
	return godog.ErrPending
}

func stepDefinition245(arg1 string) error {
	return godog.ErrPending
}

func stepDefinition246() error {
	return godog.ErrPending
}

func stepDefinition247(arg1 string) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^Я создаю новый объект МенеджерХранилищаКонфигурации$`, stepDefinition1)
	s.Step(`^Я создаю временный каталог и сохраняю его в контекст$`, stepDefinition2)
	s.Step(`^Я сохраняю значение временного каталога в переменной "([^"]*)"$`, stepDefinition3)
	s.Step(`^Я создаю временную базу в каталоге из переменной "([^"]*)"$`, stepDefinition4)
	s.Step(`^я устанавливаю контекст выполнения конфигуратора$`, stepDefinition5)
	s.Step(`^Я устанавливаю параметры авторитизации пользователя "([^"]*)" и пароль "([^"]*)"$`, stepDefinition6)
	s.Step(`^Я устанавливаю каталог хранилища из переменной "([^"]*)"$`, stepDefinition9)
	s.Step(`^Я загружаю файл конфигурации "([^"]*)" в базу данных$`, stepDefinition10)
	s.Step(`^Я создаю файловое хранилище с параметром подключения базы к хранилищу "([^"]*)"$`, stepDefinition11)
	s.Step(`^Вывод лога содержит "([^"]*)"$`, stepDefinition12)
	s.Step(`^Я выполняю отключение от хранилища конфигурации$`, stepDefinition13)
	s.Step(`^Я копирую тестовое хранилище во временный каталог$`, stepDefinition21)
	s.Step(`^Я подключаю базу к хранилищу с параметром замены конфигурации "([^"]*)"$`, stepDefinition25)
	s.Step(`^Я устанавливаю каталог хранилища во временный каталог$`, stepDefinition44)
	s.Step(`^Я получаю отчет из хранилища$`, stepDefinition46)
	s.Step(`^Файл отчета существует$`, stepDefinition47)
	s.Step(`^Я получаю отчет из хранилища начиная с версии "([^"]*)"$`, stepDefinition54)
	s.Step(`^Я получаю отчет из хранилища начиная с "([^"]*)" по "([^"]*)" версию$`, stepDefinition62)
	s.Step(`^Я конвертирую файл отчета в json$`, json)
	s.Step(`^Файл отчета в формате json существует$`, json)
	s.Step(`^Я получаю файл конфигурации версии "([^"]*)" из хранилища$`, stepDefinition77)
	s.Step(`^Файл конфигурации существует$`, stepDefinition78)
	s.Step(`^Я получаю файл конфигурации последней версии из хранилища$`, stepDefinition85)
	s.Step(`^Я создаю новый объект СписокОбъектовКонфигурации$`, stepDefinition88)
	s.Step(`^Я добавляю пользователя "([^"]*)" с паролем "([^"]*)" и правом "([^"]*)"$`, stepDefinition98)
	s.Step(`^Я добавляю в список объектов захват корня конфигурации с включением подчиненных "([^"]*)"$`, stepDefinition102)
	s.Step(`^Я записываю список объектов конфигурации во временный файл и сохраняю значение в переменной "([^"]*)"$`, stepDefinition103)
	s.Step(`^Я выполняю ошибочную попытку захват объектов в хранилище конфигурации по файлу списка объектов$`, stepDefinition104)
	s.Step(`^Я выполняю захват объектов в хранилище конфигурации по файлу списка объектов$`, stepDefinition124)
	s.Step(`^Я копирую пользователей из хранилища "([^"]*)" с пользователей "([^"]*)" и паролем "([^"]*)"$`, stepDefinition163)
	s.Step(`^Я выполняю попытку подключения базы к хранилищу с параметром замены конфигурации "([^"]*)"$`, stepDefinition168)
	s.Step(`^Я подключаю базу к хранилищу$`, stepDefinition181)
	s.Step(`^Я выгружаю конфигурацию в каталог из переменной "([^"]*)"$`, stepDefinition184)
	s.Step(`^Я формирую список объектов конфигурации для каталога из переменной "([^"]*)"$`, stepDefinition185)
	s.Step(`^Все объекты успешно захвачены$`, stepDefinition188)
	s.Step(`^Я отменяю захват в хранилище$`, stepDefinition189)
	s.Step(`^Я выключаю отладку лога с именем "([^"]*)"$`, stepDefinition230)
	s.Step(`^Я очищаю параметры команды "([^"]*)" в контексте$`, stepDefinition231)
	s.Step(`^Я добавляю параметр "([^"]*)" для команды "([^"]*)"$`, stepDefinition232)
	s.Step(`^Я выполняю команду "([^"]*)"$`, stepDefinition233)
	s.Step(`^Вывод команды "([^"]*)" содержит "([^"]*)"$`, stepDefinition234)
	s.Step(`^Вывод команды "([^"]*)" не содержит "([^"]*)"$`, stepDefinition235)
	s.Step(`^Код возврата команды "([^"]*)" равен (\d+)$`, stepDefinition236)
	s.Step(`^Я читаю данные из хранилища$`, stepDefinition243)
	s.Step(`^Я получаю таблицу версий хранилища$`, stepDefinition244)
	s.Step(`^Таблица версий содержит "([^"]*)" записи$`, stepDefinition245)
	s.Step(`^Я получаю массив авторов хранилища$`, stepDefinition246)
	s.Step(`^Количество в массиве авторов равно "([^"]*)"$`, stepDefinition247)
}

