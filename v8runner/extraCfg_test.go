package v8runner

import (
	"testing"
	"github.com/stretchr/testify/suite"
	"path"
	log "github.com/sirupsen/logrus"

)

type тестыНаДополнительныйФункционал struct {
	suite.Suite
	conf                            *Конфигуратор
	КаталогЗагрузки                 string
	ПутьКФайлуОбработки             string

}


func (s *тестыНаДополнительныйФункционал) SetupSuite() {


	log.SetLevel(log.DebugLevel)

	s.КаталогЗагрузки = path.Join(pwd, "epf/ОбработкаКонвертацииMXLJSON/ОбработкаКонвертацииMXLJSON/ОбработкаКонвертацииMXLJSON.xml")

}

func (s *тестыНаДополнительныйФункционал) SetupTest() {

	s.conf = НовыйКонфигуратор()
	s.ПутьКФайлуОбработки = НовыйВременныйФайл("Обработка", ".epf")
}


func (s *тестыНаДополнительныйФункционал) TearDownSuite() {
	//ОчиститьВременныйКаталог()
}

func Test_ТестыНаДополнительныйФункционал(t *testing.T) {

	suiteTester := new(тестыНаДополнительныйФункционал)
	suite.Run(t, suiteTester)

}


func (s *тестыНаДополнительныйФункционал) TestКонфигуратор_СобратьОбработкуОтчетИзФайлов() {

	err := s.conf.СобратьОбработкуОтчетИзФайлов(s.КаталогЗагрузки, s.ПутьКФайлуОбработки)
	s.NoError(err, "TestКонфигуратор_СобратьОбработкуОтчетИзФайлов: %v", err)

}
