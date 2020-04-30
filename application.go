package gluon

import "fmt"

type Application struct {
	name     string
	parent   *Project
	services []*Service
}

func NewApplication(name string) *Application {
	return &Application{name: name}
}

func (a *Application) AddServices(services ...*Service) *Application {
	for _, service := range services {
		a.services = append(a.services, service)
		service.parent = a
	}
	return a
}

func (a *Application) Generate(ws *Workspace) error {
	ws.Debug("generate application %s", a.name)
	appName := lcase(a.name)
	if !ws.IsGoModuleRoot(appName) {
		err := ws.CreateGoModule(appName, join("/", a.parent.modulePrefix, appName))
		if err != nil {
			return fmt.Errorf("unable to create go module: %w", err)
		}
	}

	for _, service := range a.services {
		err := service.Generate(ws)
		if err != nil {
			return fmt.Errorf("failed to generate service '%s':%w", service.name, err)
		}
	}

	return nil
}
