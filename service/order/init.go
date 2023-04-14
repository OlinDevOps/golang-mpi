package order

import (
	"mpiolinprojek/domain/model/general"
	"mpiolinprojek/infra"
	"mpiolinprojek/repository"

	"github.com/sirupsen/logrus"
)

type ServiceOrder struct {
	Order OrderService
}

func NewOrderService(repo repository.Repo, conf general.AppService, dbList *infra.DatabaseList, logger *logrus.Logger) ServiceOrder {
	return ServiceOrder{
		Order: newOrderService(repo.Order, conf, dbList, logger),
	}
}
