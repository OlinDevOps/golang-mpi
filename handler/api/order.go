package api

import (
	"encoding/json"
	"io/ioutil"
	cg "mpiolinprojek/domain/constants/general"
	"mpiolinprojek/domain/model/general"
	"mpiolinprojek/domain/model/general/order"
	"mpiolinprojek/domain/utils"
	"net/http"

	so "mpiolinprojek/service/order"

	"github.com/sirupsen/logrus"
)

type OrderHandler struct {
	order so.Order
	conf        general.AppService
	log         *logrus.Logger
}

func newOrderHandler(order so.Order, conf general.AppService, logger *logrus.Logger) OrderHandler {
	return OrderHandler{
		order: order,
		conf:  conf,
		log:   logger,
	}
}


func (oh OrderHandler) CreateSalesOrder(res http.ResponseWriter, req *http.Request) {
	respData := &utils.ResponseDataV2{
		Status: cg.Fail,
	}

	var param order.CreateSalesOrder

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		respData.Message = map[string]string{
			"en": cg.HandlerErrorRequestDataEmpty,
			"id": cg.HandlerErrorRequestDataEmptyID,
		}
		utils.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(reqBody, &param)
	if err != nil {
		respData.Message = map[string]string{
			"en": cg.HandlerErrorRequestDataNotValid,
			"id": cg.HandlerErrorRequestDataNotValidID,
		}
		utils.WriteResponse(res, respData, http.StatusBadRequest)
		return
	}

	message, err := oh.order.CreateSalesOrder(req.Context(), param)
	if err != nil {
		respData.Message = message
		utils.WriteResponse(res, respData, http.StatusInternalServerError)
		return
	}

	respData = &utils.ResponseDataV2{
		Status:  cg.Success,
		Message: message,
	}

	utils.WriteResponse(res, respData, http.StatusOK)
	return
}