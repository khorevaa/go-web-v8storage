package v8runner

// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type Error interface {
	Error() string
}

const (
	prefix = "v8runner"
)

var (
	tempFiles []string
	tempDir string = ИницализороватьВременныйКаталог()
	workDir string = ВременныйКаталог()
)

func ФайлИнформации(файлИнформации string) func(*Конфигуратор) {
	return func(s *Конфигуратор) {
		s.ФайлИнформации = файлИнформации
	}
}
func УстановитьКонтекст(контекст *Контекст) func(*Конфигуратор) {
	return func(s *Конфигуратор) {
		s.Контекст = контекст
	}
}
