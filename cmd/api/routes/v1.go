package routes

import (
	"mpiolinprojek/domain/model/general"
	"mpiolinprojek/handler/api"
	"net/http"

	"github.com/gorilla/mux"
)

func getV1(routerIntegrations, router, routerJWT *mux.Router, conf *general.AppService, handler api.Handler) {
	routerIntegrations.HandleFunc("/mpi/change-status", handler.Account.UpdateStatusCompany).Methods(http.MethodPost)
	routerIntegrations.HandleFunc("/mpi/registration", handler.Account.Registration).Methods(http.MethodPost)
	routerIntegrations.HandleFunc("/mpi/sales-order", handler.Order.CreateSalesOrder).Methods(http.MethodPost)
}
