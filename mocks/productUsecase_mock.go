// Code generated by MockGen. DO NOT EDIT.
// Source: productUsecase.go

// Package mock_usecase is a generated GoMock package.
package mocks

import (
        reflect "reflect"
        model "serv/domain/model"
        orders "serv/microservices/orders/gen_files"

        gomock "github.com/golang/mock/gomock"
)

// MockProductUsecaseInterface is a mock of ProductUsecaseInterface interface.
type MockProductUsecaseInterface struct {
        ctrl     *gomock.Controller
        recorder *MockProductUsecaseInterfaceMockRecorder
}

// MockProductUsecaseInterfaceMockRecorder is the mock recorder for MockProductUsecaseInterface.
type MockProductUsecaseInterfaceMockRecorder struct {
        mock *MockProductUsecaseInterface
}

// NewMockProductUsecaseInterface creates a new mock instance.
func NewMockProductUsecaseInterface(ctrl *gomock.Controller) *MockProductUsecaseInterface {
        mock := &MockProductUsecaseInterface{ctrl: ctrl}
        mock.recorder = &MockProductUsecaseInterfaceMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProductUsecaseInterface) EXPECT() *MockProductUsecaseInterfaceMockRecorder {
        return m.recorder
}

// AddToOrder mocks base method.
func (m *MockProductUsecaseInterface) AddToOrder(userID, itemID int) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "AddToOrder", userID, itemID)
        ret0, _ := ret[0].(error)
        return ret0
}

// AddToOrder indicates an expected call of AddToOrder.
func (mr *MockProductUsecaseInterfaceMockRecorder) AddToOrder(userID, itemID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToOrder", reflect.TypeOf((*MockProductUsecaseInterface)(nil).AddToOrder), userID, itemID)
}

// ChangeOrderStatus mocks base method.
func (m *MockProductUsecaseInterface) ChangeOrderStatus(userID int, in *model.ChangeOrderStatus) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "ChangeOrderStatus", userID, in)
        ret0, _ := ret[0].(error)
        return ret0
}

// ChangeOrderStatus indicates an expected call of ChangeOrderStatus.
func (mr *MockProductUsecaseInterfaceMockRecorder) ChangeOrderStatus(userID, in interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeOrderStatus", reflect.TypeOf((*MockProductUsecaseInterface)(nil).ChangeOrderStatus), userID, in)
}

// CreateComment mocks base method.
func (m *MockProductUsecaseInterface) CreateComment(in *model.CreateComment) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CreateComment", in)
        ret0, _ := ret[0].(error)
        return ret0
}

// CreateComment indicates an expected call of CreateComment.
func (mr *MockProductUsecaseInterfaceMockRecorder) CreateComment(in interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateComment", reflect.TypeOf((*MockProductUsecaseInterface)(nil).CreateComment), in)
}

// DeleteFromOrder mocks base method.
func (m *MockProductUsecaseInterface) DeleteFromOrder(userID, itemID int) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "DeleteFromOrder", userID, itemID)
        ret0, _ := ret[0].(error)
        return ret0
}

// DeleteFromOrder indicates an expected call of DeleteFromOrder.
func (mr *MockProductUsecaseInterfaceMockRecorder) DeleteFromOrder(userID, itemID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFromOrder", reflect.TypeOf((*MockProductUsecaseInterface)(nil).DeleteFromOrder), userID, itemID)
}

// DeleteItemFromFavorites mocks base method.
func (m *MockProductUsecaseInterface) DeleteItemFromFavorites(userID, itemID int) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "DeleteItemFromFavorites", userID, itemID)
        ret0, _ := ret[0].(error)
        return ret0
}

// DeleteItemFromFavorites indicates an expected call of DeleteItemFromFavorites.
func (mr *MockProductUsecaseInterfaceMockRecorder) DeleteItemFromFavorites(userID, itemID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteItemFromFavorites", reflect.TypeOf((*MockProductUsecaseInterface)(nil).DeleteItemFromFavorites), userID, itemID)
}

// GetBestProductInCategory mocks base method.
func (m *MockProductUsecaseInterface) GetBestProductInCategory(category string, userID int) (*model.Product, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetBestProductInCategory", category, userID)
        ret0, _ := ret[0].(*model.Product)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetBestProductInCategory indicates an expected call of GetBestProductInCategory.
func (mr *MockProductUsecaseInterfaceMockRecorder) GetBestProductInCategory(category, userID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBestProductInCategory", reflect.TypeOf((*MockProductUsecaseInterface)(nil).GetBestProductInCategory), category, userID)
}

// GetCart mocks base method.
func (m *MockProductUsecaseInterface) GetCart(userID int) (*model.Order, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetCart", userID)
        ret0, _ := ret[0].(*model.Order)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetCart indicates an expected call of GetCart.
func (mr *MockProductUsecaseInterfaceMockRecorder) GetCart(userID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCart", reflect.TypeOf((*MockProductUsecaseInterface)(nil).GetCart), userID)
}

// GetComments mocks base method.
func (m *MockProductUsecaseInterface) GetComments(productID int) ([]*model.CommentDB, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetComments", productID)
        ret0, _ := ret[0].([]*model.CommentDB)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetComments indicates an expected call of GetComments.
func (mr *MockProductUsecaseInterfaceMockRecorder) GetComments(productID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComments", reflect.TypeOf((*MockProductUsecaseInterface)(nil).GetComments), productID)
}

// GetFavorites mocks base method.
func (m *MockProductUsecaseInterface) GetFavorites(userID, lastitemid, count int, sort string) ([]*model.Product, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetFavorites", userID, lastitemid, count, sort)
        ret0, _ := ret[0].([]*model.Product)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetFavorites indicates an expected call of GetFavorites.
func (mr *MockProductUsecaseInterfaceMockRecorder) GetFavorites(userID, lastitemid, count, sort interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFavorites", reflect.TypeOf((*MockProductUsecaseInterface)(nil).GetFavorites), userID, lastitemid, count, sort)
}

// GetOrders mocks base method.
func (m *MockProductUsecaseInterface) GetOrders(userID int) (*orders.OrdersResponse, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetOrders", userID)
        ret0, _ := ret[0].(*orders.OrdersResponse)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockProductUsecaseInterfaceMockRecorder) GetOrders(userID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockProductUsecaseInterface)(nil).GetOrders), userID)
}

// GetProductByID mocks base method.
func (m *MockProductUsecaseInterface) GetProductByID(id, userID int) (*model.Product, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetProductByID", id, userID)
        ret0, _ := ret[0].(*model.Product)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetProductByID indicates an expected call of GetProductByID.
func (mr *MockProductUsecaseInterfaceMockRecorder) GetProductByID(id, userID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductByID", reflect.TypeOf((*MockProductUsecaseInterface)(nil).GetProductByID), id, userID)
}

// GetProducts mocks base method.
func (m *MockProductUsecaseInterface) GetProducts(lastitemid, count int, sort string, userID int) ([]*model.Product, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetProducts", lastitemid, count, sort, userID)
        ret0, _ := ret[0].([]*model.Product)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetProducts indicates an expected call of GetProducts.
func (mr *MockProductUsecaseInterfaceMockRecorder) GetProducts(lastitemid, count, sort, userID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProducts", reflect.TypeOf((*MockProductUsecaseInterface)(nil).GetProducts), lastitemid, count, sort, userID)
}

// GetProductsBySearch mocks base method.
func (m *MockProductUsecaseInterface) GetProductsBySearch(search string, userID int) ([]*model.Product, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetProductsBySearch", search, userID)
        ret0, _ := ret[0].([]*model.Product)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetProductsBySearch indicates an expected call of GetProductsBySearch.
func (mr *MockProductUsecaseInterfaceMockRecorder) GetProductsBySearch(search, userID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsBySearch", reflect.TypeOf((*MockProductUsecaseInterface)(nil).GetProductsBySearch), search, userID)
}

// GetProductsWithBiggestDiscount mocks base method.
func (m *MockProductUsecaseInterface) GetProductsWithBiggestDiscount(lastitemid, count, userID int) ([]*model.Product, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetProductsWithBiggestDiscount", lastitemid, count, userID)
        ret0, _ := ret[0].([]*model.Product)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetProductsWithBiggestDiscount indicates an expected call of GetProductsWithBiggestDiscount.
func (mr *MockProductUsecaseInterfaceMockRecorder) GetProductsWithBiggestDiscount(lastitemid, count, userID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsWithBiggestDiscount", reflect.TypeOf((*MockProductUsecaseInterface)(nil).GetProductsWithBiggestDiscount), lastitemid, count, userID)
}

// GetProductsWithCategory mocks base method.
func (m *MockProductUsecaseInterface) GetProductsWithCategory(cat string, lastitemid, count int, sort string, userID int) ([]*model.Product, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetProductsWithCategory", cat, lastitemid, count, sort, userID)
        ret0, _ := ret[0].([]*model.Product)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetProductsWithCategory indicates an expected call of GetProductsWithCategory.
func (mr *MockProductUsecaseInterfaceMockRecorder) GetProductsWithCategory(cat, lastitemid, count, sort, userID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductsWithCategory", reflect.TypeOf((*MockProductUsecaseInterface)(nil).GetProductsWithCategory), cat, lastitemid, count, sort, userID)
}

// GetRecommendationProducts mocks base method.
func (m *MockProductUsecaseInterface) GetRecommendationProducts(itemID, userID int) ([]*model.Product, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetRecommendationProducts", itemID, userID)
        ret0, _ := ret[0].([]*model.Product)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetRecommendationProducts indicates an expected call of GetRecommendationProducts.
func (mr *MockProductUsecaseInterfaceMockRecorder) GetRecommendationProducts(itemID, userID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecommendationProducts", reflect.TypeOf((*MockProductUsecaseInterface)(nil).GetRecommendationProducts), itemID, userID)
}

// GetSuggestions mocks base method.
func (m *MockProductUsecaseInterface) GetSuggestions(search string) ([]string, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "GetSuggestions", search)
        ret0, _ := ret[0].([]string)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// GetSuggestions indicates an expected call of GetSuggestions.
func (mr *MockProductUsecaseInterfaceMockRecorder) GetSuggestions(search interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSuggestions", reflect.TypeOf((*MockProductUsecaseInterface)(nil).GetSuggestions), search)
}

// InsertItemIntoFavorites mocks base method.
func (m *MockProductUsecaseInterface) InsertItemIntoFavorites(userID, itemID int) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "InsertItemIntoFavorites", userID, itemID)
        ret0, _ := ret[0].(error)
        return ret0
}

// InsertItemIntoFavorites indicates an expected call of InsertItemIntoFavorites.
func (mr *MockProductUsecaseInterfaceMockRecorder) InsertItemIntoFavorites(userID, itemID interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertItemIntoFavorites", reflect.TypeOf((*MockProductUsecaseInterface)(nil).InsertItemIntoFavorites), userID, itemID)
}

// MakeOrder mocks base method.
func (m *MockProductUsecaseInterface) MakeOrder(in *model.MakeOrder) (int, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "MakeOrder", in)
        ret0, _ := ret[0].(int)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// MakeOrder indicates an expected call of MakeOrder.
func (mr *MockProductUsecaseInterfaceMockRecorder) MakeOrder(in interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MakeOrder", reflect.TypeOf((*MockProductUsecaseInterface)(nil).MakeOrder), in)
}

// RecalculatePrices mocks base method.
func (m *MockProductUsecaseInterface) RecalculatePrices(userID int, promocode string) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "RecalculatePrices", userID, promocode)
        ret0, _ := ret[0].(error)
        return ret0
}

// RecalculatePrices indicates an expected call of RecalculatePrices.
func (mr *MockProductUsecaseInterfaceMockRecorder) RecalculatePrices(userID, promocode interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecalculatePrices", reflect.TypeOf((*MockProductUsecaseInterface)(nil).RecalculatePrices), userID, promocode)
}

// SetPromocode mocks base method.
func (m *MockProductUsecaseInterface) SetPromocode(userID int, promocode string) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "SetPromocode", userID, promocode)
        ret0, _ := ret[0].(error)
        return ret0
}

// SetPromocode indicates an expected call of SetPromocode.
func (mr *MockProductUsecaseInterfaceMockRecorder) SetPromocode(userID, promocode interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPromocode", reflect.TypeOf((*MockProductUsecaseInterface)(nil).SetPromocode), userID, promocode)
}

// UpdateOrder mocks base method.
func (m *MockProductUsecaseInterface) UpdateOrder(userID int, items *[]int) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "UpdateOrder", userID, items)
        ret0, _ := ret[0].(error)
        return ret0
}

// UpdateOrder indicates an expected call of UpdateOrder.
func (mr *MockProductUsecaseInterfaceMockRecorder) UpdateOrder(userID, items interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockProductUsecaseInterface)(nil).UpdateOrder), userID, items)
}