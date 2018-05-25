// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/spiffe/spire/proto/server/noderesolver (interfaces: NodeResolver,NodeResolverClient,NodeResolverServer)

// Package mock_noderesolver is a generated GoMock package.
package mock_noderesolver

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	plugin "github.com/spiffe/spire/proto/common/plugin"
	noderesolver "github.com/spiffe/spire/proto/server/noderesolver"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockNodeResolver is a mock of NodeResolver interface
type MockNodeResolver struct {
	ctrl     *gomock.Controller
	recorder *MockNodeResolverMockRecorder
}

// MockNodeResolverMockRecorder is the mock recorder for MockNodeResolver
type MockNodeResolverMockRecorder struct {
	mock *MockNodeResolver
}

// NewMockNodeResolver creates a new mock instance
func NewMockNodeResolver(ctrl *gomock.Controller) *MockNodeResolver {
	mock := &MockNodeResolver{ctrl: ctrl}
	mock.recorder = &MockNodeResolverMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeResolver) EXPECT() *MockNodeResolverMockRecorder {
	return m.recorder
}

// Configure mocks base method
func (m *MockNodeResolver) Configure(arg0 context.Context, arg1 *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error) {
	ret := m.ctrl.Call(m, "Configure", arg0, arg1)
	ret0, _ := ret[0].(*plugin.ConfigureResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Configure indicates an expected call of Configure
func (mr *MockNodeResolverMockRecorder) Configure(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*MockNodeResolver)(nil).Configure), arg0, arg1)
}

// GetPluginInfo mocks base method
func (m *MockNodeResolver) GetPluginInfo(arg0 context.Context, arg1 *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error) {
	ret := m.ctrl.Call(m, "GetPluginInfo", arg0, arg1)
	ret0, _ := ret[0].(*plugin.GetPluginInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPluginInfo indicates an expected call of GetPluginInfo
func (mr *MockNodeResolverMockRecorder) GetPluginInfo(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPluginInfo", reflect.TypeOf((*MockNodeResolver)(nil).GetPluginInfo), arg0, arg1)
}

// Resolve mocks base method
func (m *MockNodeResolver) Resolve(arg0 context.Context, arg1 *noderesolver.ResolveRequest) (*noderesolver.ResolveResponse, error) {
	ret := m.ctrl.Call(m, "Resolve", arg0, arg1)
	ret0, _ := ret[0].(*noderesolver.ResolveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resolve indicates an expected call of Resolve
func (mr *MockNodeResolverMockRecorder) Resolve(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockNodeResolver)(nil).Resolve), arg0, arg1)
}

// MockNodeResolverClient is a mock of NodeResolverClient interface
type MockNodeResolverClient struct {
	ctrl     *gomock.Controller
	recorder *MockNodeResolverClientMockRecorder
}

// MockNodeResolverClientMockRecorder is the mock recorder for MockNodeResolverClient
type MockNodeResolverClientMockRecorder struct {
	mock *MockNodeResolverClient
}

// NewMockNodeResolverClient creates a new mock instance
func NewMockNodeResolverClient(ctrl *gomock.Controller) *MockNodeResolverClient {
	mock := &MockNodeResolverClient{ctrl: ctrl}
	mock.recorder = &MockNodeResolverClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeResolverClient) EXPECT() *MockNodeResolverClientMockRecorder {
	return m.recorder
}

// Configure mocks base method
func (m *MockNodeResolverClient) Configure(arg0 context.Context, arg1 *plugin.ConfigureRequest, arg2 ...grpc.CallOption) (*plugin.ConfigureResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Configure", varargs...)
	ret0, _ := ret[0].(*plugin.ConfigureResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Configure indicates an expected call of Configure
func (mr *MockNodeResolverClientMockRecorder) Configure(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*MockNodeResolverClient)(nil).Configure), varargs...)
}

// GetPluginInfo mocks base method
func (m *MockNodeResolverClient) GetPluginInfo(arg0 context.Context, arg1 *plugin.GetPluginInfoRequest, arg2 ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPluginInfo", varargs...)
	ret0, _ := ret[0].(*plugin.GetPluginInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPluginInfo indicates an expected call of GetPluginInfo
func (mr *MockNodeResolverClientMockRecorder) GetPluginInfo(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPluginInfo", reflect.TypeOf((*MockNodeResolverClient)(nil).GetPluginInfo), varargs...)
}

// Resolve mocks base method
func (m *MockNodeResolverClient) Resolve(arg0 context.Context, arg1 *noderesolver.ResolveRequest, arg2 ...grpc.CallOption) (*noderesolver.ResolveResponse, error) {
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Resolve", varargs...)
	ret0, _ := ret[0].(*noderesolver.ResolveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resolve indicates an expected call of Resolve
func (mr *MockNodeResolverClientMockRecorder) Resolve(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockNodeResolverClient)(nil).Resolve), varargs...)
}

// MockNodeResolverServer is a mock of NodeResolverServer interface
type MockNodeResolverServer struct {
	ctrl     *gomock.Controller
	recorder *MockNodeResolverServerMockRecorder
}

// MockNodeResolverServerMockRecorder is the mock recorder for MockNodeResolverServer
type MockNodeResolverServerMockRecorder struct {
	mock *MockNodeResolverServer
}

// NewMockNodeResolverServer creates a new mock instance
func NewMockNodeResolverServer(ctrl *gomock.Controller) *MockNodeResolverServer {
	mock := &MockNodeResolverServer{ctrl: ctrl}
	mock.recorder = &MockNodeResolverServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeResolverServer) EXPECT() *MockNodeResolverServerMockRecorder {
	return m.recorder
}

// Configure mocks base method
func (m *MockNodeResolverServer) Configure(arg0 context.Context, arg1 *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error) {
	ret := m.ctrl.Call(m, "Configure", arg0, arg1)
	ret0, _ := ret[0].(*plugin.ConfigureResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Configure indicates an expected call of Configure
func (mr *MockNodeResolverServerMockRecorder) Configure(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Configure", reflect.TypeOf((*MockNodeResolverServer)(nil).Configure), arg0, arg1)
}

// GetPluginInfo mocks base method
func (m *MockNodeResolverServer) GetPluginInfo(arg0 context.Context, arg1 *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error) {
	ret := m.ctrl.Call(m, "GetPluginInfo", arg0, arg1)
	ret0, _ := ret[0].(*plugin.GetPluginInfoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPluginInfo indicates an expected call of GetPluginInfo
func (mr *MockNodeResolverServerMockRecorder) GetPluginInfo(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPluginInfo", reflect.TypeOf((*MockNodeResolverServer)(nil).GetPluginInfo), arg0, arg1)
}

// Resolve mocks base method
func (m *MockNodeResolverServer) Resolve(arg0 context.Context, arg1 *noderesolver.ResolveRequest) (*noderesolver.ResolveResponse, error) {
	ret := m.ctrl.Call(m, "Resolve", arg0, arg1)
	ret0, _ := ret[0].(*noderesolver.ResolveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Resolve indicates an expected call of Resolve
func (mr *MockNodeResolverServerMockRecorder) Resolve(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Resolve", reflect.TypeOf((*MockNodeResolverServer)(nil).Resolve), arg0, arg1)
}
