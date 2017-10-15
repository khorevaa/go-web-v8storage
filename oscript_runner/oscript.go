package oscript

import (
//"fmt"
)

// Запускатель
type Runner struct {
	Path string
	Arg  string
}

func (r *Runner) Run() *Runner {

	return r

}

func Run() {
}
