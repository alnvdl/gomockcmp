package mock

import "github.com/alnvdl/gomockcmp/external"

// This example handwritten mock is not a suggestion or anything. This is just
// a mish-mash of practices typically followed when writing "simple" mocks
// manually. In mockland, all mocks start simple. Until they aren't anymore.

type DoSomethingCall struct {
	ID string
}

type DoSomethingReturn struct {
	external.ServiceClientObject
	Error error
}

type MockServiceClientHandwritten struct {
	// PROBLEM: in many handwritten mocks, return values are not actually
	// stored per-call like in here. In many cases, we end up having a generic
	// "external.ServiceClientObject" value that gets returned by all related
	// calls (e.g., DoSomething and DoSomethingSlightlyDifferent). In that
	// case, we can't control behavior specifically per call when more than
	// one call is involved.
	DoSomethingReturns                  map[string][]DoSomethingReturn
	DoSomethingSlightlyDifferentReturns map[string][]DoSomethingReturn
	DoSomethingNCalls                   int
	DoSomethingSlightlyDifferentNCalls  int

	DoSomethingCalls                  []DoSomethingCall
	DoSomethingSlightlyDifferentCalls []DoSomethingCall
}

func NewMockServiceClientHandwritten() *MockServiceClientHandwritten {
	// PROBLEM: we need to manually initialize mocks.
	return &MockServiceClientHandwritten{
		DoSomethingReturns:                  make(map[string][]DoSomethingReturn),
		DoSomethingSlightlyDifferentReturns: make(map[string][]DoSomethingReturn),
	}
}

func (m *MockServiceClientHandwritten) DoSomething(id string) (external.ServiceClientObject, error) {
	// PROBLEM: We need to keep track of calls; we need special objects to track it.
	m.DoSomethingCalls = append(m.DoSomethingCalls, DoSomethingCall{id})
	defer func() { m.DoSomethingNCalls++ }()
	// PROBLEM: We need to store and parse information to be returned.

	// PROBLEM: needing to perform any extra logic here requires changing the
	// mock and adding more configuration parameters (not all tests may need
	// the same logic).
	ret := m.DoSomethingReturns[id][m.DoSomethingNCalls]
	return ret.ServiceClientObject, ret.Error
}

func (m *MockServiceClientHandwritten) DoSomethingSlightlyDifferent(id string) (external.ServiceClientObject, error) {
	m.DoSomethingSlightlyDifferentCalls = append(m.DoSomethingSlightlyDifferentCalls, DoSomethingCall{id})
	defer func() { m.DoSomethingSlightlyDifferentNCalls++ }()
	// PROBLEM: if we forget to setup a call before the mock is used, we get a
	// panic (or we have to treat this condition in all places).
	ret := m.DoSomethingSlightlyDifferentReturns[id][m.DoSomethingSlightlyDifferentNCalls]
	return ret.ServiceClientObject, ret.Error
}

// BIG PROBLEM: we are reimplementing a mocking framework, and it looks
// repetitive.
//
// BIG PROBLEM: this mock doesn't even provide us automatic verification, we
// need to perform it in each test. (well, we could have a function here, but
// that's more code).
