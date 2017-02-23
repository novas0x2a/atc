// This file was generated by counterfeiter
package lockfakes

import (
	"sync"

	"github.com/concourse/atc/db/lock"
)

type FakeLock struct {
	AcquireStub        func() (bool, error)
	acquireMutex       sync.RWMutex
	acquireArgsForCall []struct{}
	acquireReturns     struct {
		result1 bool
		result2 error
	}
	ReleaseStub        func() error
	releaseMutex       sync.RWMutex
	releaseArgsForCall []struct{}
	releaseReturns     struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLock) Acquire() (bool, error) {
	fake.acquireMutex.Lock()
	fake.acquireArgsForCall = append(fake.acquireArgsForCall, struct{}{})
	fake.recordInvocation("Acquire", []interface{}{})
	fake.acquireMutex.Unlock()
	if fake.AcquireStub != nil {
		return fake.AcquireStub()
	}
	return fake.acquireReturns.result1, fake.acquireReturns.result2
}

func (fake *FakeLock) AcquireCallCount() int {
	fake.acquireMutex.RLock()
	defer fake.acquireMutex.RUnlock()
	return len(fake.acquireArgsForCall)
}

func (fake *FakeLock) AcquireReturns(result1 bool, result2 error) {
	fake.AcquireStub = nil
	fake.acquireReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeLock) Release() error {
	fake.releaseMutex.Lock()
	fake.releaseArgsForCall = append(fake.releaseArgsForCall, struct{}{})
	fake.recordInvocation("Release", []interface{}{})
	fake.releaseMutex.Unlock()
	if fake.ReleaseStub != nil {
		return fake.ReleaseStub()
	}
	return fake.releaseReturns.result1
}

func (fake *FakeLock) ReleaseCallCount() int {
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	return len(fake.releaseArgsForCall)
}

func (fake *FakeLock) ReleaseReturns(result1 error) {
	fake.ReleaseStub = nil
	fake.releaseReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeLock) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.acquireMutex.RLock()
	defer fake.acquireMutex.RUnlock()
	fake.releaseMutex.RLock()
	defer fake.releaseMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeLock) recordInvocation(key string, args []interface{}) {
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

var _ lock.Lock = new(FakeLock)
