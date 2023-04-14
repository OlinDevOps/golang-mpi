package order

import "time"

type CreateSalesOrder struct {
	InternalID  int64      
	IID         string     
	WorkPlaceID string     `json:"szWorkplaceId"`
	EmployeeID  string     `json:"szEmployeeId"`
	DeviceID    string     `json:"szDeviceId"`
	SO          SalesOrder `json:"so"`
}

type SalesOrder struct {
	Item                []Item    `json:"item"`
	SzDocID string `json:"szDocId"`
	DtmDoc string `json:"dtmDoc"`
	SzCustomerID string `json:"szCustomerId"`
	SzEmployeeID string `json:"szEmployeeId"`
	DecAmount float64 `json:"decAmount"`
	BUploaded int64 `json:"bUploaded"`
	DecDiscount float64 `json:"decDiscount"`
	DecAmountEstimasi int64 `json:"decAmountEstimasi"`
	DecDiscountEstimasi int64 `json:"decDiscountEstimasi"`
	SzDescription string `json:"szDescription"`
	DzDocCallId string `json:"szDocCallId"`
	SzShipToId string `json:"szShipToId"`
	BTunai int64 `json:"bTunai"`
	SzTermId string `json:"szTermId"`
	SzOrderType string `json:"szOrderType"`
	BAdaSuratPesanan int64 `json:"bAdaSuratPesanan"`
	SzNoTOPK string `json:"szNoTOPK"`
	SzPOId string `json:"szPOId"`
	DTMPo string `json:"dtmPO"`
	DecOngkosKirim int64 `json:"decOngkosKirim"`
	SzTaxId string `json:"szTaxId"`
	SzTaxCode string `json:"szTaxCode"`
	DecTarif int64 `json:"decTarif"`
	DecTax int64 `json:"decTax"`
	BFlagMinimumOrder int64 `json:"bFlagMinimumOrder"`
	BFlagOverLimit int64 `json:"bFlagOverLimit"`
	BFlagOverDue int64 `json:"bFlagOverDue"`
	BFlagSIA int64 `json:"bFlagSIA"`
	BFlagEdSIPA int64 `json:"bFlagEdSIPA"`
}

type Item struct {
	ItemDiscount         []interface{} `json:"itemDiscount"`
	ItemBonus            []interface{} `json:"itemBonus"`
	SzDocID string `json:"szDocId"`
	IntItemNumber int64 `json:"intItemNumber"`
	SzProductID string `json:"szProductId"`
	DecQty int64 `json:"decQty"`
	DecPrice int64 `json:"decPrice"`
	DecAmount int64 `json:"decAmount"`
	DecDiscount int64 `json:"decDiscount"`
	DecDiscountPrinciple int64 `json:"decDiscountPrinciple"`
	DecDiscountOFF int64 `json:"decDiscountOff"`
	DecDiscountPrincipleOFF int64 `json:"decDiscountPrincipleOff"`
	DecDiscountCNTarik int64 `json:"decDiscountCNTarik"`
	DecDiscountAstek int64 `json:"decDiscountAstek"`
	SzUomID string `json:"szUomId"`
	SzWarehousID string `json:"szWarehouseId"`
	SzReasonID string `json:"szReasonId"`
	SzItemType string `json:"szItemType"`
	DecAmountEstimasi int64 `json:"decAmountEstimasi"`
	DecQtyEstimasi int64 `json:"decQtyEstimasi"`
	DecDiscountEstimasi int64 `json:"decDiscountEstimasi"`
	SzPriceID string `json:"szPriceId"`
	SzPromoID string `json:"szPromoId"`
}

type DocSo struct {
	InternalID int64 `json:"sz_internal_id"`
	IID string 
	Item []Item `json:"item"`
	SzDocID string `json:"sz_doc_id"`
	DtmDoc time.Time `json:"dtm_doc"`
	SzCustomerID string `json:"sz_customer_id"`
	SzEmployeeID string `json:"sz_employee_id"`
	SzOrderTypeId string `json:"sz_order_type_id"`
	BCash int32 `json:"bcash"`
	SzPaymentTermID string `json:"sz_payment_term_id"`
	SzPoID string `json:"sz_po_id"`
	DtmPO time.Time `json:"dtm_po_time"`
	DtmPOExpired time.Time `json:"dtm_po_expired"`
	SzPOSourceID string `json:"sz_po_source_id"`
	SzShipBranchID string `json:"sz_ship_branch_id"`
	SzStatus string `json:"sz_status"`
	IntPrintedCount int32 `json:"int_printed_count"`
	SzBranchID string `json:"sz_branch_id"`
	SzCompanyID string `json:"sz_company_id"`
	SzDocStatus string `json:"sz_doc_status"`
	SzUserCreatedID string `json:"sz_user_created_id"`
	SzUserUpdatedID string `json:"sz_user_updated_id"`
	DtmCreated time.Time `json:"dtm_created"`
	DmtLastUpdated time.Time `json:"dmt_last_updated"`
	DtmMobileTransaction time.Time `json:"dtm_mobile_transaction"`
	SzMobileID string `json:"sz_mobile_id"`
	SzParentDocID string `json:"sz_parent_doc_id"`
	SzDescription string `json:"sz_description"`
	SzPromoDesc string `json:"sz_promo_desc"`
	BDOTanpaHarga int32 `json:"bdotanpa_harga"`
	SzSONetSuite string `json:"sz_so_net_suite"`
	SzWarehousID string `json:"sz_warehous_id"`
	DecTotalDPP int64 `json:"dec_total_dpp"`
	DecTotalDiscount int64 `json:"dec_total_discount"`
	DecTotalPPN int64 `json:"dec_total_ppn"`
	DecTotalAmount int64 `json:"dec_total_amount"`
	BAdaSuratPesanan int64 `json:"bsurat_pesanan"`
	SzOrderTermID string `json:"sz_order_term_id"`
	BKonsinasi int64 `json:"bkonsinasi"`
	BSettlement int64 `json:"bsettlement"`
	DecOngkosKirim int64 `json:"dec_ongkos_kirim"`
	SzNoTOPK string `json:"sz_no_to_pk"`
	SzWarehouseKonsinasiID string `json:"sz_Warehouse_konsinasi_id"`
	DtmDelivery time.Time `json:"dtm_delivery"`
	SzTaxID string `json:"sz_tax_id"`
	DecTotalTax int64 `json:"dec_total_tax"`
	BFlagMinimumOrder int64 `json:"b_flag_minimum_order"`
	BFlagOverLimit int64 `json:"b_flag_over_limit"`
	BFlagSIA int64 `json:"b_flag_sia"`
	BFlagEDSIPA int64 `json:"b_flag_expired_date_sipa"`
	SzSource string `json:"sz_source"`
	SzTaxCodeID string `json:"sz_tax_code_id"`
	DecTarifTax int64 `json:"dec_tarif_tax"`
	SzRayonID string `json:"sz_rayon_id"`
	BFlagSIPA int64 `json:"b_flag_sipa"`
	BFlagOverDue int64 `json:"b_flag_overdue"`
	SzShippingMode string `json:"sz_shipping_mode"`
}

type DocSoItem struct {
	InternalID int64 `json:"iinternal_id"`
	IID string `json:"iid"`
	SzDocID string `json:"sz_doc_id"`
	IntItemNumber int64 `json:"int_item_number"`
	SzProductID string `json:"sz_product_id"`
	SzOrderItemTypeID string `json:"sz_order_item_type_id"`
	SzTrnType string `json:"sz_trn_type"`
	DecQty int64 `json:"dec_qty"`
	SzUomID string `json:"sz_uom_id"`
	DecQtyDelivered int64 `json:"dec_qty_delivered"`
	SzHasCode string `json:"sz_has_code"`
	SzQtyNote string `json:"sz_qty_note"`
	SzPriceID string `json:"sz_price_id"`
	DecQtyStock int64 `json:"dec_qty_stock"` 
}

type DocSoItemPrice struct {
	InternalID int64 `json:"internal_id"`
	IID string `json:"iid"`
	SzDocID string `json:"sz_doc_id"`
	IntItemNumber int64 `json:"int_item_number"`
	IntItemDetailNumber int64 `json:"int_item_detail_number"`
	SzPriceID string `json:"sz_price_id"`
	DecPrice int64 `json:"dec_price"`
	DecDiscount int64 `json:"dec_discount"`
	BTaxAble int64 `json:"btax_able"`
	DecAmount int64 `json:"dec_amount"`
	DecTax int64 `json:"dec_tax"`
	DecDPP int64 `json:"dec_dpp"`
	SzTaxID string `json:"sz_tax_id"`
	DecTaxRate int64 `json:"dec_tax_rate"`
	DecDiscountPrinciple int64 `json:"dec_discount_principle"`
	DecDiscountDistributor int64 `json:"dec_discount_distributor"`
	DecBruto int64 `json:"dec_bruto"`
	DecDiscTotal int64 `json:"dec_disc_total"`
	DecDiscTotalValue int64 `json:"dec_disc_total_value"`
	DecDisc1 int64 `json:"dec_disc_1"`
	DecDisc1Value int64 `json:"dec_disc_1_value"`
	DecDisc2 int64 `json:"dec_disc_2"`
	DecDisc2Value int64 `json:"dec_disc_2_value"`
	DecDisc3 int64 `json:"dec_disc_3"`
	DecDisc3Value int64 `json:"dec_disc_3_value"`
	DecDisc4 int64 `json:"dec_disc_4"`
	DecDisc4Value int64 `json:"dec_disc_4_value"`
	DecDisc5 int64 `json:"dec_disc_5"`
	DecDisc5Value int64 `json:"dec_disc_5_value"`
	DecDisc6 int64 `json:"dec_disc_6"`
	DecDisc6Value int64 `json:"dec_disc_6_value"`
}

type DocSoItemPromo struct {
	InternalID int64 `json:"internal_id"`
	IID string `json:"iid"`
	SzDocID string `json:"sz_doc_id"`
	IntItemNumber int64 `json:"int_item_number"`
	IntItemDetailNumber int64 `json:"int_item_detail_number"`
	SzPromoID string `json:"sz_promo_id"`
	BBonusProduct int64 `json:"b_bonus_product"`
	SzDiscountType string `json:"sz_discount_type"`
	DecDiscount int64 `json:"dec_discount"`
	SzProductBonusID string `json:"sz_product_bonus_id"`
	DecQtyBonus int64 `json:"dec_qty_bonus"`
	DecDiscountValue int64 `json:"dec_discount_value"`
	DecQtyBonusValue int64 `json:"dec_qty_bonus_value"`
	IntPriority int64 `json:"int_priority"`
	DecDiscountPrinciple int64 `json:"dec_discount_principle"`
	DecDiscountPrincipleValue int64 `json:"dec_discount_principle_value"`
	DecDiscountOFF int64 `json:"dec_discount_off"`
	DecDiscountOffValue int64 `json:"dec_discount_off_value"`
	DecDiscountPrincipleOFF int64 `json:"dec_discount_principle_off"`
	DecDiscountPrincipleOffValue int64 `json:"dec_discount_principle_off_value"`
	DecDiscountCNTarik int64 `json:"dec_discount_cntarik"`
	DecDiscountCNTarikValue int64 `json:"dec_discount_cntarik_value"`
	DecDiscountAstek int64 `json:"dec_discount_astek"`
	DecDiscountAstekValue int64 `json:"dec_discount_astek_value"`
}