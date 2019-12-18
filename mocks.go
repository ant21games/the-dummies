// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package lugo4go is a generated GoMock package.
package lugo4go

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	proto "github.com/lugobots/lugo4go/v2/proto"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockLogger is a mock of Logger interface
type MockLogger struct {
	ctrl     *gomock.Controller
	recorder *MockLoggerMockRecorder
}

// MockLoggerMockRecorder is the mock recorder for MockLogger
type MockLoggerMockRecorder struct {
	mock *MockLogger
}

// NewMockLogger creates a new mock instance
func NewMockLogger(ctrl *gomock.Controller) *MockLogger {
	mock := &MockLogger{ctrl: ctrl}
	mock.recorder = &MockLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLogger) EXPECT() *MockLoggerMockRecorder {
	return m.recorder
}

// Debug mocks base method
func (m *MockLogger) Debug(args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug
func (mr *MockLoggerMockRecorder) Debug(args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockLogger)(nil).Debug), args...)
}

// Debugf mocks base method
func (m *MockLogger) Debugf(template string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{template}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf
func (mr *MockLoggerMockRecorder) Debugf(template interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{template}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockLogger)(nil).Debugf), varargs...)
}

// Infof mocks base method
func (m *MockLogger) Infof(template string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{template}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof
func (mr *MockLoggerMockRecorder) Infof(template interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{template}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockLogger)(nil).Infof), varargs...)
}

// Warnf mocks base method
func (m *MockLogger) Warnf(template string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{template}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warnf", varargs...)
}

// Warnf indicates an expected call of Warnf
func (mr *MockLoggerMockRecorder) Warnf(template interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{template}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warnf", reflect.TypeOf((*MockLogger)(nil).Warnf), varargs...)
}

// Errorf mocks base method
func (m *MockLogger) Errorf(template string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{template}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf
func (mr *MockLoggerMockRecorder) Errorf(template interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{template}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockLogger)(nil).Errorf), varargs...)
}

// Fatalf mocks base method
func (m *MockLogger) Fatalf(template string, args ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{template}
	for _, a := range args {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Fatalf", varargs...)
}

// Fatalf indicates an expected call of Fatalf
func (mr *MockLoggerMockRecorder) Fatalf(template interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{template}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fatalf", reflect.TypeOf((*MockLogger)(nil).Fatalf), varargs...)
}

// MockClient is a mock of Client interface
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// OnNewTurn mocks base method
func (m *MockClient) OnNewTurn(arg0 DecisionMaker, arg1 Logger) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnNewTurn", arg0, arg1)
}

// OnNewTurn indicates an expected call of OnNewTurn
func (mr *MockClientMockRecorder) OnNewTurn(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnNewTurn", reflect.TypeOf((*MockClient)(nil).OnNewTurn), arg0, arg1)
}

// Stop mocks base method
func (m *MockClient) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockClientMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockClient)(nil).Stop))
}

// GetGRPCConn mocks base method
func (m *MockClient) GetGRPCConn() *grpc.ClientConn {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGRPCConn")
	ret0, _ := ret[0].(*grpc.ClientConn)
	return ret0
}

// GetGRPCConn indicates an expected call of GetGRPCConn
func (mr *MockClientMockRecorder) GetGRPCConn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGRPCConn", reflect.TypeOf((*MockClient)(nil).GetGRPCConn))
}

// GetServiceConn mocks base method
func (m *MockClient) GetServiceConn() proto.GameClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceConn")
	ret0, _ := ret[0].(proto.GameClient)
	return ret0
}

// GetServiceConn indicates an expected call of GetServiceConn
func (mr *MockClientMockRecorder) GetServiceConn() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceConn", reflect.TypeOf((*MockClient)(nil).GetServiceConn))
}

// SenderBuilder mocks base method
func (m *MockClient) SenderBuilder(builder func(*proto.GameSnapshot, Logger) OrderSender) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SenderBuilder", builder)
}

// SenderBuilder indicates an expected call of SenderBuilder
func (mr *MockClientMockRecorder) SenderBuilder(builder interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SenderBuilder", reflect.TypeOf((*MockClient)(nil).SenderBuilder), builder)
}

// MockOrderSender is a mock of OrderSender interface
type MockOrderSender struct {
	ctrl     *gomock.Controller
	recorder *MockOrderSenderMockRecorder
}

// MockOrderSenderMockRecorder is the mock recorder for MockOrderSender
type MockOrderSenderMockRecorder struct {
	mock *MockOrderSender
}

// NewMockOrderSender creates a new mock instance
func NewMockOrderSender(ctrl *gomock.Controller) *MockOrderSender {
	mock := &MockOrderSender{ctrl: ctrl}
	mock.recorder = &MockOrderSenderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrderSender) EXPECT() *MockOrderSenderMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockOrderSender) Send(ctx context.Context, orders []proto.PlayerOrder, debugMsg string) (*proto.OrderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", ctx, orders, debugMsg)
	ret0, _ := ret[0].(*proto.OrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send
func (mr *MockOrderSenderMockRecorder) Send(ctx, orders, debugMsg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockOrderSender)(nil).Send), ctx, orders, debugMsg)
}