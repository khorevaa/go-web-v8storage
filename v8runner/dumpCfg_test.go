package v8runner

import (
	"testing"
	"./dumpMode"
	//. "gopkg.in/check.v1"
)

func estКонфигуратор_ВыгрузитьКонфигурациюСРежимомВыгрузки(t *testing.T) {

	type fields struct {
		Контекст              *Контекст
		ФайлИнформации        string
		ОчищатьФайлИнформации bool
		ЭтоWindows            bool
		ВерсияПлатформы       *ВерсияПлатформы
		выводКоманды          string
	}
	type args struct {
		dir  string
		mode string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Тестовая проверка",
			fields{
				НовыйКонтекст(),
				НовыйФайлИнформации(),
				true,
				false,
				ПолучитьВерсиюПоУмолчанию(),
				"",
			},
			args{ВременныйКаталог(),
				РежимВыгрузкиКонфигурации.Иерархический,
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Конфигуратор{
				Контекст:              tt.fields.Контекст,
				ФайлИнформации:        tt.fields.ФайлИнформации,
				ОчищатьФайлИнформации: tt.fields.ОчищатьФайлИнформации,
				ЭтоWindows:            tt.fields.ЭтоWindows,
				ВерсияПлатформы:       tt.fields.ВерсияПлатформы,
				выводКоманды:          tt.fields.выводКоманды,
			}
			if err := conf.ВыгрузитьКонфигурациюСРежимомВыгрузки(tt.args.dir, tt.args.mode); (err != nil) != tt.wantErr {
				t.Errorf("Конфигуратор.ВыгрузитьКонфигурациюСРежимомВыгрузки() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestКонфигуратор_ВыгрузитьКонфигурациюПоУмолчанию(t *testing.T) {
	type fields struct {
		Контекст              *Контекст
		ФайлИнформации        string
		ОчищатьФайлИнформации bool
		ЭтоWindows            bool
		ВерсияПлатформы       *ВерсияПлатформы
		выводКоманды          string
	}
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Конфигуратор{
				Контекст:              tt.fields.Контекст,
				ФайлИнформации:        tt.fields.ФайлИнформации,
				ОчищатьФайлИнформации: tt.fields.ОчищатьФайлИнформации,
				ЭтоWindows:            tt.fields.ЭтоWindows,
				ВерсияПлатформы:       tt.fields.ВерсияПлатформы,
				выводКоманды:          tt.fields.выводКоманды,
			}
			if err := conf.ВыгрузитьКонфигурациюПоУмолчанию(tt.args.dir); (err != nil) != tt.wantErr {
				t.Errorf("Конфигуратор.ВыгрузитьКонфигурациюПоУмолчанию() error = %v, wantErr %v", err, tt.wantErr)
			}
			ОчиститьВременныйКаталог()
		})
	}
}

func TestКонфигуратор_ВыгрузитьКонфигурацию(t *testing.T) {
	type fields struct {
		Контекст              *Контекст
		ФайлИнформации        string
		ОчищатьФайлИнформации bool
		ЭтоWindows            bool
		ВерсияПлатформы       *ВерсияПлатформы
		выводКоманды          string
	}
	type args struct {
		КаталогВыгрузки              string
		ФорматВыгрузки               string
		ТолькоИзмененные             bool
		ПутьКФайлуИзменений          string
		ПутьКФайлуВерсийДляСравнения string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Конфигуратор{
				Контекст:              tt.fields.Контекст,
				ФайлИнформации:        tt.fields.ФайлИнформации,
				ОчищатьФайлИнформации: tt.fields.ОчищатьФайлИнформации,
				ЭтоWindows:            tt.fields.ЭтоWindows,
				ВерсияПлатформы:       tt.fields.ВерсияПлатформы,
				выводКоманды:          tt.fields.выводКоманды,
			}
			if err := conf.ВыгрузитьКонфигурацию(tt.args.КаталогВыгрузки, tt.args.ФорматВыгрузки, tt.args.ТолькоИзмененные, tt.args.ПутьКФайлуИзменений, tt.args.ПутьКФайлуВерсийДляСравнения); (err != nil) != tt.wantErr {
				t.Errorf("Конфигуратор.ВыгрузитьКонфигурацию() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestКонфигуратор_ВыгрузитьИзмененияКонфигурацииВФайл(t *testing.T) {
	type fields struct {
		Контекст              *Контекст
		ФайлИнформации        string
		ОчищатьФайлИнформации bool
		ЭтоWindows            bool
		ВерсияПлатформы       *ВерсияПлатформы
		выводКоманды          string
	}
	type args struct {
		КаталогВыгрузки              string
		ФорматВыгрузки               string
		ПутьКФайлуИзменений          string
		ПутьКФайлуВерсийДляСравнения string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Конфигуратор{
				Контекст:              tt.fields.Контекст,
				ФайлИнформации:        tt.fields.ФайлИнформации,
				ОчищатьФайлИнформации: tt.fields.ОчищатьФайлИнформации,
				ЭтоWindows:            tt.fields.ЭтоWindows,
				ВерсияПлатформы:       tt.fields.ВерсияПлатформы,
				выводКоманды:          tt.fields.выводКоманды,
			}
			if err := conf.ВыгрузитьИзмененияКонфигурацииВФайл(tt.args.КаталогВыгрузки, tt.args.ФорматВыгрузки, tt.args.ПутьКФайлуИзменений, tt.args.ПутьКФайлуВерсийДляСравнения); (err != nil) != tt.wantErr {
				t.Errorf("Конфигуратор.ВыгрузитьИзмененияКонфигурацииВФайл() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestКонфигуратор_dumpConfigToFiles(t *testing.T) {
	type fields struct {
		Контекст              *Контекст
		ФайлИнформации        string
		ОчищатьФайлИнформации bool
		ЭтоWindows            bool
		ВерсияПлатформы       *ВерсияПлатформы
		выводКоманды          string
	}
	type args struct {
		dir          string
		mode         string
		ch           bool
		pChFile      string
		pVersionFile string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conf := &Конфигуратор{
				Контекст:              tt.fields.Контекст,
				ФайлИнформации:        tt.fields.ФайлИнформации,
				ОчищатьФайлИнформации: tt.fields.ОчищатьФайлИнформации,
				ЭтоWindows:            tt.fields.ЭтоWindows,
				ВерсияПлатформы:       tt.fields.ВерсияПлатформы,
				выводКоманды:          tt.fields.выводКоманды,
			}
			if err := conf.dumpConfigToFiles(tt.args.dir, tt.args.mode, tt.args.ch, tt.args.pChFile, tt.args.pVersionFile); (err != nil) != tt.wantErr {
				t.Errorf("Конфигуратор.dumpConfigToFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
