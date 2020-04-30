package gluon

// A Stereotype specifies the purpose of a type definition
type Stereotype string

const (
	// Repository defines a service which is used for data storage
	Repository Stereotype = "repository"
	// RestController defines a service to handle rest requests
	RestController Stereotype = "restcontroller"
	// Default is service (component)
	Default Stereotype = "service"
)
