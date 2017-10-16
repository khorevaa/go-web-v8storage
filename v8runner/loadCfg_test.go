package v8runner

import (
	"testing"
)

func TestКонфигуратор_loadConfigFromFiles(t *testing.T) {
	type fields struct {
		Контекст              *Контекст
		ФайлИнформации        string
		ОчищатьФайлИнформации bool
		ЭтоWindows            bool
		ВерсияПлатформы       *ВерсияПлатформы
		выводКоманды          string
	}
	type args struct {
		dir         string
		pListFile   string
		format      string
		updDumpInfo bool
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
			if err := conf.loadConfigFromFiles(tt.args.dir, tt.args.pListFile, tt.args.format, tt.args.updDumpInfo); (err != nil) != tt.wantErr {
				t.Errorf("Конфигуратор.loadConfigFromFiles() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestКонфигуратор_loadCfg(t *testing.T) {
	type fields struct {
		Контекст              *Контекст
		ФайлИнформации        string
		ОчищатьФайлИнформации bool
		ЭтоWindows            bool
		ВерсияПлатформы       *ВерсияПлатформы
		выводКоманды          string
	}
	type args struct {
		cfg string
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
			if err := conf.loadCfg(tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("Конфигуратор.loadCfg() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestКонфигуратор_ЗагрузитьКонфигурациюИзФайла(t *testing.T) {
	type fields struct {
		Контекст              *Контекст
		ФайлИнформации        string
		ОчищатьФайлИнформации bool
		ЭтоWindows            bool
		ВерсияПлатформы       *ВерсияПлатформы
		выводКоманды          string
	}
	type args struct {
		ПутьКФайлуКонфигуации string
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
			if err := conf.ЗагрузитьКонфигурациюИзФайла(tt.args.ПутьКФайлуКонфигуации); (err != nil) != tt.wantErr {
				t.Errorf("Конфигуратор.ЗагрузитьКонфигурациюИзФайла() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestКонфигуратор_ЗагрузитьКонфигурациюИзФайлов(t *testing.T) {
	type fields struct {
		Контекст              *Контекст
		ФайлИнформации        string
		ОчищатьФайлИнформации bool
		ЭтоWindows            bool
		ВерсияПлатформы       *ВерсияПлатформы
		выводКоманды          string
	}
	type args struct {
		КаталогЗагрузки string
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
			if err := conf.ЗагрузитьКонфигурациюИзФайлов(tt.args.КаталогЗагрузки); (err != nil) != tt.wantErr {
				t.Errorf("Конфигуратор.ЗагрузитьКонфигурациюИзФайлов() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
