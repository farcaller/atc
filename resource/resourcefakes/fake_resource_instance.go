// This file was generated by counterfeiter
package resourcefakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc/resource"
	"github.com/concourse/atc/worker"
)

type FakeResourceInstance struct {
	FindOnStub        func(lager.Logger, worker.Client) (worker.Volume, bool, error)
	findOnMutex       sync.RWMutex
	findOnArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.Client
	}
	findOnReturns struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}
	FindOrCreateOnStub        func(lager.Logger, worker.Client) (worker.Volume, error)
	findOrCreateOnMutex       sync.RWMutex
	findOrCreateOnArgsForCall []struct {
		arg1 lager.Logger
		arg2 worker.Client
	}
	findOrCreateOnReturns struct {
		result1 worker.Volume
		result2 error
	}
	ResourceCacheIdentifierStub        func() worker.ResourceCacheIdentifier
	resourceCacheIdentifierMutex       sync.RWMutex
	resourceCacheIdentifierArgsForCall []struct{}
	resourceCacheIdentifierReturns     struct {
		result1 worker.ResourceCacheIdentifier
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceInstance) FindOn(arg1 lager.Logger, arg2 worker.Client) (worker.Volume, bool, error) {
	fake.findOnMutex.Lock()
	fake.findOnArgsForCall = append(fake.findOnArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.Client
	}{arg1, arg2})
	fake.recordInvocation("FindOn", []interface{}{arg1, arg2})
	fake.findOnMutex.Unlock()
	if fake.FindOnStub != nil {
		return fake.FindOnStub(arg1, arg2)
	}
	return fake.findOnReturns.result1, fake.findOnReturns.result2, fake.findOnReturns.result3
}

func (fake *FakeResourceInstance) FindOnCallCount() int {
	fake.findOnMutex.RLock()
	defer fake.findOnMutex.RUnlock()
	return len(fake.findOnArgsForCall)
}

func (fake *FakeResourceInstance) FindOnArgsForCall(i int) (lager.Logger, worker.Client) {
	fake.findOnMutex.RLock()
	defer fake.findOnMutex.RUnlock()
	return fake.findOnArgsForCall[i].arg1, fake.findOnArgsForCall[i].arg2
}

func (fake *FakeResourceInstance) FindOnReturns(result1 worker.Volume, result2 bool, result3 error) {
	fake.FindOnStub = nil
	fake.findOnReturns = struct {
		result1 worker.Volume
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResourceInstance) FindOrCreateOn(arg1 lager.Logger, arg2 worker.Client) (worker.Volume, error) {
	fake.findOrCreateOnMutex.Lock()
	fake.findOrCreateOnArgsForCall = append(fake.findOrCreateOnArgsForCall, struct {
		arg1 lager.Logger
		arg2 worker.Client
	}{arg1, arg2})
	fake.recordInvocation("FindOrCreateOn", []interface{}{arg1, arg2})
	fake.findOrCreateOnMutex.Unlock()
	if fake.FindOrCreateOnStub != nil {
		return fake.FindOrCreateOnStub(arg1, arg2)
	}
	return fake.findOrCreateOnReturns.result1, fake.findOrCreateOnReturns.result2
}

func (fake *FakeResourceInstance) FindOrCreateOnCallCount() int {
	fake.findOrCreateOnMutex.RLock()
	defer fake.findOrCreateOnMutex.RUnlock()
	return len(fake.findOrCreateOnArgsForCall)
}

func (fake *FakeResourceInstance) FindOrCreateOnArgsForCall(i int) (lager.Logger, worker.Client) {
	fake.findOrCreateOnMutex.RLock()
	defer fake.findOrCreateOnMutex.RUnlock()
	return fake.findOrCreateOnArgsForCall[i].arg1, fake.findOrCreateOnArgsForCall[i].arg2
}

func (fake *FakeResourceInstance) FindOrCreateOnReturns(result1 worker.Volume, result2 error) {
	fake.FindOrCreateOnStub = nil
	fake.findOrCreateOnReturns = struct {
		result1 worker.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceInstance) ResourceCacheIdentifier() worker.ResourceCacheIdentifier {
	fake.resourceCacheIdentifierMutex.Lock()
	fake.resourceCacheIdentifierArgsForCall = append(fake.resourceCacheIdentifierArgsForCall, struct{}{})
	fake.recordInvocation("ResourceCacheIdentifier", []interface{}{})
	fake.resourceCacheIdentifierMutex.Unlock()
	if fake.ResourceCacheIdentifierStub != nil {
		return fake.ResourceCacheIdentifierStub()
	}
	return fake.resourceCacheIdentifierReturns.result1
}

func (fake *FakeResourceInstance) ResourceCacheIdentifierCallCount() int {
	fake.resourceCacheIdentifierMutex.RLock()
	defer fake.resourceCacheIdentifierMutex.RUnlock()
	return len(fake.resourceCacheIdentifierArgsForCall)
}

func (fake *FakeResourceInstance) ResourceCacheIdentifierReturns(result1 worker.ResourceCacheIdentifier) {
	fake.ResourceCacheIdentifierStub = nil
	fake.resourceCacheIdentifierReturns = struct {
		result1 worker.ResourceCacheIdentifier
	}{result1}
}

func (fake *FakeResourceInstance) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.findOnMutex.RLock()
	defer fake.findOnMutex.RUnlock()
	fake.findOrCreateOnMutex.RLock()
	defer fake.findOrCreateOnMutex.RUnlock()
	fake.resourceCacheIdentifierMutex.RLock()
	defer fake.resourceCacheIdentifierMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeResourceInstance) recordInvocation(key string, args []interface{}) {
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

var _ resource.ResourceInstance = new(FakeResourceInstance)