package api

import (
	"encoding/json"
	"io/ioutil"
	"mpiolinprojek/domain/model/general"
	"mpiolinprojek/domain/model/general/account"
	"mpiolinprojek/domain/utils"
	"net/http"

	cg "mpiolinprojek/domain/constants/general"
	sa "mpiolinprojek/service/account"

	"github.com/sirupsen/logrus"
)

type AccountHandler struct {
	account sa.Account
	conf        general.AppService
	log         *logrus.Logger
}

func newAccountHandler(account sa.Account, conf general.AppService, logger *logrus.Logger) AccountHandler {
	return AccountHandler{
		account: account,
		conf:        conf,
		log:         logger,
	}
}

func (th AccountHandler) UpdateStatusCompany(res http.ResponseWriter, req *http.Request) {
	respData := &utils.ResponseDataV2{
		Status: cg.Fail,
	}

	SIAParam := req.URL.Query().Get("sia")

	message, err := th.account.UpdateStatusApotek(req.Context(), SIAParam)
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

func (th AccountHandler) Registration(res http.ResponseWriter, req *http.Request) {
	respData := &utils.ResponseDataV2{
		Status: cg.Fail,
	}

	var param account.Registration

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

	message, err := th.account.RegistrationApotek(req.Context(), param)
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