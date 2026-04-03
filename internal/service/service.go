package service

import (
	"salon/internal/repo"
	"salon/internal/router"
)

func New(repo *repo.Repo) *Service {
	return &Service{repo: repo}
}

type Service struct {
	repo *repo.Repo
}

func (s *Service) ListOrders(params router.ListOrdersParams) ([]*router.Order, error) {
	var data []*router.Order

	return data, nil
}

func (s *Service) CreateOrder() error {
	return nil
}

func (s *Service) GetOrderSummary(params router.GetOrderSummaryParams) (*router.OrderSummary, error) {
	var data *router.OrderSummary

	return data, nil
}

func (s *Service) GetOrder(id router.IdParam) (*router.Order, error) {
	var data *router.Order

	return data, nil
}

func (s *Service) ListPaymentMethods(params router.ListPaymentMethodsParams) ([]*router.PaymentMethod, error) {
	var data []*router.PaymentMethod

	return data, nil
}

func (s *Service) CreatePaymentMethod() error {
	return nil
}

func (s *Service) DeletePaymentMethod(id router.IdParam) error {
	return nil
}

func (s *Service) GetPaymentMethod(id router.IdParam) (*router.PaymentMethod, error) {
	var data *router.PaymentMethod

	return data, nil
}

func (s *Service) UpdatePaymentMethod(id router.IdParam) error {
	return nil
}

func (s *Service) DownloadOrderReport(params router.DownloadOrderReportParams) error {
	return nil
}

func (s *Service) ListServiceCategories(params router.ListServiceCategoriesParams) ([]*router.ServiceCategory, error) {
	var data []*router.ServiceCategory

	return data, nil
}

func (s *Service) CreateServiceCategory() error {
	return nil
}

func (s *Service) DeleteServiceCategory(id router.IdParam) error {
	return nil
}

func (s *Service) GetServiceCategory(id router.IdParam) (*router.ServiceCategory, error) {
	var data *router.ServiceCategory

	return data, nil
}

func (s *Service) UpdateServiceCategory(id router.IdParam) error {
	return nil
}

func (s *Service) ListServices(params router.ListServicesParams) ([]*router.Service, error) {
	var data []*router.Service

	return data, nil
}

func (s *Service) CreateService() error {
	return nil
}

func (s *Service) DeleteService(id router.IdParam) error {
	return nil
}

func (s *Service) GetService(id router.IdParam) (*router.Service, error) {
	var data *router.Service

	return data, nil
}

func (s *Service) UpdateService(id router.IdParam) error {
	return nil
}

func (s *Service) ListWorkers(params router.ListWorkersParams) ([]*router.Worker, error) {
	var data []*router.Worker

	return data, nil
}

func (s *Service) CreateWorker() error {
	return nil
}

func (s *Service) DeleteWorker(id router.IdParam) error {
	return nil
}

func (s *Service) GetWorker(id router.IdParam) (*router.Worker, error) {
	var data *router.Worker

	return data, nil
}

func (s *Service) UpdateWorker(id router.IdParam) error {
	return nil
}

func (s *Service) GetWorkerCommissionSummary(id router.IdParam, params router.GetWorkerCommissionSummaryParams) (*router.WorkerCommissionSummary, error) {
	var data *router.WorkerCommissionSummary

	return data, nil
}
