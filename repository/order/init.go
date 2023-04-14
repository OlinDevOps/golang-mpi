package order

import (
	"mpiolinprojek/infra"

	"github.com/sirupsen/logrus"
)

type DatabaseOrder struct {
	Order Order
}

func NewDatabaseOrder(db *infra.DatabaseList, logger *logrus.Logger) DatabaseOrder {
	return DatabaseOrder{
		Order: newOrder(db, logger),
	}
}
