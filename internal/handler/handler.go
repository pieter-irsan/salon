package handler

import (
	"net/http"
	"salon/internal/router"
	"salon/internal/service"
)

type Handler struct {
	service *service.Service
}

func New(service *service.Service) *Handler {
	return &Handler{service: service}
}

// POST /auth/login
func (h *Handler) Login(w http.ResponseWriter, r *http.Request)

// POST /auth/refresh
func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request)

// GET /orders
func (h *Handler) ListOrders(w http.ResponseWriter, r *http.Request, params router.ListOrdersParams)

// POST /orders
func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request)

// GET /orders/summary
func (h *Handler) GetOrderSummary(w http.ResponseWriter, r *http.Request, params router.GetOrderSummaryParams)

// GET /orders/{id}
func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request, id router.IdParam)

// GET /payment-methods
func (h *Handler) ListPaymentMethods(w http.ResponseWriter, r *http.Request, params router.ListPaymentMethodsParams)

// POST /payment-methods
func (h *Handler) CreatePaymentMethod(w http.ResponseWriter, r *http.Request)

// DELETE /payment-methods/{id}
func (h *Handler) DeletePaymentMethod(w http.ResponseWriter, r *http.Request, id router.IdParam)

// GET /payment-methods/{id}
func (h *Handler) GetPaymentMethod(w http.ResponseWriter, r *http.Request, id router.IdParam)

// PUT /payment-methods/{id}
func (h *Handler) UpdatePaymentMethod(w http.ResponseWriter, r *http.Request, id router.IdParam)

// GET /reports/orders
func (h *Handler) DownloadOrderReport(w http.ResponseWriter, r *http.Request, params router.DownloadOrderReportParams)

// GET /service-categories
func (h *Handler) ListServiceCategories(w http.ResponseWriter, r *http.Request, params router.ListServiceCategoriesParams)

// POST /service-categories
func (h *Handler) CreateServiceCategory(w http.ResponseWriter, r *http.Request)

// DELETE /service-categories/{id}
func (h *Handler) DeleteServiceCategory(w http.ResponseWriter, r *http.Request, id router.IdParam)

// GET /service-categories/{id}
func (h *Handler) GetServiceCategory(w http.ResponseWriter, r *http.Request, id router.IdParam)

// PUT /service-categories/{id}
func (h *Handler) UpdateServiceCategory(w http.ResponseWriter, r *http.Request, id router.IdParam)

// GET /services
func (h *Handler) ListServices(w http.ResponseWriter, r *http.Request, params router.ListServicesParams)

// POST /services
func (h *Handler) CreateService(w http.ResponseWriter, r *http.Request)

// DELETE /services/{id}
func (h *Handler) DeleteService(w http.ResponseWriter, r *http.Request, id router.IdParam)

// GET /services/{id}
func (h *Handler) GetService(w http.ResponseWriter, r *http.Request, id router.IdParam)

// PUT /services/{id}
func (h *Handler) UpdateService(w http.ResponseWriter, r *http.Request, id router.IdParam)

// GET /workers
func (h *Handler) ListWorkers(w http.ResponseWriter, r *http.Request, params router.ListWorkersParams)

// POST /workers
func (h *Handler) CreateWorker(w http.ResponseWriter, r *http.Request)

// DELETE /workers/{id}
func (h *Handler) DeleteWorkersId(w http.ResponseWriter, r *http.Request, id router.IdParam)

// GET /workers/{id}
func (h *Handler) GetWorker(w http.ResponseWriter, r *http.Request, id router.IdParam)

// PUT /workers/{id}
func (h *Handler) UpdateWorker(w http.ResponseWriter, r *http.Request, id router.IdParam)

// GET /workers/{id}/commissions/summary
func (h *Handler) GetWorkerCommissionSummary(w http.ResponseWriter, r *http.Request, id router.IdParam, params router.GetWorkerCommissionSummaryParams)
