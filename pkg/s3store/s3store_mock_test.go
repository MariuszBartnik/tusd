// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/MariuszBartnik/tusd/v3/pkg/s3store (interfaces: S3API)

// Package s3store is a generated GoMock package.
package s3store

import (
	context "context"
	reflect "reflect"

	s3 "github.com/aws/aws-sdk-go-v2/service/s3"
	gomock "github.com/golang/mock/gomock"
)

// MockS3API is a mock of S3API interface.
type MockS3API struct {
	ctrl     *gomock.Controller
	recorder *MockS3APIMockRecorder
}

// MockS3APIMockRecorder is the mock recorder for MockS3API.
type MockS3APIMockRecorder struct {
	mock *MockS3API
}

// NewMockS3API creates a new mock instance.
func NewMockS3API(ctrl *gomock.Controller) *MockS3API {
	mock := &MockS3API{ctrl: ctrl}
	mock.recorder = &MockS3APIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockS3API) EXPECT() *MockS3APIMockRecorder {
	return m.recorder
}

// AbortMultipartUpload mocks base method.
func (m *MockS3API) AbortMultipartUpload(arg0 context.Context, arg1 *s3.AbortMultipartUploadInput, arg2 ...func(*s3.Options)) (*s3.AbortMultipartUploadOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AbortMultipartUpload", varargs...)
	ret0, _ := ret[0].(*s3.AbortMultipartUploadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AbortMultipartUpload indicates an expected call of AbortMultipartUpload.
func (mr *MockS3APIMockRecorder) AbortMultipartUpload(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbortMultipartUpload", reflect.TypeOf((*MockS3API)(nil).AbortMultipartUpload), varargs...)
}

// CompleteMultipartUpload mocks base method.
func (m *MockS3API) CompleteMultipartUpload(arg0 context.Context, arg1 *s3.CompleteMultipartUploadInput, arg2 ...func(*s3.Options)) (*s3.CompleteMultipartUploadOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CompleteMultipartUpload", varargs...)
	ret0, _ := ret[0].(*s3.CompleteMultipartUploadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CompleteMultipartUpload indicates an expected call of CompleteMultipartUpload.
func (mr *MockS3APIMockRecorder) CompleteMultipartUpload(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CompleteMultipartUpload", reflect.TypeOf((*MockS3API)(nil).CompleteMultipartUpload), varargs...)
}

// CreateMultipartUpload mocks base method.
func (m *MockS3API) CreateMultipartUpload(arg0 context.Context, arg1 *s3.CreateMultipartUploadInput, arg2 ...func(*s3.Options)) (*s3.CreateMultipartUploadOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateMultipartUpload", varargs...)
	ret0, _ := ret[0].(*s3.CreateMultipartUploadOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMultipartUpload indicates an expected call of CreateMultipartUpload.
func (mr *MockS3APIMockRecorder) CreateMultipartUpload(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMultipartUpload", reflect.TypeOf((*MockS3API)(nil).CreateMultipartUpload), varargs...)
}

// DeleteObject mocks base method.
func (m *MockS3API) DeleteObject(arg0 context.Context, arg1 *s3.DeleteObjectInput, arg2 ...func(*s3.Options)) (*s3.DeleteObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteObject", varargs...)
	ret0, _ := ret[0].(*s3.DeleteObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteObject indicates an expected call of DeleteObject.
func (mr *MockS3APIMockRecorder) DeleteObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObject", reflect.TypeOf((*MockS3API)(nil).DeleteObject), varargs...)
}

// DeleteObjects mocks base method.
func (m *MockS3API) DeleteObjects(arg0 context.Context, arg1 *s3.DeleteObjectsInput, arg2 ...func(*s3.Options)) (*s3.DeleteObjectsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteObjects", varargs...)
	ret0, _ := ret[0].(*s3.DeleteObjectsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteObjects indicates an expected call of DeleteObjects.
func (mr *MockS3APIMockRecorder) DeleteObjects(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteObjects", reflect.TypeOf((*MockS3API)(nil).DeleteObjects), varargs...)
}

// GetObject mocks base method.
func (m *MockS3API) GetObject(arg0 context.Context, arg1 *s3.GetObjectInput, arg2 ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetObject", varargs...)
	ret0, _ := ret[0].(*s3.GetObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetObject indicates an expected call of GetObject.
func (mr *MockS3APIMockRecorder) GetObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockS3API)(nil).GetObject), varargs...)
}

// HeadObject mocks base method.
func (m *MockS3API) HeadObject(arg0 context.Context, arg1 *s3.HeadObjectInput, arg2 ...func(*s3.Options)) (*s3.HeadObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HeadObject", varargs...)
	ret0, _ := ret[0].(*s3.HeadObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeadObject indicates an expected call of HeadObject.
func (mr *MockS3APIMockRecorder) HeadObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadObject", reflect.TypeOf((*MockS3API)(nil).HeadObject), varargs...)
}

// ListParts mocks base method.
func (m *MockS3API) ListParts(arg0 context.Context, arg1 *s3.ListPartsInput, arg2 ...func(*s3.Options)) (*s3.ListPartsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListParts", varargs...)
	ret0, _ := ret[0].(*s3.ListPartsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListParts indicates an expected call of ListParts.
func (mr *MockS3APIMockRecorder) ListParts(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListParts", reflect.TypeOf((*MockS3API)(nil).ListParts), varargs...)
}

// PutObject mocks base method.
func (m *MockS3API) PutObject(arg0 context.Context, arg1 *s3.PutObjectInput, arg2 ...func(*s3.Options)) (*s3.PutObjectOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PutObject", varargs...)
	ret0, _ := ret[0].(*s3.PutObjectOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutObject indicates an expected call of PutObject.
func (mr *MockS3APIMockRecorder) PutObject(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutObject", reflect.TypeOf((*MockS3API)(nil).PutObject), varargs...)
}

// UploadPart mocks base method.
func (m *MockS3API) UploadPart(arg0 context.Context, arg1 *s3.UploadPartInput, arg2 ...func(*s3.Options)) (*s3.UploadPartOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UploadPart", varargs...)
	ret0, _ := ret[0].(*s3.UploadPartOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadPart indicates an expected call of UploadPart.
func (mr *MockS3APIMockRecorder) UploadPart(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadPart", reflect.TypeOf((*MockS3API)(nil).UploadPart), varargs...)
}

// UploadPartCopy mocks base method.
func (m *MockS3API) UploadPartCopy(arg0 context.Context, arg1 *s3.UploadPartCopyInput, arg2 ...func(*s3.Options)) (*s3.UploadPartCopyOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UploadPartCopy", varargs...)
	ret0, _ := ret[0].(*s3.UploadPartCopyOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadPartCopy indicates an expected call of UploadPartCopy.
func (mr *MockS3APIMockRecorder) UploadPartCopy(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadPartCopy", reflect.TypeOf((*MockS3API)(nil).UploadPartCopy), varargs...)
}
