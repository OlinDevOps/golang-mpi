package account

import (
	"context"
	"errors"
	"log"
	"mpiolinprojek/domain/model/general"
	"mpiolinprojek/domain/model/general/account"
	"mpiolinprojek/infra"
	ra "mpiolinprojek/repository/account"
	"time"

	"github.com/sirupsen/logrus"
)

type AccountService struct {
	dbAccount ra.DatabaseAccount
	conf          general.AppService
	dbConn        *infra.DatabaseList
	log           *logrus.Logger
	// redis *redis.Client
}

// redis *redis.Client add this on newTransactionService
func newAccountService(dbAccount ra.DatabaseAccount, conf general.AppService, dbConn *infra.DatabaseList, logger *logrus.Logger) AccountService {
	return AccountService{
		dbAccount: dbAccount,
		conf:          conf,
		dbConn:        dbConn,
		log:           logger,
	}
}

type Account interface {
	UpdateStatusApotek(ctx context.Context, SIA string) (map[string]string, error)
	RegistrationApotek(ctx context.Context, data account.Registration) (map[string]string, error)
}

func (as AccountService) UpdateStatusApotek(ctx context.Context, SIA string) (map[string]string, error) {
	if SIA == "" {
		return map[string]string{
			"en" : "Parameter Required",
			"id" : "Parameter Required",
		}, errors.New("Parameter Required")
	}

	err := as.dbAccount.Account.UpdateStatusApotek(ctx, nil, SIA)
	if err != nil {
		log.Println(err.Error())
		return map[string]string{
			"en" : "Failed to change status company",
			"id" : "Maaf ada kesalahan dalam memperbarui status perusahaan",
		}, err
	}

	return map[string]string{
		"en" : "successfully change status company",
		"id" : "Berhasil memperbarui status perusahaan",
	}, nil
}

func (as AccountService) RegistrationApotek(ctx context.Context, data account.Registration) (map[string]string, error) {

	done := make(chan struct{})
	defer close(done)

	errChan := make(chan error)
	defer close(errChan)

	go func() {
		err := as.dbAccount.Account.RegistrationRepo(ctx, nil, account.Registration{
			PartyName: data.PartyName,
			Alamat: data.Alamat,
			PartySiteName: data.PartySiteName,
			ShortCode: data.ShortCode,
			Attribute2: data.Attribute2,
			Attribute3: data.Attribute3,
			Attribute4: data.Attribute4,
			Attribute5: data.Attribute5,
			Attribute6: data.Attribute6,
			Attribute12: data.Attribute12,
			Attribute13: data.Attribute13,
			Attribute14: data.Attribute14,
			Attribute18: data.Attribute18,
			Attribute19: data.Attribute19,
			AttributeNumber12: data.AttributeNumber12,
		})
		if err != nil {
			errChan <- err
			return
		}
		done <- struct{}{}
	}()

	select {
	case <-done:
		return map[string]string{
			"en": "successfully change status company",
			"id": "Berhasil memperbarui status perusahaan",
		}, nil
	case err := <-errChan:
		log.Println(err.Error())
		return map[string]string{
			"en": "Failed to change status company",
			"id": "Maaf ada kesalahan dalam memperbarui status perusahaan",
		}, errors.New("error")
	case <-time.After(time.Second * 30):
		return map[string]string{
			"en": "Failed to change status company",
			"id": "Maaf ada kesalahan dalam memperbarui status perusahaan",
		}, errors.New("timeout error")
	}
}
