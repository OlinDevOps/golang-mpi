package api

import (
	"mpiolinprojek/domain/model/general"
	"mpiolinprojek/handler/api/authorization"
	"mpiolinprojek/service"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	Token       authorization.TokenHandler
	Public      authorization.PublicHandler
	Account AccountHandler
	Order OrderHandler
}

func NewHandler(sv service.Service, conf general.AppService, logger *logrus.Logger) Handler {
	return Handler{
		Token:       authorization.NewTokenHandler(conf, logger),
		Public:      authorization.NewPublicHandler(conf, logger),
		Account:	newAccountHandler(sv.Account.Account, conf, logger),
		Order: newOrderHandler(sv.Order.Order, conf, logger),
	}
}
