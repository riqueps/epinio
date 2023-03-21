// Code generated by counterfeiter. DO NOT EDIT.
package applicationfakes

import (
	"context"
	"sync"

	"github.com/epinio/epinio/internal/application"
	v1 "k8s.io/api/batch/v1"
)

type FakeJobLister struct {
	ListJobsStub        func(context.Context, string, string) (*v1.JobList, error)
	listJobsMutex       sync.RWMutex
	listJobsArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}
	listJobsReturns struct {
		result1 *v1.JobList
		result2 error
	}
	listJobsReturnsOnCall map[int]struct {
		result1 *v1.JobList
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeJobLister) ListJobs(arg1 context.Context, arg2 string, arg3 string) (*v1.JobList, error) {
	fake.listJobsMutex.Lock()
	ret, specificReturn := fake.listJobsReturnsOnCall[len(fake.listJobsArgsForCall)]
	fake.listJobsArgsForCall = append(fake.listJobsArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.ListJobsStub
	fakeReturns := fake.listJobsReturns
	fake.recordInvocation("ListJobs", []interface{}{arg1, arg2, arg3})
	fake.listJobsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeJobLister) ListJobsCallCount() int {
	fake.listJobsMutex.RLock()
	defer fake.listJobsMutex.RUnlock()
	return len(fake.listJobsArgsForCall)
}

func (fake *FakeJobLister) ListJobsCalls(stub func(context.Context, string, string) (*v1.JobList, error)) {
	fake.listJobsMutex.Lock()
	defer fake.listJobsMutex.Unlock()
	fake.ListJobsStub = stub
}

func (fake *FakeJobLister) ListJobsArgsForCall(i int) (context.Context, string, string) {
	fake.listJobsMutex.RLock()
	defer fake.listJobsMutex.RUnlock()
	argsForCall := fake.listJobsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeJobLister) ListJobsReturns(result1 *v1.JobList, result2 error) {
	fake.listJobsMutex.Lock()
	defer fake.listJobsMutex.Unlock()
	fake.ListJobsStub = nil
	fake.listJobsReturns = struct {
		result1 *v1.JobList
		result2 error
	}{result1, result2}
}

func (fake *FakeJobLister) ListJobsReturnsOnCall(i int, result1 *v1.JobList, result2 error) {
	fake.listJobsMutex.Lock()
	defer fake.listJobsMutex.Unlock()
	fake.ListJobsStub = nil
	if fake.listJobsReturnsOnCall == nil {
		fake.listJobsReturnsOnCall = make(map[int]struct {
			result1 *v1.JobList
			result2 error
		})
	}
	fake.listJobsReturnsOnCall[i] = struct {
		result1 *v1.JobList
		result2 error
	}{result1, result2}
}

func (fake *FakeJobLister) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listJobsMutex.RLock()
	defer fake.listJobsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeJobLister) recordInvocation(key string, args []interface{}) {
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

var _ application.JobLister = new(FakeJobLister)