// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotal-cf/om/api"
)

type SetupService struct {
	SetupStub        func(api.SetupInput) (api.SetupOutput, error)
	setupMutex       sync.RWMutex
	setupArgsForCall []struct {
		arg1 api.SetupInput
	}
	setupReturns struct {
		result1 api.SetupOutput
		result2 error
	}
	setupReturnsOnCall map[int]struct {
		result1 api.SetupOutput
		result2 error
	}
	EnsureAvailabilityStub        func(api.EnsureAvailabilityInput) (api.EnsureAvailabilityOutput, error)
	ensureAvailabilityMutex       sync.RWMutex
	ensureAvailabilityArgsForCall []struct {
		arg1 api.EnsureAvailabilityInput
	}
	ensureAvailabilityReturns struct {
		result1 api.EnsureAvailabilityOutput
		result2 error
	}
	ensureAvailabilityReturnsOnCall map[int]struct {
		result1 api.EnsureAvailabilityOutput
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *SetupService) Setup(arg1 api.SetupInput) (api.SetupOutput, error) {
	fake.setupMutex.Lock()
	ret, specificReturn := fake.setupReturnsOnCall[len(fake.setupArgsForCall)]
	fake.setupArgsForCall = append(fake.setupArgsForCall, struct {
		arg1 api.SetupInput
	}{arg1})
	fake.recordInvocation("Setup", []interface{}{arg1})
	fake.setupMutex.Unlock()
	if fake.SetupStub != nil {
		return fake.SetupStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.setupReturns.result1, fake.setupReturns.result2
}

func (fake *SetupService) SetupCallCount() int {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	return len(fake.setupArgsForCall)
}

func (fake *SetupService) SetupArgsForCall(i int) api.SetupInput {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	return fake.setupArgsForCall[i].arg1
}

func (fake *SetupService) SetupReturns(result1 api.SetupOutput, result2 error) {
	fake.SetupStub = nil
	fake.setupReturns = struct {
		result1 api.SetupOutput
		result2 error
	}{result1, result2}
}

func (fake *SetupService) SetupReturnsOnCall(i int, result1 api.SetupOutput, result2 error) {
	fake.SetupStub = nil
	if fake.setupReturnsOnCall == nil {
		fake.setupReturnsOnCall = make(map[int]struct {
			result1 api.SetupOutput
			result2 error
		})
	}
	fake.setupReturnsOnCall[i] = struct {
		result1 api.SetupOutput
		result2 error
	}{result1, result2}
}

func (fake *SetupService) EnsureAvailability(arg1 api.EnsureAvailabilityInput) (api.EnsureAvailabilityOutput, error) {
	fake.ensureAvailabilityMutex.Lock()
	ret, specificReturn := fake.ensureAvailabilityReturnsOnCall[len(fake.ensureAvailabilityArgsForCall)]
	fake.ensureAvailabilityArgsForCall = append(fake.ensureAvailabilityArgsForCall, struct {
		arg1 api.EnsureAvailabilityInput
	}{arg1})
	fake.recordInvocation("EnsureAvailability", []interface{}{arg1})
	fake.ensureAvailabilityMutex.Unlock()
	if fake.EnsureAvailabilityStub != nil {
		return fake.EnsureAvailabilityStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.ensureAvailabilityReturns.result1, fake.ensureAvailabilityReturns.result2
}

func (fake *SetupService) EnsureAvailabilityCallCount() int {
	fake.ensureAvailabilityMutex.RLock()
	defer fake.ensureAvailabilityMutex.RUnlock()
	return len(fake.ensureAvailabilityArgsForCall)
}

func (fake *SetupService) EnsureAvailabilityArgsForCall(i int) api.EnsureAvailabilityInput {
	fake.ensureAvailabilityMutex.RLock()
	defer fake.ensureAvailabilityMutex.RUnlock()
	return fake.ensureAvailabilityArgsForCall[i].arg1
}

func (fake *SetupService) EnsureAvailabilityReturns(result1 api.EnsureAvailabilityOutput, result2 error) {
	fake.EnsureAvailabilityStub = nil
	fake.ensureAvailabilityReturns = struct {
		result1 api.EnsureAvailabilityOutput
		result2 error
	}{result1, result2}
}

func (fake *SetupService) EnsureAvailabilityReturnsOnCall(i int, result1 api.EnsureAvailabilityOutput, result2 error) {
	fake.EnsureAvailabilityStub = nil
	if fake.ensureAvailabilityReturnsOnCall == nil {
		fake.ensureAvailabilityReturnsOnCall = make(map[int]struct {
			result1 api.EnsureAvailabilityOutput
			result2 error
		})
	}
	fake.ensureAvailabilityReturnsOnCall[i] = struct {
		result1 api.EnsureAvailabilityOutput
		result2 error
	}{result1, result2}
}

func (fake *SetupService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	fake.ensureAvailabilityMutex.RLock()
	defer fake.ensureAvailabilityMutex.RUnlock()
	return fake.invocations
}

func (fake *SetupService) recordInvocation(key string, args []interface{}) {
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
