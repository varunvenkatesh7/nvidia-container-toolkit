// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package ldcache

import (
	"sync"
)

// Ensure, that LDCacheMock does implement LDCache.
// If this is not the case, regenerate this file with moq.
var _ LDCache = &LDCacheMock{}

// LDCacheMock is a mock implementation of LDCache.
//
//	func TestSomethingThatUsesLDCache(t *testing.T) {
//
//		// make and configure a mocked LDCache
//		mockedLDCache := &LDCacheMock{
//			ListFunc: func() ([]string, []string) {
//				panic("mock out the List method")
//			},
//			LookupFunc: func(strings ...string) ([]string, []string) {
//				panic("mock out the Lookup method")
//			},
//		}
//
//		// use mockedLDCache in code that requires LDCache
//		// and then make assertions.
//
//	}
type LDCacheMock struct {
	// ListFunc mocks the List method.
	ListFunc func() ([]string, []string)

	// LookupFunc mocks the Lookup method.
	LookupFunc func(strings ...string) ([]string, []string)

	// calls tracks calls to the methods.
	calls struct {
		// List holds details about calls to the List method.
		List []struct {
		}
		// Lookup holds details about calls to the Lookup method.
		Lookup []struct {
			// Strings is the strings argument value.
			Strings []string
		}
	}
	lockList   sync.RWMutex
	lockLookup sync.RWMutex
}

// List calls ListFunc.
func (mock *LDCacheMock) List() ([]string, []string) {
	if mock.ListFunc == nil {
		panic("LDCacheMock.ListFunc: method is nil but LDCache.List was just called")
	}
	callInfo := struct {
	}{}
	mock.lockList.Lock()
	mock.calls.List = append(mock.calls.List, callInfo)
	mock.lockList.Unlock()
	return mock.ListFunc()
}

// ListCalls gets all the calls that were made to List.
// Check the length with:
//
//	len(mockedLDCache.ListCalls())
func (mock *LDCacheMock) ListCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockList.RLock()
	calls = mock.calls.List
	mock.lockList.RUnlock()
	return calls
}

// Lookup calls LookupFunc.
func (mock *LDCacheMock) Lookup(strings ...string) ([]string, []string) {
	if mock.LookupFunc == nil {
		panic("LDCacheMock.LookupFunc: method is nil but LDCache.Lookup was just called")
	}
	callInfo := struct {
		Strings []string
	}{
		Strings: strings,
	}
	mock.lockLookup.Lock()
	mock.calls.Lookup = append(mock.calls.Lookup, callInfo)
	mock.lockLookup.Unlock()
	return mock.LookupFunc(strings...)
}

// LookupCalls gets all the calls that were made to Lookup.
// Check the length with:
//
//	len(mockedLDCache.LookupCalls())
func (mock *LDCacheMock) LookupCalls() []struct {
	Strings []string
} {
	var calls []struct {
		Strings []string
	}
	mock.lockLookup.RLock()
	calls = mock.calls.Lookup
	mock.lockLookup.RUnlock()
	return calls
}