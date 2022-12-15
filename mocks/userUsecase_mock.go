// Code generated by MockGen. DO NOT EDIT.
// Source: userUsecase.go

// Package mock_usecase is a generated GoMock package.
package mocks

import (
        multipart "mime/multipart"
        reflect "reflect"
        model "serv/domain/model"
        auth "serv/microservices/auth/gen_files"

        gomock "github.com/golang/mock/gomock"
)

// MockUserUsecaseInterface is a mock of UserUsecaseInterface interface.
type MockUserUsecaseInterface struct {
        ctrl     *gomock.Controller
        recorder *MockUserUsecaseInterfaceMockRecorder
}

// MockUserUsecaseInterfaceMockRecorder is the mock recorder for MockUserUsecaseInterface.
type MockUserUsecaseInterfaceMockRecorder struct {
        mock *MockUserUsecaseInterface
}

// NewMockUserUsecaseInterface creates a new mock instance.
func NewMockUserUsecaseInterface(ctrl *gomock.Controller) *MockUserUsecaseInterface {
        mock := &MockUserUsecaseInterface{ctrl: ctrl}
        mock.recorder = &MockUserUsecaseInterfaceMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserUsecaseInterface) EXPECT() *MockUserUsecaseInterfaceMockRecorder {
        return m.recorder
}

// AddUser mocks base method.
func (m *MockUserUsecaseInterface) AddUser(params *model.UserCreateParams) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "AddUser", params)
        ret0, _ := ret[0].(error)
        return ret0
}

// AddUser indicates an expected call of AddUser.
func (mr *MockUserUsecaseInterfaceMockRecorder) AddUser(params interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockUserUsecaseInterface)(nil).AddUser), params)
}

// ChangeEmail mocks base method.
func (m *MockUserUsecaseInterface) ChangeEmail(sessID, newEmail string) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "ChangeEmail", sessID, newEmail)
        ret0, _ := ret[0].(error)
        return ret0
}

// ChangeEmail indicates an expected call of ChangeEmail.
func (mr *MockUserUsecaseInterfaceMockRecorder) ChangeEmail(sessID, newEmail interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeEmail", reflect.TypeOf((*MockUserUsecaseInterface)(nil).ChangeEmail), sessID, newEmail)
}

// ChangeUser mocks base method.
func (m *MockUserUsecaseInterface) ChangeUser(oldUserData, params *model.UserProfile) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "ChangeUser", oldUserData, params)
        ret0, _ := ret[0].(error)
        return ret0
}

// ChangeUser indicates an expected call of ChangeUser.
func (mr *MockUserUsecaseInterfaceMockRecorder) ChangeUser(oldUserData, params interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUser", reflect.TypeOf((*MockUserUsecaseInterface)(nil).ChangeUser), oldUserData, params)
}

// ChangeUserAddresses mocks base method.
func (m *MockUserUsecaseInterface) ChangeUserAddresses(userID int, userAddresses, queryAddresses []*model.Address) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "ChangeUserAddresses", userID, userAddresses, queryAddresses)
        ret0, _ := ret[0].(error)
        return ret0
}

// ChangeUserAddresses indicates an expected call of ChangeUserAddresses.
func (mr *MockUserUsecaseInterfaceMockRecorder) ChangeUserAddresses(userID, userAddresses, queryAddresses interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserAddresses", reflect.TypeOf((*MockUserUsecaseInterface)(nil).ChangeUserAddresses), userID, userAddresses, queryAddresses)
}

// ChangeUserPassword mocks base method.
func (m *MockUserUsecaseInterface) ChangeUserPassword(userID int, newPass string) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "ChangeUserPassword", userID, newPass)
        ret0, _ := ret[0].(error)
        return ret0
}

// ChangeUserPassword indicates an expected call of ChangeUserPassword.
func (mr *MockUserUsecaseInterfaceMockRecorder) ChangeUserPassword(userID, newPass interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserPassword", reflect.TypeOf((*MockUserUsecaseInterface)(nil).ChangeUserPassword), userID, newPass)
}

// ChangeUserPayments mocks base method.
func (m *MockUserUsecaseInterface) ChangeUserPayments(userID int, userPayments, queryPayments []*model.PaymentMethod) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "ChangeUserPayments", userID, userPayments, queryPayments)
        ret0, _ := ret[0].(error)
        return ret0
}

// ChangeUserPayments indicates an expected call of ChangeUserPayments.
func (mr *MockUserUsecaseInterfaceMockRecorder) ChangeUserPayments(userID, userPayments, queryPayments interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeUserPayments", reflect.TypeOf((*MockUserUsecaseInterface)(nil).ChangeUserPayments), userID, userPayments, queryPayments)
}

// CheckSession mocks base method.
func (m *MockUserUsecaseInterface) CheckSession(sessID string) (string, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CheckSession", sessID)
        ret0, _ := ret[0].(string)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CheckSession indicates an expected call of CheckSession.
func (mr *MockUserUsecaseInterfaceMockRecorder) CheckSession(sessID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSession", reflect.TypeOf((*MockUserUsecaseInterface)(nil).CheckSession), sessID)
}

// DeleteSession mocks base method.
func (m *MockUserUsecaseInterface) DeleteSession(sessID string) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "DeleteSession", sessID)
        ret0, _ := ret[0].(error)
        return ret0
}

// DeleteSession indicates an expected call of DeleteSession.
func (mr *MockUserUsecaseInterfaceMockRecorder) DeleteSession(sessID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSession", reflect.TypeOf((*MockUserUsecaseInterface)(nil).DeleteSession), sessID)
}

// GetAddressesByUserID mocks base method.
func (m *MockUserUsecaseInterface) GetAddressesByUserID(userID int) ([]*model.Address, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetAddressesByUserID", userID)
        ret0, _ := ret[0].([]*model.Address)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetAddressesByUserID indicates an expected call of GetAddressesByUserID.
func (mr *MockUserUsecaseInterfaceMockRecorder) GetAddressesByUserID(userID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddressesByUserID", reflect.TypeOf((*MockUserUsecaseInterface)(nil).GetAddressesByUserID), userID)
}

// GetPaymentMethodByUserID mocks base method.
func (m *MockUserUsecaseInterface) GetPaymentMethodByUserID(userID int) ([]*model.PaymentMethod, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetPaymentMethodByUserID", userID)
        ret0, _ := ret[0].([]*model.PaymentMethod)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetPaymentMethodByUserID indicates an expected call of GetPaymentMethodByUserID.
func (mr *MockUserUsecaseInterfaceMockRecorder) GetPaymentMethodByUserID(userID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaymentMethodByUserID", reflect.TypeOf((*MockUserUsecaseInterface)(nil).GetPaymentMethodByUserID), userID)
}

// GetUserByUsername mocks base method.
func (m *MockUserUsecaseInterface) GetUserByUsername(email string) (model.UserDB, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetUserByUsername", email)
        ret0, _ := ret[0].(model.UserDB)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetUserByUsername indicates an expected call of GetUserByUsername.
func (mr *MockUserUsecaseInterfaceMockRecorder) GetUserByUsername(email interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByUsername", reflect.TypeOf((*MockUserUsecaseInterface)(nil).GetUserByUsername), email)
}

// SetAvatar mocks base method.
func (m *MockUserUsecaseInterface) SetAvatar(usedID int, file multipart.File) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "SetAvatar", usedID, file)
        ret0, _ := ret[0].(error)
        return ret0
}

// SetAvatar indicates an expected call of SetAvatar.
func (mr *MockUserUsecaseInterfaceMockRecorder) SetAvatar(usedID, file interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAvatar", reflect.TypeOf((*MockUserUsecaseInterface)(nil).SetAvatar), usedID, file)
}

// SetSession mocks base method.
func (m *MockUserUsecaseInterface) SetSession(userEmail string) (*auth.SessionID, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "SetSession", userEmail)
        ret0, _ := ret[0].(*auth.SessionID)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// SetSession indicates an expected call of SetSession.
func (mr *MockUserUsecaseInterfaceMockRecorder) SetSession(userEmail interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSession", reflect.TypeOf((*MockUserUsecaseInterface)(nil).SetSession), userEmail)
}

// SetUsernamesForComments mocks base method.
func (m *MockUserUsecaseInterface) SetUsernamesForComments(comms []*model.CommentDB) ([]*model.Comment, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "SetUsernamesForComments", comms)
        ret0, _ := ret[0].([]*model.Comment)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// SetUsernamesForComments indicates an expected call of SetUsernamesForComments.
func (mr *MockUserUsecaseInterfaceMockRecorder) SetUsernamesForComments(comms interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetUsernamesForComments", reflect.TypeOf((*MockUserUsecaseInterface)(nil).SetUsernamesForComments), comms)
}