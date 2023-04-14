package repository

import (
	"mpiolinprojek/infra"
	ra "mpiolinprojek/repository/account"
	ro "mpiolinprojek/repository/order"

	"github.com/sirupsen/logrus"
)

type Repo struct {
	Account ra.DatabaseAccount
	Order ro.DatabaseOrder
}

func NewRepo(database *infra.DatabaseList, logger *logrus.Logger) Repo {
	return Repo{
		Account: ra.NewDatabaseAccount(database, logger),
		Order: ro.NewDatabaseOrder(database, logger),
	}
}
