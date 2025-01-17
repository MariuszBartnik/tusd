// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/MariuszBartnik/tusd/v3/pkg/azurestore (interfaces: AzService,AzBlob)

// Package azurestore_test is a generated GoMock package.
package azurestore_test

import (
	context "context"
	io "io"
	reflect "reflect"

	azurestore "github.com/MariuszBartnik/tusd/v3/pkg/azurestore"
	gomock "github.com/golang/mock/gomock"
)

// MockAzService is a mock of AzService interface.
type MockAzService struct {
	ctrl     *gomock.Controller
	recorder *MockAzServiceMockRecorder
}

// MockAzServiceMockRecorder is the mock recorder for MockAzService.
type MockAzServiceMockRecorder struct {
	mock *MockAzService
}

// NewMockAzService creates a new mock instance.
func NewMockAzService(ctrl *gomock.Controller) *MockAzService {
	mock := &MockAzService{ctrl: ctrl}
	mock.recorder = &MockAzServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAzService) EXPECT() *MockAzServiceMockRecorder {
	return m.recorder
}

// NewBlob mocks base method.
func (m *MockAzService) NewBlob(arg0 context.Context, arg1 string) (azurestore.AzBlob, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewBlob", arg0, arg1)
	ret0, _ := ret[0].(azurestore.AzBlob)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewBlob indicates an expected call of NewBlob.
func (mr *MockAzServiceMockRecorder) NewBlob(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewBlob", reflect.TypeOf((*MockAzService)(nil).NewBlob), arg0, arg1)
}

// MockAzBlob is a mock of AzBlob interface.
type MockAzBlob struct {
	ctrl     *gomock.Controller
	recorder *MockAzBlobMockRecorder
}

// MockAzBlobMockRecorder is the mock recorder for MockAzBlob.
type MockAzBlobMockRecorder struct {
	mock *MockAzBlob
}

// NewMockAzBlob creates a new mock instance.
func NewMockAzBlob(ctrl *gomock.Controller) *MockAzBlob {
	mock := &MockAzBlob{ctrl: ctrl}
	mock.recorder = &MockAzBlobMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAzBlob) EXPECT() *MockAzBlobMockRecorder {
	return m.recorder
}

// Commit mocks base method.
func (m *MockAzBlob) Commit(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockAzBlobMockRecorder) Commit(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockAzBlob)(nil).Commit), arg0)
}

// Delete mocks base method.
func (m *MockAzBlob) Delete(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAzBlobMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAzBlob)(nil).Delete), arg0)
}

// Download mocks base method.
func (m *MockAzBlob) Download(arg0 context.Context) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Download", arg0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Download indicates an expected call of Download.
func (mr *MockAzBlobMockRecorder) Download(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Download", reflect.TypeOf((*MockAzBlob)(nil).Download), arg0)
}

// GetOffset mocks base method.
func (m *MockAzBlob) GetOffset(arg0 context.Context) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOffset", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOffset indicates an expected call of GetOffset.
func (mr *MockAzBlobMockRecorder) GetOffset(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOffset", reflect.TypeOf((*MockAzBlob)(nil).GetOffset), arg0)
}

// Upload mocks base method.
func (m *MockAzBlob) Upload(arg0 context.Context, arg1 io.ReadSeeker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upload", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upload indicates an expected call of Upload.
func (mr *MockAzBlobMockRecorder) Upload(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upload", reflect.TypeOf((*MockAzBlob)(nil).Upload), arg0, arg1)
}
