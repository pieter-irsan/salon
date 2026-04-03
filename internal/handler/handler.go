package handler

import (
	"net/http"
	"salon/internal/router"
	"salon/internal/service"
)

func New(service *service.Service) *handler {
	return &handler{service: service}
}

type handler struct {
	service *service.Service
}

func (h *handler) ListOrders(w http.ResponseWriter, r *http.Request, params router.ListOrdersParams)

func (h *handler) CreateOrder(w http.ResponseWriter, r *http.Request)

func (h *handler) GetOrderSummary(w http.ResponseWriter, r *http.Request, params router.GetOrderSummaryParams)

func (h *handler) GetOrder(w http.ResponseWriter, r *http.Request, id router.IdParam)

func (h *handler) ListPaymentMethods(w http.ResponseWriter, r *http.Request, params router.ListPaymentMethodsParams)

func (h *handler) CreatePaymentMethod(w http.ResponseWriter, r *http.Request)

func (h *handler) DeletePaymentMethod(w http.ResponseWriter, r *http.Request, id router.IdParam)

func (h *handler) GetPaymentMethod(w http.ResponseWriter, r *http.Request, id router.IdParam)

func (h *handler) UpdatePaymentMethod(w http.ResponseWriter, r *http.Request, id router.IdParam)

func (h *handler) DownloadOrderReport(w http.ResponseWriter, r *http.Request, params router.DownloadOrderReportParams)

func (h *handler) ListServiceCategories(w http.ResponseWriter, r *http.Request, params router.ListServiceCategoriesParams)

func (h *handler) CreateServiceCategory(w http.ResponseWriter, r *http.Request)

func (h *handler) DeleteServiceCategory(w http.ResponseWriter, r *http.Request, id router.IdParam)

func (h *handler) GetServiceCategory(w http.ResponseWriter, r *http.Request, id router.IdParam)

func (h *handler) UpdateServiceCategory(w http.ResponseWriter, r *http.Request, id router.IdParam)

func (h *handler) ListServices(w http.ResponseWriter, r *http.Request, params router.ListServicesParams)

func (h *handler) CreateService(w http.ResponseWriter, r *http.Request)

func (h *handler) DeleteService(w http.ResponseWriter, r *http.Request, id router.IdParam)

func (h *handler) GetService(w http.ResponseWriter, r *http.Request, id router.IdParam)

func (h *handler) UpdateService(w http.ResponseWriter, r *http.Request, id router.IdParam)

func (h *handler) ListWorkers(w http.ResponseWriter, r *http.Request, params router.ListWorkersParams)

func (h *handler) CreateWorker(w http.ResponseWriter, r *http.Request)

func (h *handler) DeleteWorker(w http.ResponseWriter, r *http.Request, id router.IdParam)

func (h *handler) GetWorker(w http.ResponseWriter, r *http.Request, id router.IdParam)

func (h *handler) UpdateWorker(w http.ResponseWriter, r *http.Request, id router.IdParam)

func (h *handler) GetWorkerCommissionSummary(w http.ResponseWriter, r *http.Request, id router.IdParam, params router.GetWorkerCommissionSummaryParams)
