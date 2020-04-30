package gluon

// A Service has a different definition than usually given in literature. For us, it can be a mixture of a component,
// module or (remote) service. The purpose is to create a contract for communication. If this service is run
// remotely or compiled statically and how it will be instantiated is not defined here. So the definitions from
// literature do not apply properly. A component needs a component model (we may not have that at runtime), a service
// must be stateless (usually not realistically) and a module suggests replaceability at compile- or runtime.
//
// The expected properties of a service are as follows:
//  - so small, that a single developer can handle it within a few days or weeks
//  - may require other services to be functional
//  - understanding in a singleton way, however does not keep an internal static state and may be spawned in a network
//    to scale
//  - contained within a single package, but may have other internal packages
//  - only provides public interfaces and data classes and a factory method, e.g. myservicepackage.New(...) to
//    inject other services as dependencies
//  - it is not SOA in the sense that it must be on a network (which is usually not the case for monoliths)
//  - everything is designed by contract
//  - important: there is never(!) a dependency between services other than a contract specification.
//  - probably a single domain or a single use case with a single responsibility
//
// Go - implementation note
//
// Go has structural (or kind of duck) typing to allow dynamic dispatch and create abstractions. Instead of
// accepting interfaces and returning concrete types an "all-interface" concept is favored to allow dependency
// free service contracts using anonymous interface
// aliases. However this won't work for recursion and is still more a workaround, see discussion at
// https://github.com/golang/go/issues/8082
//
// Java - implementation thoughts
//
// Java has no structural or duck typing so one must
//  - either introduce a dependency
//  - write an adapter or decorator at a top level, which has both dependencies
//  - repeat yourself and abstract over a type-unsafe message serialization step
type Service struct {
	stereotype Stereotype
	name       string
	contract   *Interface
	parent     *Application
}

func NewService(name string) *Service {
	return &Service{
		stereotype: Default,
		name:       name,
		contract:   NewInterface(name),
	}
}

func (s *Service) Generate(ws *Workspace) error {
	return ws.NewGoFile(s.name, func(w *Writer) error {
		return nil
	})

}
