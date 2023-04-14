package order

import (
	"context"
	"database/sql"
	"fmt"
	"mpiolinprojek/domain/model/general/order"
	"mpiolinprojek/infra"

	"github.com/sirupsen/logrus"
)

type OrderConfig struct {
	db  *infra.DatabaseList
	log *logrus.Logger
}

func newOrder(db *infra.DatabaseList, logger *logrus.Logger) OrderConfig {
	return OrderConfig{
		db:  db,
		log: logger,
	}
}

type Order interface {
	CreateSalesOrder(ctx context.Context, tx *sql.Tx, data order.DocSo) error
	CreateSalesOrderItem(ctx context.Context, tx *sql.Tx, data order.DocSoItem) error
	CreateSalesOrderItemPrice(ctx context.Context, tx *sql.Tx, data order.DocSoItemPrice) error
}

const (

	oQsalesOrder = `INSERT INTO MPI_PROD_NEW.dbo.DMS_SD_DocSo
	( iInternalId,
		iId,
		szWarehouseId,
		szEmployeeId, 
		szDocId, 
		dtmDoc, 
		szCustomerId, 
		szDescription, 
		bCash,
		szOrderTermId, 
		szOrderTypeId,
		bAdaSuratPesanan,
		szNoTOPK,
		szPOId,
		dtmPO,
		decOngkosKirim,
		szTaxId,
		szTaxCodeId,
		decTarifTax,
		decTotalTax,
		bFlagMinimumOrder,
		bFlagOverLimit,
		bFlagOverDue,
		bFlagSIA,
		bFlagSIPA,
		szBranchId,
		szCompanyId, 
		dtmCreated,
		szShippingMode,
		szPaymentTermId,
		dtmPOExpired,
		szPOSourceId,
		szShipBranchId,
		szStatus,
		intPrintedCount,
		szDocStatus,
		szUserCreatedId,
		szUserUpdatedId,
		dtmLastUpdated,
		dtmMobileTransaction,
		szMobileId,
		szParentDocId,
		szPromoDesc,
		bDOTanpaHarga,
		szSONetsuite,
		decTotalDPP,
		decTotalDiscount,
		decTotalPPN,
		decTotalAmount,
		bKonsinasi,
		bSettlement,
		szWarehouseKonsinasiId,
		dtmDelivery,
		bFlagEdSIPA,
		szSource,
		szRayonId
	)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`

	oQSalesOrderItem = `INSERT INTO MPI_PROD_NEW.dbo.DMS_SD_DocSoItem
	(iInternalId, 
		iId, 
		szDocId, 
		intItemNumber, 
		szProductId, 
		szOrderItemTypeId, 
		szTrnType, 
		decQty, 
		szUomId, 
		decQtyDelivered, 
		szHasCode, 
		szQtyNote, 
		szPriceId, 
		decQtyStock)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`

	oQSalesOrderItemPrice = `INSERT INTO MPI_PROD_NEW.dbo.DMS_SD_DocSoItemPrice
	(iInternalId, 
		iId, 
		szDocId, 
		intItemNumber, 
		intItemDetailNumber, 
		szPriceId, 
		decPrice, 
		decDiscount, 
		bTaxable, 
		decAmount, 
		decTax, 
		decDpp, 
		szTaxId, 
		decTaxRate, 
		decDiscPrinciple, 
		decDiscDistributor, 
		decBruto, 
		decDiscTotal, 
		decDiscTotalValue, 
		decDisc1, 
		decDisc1Value, 
		decDisc2, 
		decDisc2Value, 
		decDisc3, 
		decDisc3Value, 
		decDisc4, 
		decDisc4Value, 
		decDisc5, 
		decDisc5Value, 
		decDisc6, 
		decDisc6Value)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
	`

)

func (oc OrderConfig) CreateSalesOrder(ctx context.Context, tx *sql.Tx, data order.DocSo) error {
	param := make([]interface{}, 0)

	param = append(param, data.InternalID)
	param = append(param, data.IID)
	param = append(param, data.SzWarehousID)
	param = append(param, data.SzEmployeeID)
	param = append(param, data.SzDocID)
	param = append(param, data.DtmDoc)
	param = append(param, data.SzCustomerID)
	param = append(param, data.SzDescription)
	param = append(param, data.BCash)
	param = append(param, data.SzOrderTermID)
	param = append(param, data.SzOrderTypeId)
	param = append(param, data.BAdaSuratPesanan)
	param = append(param, data.SzNoTOPK)
	param = append(param, data.SzPoID)
	param = append(param, data.DtmPO)
	param = append(param, data.DecOngkosKirim)
	param = append(param, data.SzTaxID)
	param = append(param, data.SzTaxCodeID)
	param = append(param, data.DecTarifTax)
	param = append(param, data.DecTotalTax)
	param = append(param, data.BFlagMinimumOrder)
	param = append(param, data.BFlagOverLimit)
	param = append(param, data.BFlagOverDue)
	param = append(param, data.BFlagSIA)
	param = append(param, data.BFlagSIPA)
	param = append(param, data.SzBranchID)
	param = append(param, data.SzCompanyID)
	param = append(param, data.DtmCreated)
	param = append(param, data.SzShippingMode)
	param = append(param, data.SzPaymentTermID)
	param = append(param, data.DtmPOExpired)
	param = append(param, data.SzPOSourceID)
	param = append(param, data.SzShipBranchID)
	param = append(param, data.SzStatus)
	param = append(param, data.IntPrintedCount)
	param = append(param, data.SzDocStatus)
	param = append(param, data.SzUserCreatedID)
	param = append(param, data.SzUserUpdatedID)
	param = append(param, data.DmtLastUpdated)
	param = append(param, data.DtmMobileTransaction)
	param = append(param, data.SzMobileID)
	param = append(param, data.SzParentDocID)
	param = append(param, data.SzPromoDesc)
	param = append(param, data.BDOTanpaHarga)
	param = append(param, data.SzSONetSuite)
	param = append(param, data.DecTotalDPP)
	param = append(param, data.DecTotalDiscount)
	param = append(param, data.DecTotalPPN)
	param = append(param, data.DecTotalAmount)
	param = append(param, data.BKonsinasi)
	param = append(param, data.BSettlement)
	param = append(param, data.SzWarehouseKonsinasiID)
	param = append(param, data.DtmDelivery)
	param = append(param, data.BFlagEDSIPA)
	param = append(param, data.SzSource)
	param = append(param, data.SzRayonID)



		

	query, args, err := oc.db.Backend.Write.In(oQsalesOrder, param...)
	if err != nil {
		return  err
	}

	fmt.Println("Query:", query)
	fmt.Println("Args:", args)

	query = oc.db.Backend.Write.Rebind(query)

	var res *sql.Row
	if tx == nil {
		res = oc.db.Backend.Write.QueryRow(ctx, query, args...)
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

func (oc OrderConfig) CreateSalesOrderItem(ctx context.Context, tx *sql.Tx, data order.DocSoItem) error {
	param := make([]interface{}, 0)

	param = append(param, data.InternalID)
	param = append(param, data.IID)
	param = append(param, data.SzDocID)
	param = append(param, data.IntItemNumber)
	param = append(param, data.SzProductID)
	param = append(param, data.SzOrderItemTypeID)
	param = append(param, data.SzTrnType)
	param = append(param, data.DecQty)
	param = append(param, data.SzUomID)
	param = append(param, data.DecQtyDelivered)
	param = append(param, data.SzHasCode)
	param = append(param, data.SzQtyNote)
	param = append(param, data.SzPriceID)
	param = append(param, data.DecQtyStock)

	query, args, err := oc.db.Backend.Write.In(oQSalesOrderItem, param...)
	if err != nil {
		return  err
	}

	fmt.Println("Query:", query)
	fmt.Println("Args:", args)

	query = oc.db.Backend.Write.Rebind(query)

	var res *sql.Row
	if tx == nil {
		res = oc.db.Backend.Write.QueryRow(ctx, query, args...)
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

func (oc OrderConfig) CreateSalesOrderItemPrice(ctx context.Context, tx *sql.Tx, data order.DocSoItemPrice) error {
	param := make([]interface{}, 0)

	param = append(param, data.InternalID)
	param = append(param, data.IID)
	param = append(param, data.SzDocID)
	param = append(param, data.IntItemNumber)
	param = append(param, data.IntItemDetailNumber)
	param = append(param, data.SzPriceID)
	param = append(param, data.DecPrice)
	param = append(param, data.DecDiscount)
	param = append(param, data.BTaxAble)
	param = append(param, data.DecAmount)
	param = append(param, data.DecTax)
	param = append(param, data.DecDPP)
	param = append(param, data.SzTaxID)
	param = append(param, data.DecTaxRate)
	param = append(param, data.DecDiscountPrinciple)
	param = append(param, data.DecDiscountDistributor)
	param = append(param, data.DecBruto)
	param = append(param, data.DecDiscTotal)
	param = append(param, data.DecDiscTotalValue)
	param = append(param, data.DecDisc1)
	param = append(param, data.DecDisc1Value)
	param = append(param, data.DecDisc2)
	param = append(param, data.DecDisc2Value)
	param = append(param, data.DecDisc3)
	param = append(param, data.DecDisc3Value)
	param = append(param, data.DecDisc4)
	param = append(param, data.DecDisc4Value)
	param = append(param, data.DecDisc5)
	param = append(param, data.DecDisc5Value)
	param = append(param, data.DecDisc6)
	param = append(param, data.DecDisc6Value)

	query, args, err := oc.db.Backend.Write.In(oQSalesOrderItemPrice, param...)
	if err != nil {
		return  err
	}

	fmt.Println("Query:", query)
	fmt.Println("Args:", args)

	query = oc.db.Backend.Write.Rebind(query)

	var res *sql.Row
	if tx == nil {
		res = oc.db.Backend.Write.QueryRow(ctx, query, args...)
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