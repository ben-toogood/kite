// Code generated by counterfeiter. DO NOT EDIT.
package postsfakes

import (
	"context"
	"sync"

	"github.com/ben-toogood/kite/posts"
	"google.golang.org/grpc"
)

type FakePostsServiceClient struct {
	CreateStub        func(context.Context, *posts.CreateRequest, ...grpc.CallOption) (*posts.CreateResponse, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 context.Context
		arg2 *posts.CreateRequest
		arg3 []grpc.CallOption
	}
	createReturns struct {
		result1 *posts.CreateResponse
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 *posts.CreateResponse
		result2 error
	}
	ListStub        func(context.Context, *posts.ListRequest, ...grpc.CallOption) (*posts.ListResponse, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
		arg1 context.Context
		arg2 *posts.ListRequest
		arg3 []grpc.CallOption
	}
	listReturns struct {
		result1 *posts.ListResponse
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 *posts.ListResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePostsServiceClient) Create(arg1 context.Context, arg2 *posts.CreateRequest, arg3 ...grpc.CallOption) (*posts.CreateResponse, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 context.Context
		arg2 *posts.CreateRequest
		arg3 []grpc.CallOption
	}{arg1, arg2, arg3})
	stub := fake.CreateStub
	fakeReturns := fake.createReturns
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3})
	fake.createMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePostsServiceClient) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakePostsServiceClient) CreateCalls(stub func(context.Context, *posts.CreateRequest, ...grpc.CallOption) (*posts.CreateResponse, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakePostsServiceClient) CreateArgsForCall(i int) (context.Context, *posts.CreateRequest, []grpc.CallOption) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePostsServiceClient) CreateReturns(result1 *posts.CreateResponse, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *posts.CreateResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePostsServiceClient) CreateReturnsOnCall(i int, result1 *posts.CreateResponse, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 *posts.CreateResponse
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 *posts.CreateResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePostsServiceClient) List(arg1 context.Context, arg2 *posts.ListRequest, arg3 ...grpc.CallOption) (*posts.ListResponse, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
		arg1 context.Context
		arg2 *posts.ListRequest
		arg3 []grpc.CallOption
	}{arg1, arg2, arg3})
	stub := fake.ListStub
	fakeReturns := fake.listReturns
	fake.recordInvocation("List", []interface{}{arg1, arg2, arg3})
	fake.listMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePostsServiceClient) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakePostsServiceClient) ListCalls(stub func(context.Context, *posts.ListRequest, ...grpc.CallOption) (*posts.ListResponse, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakePostsServiceClient) ListArgsForCall(i int) (context.Context, *posts.ListRequest, []grpc.CallOption) {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	argsForCall := fake.listArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakePostsServiceClient) ListReturns(result1 *posts.ListResponse, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 *posts.ListResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePostsServiceClient) ListReturnsOnCall(i int, result1 *posts.ListResponse, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 *posts.ListResponse
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 *posts.ListResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePostsServiceClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePostsServiceClient) recordInvocation(key string, args []interface{}) {
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

var _ posts.PostsServiceClient = new(FakePostsServiceClient)
