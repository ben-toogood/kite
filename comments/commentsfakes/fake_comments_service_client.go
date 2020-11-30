// Code generated by counterfeiter. DO NOT EDIT.
package commentsfakes

import (
	"context"
	"sync"

	"github.com/ben-toogood/kite/comments"
	"google.golang.org/grpc"
)

type FakeCommentsServiceClient struct {
	CreateStub        func(context.Context, *comments.CreateRequest, ...grpc.CallOption) (*comments.CreateResponse, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 context.Context
		arg2 *comments.CreateRequest
		arg3 []grpc.CallOption
	}
	createReturns struct {
		result1 *comments.CreateResponse
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 *comments.CreateResponse
		result2 error
	}
	DeleteStub        func(context.Context, *comments.DeleteRequest, ...grpc.CallOption) (*comments.DeleteResponse, error)
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 context.Context
		arg2 *comments.DeleteRequest
		arg3 []grpc.CallOption
	}
	deleteReturns struct {
		result1 *comments.DeleteResponse
		result2 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 *comments.DeleteResponse
		result2 error
	}
	GetStub        func(context.Context, *comments.GetRequest, ...grpc.CallOption) (*comments.GetResponse, error)
	getMutex       sync.RWMutex
	getArgsForCall []struct {
		arg1 context.Context
		arg2 *comments.GetRequest
		arg3 []grpc.CallOption
	}
	getReturns struct {
		result1 *comments.GetResponse
		result2 error
	}
	getReturnsOnCall map[int]struct {
		result1 *comments.GetResponse
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCommentsServiceClient) Create(arg1 context.Context, arg2 *comments.CreateRequest, arg3 ...grpc.CallOption) (*comments.CreateResponse, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 context.Context
		arg2 *comments.CreateRequest
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

func (fake *FakeCommentsServiceClient) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeCommentsServiceClient) CreateCalls(stub func(context.Context, *comments.CreateRequest, ...grpc.CallOption) (*comments.CreateResponse, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeCommentsServiceClient) CreateArgsForCall(i int) (context.Context, *comments.CreateRequest, []grpc.CallOption) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCommentsServiceClient) CreateReturns(result1 *comments.CreateResponse, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 *comments.CreateResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCommentsServiceClient) CreateReturnsOnCall(i int, result1 *comments.CreateResponse, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 *comments.CreateResponse
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 *comments.CreateResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCommentsServiceClient) Delete(arg1 context.Context, arg2 *comments.DeleteRequest, arg3 ...grpc.CallOption) (*comments.DeleteResponse, error) {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 context.Context
		arg2 *comments.DeleteRequest
		arg3 []grpc.CallOption
	}{arg1, arg2, arg3})
	stub := fake.DeleteStub
	fakeReturns := fake.deleteReturns
	fake.recordInvocation("Delete", []interface{}{arg1, arg2, arg3})
	fake.deleteMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCommentsServiceClient) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeCommentsServiceClient) DeleteCalls(stub func(context.Context, *comments.DeleteRequest, ...grpc.CallOption) (*comments.DeleteResponse, error)) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeCommentsServiceClient) DeleteArgsForCall(i int) (context.Context, *comments.DeleteRequest, []grpc.CallOption) {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCommentsServiceClient) DeleteReturns(result1 *comments.DeleteResponse, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 *comments.DeleteResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCommentsServiceClient) DeleteReturnsOnCall(i int, result1 *comments.DeleteResponse, result2 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 *comments.DeleteResponse
			result2 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 *comments.DeleteResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCommentsServiceClient) Get(arg1 context.Context, arg2 *comments.GetRequest, arg3 ...grpc.CallOption) (*comments.GetResponse, error) {
	fake.getMutex.Lock()
	ret, specificReturn := fake.getReturnsOnCall[len(fake.getArgsForCall)]
	fake.getArgsForCall = append(fake.getArgsForCall, struct {
		arg1 context.Context
		arg2 *comments.GetRequest
		arg3 []grpc.CallOption
	}{arg1, arg2, arg3})
	stub := fake.GetStub
	fakeReturns := fake.getReturns
	fake.recordInvocation("Get", []interface{}{arg1, arg2, arg3})
	fake.getMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCommentsServiceClient) GetCallCount() int {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	return len(fake.getArgsForCall)
}

func (fake *FakeCommentsServiceClient) GetCalls(stub func(context.Context, *comments.GetRequest, ...grpc.CallOption) (*comments.GetResponse, error)) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = stub
}

func (fake *FakeCommentsServiceClient) GetArgsForCall(i int) (context.Context, *comments.GetRequest, []grpc.CallOption) {
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	argsForCall := fake.getArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeCommentsServiceClient) GetReturns(result1 *comments.GetResponse, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	fake.getReturns = struct {
		result1 *comments.GetResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCommentsServiceClient) GetReturnsOnCall(i int, result1 *comments.GetResponse, result2 error) {
	fake.getMutex.Lock()
	defer fake.getMutex.Unlock()
	fake.GetStub = nil
	if fake.getReturnsOnCall == nil {
		fake.getReturnsOnCall = make(map[int]struct {
			result1 *comments.GetResponse
			result2 error
		})
	}
	fake.getReturnsOnCall[i] = struct {
		result1 *comments.GetResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeCommentsServiceClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.getMutex.RLock()
	defer fake.getMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCommentsServiceClient) recordInvocation(key string, args []interface{}) {
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

var _ comments.CommentsServiceClient = new(FakeCommentsServiceClient)
