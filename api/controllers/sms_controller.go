package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/yuriadams/bluestreak/api/models"
	"github.com/yuriadams/gocomtele"
)

type (
	SMSController struct {
	}
)

type FormValues map[string]string

func NewSMSController() *SMSController {
	return &SMSController{}
}

func (smsc *SMSController) GetSMSHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	sms := models.SMS{"#23", "85999350334", "EEEI MAHH", time.Now()}
	json, _ := json.Marshal(sms)
	RespondWithJSON(w, string(json))
}

func (smsc *SMSController) SendSMSHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rawBody := make([]byte, r.ContentLength, r.ContentLength)
	r.Body.Read(rawBody)

	var fv FormValues
	json.Unmarshal(rawBody, &fv)

	log.Printf(fmt.Sprintf("%s\n", fv))

	switch fv["smsDispatcher"] {
	case "comtele":
		comtele := comtele.NewComteleClient(fv["authToken"])
		from := fv["from"]
		to := fv["to"]
		message := fv["message"]
		comtele.SendSMS(from, to, message)
		RespondWithJSON(w, fmt.Sprintf("message: %s sended with success to %s from %s\n", message, to, from))
	}
}
