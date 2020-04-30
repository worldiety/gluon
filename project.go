package gluon

import "fmt"

type Project struct {
	name         string
	applications []*Application
	modulePrefix string
}

func NewProject(name string) *Project {
	return &Project{name: name}
}

func (p *Project) SetModulePrefix(s string) *Project {
	p.modulePrefix = lcase(s)
	return p
}

func (p *Project) AddApplications(apps ...*Application) *Project {
	for _, a := range apps {
		p.applications = append(p.applications, a)
		a.parent = p
	}
	return p
}

func (p *Project) Generate(ws *Workspace) error {
	ws.Debug("generate project %s", p.name)
	err := ws.Clear()
	if err != nil {
		return fmt.Errorf("failed to clear: %w", err)
	}

	for _, app := range p.applications {
		err := app.Generate(ws)
		if err != nil {
			return fmt.Errorf("failed to generate app '%s':%w", app.name, err)
		}
	}
	return nil
}
