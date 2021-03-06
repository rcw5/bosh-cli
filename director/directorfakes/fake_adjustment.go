// Code generated by counterfeiter. DO NOT EDIT.
package directorfakes

import (
	"net/http"
	"sync"

	"github.com/cloudfoundry/bosh-cli/director"
)

type FakeAdjustment struct {
	AdjustStub        func(req *http.Request, retried bool) error
	adjustMutex       sync.RWMutex
	adjustArgsForCall []struct {
		req     *http.Request
		retried bool
	}
	adjustReturns struct {
		result1 error
	}
	adjustReturnsOnCall map[int]struct {
		result1 error
	}
	NeedsReadjustmentStub        func(*http.Response) bool
	needsReadjustmentMutex       sync.RWMutex
	needsReadjustmentArgsForCall []struct {
		arg1 *http.Response
	}
	needsReadjustmentReturns struct {
		result1 bool
	}
	needsReadjustmentReturnsOnCall map[int]struct {
		result1 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAdjustment) Adjust(req *http.Request, retried bool) error {
	fake.adjustMutex.Lock()
	ret, specificReturn := fake.adjustReturnsOnCall[len(fake.adjustArgsForCall)]
	fake.adjustArgsForCall = append(fake.adjustArgsForCall, struct {
		req     *http.Request
		retried bool
	}{req, retried})
	fake.recordInvocation("Adjust", []interface{}{req, retried})
	fake.adjustMutex.Unlock()
	if fake.AdjustStub != nil {
		return fake.AdjustStub(req, retried)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.adjustReturns.result1
}

func (fake *FakeAdjustment) AdjustCallCount() int {
	fake.adjustMutex.RLock()
	defer fake.adjustMutex.RUnlock()
	return len(fake.adjustArgsForCall)
}

func (fake *FakeAdjustment) AdjustArgsForCall(i int) (*http.Request, bool) {
	fake.adjustMutex.RLock()
	defer fake.adjustMutex.RUnlock()
	return fake.adjustArgsForCall[i].req, fake.adjustArgsForCall[i].retried
}

func (fake *FakeAdjustment) AdjustReturns(result1 error) {
	fake.AdjustStub = nil
	fake.adjustReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeAdjustment) AdjustReturnsOnCall(i int, result1 error) {
	fake.AdjustStub = nil
	if fake.adjustReturnsOnCall == nil {
		fake.adjustReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.adjustReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeAdjustment) NeedsReadjustment(arg1 *http.Response) bool {
	fake.needsReadjustmentMutex.Lock()
	ret, specificReturn := fake.needsReadjustmentReturnsOnCall[len(fake.needsReadjustmentArgsForCall)]
	fake.needsReadjustmentArgsForCall = append(fake.needsReadjustmentArgsForCall, struct {
		arg1 *http.Response
	}{arg1})
	fake.recordInvocation("NeedsReadjustment", []interface{}{arg1})
	fake.needsReadjustmentMutex.Unlock()
	if fake.NeedsReadjustmentStub != nil {
		return fake.NeedsReadjustmentStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.needsReadjustmentReturns.result1
}

func (fake *FakeAdjustment) NeedsReadjustmentCallCount() int {
	fake.needsReadjustmentMutex.RLock()
	defer fake.needsReadjustmentMutex.RUnlock()
	return len(fake.needsReadjustmentArgsForCall)
}

func (fake *FakeAdjustment) NeedsReadjustmentArgsForCall(i int) *http.Response {
	fake.needsReadjustmentMutex.RLock()
	defer fake.needsReadjustmentMutex.RUnlock()
	return fake.needsReadjustmentArgsForCall[i].arg1
}

func (fake *FakeAdjustment) NeedsReadjustmentReturns(result1 bool) {
	fake.NeedsReadjustmentStub = nil
	fake.needsReadjustmentReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeAdjustment) NeedsReadjustmentReturnsOnCall(i int, result1 bool) {
	fake.NeedsReadjustmentStub = nil
	if fake.needsReadjustmentReturnsOnCall == nil {
		fake.needsReadjustmentReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.needsReadjustmentReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeAdjustment) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.adjustMutex.RLock()
	defer fake.adjustMutex.RUnlock()
	fake.needsReadjustmentMutex.RLock()
	defer fake.needsReadjustmentMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAdjustment) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ director.Adjustment = new(FakeAdjustment)
