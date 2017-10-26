package bootstrap

import (
	"fmt"
	"html/template"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	"github.com/yosssi/ace"
)

type Configurator func(*Bootstrapper)

var GlobalSessions *session.Manager

type Bootstrapper struct {
	*beego.App
	sessionConfig *session.ManagerConfig
}

// New returns a new Bootstrapper.
func New(cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		App: beego.BeeApp,
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}

// Bootstrap prepares our application.
//
// Returns itself.
func (b *Bootstrapper) Bootstrap() *Bootstrapper {

	beego.AddViewPath("views")

	b.sessionConfig = &session.ManagerConfig{
		CookieName: "jigsessionid",
		Gclifetime: 3600,
	}
	beego.BConfig.WebConfig.Session.SessionOn = true

	beego.AddTemplateEngine("ace", func(root, path string, funcs template.FuncMap) (*template.Template, error) {
		aceOptions := &ace.Options{DynamicReload: true, FuncMap: funcs}
		aceBasePath := filepath.Join(root, "base/base")
		aceInnerPath := filepath.Join(root, strings.TrimSuffix(path, ".ace"))

		tpl, err := ace.Load(aceBasePath, aceInnerPath, aceOptions)
		fmt.Errorf("error loading ace template: %v", tpl)
		if err != nil {
			return nil, fmt.Errorf("error loading ace template: %v", err)
		}

		return tpl, nil
	})

	// Set up global session management
	//beego.LoadAppConfig("yaml", "conf/app.yaml")
	//beego.BConfig.WebConfig.Session.SessionOn = true
	//beego.BConfig.WebConfig.Session.SessionProvider = "file"
	//beego.BConfig.WebConfig.Session.SessionProviderConfig = "./tmp"

	return b
}

// Listen starts the http server with the specified "addr".
func (b *Bootstrapper) Listen() {

	beego.GlobalSessions, _ = session.NewManager("memory", b.sessionConfig)
	//go beego.GlobalSessions.GC()

	b.Run()
}
