package service

import service "github.com/kardianos/service"

var _ service.Interface = (*program)(nil)

type program struct {
	version string
	commit  string
}

func NewProgam(version, commit string) service.Interface {
	return &program{
		version: version,
		commit:  commit,
	}
}

func (p *program) Start(s service.Service) (err error) {
	return
}

func (p *program) Stop(s service.Service) (err error) {
	return
}
