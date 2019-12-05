// Code generated by MockGen. DO NOT EDIT.
// Source: infra.go

// Package mock_infra is a generated GoMock package.
package mock_infra

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	types "github.com/metrue/fx/types"
	reflect "reflect"
)

// MockProvisioner is a mock of Provisioner interface
type MockProvisioner struct {
	ctrl     *gomock.Controller
	recorder *MockProvisionerMockRecorder
}

// MockProvisionerMockRecorder is the mock recorder for MockProvisioner
type MockProvisionerMockRecorder struct {
	mock *MockProvisioner
}

// NewMockProvisioner creates a new mock instance
func NewMockProvisioner(ctrl *gomock.Controller) *MockProvisioner {
	mock := &MockProvisioner{ctrl: ctrl}
	mock.recorder = &MockProvisionerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProvisioner) EXPECT() *MockProvisionerMockRecorder {
	return m.recorder
}

// Provision mocks base method
func (m *MockProvisioner) Provision() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Provision")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Provision indicates an expected call of Provision
func (mr *MockProvisionerMockRecorder) Provision() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provision", reflect.TypeOf((*MockProvisioner)(nil).Provision))
}

// HealthCheck mocks base method
func (m *MockProvisioner) HealthCheck() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HealthCheck indicates an expected call of HealthCheck
func (mr *MockProvisionerMockRecorder) HealthCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockProvisioner)(nil).HealthCheck))
}

// MockDeployer is a mock of Deployer interface
type MockDeployer struct {
	ctrl     *gomock.Controller
	recorder *MockDeployerMockRecorder
}

// MockDeployerMockRecorder is the mock recorder for MockDeployer
type MockDeployerMockRecorder struct {
	mock *MockDeployer
}

// NewMockDeployer creates a new mock instance
func NewMockDeployer(ctrl *gomock.Controller) *MockDeployer {
	mock := &MockDeployer{ctrl: ctrl}
	mock.recorder = &MockDeployerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDeployer) EXPECT() *MockDeployerMockRecorder {
	return m.recorder
}

// Deploy mocks base method
func (m *MockDeployer) Deploy(ctx context.Context, fn, name, image string, bindings []types.PortBinding) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", ctx, fn, name, image, bindings)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deploy indicates an expected call of Deploy
func (mr *MockDeployerMockRecorder) Deploy(ctx, fn, name, image, bindings interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockDeployer)(nil).Deploy), ctx, fn, name, image, bindings)
}

// Destroy mocks base method
func (m *MockDeployer) Destroy(ctx context.Context, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy
func (mr *MockDeployerMockRecorder) Destroy(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockDeployer)(nil).Destroy), ctx, name)
}

// Update mocks base method
func (m *MockDeployer) Update(ctx context.Context, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockDeployerMockRecorder) Update(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDeployer)(nil).Update), ctx, name)
}

// GetStatus mocks base method
func (m *MockDeployer) GetStatus(ctx context.Context, name string) (types.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus", ctx, name)
	ret0, _ := ret[0].(types.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatus indicates an expected call of GetStatus
func (mr *MockDeployerMockRecorder) GetStatus(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockDeployer)(nil).GetStatus), ctx, name)
}

// List mocks base method
func (m *MockDeployer) List(ctx context.Context, name string) ([]types.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, name)
	ret0, _ := ret[0].([]types.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockDeployerMockRecorder) List(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockDeployer)(nil).List), ctx, name)
}

// Ping mocks base method
func (m *MockDeployer) Ping(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping
func (mr *MockDeployerMockRecorder) Ping(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockDeployer)(nil).Ping), ctx)
}

// MockInfra is a mock of Infra interface
type MockInfra struct {
	ctrl     *gomock.Controller
	recorder *MockInfraMockRecorder
}

// MockInfraMockRecorder is the mock recorder for MockInfra
type MockInfraMockRecorder struct {
	mock *MockInfra
}

// NewMockInfra creates a new mock instance
func NewMockInfra(ctrl *gomock.Controller) *MockInfra {
	mock := &MockInfra{ctrl: ctrl}
	mock.recorder = &MockInfraMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInfra) EXPECT() *MockInfraMockRecorder {
	return m.recorder
}

// Provision mocks base method
func (m *MockInfra) Provision() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Provision")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Provision indicates an expected call of Provision
func (mr *MockInfraMockRecorder) Provision() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Provision", reflect.TypeOf((*MockInfra)(nil).Provision))
}

// HealthCheck mocks base method
func (m *MockInfra) HealthCheck() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HealthCheck indicates an expected call of HealthCheck
func (mr *MockInfraMockRecorder) HealthCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockInfra)(nil).HealthCheck))
}

// Deploy mocks base method
func (m *MockInfra) Deploy(ctx context.Context, fn, name, image string, bindings []types.PortBinding) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", ctx, fn, name, image, bindings)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deploy indicates an expected call of Deploy
func (mr *MockInfraMockRecorder) Deploy(ctx, fn, name, image, bindings interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockInfra)(nil).Deploy), ctx, fn, name, image, bindings)
}

// Destroy mocks base method
func (m *MockInfra) Destroy(ctx context.Context, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy
func (mr *MockInfraMockRecorder) Destroy(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockInfra)(nil).Destroy), ctx, name)
}

// Update mocks base method
func (m *MockInfra) Update(ctx context.Context, name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, name)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockInfraMockRecorder) Update(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockInfra)(nil).Update), ctx, name)
}

// GetStatus mocks base method
func (m *MockInfra) GetStatus(ctx context.Context, name string) (types.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus", ctx, name)
	ret0, _ := ret[0].(types.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatus indicates an expected call of GetStatus
func (mr *MockInfraMockRecorder) GetStatus(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockInfra)(nil).GetStatus), ctx, name)
}

// List mocks base method
func (m *MockInfra) List(ctx context.Context, name string) ([]types.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, name)
	ret0, _ := ret[0].([]types.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockInfraMockRecorder) List(ctx, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockInfra)(nil).List), ctx, name)
}

// Ping mocks base method
func (m *MockInfra) Ping(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping
func (mr *MockInfraMockRecorder) Ping(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockInfra)(nil).Ping), ctx)
}
