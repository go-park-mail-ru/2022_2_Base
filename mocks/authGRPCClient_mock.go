// Code generated by MockGen. DO NOT EDIT.
// Source: auth_grpc.pb.go

// Package mock_auth is a generated GoMock package.
package mocks

import (
        context "context"
        reflect "reflect"
        auth "serv/microservices/auth/gen_files"

        gomock "github.com/golang/mock/gomock"
        grpc "google.golang.org/grpc"
)

// MockAuthCheckerClient is a mock of AuthCheckerClient interface.
type MockAuthCheckerClient struct {
        ctrl     *gomock.Controller
        recorder *MockAuthCheckerClientMockRecorder
}

// MockAuthCheckerClientMockRecorder is the mock recorder for MockAuthCheckerClient.
type MockAuthCheckerClientMockRecorder struct {
        mock *MockAuthCheckerClient
}

// NewMockAuthCheckerClient creates a new mock instance.
func NewMockAuthCheckerClient(ctrl *gomock.Controller) *MockAuthCheckerClient {
        mock := &MockAuthCheckerClient{ctrl: ctrl}
        mock.recorder = &MockAuthCheckerClientMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthCheckerClient) EXPECT() *MockAuthCheckerClientMockRecorder {
        return m.recorder
}

// ChangeEmail mocks base method.
func (m *MockAuthCheckerClient) ChangeEmail(ctx context.Context, in *auth.NewLogin, opts ...grpc.CallOption) (*auth.Nothing, error) {
        m.ctrl.T.Helper()
        varargs := []interface{}{ctx, in}
        for _, a := range opts {
                varargs = append(varargs, a)
        }
        ret := m.ctrl.Call(m, "ChangeEmail", varargs...)
        ret0, _ := ret[0].(*auth.Nothing)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// ChangeEmail indicates an expected call of ChangeEmail.
func (mr *MockAuthCheckerClientMockRecorder) ChangeEmail(ctx, in interface{}, opts ...interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        varargs := append([]interface{}{ctx, in}, opts...)
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeEmail", reflect.TypeOf((*MockAuthCheckerClient)(nil).ChangeEmail), varargs...)
}

// Check mocks base method.
func (m *MockAuthCheckerClient) Check(ctx context.Context, in *auth.SessionID, opts ...grpc.CallOption) (*auth.Session, error) {
        m.ctrl.T.Helper()
        varargs := []interface{}{ctx, in}
        for _, a := range opts {
                varargs = append(varargs, a)
        }
        ret := m.ctrl.Call(m, "Check", varargs...)
        ret0, _ := ret[0].(*auth.Session)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Check indicates an expected call of Check.
func (mr *MockAuthCheckerClientMockRecorder) Check(ctx, in interface{}, opts ...interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        varargs := append([]interface{}{ctx, in}, opts...)
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockAuthCheckerClient)(nil).Check), varargs...)
}

// Create mocks base method.
func (m *MockAuthCheckerClient) Create(ctx context.Context, in *auth.Session, opts ...grpc.CallOption) (*auth.SessionID, error) {
        m.ctrl.T.Helper()
        varargs := []interface{}{ctx, in}
        for _, a := range opts {
                varargs = append(varargs, a)
        }
        ret := m.ctrl.Call(m, "Create", varargs...)
        ret0, _ := ret[0].(*auth.SessionID)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAuthCheckerClientMockRecorder) Create(ctx, in interface{}, opts ...interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        varargs := append([]interface{}{ctx, in}, opts...)
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuthCheckerClient)(nil).Create), varargs...)
}

// Delete mocks base method.
func (m *MockAuthCheckerClient) Delete(ctx context.Context, in *auth.SessionID, opts ...grpc.CallOption) (*auth.Nothing, error) {
        m.ctrl.T.Helper()
        varargs := []interface{}{ctx, in}
        for _, a := range opts {
                varargs = append(varargs, a)
        }
        ret := m.ctrl.Call(m, "Delete", varargs...)
        ret0, _ := ret[0].(*auth.Nothing)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockAuthCheckerClientMockRecorder) Delete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        varargs := append([]interface{}{ctx, in}, opts...)
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAuthCheckerClient)(nil).Delete), varargs...)
}

// MockAuthCheckerServer is a mock of AuthCheckerServer interface.
type MockAuthCheckerServer struct {
        ctrl     *gomock.Controller
        recorder *MockAuthCheckerServerMockRecorder
}

// MockAuthCheckerServerMockRecorder is the mock recorder for MockAuthCheckerServer.
type MockAuthCheckerServerMockRecorder struct {
        mock *MockAuthCheckerServer
}

// NewMockAuthCheckerServer creates a new mock instance.
func NewMockAuthCheckerServer(ctrl *gomock.Controller) *MockAuthCheckerServer {
        mock := &MockAuthCheckerServer{ctrl: ctrl}
        mock.recorder = &MockAuthCheckerServerMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthCheckerServer) EXPECT() *MockAuthCheckerServerMockRecorder {
        return m.recorder
}

// ChangeEmail mocks base method.
func (m *MockAuthCheckerServer) ChangeEmail(arg0 context.Context, arg1 *auth.NewLogin) (*auth.Nothing, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "ChangeEmail", arg0, arg1)
        ret0, _ := ret[0].(*auth.Nothing)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// ChangeEmail indicates an expected call of ChangeEmail.
func (mr *MockAuthCheckerServerMockRecorder) ChangeEmail(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeEmail", reflect.TypeOf((*MockAuthCheckerServer)(nil).ChangeEmail), arg0, arg1)
}

// Check mocks base method.
func (m *MockAuthCheckerServer) Check(arg0 context.Context, arg1 *auth.SessionID) (*auth.Session, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Check", arg0, arg1)
        ret0, _ := ret[0].(*auth.Session)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Check indicates an expected call of Check.
func (mr *MockAuthCheckerServerMockRecorder) Check(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Check", reflect.TypeOf((*MockAuthCheckerServer)(nil).Check), arg0, arg1)
}

// Create mocks base method.
func (m *MockAuthCheckerServer) Create(arg0 context.Context, arg1 *auth.Session) (*auth.SessionID, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Create", arg0, arg1)
        ret0, _ := ret[0].(*auth.SessionID)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAuthCheckerServerMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAuthCheckerServer)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockAuthCheckerServer) Delete(arg0 context.Context, arg1 *auth.SessionID) (*auth.Nothing, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Delete", arg0, arg1)
        ret0, _ := ret[0].(*auth.Nothing)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockAuthCheckerServerMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAuthCheckerServer)(nil).Delete), arg0, arg1)
}

// mustEmbedUnimplementedAuthCheckerServer mocks base method.
func (m *MockAuthCheckerServer) mustEmbedUnimplementedAuthCheckerServer() {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "mustEmbedUnimplementedAuthCheckerServer")
}

// mustEmbedUnimplementedAuthCheckerServer indicates an expected call of mustEmbedUnimplementedAuthCheckerServer.
func (mr *MockAuthCheckerServerMockRecorder) mustEmbedUnimplementedAuthCheckerServer() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAuthCheckerServer", reflect.TypeOf((*MockAuthCheckerServer)(nil).mustEmbedUnimplementedAuthCheckerServer))
}

// MockUnsafeAuthCheckerServer is a mock of UnsafeAuthCheckerServer interface.
type MockUnsafeAuthCheckerServer struct {
        ctrl     *gomock.Controller
        recorder *MockUnsafeAuthCheckerServerMockRecorder
}

// MockUnsafeAuthCheckerServerMockRecorder is the mock recorder for MockUnsafeAuthCheckerServer.
type MockUnsafeAuthCheckerServerMockRecorder struct {
        mock *MockUnsafeAuthCheckerServer
}

// NewMockUnsafeAuthCheckerServer creates a new mock instance.
func NewMockUnsafeAuthCheckerServer(ctrl *gomock.Controller) *MockUnsafeAuthCheckerServer {
        mock := &MockUnsafeAuthCheckerServer{ctrl: ctrl}
        mock.recorder = &MockUnsafeAuthCheckerServerMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeAuthCheckerServer) EXPECT() *MockUnsafeAuthCheckerServerMockRecorder {
        return m.recorder
}

// mustEmbedUnimplementedAuthCheckerServer mocks base method.
func (m *MockUnsafeAuthCheckerServer) mustEmbedUnimplementedAuthCheckerServer() {
        m.ctrl.T.Helper()
        m.ctrl.Call(m, "mustEmbedUnimplementedAuthCheckerServer")
}

// mustEmbedUnimplementedAuthCheckerServer indicates an expected call of mustEmbedUnimplementedAuthCheckerServer.
func (mr *MockUnsafeAuthCheckerServerMockRecorder) mustEmbedUnimplementedAuthCheckerServer() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedAuthCheckerServer", reflect.TypeOf((*MockUnsafeAuthCheckerServer)(nil).mustEmbedUnimplementedAuthCheckerServer))
}