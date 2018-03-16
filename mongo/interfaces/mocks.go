// Code generated by MockGen. DO NOT EDIT.
// Source: mongo/interfaces/interfaces.go

// Package interfaces is a generated GoMock package.
package interfaces

import (
	gomock "github.com/golang/mock/gomock"
	mgo_v2 "gopkg.in/mgo.v2"
	reflect "reflect"
)

// MockMongoDB is a mock of MongoDB interface
type MockMongoDB struct {
	ctrl     *gomock.Controller
	recorder *MockMongoDBMockRecorder
}

// MockMongoDBMockRecorder is the mock recorder for MockMongoDB
type MockMongoDBMockRecorder struct {
	mock *MockMongoDB
}

// NewMockMongoDB creates a new mock instance
func NewMockMongoDB(ctrl *gomock.Controller) *MockMongoDB {
	mock := &MockMongoDB{ctrl: ctrl}
	mock.recorder = &MockMongoDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMongoDB) EXPECT() *MockMongoDBMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockMongoDB) Run(cmd, result interface{}) error {
	ret := m.ctrl.Call(m, "Run", cmd, result)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockMongoDBMockRecorder) Run(cmd, result interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockMongoDB)(nil).Run), cmd, result)
}

// C mocks base method
func (m *MockMongoDB) C(name string) (Collection, Session) {
	ret := m.ctrl.Call(m, "C", name)
	ret0, _ := ret[0].(Collection)
	ret1, _ := ret[1].(Session)
	return ret0, ret1
}

// C indicates an expected call of C
func (mr *MockMongoDBMockRecorder) C(name interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "C", reflect.TypeOf((*MockMongoDB)(nil).C), name)
}

// Close mocks base method
func (m *MockMongoDB) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockMongoDBMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMongoDB)(nil).Close))
}

// MockCollection is a mock of Collection interface
type MockCollection struct {
	ctrl     *gomock.Controller
	recorder *MockCollectionMockRecorder
}

// MockCollectionMockRecorder is the mock recorder for MockCollection
type MockCollectionMockRecorder struct {
	mock *MockCollection
}

// NewMockCollection creates a new mock instance
func NewMockCollection(ctrl *gomock.Controller) *MockCollection {
	mock := &MockCollection{ctrl: ctrl}
	mock.recorder = &MockCollectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCollection) EXPECT() *MockCollectionMockRecorder {
	return m.recorder
}

// Find mocks base method
func (m *MockCollection) Find(query interface{}) Query {
	ret := m.ctrl.Call(m, "Find", query)
	ret0, _ := ret[0].(Query)
	return ret0
}

// Find indicates an expected call of Find
func (mr *MockCollectionMockRecorder) Find(query interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockCollection)(nil).Find), query)
}

// FindId mocks base method
func (m *MockCollection) FindId(id interface{}) Query {
	ret := m.ctrl.Call(m, "FindId", id)
	ret0, _ := ret[0].(Query)
	return ret0
}

// FindId indicates an expected call of FindId
func (mr *MockCollectionMockRecorder) FindId(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindId", reflect.TypeOf((*MockCollection)(nil).FindId), id)
}

// Insert mocks base method
func (m *MockCollection) Insert(docs ...interface{}) error {
	varargs := []interface{}{}
	for _, a := range docs {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Insert", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert
func (mr *MockCollectionMockRecorder) Insert(docs ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockCollection)(nil).Insert), docs...)
}

// UpsertId mocks base method
func (m *MockCollection) UpsertId(id, update interface{}) (*mgo_v2.ChangeInfo, error) {
	ret := m.ctrl.Call(m, "UpsertId", id, update)
	ret0, _ := ret[0].(*mgo_v2.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpsertId indicates an expected call of UpsertId
func (mr *MockCollectionMockRecorder) UpsertId(id, update interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertId", reflect.TypeOf((*MockCollection)(nil).UpsertId), id, update)
}

// RemoveId mocks base method
func (m *MockCollection) RemoveId(id interface{}) error {
	ret := m.ctrl.Call(m, "RemoveId", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveId indicates an expected call of RemoveId
func (mr *MockCollectionMockRecorder) RemoveId(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveId", reflect.TypeOf((*MockCollection)(nil).RemoveId), id)
}

// MockSession is a mock of Session interface
type MockSession struct {
	ctrl     *gomock.Controller
	recorder *MockSessionMockRecorder
}

// MockSessionMockRecorder is the mock recorder for MockSession
type MockSessionMockRecorder struct {
	mock *MockSession
}

// NewMockSession creates a new mock instance
func NewMockSession(ctrl *gomock.Controller) *MockSession {
	mock := &MockSession{ctrl: ctrl}
	mock.recorder = &MockSessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSession) EXPECT() *MockSessionMockRecorder {
	return m.recorder
}

// Copy mocks base method
func (m *MockSession) Copy() *mgo_v2.Session {
	ret := m.ctrl.Call(m, "Copy")
	ret0, _ := ret[0].(*mgo_v2.Session)
	return ret0
}

// Copy indicates an expected call of Copy
func (mr *MockSessionMockRecorder) Copy() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Copy", reflect.TypeOf((*MockSession)(nil).Copy))
}

// Close mocks base method
func (m *MockSession) Close() {
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close
func (mr *MockSessionMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockSession)(nil).Close))
}

// MockQuery is a mock of Query interface
type MockQuery struct {
	ctrl     *gomock.Controller
	recorder *MockQueryMockRecorder
}

// MockQueryMockRecorder is the mock recorder for MockQuery
type MockQueryMockRecorder struct {
	mock *MockQuery
}

// NewMockQuery creates a new mock instance
func NewMockQuery(ctrl *gomock.Controller) *MockQuery {
	mock := &MockQuery{ctrl: ctrl}
	mock.recorder = &MockQueryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQuery) EXPECT() *MockQueryMockRecorder {
	return m.recorder
}

// Iter mocks base method
func (m *MockQuery) Iter() Iter {
	ret := m.ctrl.Call(m, "Iter")
	ret0, _ := ret[0].(Iter)
	return ret0
}

// Iter indicates an expected call of Iter
func (mr *MockQueryMockRecorder) Iter() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iter", reflect.TypeOf((*MockQuery)(nil).Iter))
}

// All mocks base method
func (m *MockQuery) All(result interface{}) error {
	ret := m.ctrl.Call(m, "All", result)
	ret0, _ := ret[0].(error)
	return ret0
}

// All indicates an expected call of All
func (mr *MockQueryMockRecorder) All(result interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockQuery)(nil).All), result)
}

// One mocks base method
func (m *MockQuery) One(result interface{}) error {
	ret := m.ctrl.Call(m, "One", result)
	ret0, _ := ret[0].(error)
	return ret0
}

// One indicates an expected call of One
func (mr *MockQueryMockRecorder) One(result interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockQuery)(nil).One), result)
}

// MockIter is a mock of Iter interface
type MockIter struct {
	ctrl     *gomock.Controller
	recorder *MockIterMockRecorder
}

// MockIterMockRecorder is the mock recorder for MockIter
type MockIterMockRecorder struct {
	mock *MockIter
}

// NewMockIter creates a new mock instance
func NewMockIter(ctrl *gomock.Controller) *MockIter {
	mock := &MockIter{ctrl: ctrl}
	mock.recorder = &MockIterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIter) EXPECT() *MockIterMockRecorder {
	return m.recorder
}

// Next mocks base method
func (m *MockIter) Next(result interface{}) bool {
	ret := m.ctrl.Call(m, "Next", result)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Next indicates an expected call of Next
func (mr *MockIterMockRecorder) Next(result interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockIter)(nil).Next), result)
}

// Close mocks base method
func (m *MockIter) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockIterMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockIter)(nil).Close))
}
