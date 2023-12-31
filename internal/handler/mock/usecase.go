// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go
//
// Generated by this command:
//
//	mockgen -source=usecase.go -destination=mock/usecase.go -package=mock
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	domain "github.com/k-akari/golang-rest-api-sample/internal/domain"
	gomock "go.uber.org/mock/gomock"
)

// MockcompanyUsecase is a mock of companyUsecase interface.
type MockcompanyUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockcompanyUsecaseMockRecorder
}

// MockcompanyUsecaseMockRecorder is the mock recorder for MockcompanyUsecase.
type MockcompanyUsecaseMockRecorder struct {
	mock *MockcompanyUsecase
}

// NewMockcompanyUsecase creates a new mock instance.
func NewMockcompanyUsecase(ctrl *gomock.Controller) *MockcompanyUsecase {
	mock := &MockcompanyUsecase{ctrl: ctrl}
	mock.recorder = &MockcompanyUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockcompanyUsecase) EXPECT() *MockcompanyUsecaseMockRecorder {
	return m.recorder
}

// CreateCompany mocks base method.
func (m *MockcompanyUsecase) CreateCompany(ctx context.Context, company *domain.Company) (domain.CompanyID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCompany", ctx, company)
	ret0, _ := ret[0].(domain.CompanyID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCompany indicates an expected call of CreateCompany.
func (mr *MockcompanyUsecaseMockRecorder) CreateCompany(ctx, company any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompany", reflect.TypeOf((*MockcompanyUsecase)(nil).CreateCompany), ctx, company)
}

// GetCompanyByID mocks base method.
func (m *MockcompanyUsecase) GetCompanyByID(ctx context.Context, companyID domain.CompanyID) (*domain.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompanyByID", ctx, companyID)
	ret0, _ := ret[0].(*domain.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompanyByID indicates an expected call of GetCompanyByID.
func (mr *MockcompanyUsecaseMockRecorder) GetCompanyByID(ctx, companyID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompanyByID", reflect.TypeOf((*MockcompanyUsecase)(nil).GetCompanyByID), ctx, companyID)
}

// MockclientUsecase is a mock of clientUsecase interface.
type MockclientUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockclientUsecaseMockRecorder
}

// MockclientUsecaseMockRecorder is the mock recorder for MockclientUsecase.
type MockclientUsecaseMockRecorder struct {
	mock *MockclientUsecase
}

// NewMockclientUsecase creates a new mock instance.
func NewMockclientUsecase(ctrl *gomock.Controller) *MockclientUsecase {
	mock := &MockclientUsecase{ctrl: ctrl}
	mock.recorder = &MockclientUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockclientUsecase) EXPECT() *MockclientUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockclientUsecase) Create(ctx context.Context, client *domain.Client) (domain.ClientID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, client)
	ret0, _ := ret[0].(domain.ClientID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockclientUsecaseMockRecorder) Create(ctx, client any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockclientUsecase)(nil).Create), ctx, client)
}

// GetByID mocks base method.
func (m *MockclientUsecase) GetByID(ctx context.Context, clientID domain.ClientID) (*domain.Client, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, clientID)
	ret0, _ := ret[0].(*domain.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockclientUsecaseMockRecorder) GetByID(ctx, clientID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockclientUsecase)(nil).GetByID), ctx, clientID)
}

// MockinvoiceUsecase is a mock of invoiceUsecase interface.
type MockinvoiceUsecase struct {
	ctrl     *gomock.Controller
	recorder *MockinvoiceUsecaseMockRecorder
}

// MockinvoiceUsecaseMockRecorder is the mock recorder for MockinvoiceUsecase.
type MockinvoiceUsecaseMockRecorder struct {
	mock *MockinvoiceUsecase
}

// NewMockinvoiceUsecase creates a new mock instance.
func NewMockinvoiceUsecase(ctrl *gomock.Controller) *MockinvoiceUsecase {
	mock := &MockinvoiceUsecase{ctrl: ctrl}
	mock.recorder = &MockinvoiceUsecaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockinvoiceUsecase) EXPECT() *MockinvoiceUsecaseMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockinvoiceUsecase) Create(ctx context.Context, invoice *domain.Invoice) (domain.InvoiceID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, invoice)
	ret0, _ := ret[0].(domain.InvoiceID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockinvoiceUsecaseMockRecorder) Create(ctx, invoice any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockinvoiceUsecase)(nil).Create), ctx, invoice)
}

// ListByPaymentDueDateBetween mocks base method.
func (m *MockinvoiceUsecase) ListByPaymentDueDateBetween(ctx context.Context, coid domain.CompanyID, from, to *time.Time) ([]*domain.Invoice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListByPaymentDueDateBetween", ctx, coid, from, to)
	ret0, _ := ret[0].([]*domain.Invoice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListByPaymentDueDateBetween indicates an expected call of ListByPaymentDueDateBetween.
func (mr *MockinvoiceUsecaseMockRecorder) ListByPaymentDueDateBetween(ctx, coid, from, to any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListByPaymentDueDateBetween", reflect.TypeOf((*MockinvoiceUsecase)(nil).ListByPaymentDueDateBetween), ctx, coid, from, to)
}
