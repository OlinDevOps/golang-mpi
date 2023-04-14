package account

import (
	"mpiolinprojek/domain/model/general"
	"mpiolinprojek/infra"
	"mpiolinprojek/repository"

	"github.com/sirupsen/logrus"
)

type ServiceAccount struct {
	Account AccountService
}

func NewAccountService(repo repository.Repo, conf general.AppService, dbList *infra.DatabaseList, logger *logrus.Logger) ServiceAccount {
	return ServiceAccount{
		Account: newAccountService(repo.Account, conf, dbList, logger),
	}
}
