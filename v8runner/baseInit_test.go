package v8runner

import (
	"testing"
	"github.com/stretchr/testify/suite"
	"os"
	"path"

	//"github.com/stretchr/testify/assert"
)
var pwd, _ = os.Getwd()



type MySuite struct{
	suite.Suite
	conf Конфигуратор
	КаталогБазы string
	ПутьКШаблону string
}

func (s *MySuite) SetupSuite() {

	s.conf = Конфигуратор{
		НовыйКонтекст(),
		НовыйФайлИнформации(),
		true,
		false,
		ПолучитьВерсиюПоУмолчанию(),
		"",
	}
	s.КаталогБазы = ВременныйКаталог()
	s.ПутьКШаблону = path.Join( pwd,"fixtures/ТестовыйФайлКонфигурации.cf")


// Use s.dir to prepare some data.
}

func (s *MySuite) TestКонфигуратор_СоздатьФайловуюБазуПоШаблону() {

	s.Assert().NotEmpty(s.conf.ВерсияПлатформы, "Версия не определилась")


	err := s.conf.СоздатьФайловуюБазуПоШаблону(s.КаталогБазы, s.ПутьКШаблону)

	//s.T().Errorf(err, "TestКонфигуратор_СоздатьФайловуюБазуПоШаблону")
	s.NoErrorf(err, "TestКонфигуратор_СоздатьФайловуюБазуПоШаблону: %v", err)
	//s.Zero(1, "Нули")
	s.Assert().NotEmptyf(s.conf.выводКоманды, "Вывод команды %s", s.conf.выводКоманды)
	}


// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestМойБлокТестов(t *testing.T) {
	suiteTester := new(MySuite)
	suite.Run(t, suiteTester)

}
