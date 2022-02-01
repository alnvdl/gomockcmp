package external

// ServiceClientObject represents a data type used in an external service.
type ServiceClientObject struct {
	ID string
	A  int
	B  string
}

// ServiceClient represents the interface for a client of an external service.
type ServiceClient interface {
	DoSomething(id string) (ServiceClientObject, error)
	DoSomethingSlightlyDifferent(id string) (ServiceClientObject, error)
}
