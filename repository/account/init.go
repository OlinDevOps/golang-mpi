package account

import (
	"mpiolinprojek/infra"

	"github.com/sirupsen/logrus"
)

type DatabaseAccount struct {
	Account Account
}

func NewDatabaseAccount(db *infra.DatabaseList, logger *logrus.Logger) DatabaseAccount {
	return DatabaseAccount{
		Account: newAccount(db, logger),
	}
}
