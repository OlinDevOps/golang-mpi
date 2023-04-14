package order

import (
	"context"
	"fmt"
	"log"
	"mpiolinprojek/domain/model/general"
	"mpiolinprojek/domain/model/general/order"
	"mpiolinprojek/infra"
	ro "mpiolinprojek/repository/order"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type OrderService struct {
	dbOrder ro.DatabaseOrder
	conf          general.AppService
	dbConn        *infra.DatabaseList
	log           *logrus.Logger
	// redis *redis.Client
}

// redis *redis.Client add this on newTransactionService
func newOrderService(dbOrder ro.DatabaseOrder, conf general.AppService, dbConn *infra.DatabaseList, logger *logrus.Logger) OrderService {
	return OrderService{
		dbOrder: dbOrder,
		conf:          conf,
		dbConn:        dbConn,
		log:           logger,
	}
}

type Order interface {
	CreateSalesOrder(ctx context.Context, data order.CreateSalesOrder) (map[string]string, error)
}

func (os OrderService) CreateSalesOrder(ctx context.Context, data order.CreateSalesOrder) (map[string]string, error) {

	// Generate a new UUID
	id := uuid.New()
	// Convert the UUID to a string in the correct format for SQL Server
	idString := id.String()

	tx, err := os.dbConn.Backend.Write.Begin()
	if err != nil {
		return map[string]string{
			"en": "fail to create so, please retry later",
			"id": "sales order gagal, silahkan coba dilain waktu",
		}, err
	}

	dateString := data.SO.DTMPo
    layout := "2006-01-02 15:04:05"
    parsedTime, err := time.Parse(layout, dateString)
    if err != nil {
        fmt.Println("Error parsing time:", err)
    } else {
        fmt.Println("Parsed time:", parsedTime)
    }

	err = os.dbOrder.Order.CreateSalesOrder(ctx, tx, order.DocSo{
		IID: idString,
		SzDocID: data.SO.SzDocID,
		DtmDoc: time.Now().UTC(),
		SzCustomerID: data.SO.SzCustomerID,
		SzEmployeeID: data.SO.SzEmployeeID,
		SzOrderTypeId: "",
		BCash: int32(data.SO.BTunai),
		SzPaymentTermID: "", // 1 hari
		SzPoID: data.SO.SzPOId,
		DtmPO: parsedTime,
		DtmPOExpired: time.Now().UTC(),
		SzPOSourceID: "",
		SzShipBranchID: "", // work
		SzStatus: "OPEN",
		IntPrintedCount: 0,
		SzBranchID: data.WorkPlaceID,
		SzCompanyID: "",
		SzDocStatus: "",
		SzUserCreatedID: "",
		SzUserUpdatedID: "",
		DtmCreated: time.Now().UTC(),
		DmtLastUpdated: time.Now().UTC(),
		DtmMobileTransaction: time.Now().UTC(),
		SzMobileID: "",
		SzParentDocID: "",
		SzDescription: data.SO.SzDescription,
		SzPromoDesc: "",
		BDOTanpaHarga: 0,
		SzSONetSuite: "",
		SzWarehousID: "",
		DecTotalDPP: 0,
		DecTotalDiscount: 0,
		DecTotalPPN: 0,
		DecTotalAmount: 0,
		BAdaSuratPesanan: data.SO.BAdaSuratPesanan,
		SzOrderTermID: "",
		BKonsinasi: 0,
		BSettlement: 0,
		DecOngkosKirim: int64(data.SO.DecOngkosKirim),
		SzNoTOPK: data.SO.SzNoTOPK,
		SzWarehouseKonsinasiID: "",
		DtmDelivery: time.Now().UTC().Add(0),
		SzTaxID: "",
		DecTotalTax: 0,
		BFlagMinimumOrder: data.SO.BFlagMinimumOrder,
		BFlagOverLimit: data.SO.BFlagOverLimit,
		BFlagSIA: data.SO.BFlagSIA,
		BFlagEDSIPA: data.SO.BFlagEdSIPA,
		SzSource: "",
		SzTaxCodeID: data.SO.SzTaxCode,
		DecTarifTax: int64(data.SO.DecTarif),
		SzRayonID: "",
		BFlagSIPA: 0,
		BFlagOverDue: data.SO.BFlagOverDue,
		SzShippingMode: "",
		},)
	if err != nil {
		log.Println("ERROR | error create so", err.Error())
		return map[string]string{
			"en": "fail to create so, please retry later",
			"id": "sales order gagal, silahkan coba dilain waktu",
		}, err
	}

	for _, item := range data.SO.Item {
		err = os.dbOrder.Order.CreateSalesOrderItem(ctx, tx, order.DocSoItem{
			// InternalID: item.
			IID: idString,
			SzDocID: item.SzDocID,
			IntItemNumber: item.IntItemNumber,
			SzProductID: item.SzProductID,
			SzOrderItemTypeID: "",
			SzTrnType: "",
			DecQty: item.DecQty,
			SzUomID: item.SzUomID,
			DecQtyDelivered: 0,
			SzHasCode: "",
			SzQtyNote: "",
			SzPriceID: item.SzPriceID,
			DecQtyStock: 0,
		})
		if err != nil {
			log.Println("ERROR | error create so item", err.Error())
			return map[string]string{
				"en": "fail to create so, please retry later",
				"id": "sales order gagal, silahkan coba dilain waktu",
			}, err
		}

		err = os.dbOrder.Order.CreateSalesOrderItemPrice(ctx, tx, order.DocSoItemPrice{
			// InternalID
			IID: idString,
			SzDocID: item.SzDocID,
			IntItemNumber: item.IntItemNumber,
			IntItemDetailNumber: 0,
			SzPriceID: item.SzPriceID,
			DecPrice: item.DecPrice,
			DecDiscount: item.DecDiscount,
			BTaxAble: 1,
			DecAmount: item.DecAmount,
			DecTax: data.SO.DecTax,
			DecDPP: 0, //harga asal sebelum kena pajak / hna
			SzTaxID: "0",
			DecTaxRate: 0,
			DecDiscountPrinciple: item.DecDiscountPrinciple,
			DecDiscountDistributor: 0,
			DecBruto: 0,
			DecDiscTotal: 0,
			DecDiscTotalValue: 0,
			DecDisc1: 0,
			DecDisc1Value: 0,
			DecDisc2: 0,
			DecDisc2Value: 0,
			DecDisc3: 0,
			DecDisc3Value: 0,
			DecDisc4: 0,
			DecDisc4Value: 0,
			DecDisc5: 0,
			DecDisc5Value: 0,
			DecDisc6: 0,
			DecDisc6Value: 0,
		})	
	}
	if err != nil {
			log.Println("ERROR | error create so item price", err.Error())
			return map[string]string{
				"en": "fail to create so, please retry later",
				"id": "sales order gagal, silahkan coba dilain waktu",
			}, err
		}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return map[string]string{
			"en": "fail to create so, please retry later",
			"id": "sales order gagal, silahkan coba dilain waktu",
		}, err
	}

	return map[string]string{	
		"en" : "successfully created sales order",
		"id" : "berhasil membuat order penjualan",
	}, nil
}
