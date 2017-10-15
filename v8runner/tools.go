package v8runner

import (
	"os"
	"io/ioutil"
	"github.com/mash/go-tempfile-suffix"
	"github.com/shomali11/util/strings"
)

func ВременныйКаталог() string {

	userTmpDir := os.TempDir()
	tmpDir, err := ioutil.TempDir(userTmpDir, prefix)
	if err != nil {
		panic(err)
	}
	return tmpDir
}

func ЗначениеЗаполнено(Значение string) bool {
	return !strings.IsEmpty(Значение)
}

func ПустаяСтрока(Значение string) bool {
	return strings.IsEmpty(Значение)
}

func НовыйФайлИнформации() string {

	userTmpDir := os.TempDir()
	tmpDir, err := ioutil.TempDir(userTmpDir, prefix)
	if err != nil {
		panic(err)
	}

	f, err := tempfile.TempFileWithSuffix(tmpDir, "", ".txt")
	if err != nil {
		panic(err)
	}

	return f.Name()
}

func Exists(name string) (bool, error) {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false, err
	}
	return true, nil
}
