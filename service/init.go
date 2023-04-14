package service

import (
	"mpiolinprojek/domain/model/general"
	"mpiolinprojek/infra"
	"mpiolinprojek/repository"
	sa "mpiolinprojek/service/account"
	so "mpiolinprojek/service/order"

	"github.com/sirupsen/logrus"
)
type Service struct {
	Account        sa.ServiceAccount
	Order so.ServiceOrder
}

func NewService(repo repository.Repo, conf general.AppService, dbList *infra.DatabaseList, logger *logrus.Logger) Service {
	return Service{
		Account: sa.NewAccountService(repo, conf, dbList, logger),
		Order: so.NewOrderService(repo, conf, dbList, logger),
	}
}
