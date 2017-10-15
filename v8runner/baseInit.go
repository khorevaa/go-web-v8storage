package v8runner

import "fmt"

func (conf *Конфигуратор) СоздатьФайловуюБазуПоУмолчанию(КаталогБазы string) error {
	return conf.createFileBase(КаталогБазы, "", "")
}

func (conf *Конфигуратор) СоздатьФайловуюБазуПоШаблону(КаталогБазы string, ПутьКШаблону string) error {
	return conf.createFileBase(КаталогБазы, ПутьКШаблону, "")
}

func (conf *Конфигуратор) СоздатьИменнуюФайловуюБазу(КаталогБазы string, ИмяБазыВСписке string) error {
	return conf.createFileBase(КаталогБазы, "", ИмяБазыВСписке)
}

func (conf *Конфигуратор) СоздатьИменнуюФайловуюБазуПоШаблону(КаталогБазы string, ПутьКШаблону string, ИмяБазыВСписке string) error {
	return conf.createFileBase(КаталогБазы, ПутьКШаблону, ИмяБазыВСписке)
}

func (conf *Конфигуратор) СоздатьФайловуюБазу(КаталогБазы string, ПутьКШаблону string, ИмяБазыВСписке string) error {
	return conf.createFileBase(КаталогБазы, ПутьКШаблону, ИмяБазыВСписке)
}

//
func (conf *Конфигуратор) createFileBase(dir string, pTemplate string, lName string) error {

	var p []string
	p = append(p, "CREATEINFOBASE")
	p = append(p, fmt.Sprintf("File=%s", dir))

	if ЗначениеЗаполнено(pTemplate) {
		p = append(p, fmt.Sprintf("/UseTemplate %s", pTemplate))
	}

	if ЗначениеЗаполнено(lName) {
		p = append(p, fmt.Sprintf("/AddInList %s", lName))
	}

	p = append(p, "/Out", НовыйФайлИнформации())

	return conf.выполнить(p)
}
