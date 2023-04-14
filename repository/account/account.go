package account

import (
	"context"
	"database/sql"
	"fmt"
	"mpiolinprojek/domain/model/general/account"
	"mpiolinprojek/infra"

	"github.com/sirupsen/logrus"
)

type AccountConfig struct {
	db  *infra.DatabaseList
	log *logrus.Logger
}

func newAccount(db *infra.DatabaseList, logger *logrus.Logger) AccountConfig {
	return AccountConfig{
		db:  db,
		log: logger,
	}
}

type Account interface {
	UpdateStatusApotek(ctx context.Context, tx *sql.Tx, SIA string) error
	RegistrationRepo(ctx context.Context, tx *sql.Tx, data account.Registration) error
}

func (as AccountConfig) UpdateStatusApotek(ctx context.Context, tx *sql.Tx,SIA string) error {
	tcUpdateStatusTrx := `UPDATE master.dbo.outlet
SET  ATTRIBUTE_NUMBER12=? where ATTRIBUTE18=?;
	`

	query, args, err := as.db.Backend.Read.In(tcUpdateStatusTrx, 1, SIA)
	if err != nil {
		return err
	}

	query = as.db.Backend.Read.Rebind(query)
	res, err := as.db.Backend.Write.ExecContext(ctx, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return err
	}

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		err = fmt.Errorf("no rows inserted")
		return err
	}

	return nil
}

func (as AccountConfig) RegistrationRepo(ctx context.Context, tx *sql.Tx, data account.Registration) error {
	qInsertProduct := `INSERT INTO master.dbo.outlet(party_name,
alamat,
party_site_name,
short_code,
attribute1,
attribute3,
attribute4,
attribute5,
attribute6,
attribute12,
attribute13, 
attribute14,
attribute18,
attribute19,
attribute_number12) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`



	param := make([]interface{}, 0)


	param = append(param, data.PartyName)
	param = append(param, data.Alamat)
	param = append(param, data.PartySiteName)
	param = append(param, data.ShortCode)
	param = append(param, data.Attribute2)
	param = append(param, data.Attribute3)
	param = append(param, data.Attribute4)
	param = append(param, data.Attribute5)
	param = append(param, data.Attribute6)
	param = append(param, data.Attribute12)
	param = append(param, data.Attribute13)
	param = append(param, data.Attribute14)
	param = append(param, data.Attribute18)
	param = append(param, data.Attribute19)
	param = append(param, data.AttributeNumber12)

	query, args, err := as.db.Backend.Write.In(qInsertProduct, param...)
	if err != nil {
		return  err
	}

	query = as.db.Backend.Write.Rebind(query)

	var res *sql.Row
	if tx == nil {
		res = as.db.Backend.Write.QueryRow(ctx, query, args...)
	} else {
		res = tx.QueryRowContext(ctx, query, args...)
	}

	if err != nil {
		return  err
	}

	err = res.Err()
	if err != nil {
		return  err
	}

	return nil
}