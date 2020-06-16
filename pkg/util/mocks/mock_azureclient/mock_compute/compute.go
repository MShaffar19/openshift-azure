// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/openshift-azure/pkg/util/azureclient/compute (interfaces: VirtualMachineScaleSetExtensionsClient,VirtualMachineScaleSetsClient,VirtualMachineScaleSetVMsClient)

// Package mock_compute is a generated GoMock package.
package mock_compute

import (
	context "context"
	reflect "reflect"

	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	autorest "github.com/Azure/go-autorest/autorest"
	gomock "github.com/golang/mock/gomock"
)

// MockVirtualMachineScaleSetExtensionsClient is a mock of VirtualMachineScaleSetExtensionsClient interface.
type MockVirtualMachineScaleSetExtensionsClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMachineScaleSetExtensionsClientMockRecorder
}

// MockVirtualMachineScaleSetExtensionsClientMockRecorder is the mock recorder for MockVirtualMachineScaleSetExtensionsClient.
type MockVirtualMachineScaleSetExtensionsClientMockRecorder struct {
	mock *MockVirtualMachineScaleSetExtensionsClient
}

// NewMockVirtualMachineScaleSetExtensionsClient creates a new mock instance.
func NewMockVirtualMachineScaleSetExtensionsClient(ctrl *gomock.Controller) *MockVirtualMachineScaleSetExtensionsClient {
	mock := &MockVirtualMachineScaleSetExtensionsClient{ctrl: ctrl}
	mock.recorder = &MockVirtualMachineScaleSetExtensionsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMachineScaleSetExtensionsClient) EXPECT() *MockVirtualMachineScaleSetExtensionsClientMockRecorder {
	return m.recorder
}

// Client mocks base method.
func (m *MockVirtualMachineScaleSetExtensionsClient) Client() autorest.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Client")
	ret0, _ := ret[0].(autorest.Client)
	return ret0
}

// Client indicates an expected call of Client.
func (mr *MockVirtualMachineScaleSetExtensionsClientMockRecorder) Client() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Client", reflect.TypeOf((*MockVirtualMachineScaleSetExtensionsClient)(nil).Client))
}

// CreateOrUpdate mocks base method.
func (m *MockVirtualMachineScaleSetExtensionsClient) CreateOrUpdate(arg0 context.Context, arg1, arg2, arg3 string, arg4 compute.VirtualMachineScaleSetExtension) (compute.VirtualMachineScaleSetExtensionsCreateOrUpdateFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSetExtensionsCreateOrUpdateFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockVirtualMachineScaleSetExtensionsClientMockRecorder) CreateOrUpdate(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockVirtualMachineScaleSetExtensionsClient)(nil).CreateOrUpdate), arg0, arg1, arg2, arg3, arg4)
}

// Get mocks base method.
func (m *MockVirtualMachineScaleSetExtensionsClient) Get(arg0 context.Context, arg1, arg2, arg3, arg4 string) (compute.VirtualMachineScaleSetExtension, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSetExtension)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockVirtualMachineScaleSetExtensionsClientMockRecorder) Get(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVirtualMachineScaleSetExtensionsClient)(nil).Get), arg0, arg1, arg2, arg3, arg4)
}

// List mocks base method.
func (m *MockVirtualMachineScaleSetExtensionsClient) List(arg0 context.Context, arg1, arg2 string) (compute.VirtualMachineScaleSetExtensionListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSetExtensionListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockVirtualMachineScaleSetExtensionsClientMockRecorder) List(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVirtualMachineScaleSetExtensionsClient)(nil).List), arg0, arg1, arg2)
}

// MockVirtualMachineScaleSetsClient is a mock of VirtualMachineScaleSetsClient interface.
type MockVirtualMachineScaleSetsClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMachineScaleSetsClientMockRecorder
}

// MockVirtualMachineScaleSetsClientMockRecorder is the mock recorder for MockVirtualMachineScaleSetsClient.
type MockVirtualMachineScaleSetsClientMockRecorder struct {
	mock *MockVirtualMachineScaleSetsClient
}

// NewMockVirtualMachineScaleSetsClient creates a new mock instance.
func NewMockVirtualMachineScaleSetsClient(ctrl *gomock.Controller) *MockVirtualMachineScaleSetsClient {
	mock := &MockVirtualMachineScaleSetsClient{ctrl: ctrl}
	mock.recorder = &MockVirtualMachineScaleSetsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMachineScaleSetsClient) EXPECT() *MockVirtualMachineScaleSetsClientMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method.
func (m *MockVirtualMachineScaleSetsClient) CreateOrUpdate(arg0 context.Context, arg1, arg2 string, arg3 compute.VirtualMachineScaleSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate.
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) CreateOrUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).CreateOrUpdate), arg0, arg1, arg2, arg3)
}

// Delete mocks base method.
func (m *MockVirtualMachineScaleSetsClient) Delete(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).Delete), arg0, arg1, arg2)
}

// Get mocks base method.
func (m *MockVirtualMachineScaleSetsClient) Get(arg0 context.Context, arg1, arg2 string) (compute.VirtualMachineScaleSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).Get), arg0, arg1, arg2)
}

// GetInstanceView mocks base method.
func (m *MockVirtualMachineScaleSetsClient) GetInstanceView(arg0 context.Context, arg1, arg2 string) (compute.VirtualMachineScaleSetInstanceView, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetInstanceView", arg0, arg1, arg2)
	ret0, _ := ret[0].(compute.VirtualMachineScaleSetInstanceView)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstanceView indicates an expected call of GetInstanceView.
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) GetInstanceView(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstanceView", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).GetInstanceView), arg0, arg1, arg2)
}

// List mocks base method.
func (m *MockVirtualMachineScaleSetsClient) List(arg0 context.Context, arg1 string) ([]compute.VirtualMachineScaleSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].([]compute.VirtualMachineScaleSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).List), arg0, arg1)
}

// Update mocks base method.
func (m *MockVirtualMachineScaleSetsClient) Update(arg0 context.Context, arg1, arg2 string, arg3 compute.VirtualMachineScaleSetUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) Update(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).Update), arg0, arg1, arg2, arg3)
}

// UpdateInstances mocks base method.
func (m *MockVirtualMachineScaleSetsClient) UpdateInstances(arg0 context.Context, arg1, arg2 string, arg3 compute.VirtualMachineScaleSetVMInstanceRequiredIDs) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateInstances", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateInstances indicates an expected call of UpdateInstances.
func (mr *MockVirtualMachineScaleSetsClientMockRecorder) UpdateInstances(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateInstances", reflect.TypeOf((*MockVirtualMachineScaleSetsClient)(nil).UpdateInstances), arg0, arg1, arg2, arg3)
}

// MockVirtualMachineScaleSetVMsClient is a mock of VirtualMachineScaleSetVMsClient interface.
type MockVirtualMachineScaleSetVMsClient struct {
	ctrl     *gomock.Controller
	recorder *MockVirtualMachineScaleSetVMsClientMockRecorder
}

// MockVirtualMachineScaleSetVMsClientMockRecorder is the mock recorder for MockVirtualMachineScaleSetVMsClient.
type MockVirtualMachineScaleSetVMsClientMockRecorder struct {
	mock *MockVirtualMachineScaleSetVMsClient
}

// NewMockVirtualMachineScaleSetVMsClient creates a new mock instance.
func NewMockVirtualMachineScaleSetVMsClient(ctrl *gomock.Controller) *MockVirtualMachineScaleSetVMsClient {
	mock := &MockVirtualMachineScaleSetVMsClient{ctrl: ctrl}
	mock.recorder = &MockVirtualMachineScaleSetVMsClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVirtualMachineScaleSetVMsClient) EXPECT() *MockVirtualMachineScaleSetVMsClientMockRecorder {
	return m.recorder
}

// Deallocate mocks base method.
func (m *MockVirtualMachineScaleSetVMsClient) Deallocate(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deallocate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deallocate indicates an expected call of Deallocate.
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) Deallocate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deallocate", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).Deallocate), arg0, arg1, arg2, arg3)
}

// Delete mocks base method.
func (m *MockVirtualMachineScaleSetVMsClient) Delete(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) Delete(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).Delete), arg0, arg1, arg2, arg3)
}

// List mocks base method.
func (m *MockVirtualMachineScaleSetVMsClient) List(arg0 context.Context, arg1, arg2, arg3, arg4, arg5 string) ([]compute.VirtualMachineScaleSetVM, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]compute.VirtualMachineScaleSetVM)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) List(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).List), arg0, arg1, arg2, arg3, arg4, arg5)
}

// Reimage mocks base method.
func (m *MockVirtualMachineScaleSetVMsClient) Reimage(arg0 context.Context, arg1, arg2, arg3 string, arg4 *compute.VirtualMachineScaleSetVMReimageParameters) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reimage", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// Reimage indicates an expected call of Reimage.
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) Reimage(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reimage", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).Reimage), arg0, arg1, arg2, arg3, arg4)
}

// Restart mocks base method.
func (m *MockVirtualMachineScaleSetVMsClient) Restart(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restart", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Restart indicates an expected call of Restart.
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) Restart(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restart", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).Restart), arg0, arg1, arg2, arg3)
}

// RunCommand mocks base method.
func (m *MockVirtualMachineScaleSetVMsClient) RunCommand(arg0 context.Context, arg1, arg2, arg3 string, arg4 compute.RunCommandInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RunCommand", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// RunCommand indicates an expected call of RunCommand.
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) RunCommand(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RunCommand", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).RunCommand), arg0, arg1, arg2, arg3, arg4)
}

// Start mocks base method.
func (m *MockVirtualMachineScaleSetVMsClient) Start(arg0 context.Context, arg1, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockVirtualMachineScaleSetVMsClientMockRecorder) Start(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockVirtualMachineScaleSetVMsClient)(nil).Start), arg0, arg1, arg2, arg3)
}
