package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"local/jul/bankiso/iso"
	"log"
	"net/http"

	"github.com/j03hanafi/bankiso/iso20022/head"
	"github.com/j03hanafi/bankiso/iso20022/pacs"
)

// type tempDocumentXML interface{}
type ChannelInput struct {
	BusMsg BusMsg `xml:"BusMsg" json:"BusMsg"`
}

type BusMsg struct {
	AppHdr   *head.BusinessApplicationHeaderV01 `xml:"AppHdr" json:"AppHdr"`
	Document json.RawMessage                    `xml:"Document" json:"Document"`
}
type BusMsgAfterUnmarshal struct {
	AppHdr   *head.BusinessApplicationHeaderV01 `xml:"AppHdr" json:"AppHdr"`
	Document *interface{}                       `xml:"Document" json:"Document"`
}

type BusMsgAfterUnmarshal2 struct {
	AppHdr   *head.BusinessApplicationHeaderV01 `xml:"AppHdr" json:"AppHdr"`
	Document *pacs.Document00200110             `xml:"Document" json:"Document"`
}

func main() {

	content, err := ioutil.ReadFile("test.json") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	jsonReq := PACS008AccEnq{}
	err = json.Unmarshal(content, &jsonReq)
	if err != nil {
		fmt.Println("Err")
	}

	bodyRequest := convJsonIso(jsonReq)

	request := hitBiller(bodyRequest)

	requestMessageJson := ChannelInput{}
	err = json.Unmarshal(request, &requestMessageJson)

	if err != nil {
		log.Fatal("json unmarshall err:", err)
	}

	response := convIsoJson(requestMessageJson)

	result, _ := json.Marshal(response)
	fmt.Println(string(result))

}

func convIsoJson(jsonRes ChannelInput) PACS002AccEnq {
	MsgDefIdn := string(*jsonRes.BusMsg.AppHdr.MessageDefinitionIdentifier)
	log.Println("MsgDefIdn:", MsgDefIdn)
	val, ok := iso.ISO20022Registry[MsgDefIdn]
	if !ok {
		log.Fatal("invalid ISO20022 code ", MsgDefIdn)
	}
	err := json.Unmarshal(jsonRes.BusMsg.Document, &val)
	if err != nil {
		log.Fatal("documentType unmarshall err:", err)
	}
	response := BusMsgAfterUnmarshal{jsonRes.BusMsg.AppHdr, &val}

	jsonResConv := PACS002AccEnq{}
	jsonResConv.MessageId = string(*response.AppHdr.BusinessMessageIdentifier)
	jsonResConv.CreationDateTime = string(*response.AppHdr.CreationDate)

	result, _ := json.Marshal(val)

	res := new(pacs.Document00200110)
	err = json.Unmarshal(result, &res)
	if err != nil {
		log.Fatal("documentType unmarshall err:", err)
	}

	jsonResConv.OriginalEndToEndId = string(*res.Message.TransactionInformationAndStatus[0].OrgnlEndToEndId)

	return jsonResConv

}

func convJsonIso(jsonRes PACS008AccEnq) []byte {

	head := new(head.BusinessApplicationHeaderV01)
	head.SetBusinessMessageIdentifier(jsonRes.BusinessMessageId)
	head.SetMessageDefinitionIdentifier(jsonRes.MessageDefinitionId)
	head.SetCreationDate(jsonRes.CreationDate)

	headFr := head.AddFrom().AddFinancialInstitutionIdentification().AddFinancialInstitutionIdentification().AddOther()
	headFr.SetIdentification(jsonRes.FromId)

	headTo := head.AddTo().AddFinancialInstitutionIdentification().AddFinancialInstitutionIdentification().AddOther()
	headTo.SetIdentification(jsonRes.ToId)

	res := new(pacs.Document00800108)
	resAdd := res.AddMessage()

	resGrpHeader := resAdd.AddGroupHeader()
	resGrpHeader.SetMsgId(jsonRes.Messageid)
	resGrpHeader.SetCreDtTm(jsonRes.Creationdatetime)
	resGrpHeader.SetNbOfTxs(jsonRes.Numbertransaction)

	resSttlntInfo := resGrpHeader.AddSttlmInf()
	resSttlntInfo.SetSttlmMtd(jsonRes.Settlementmethod)

	resCtf := resAdd.AddCreditTransferTransactionInformation()
	resCtf.AddDbtr()
	resCtf.AddCdtr()
	resCtf.SetChrgBr(jsonRes.Chargebearer)

	resPayment := resCtf.AddPmtId()
	resPayment.SetEndToEndId(jsonRes.Endtoendid)
	resPayment.SetTxId(jsonRes.Transactionid)

	resDbtr := resCtf.AddDbtrAgt()
	resDbtrFinInstnId := resDbtr.AddFinInstnId()
	resDbtrOthr := resDbtrFinInstnId.AddOthr()
	resDbtrOthr.SetIdentification(jsonRes.Debtorbankid)

	resCdtr := resCtf.AddCdtrAgt()
	resCdtrFinInstnId := resCdtr.AddFinInstnId()
	resCdtrOthr := resCdtrFinInstnId.AddOthr()
	resCdtrOthr.SetIdentification(jsonRes.Debtorbankid)

	resCtf.SetInterbankSettlementAmount(jsonRes.Interbanksettlementamount, jsonRes.Currencycode)

	a, _ := json.Marshal(res)

	msg := ChannelInput{}
	msg.BusMsg.AppHdr = head
	msg.BusMsg.Document = a

	b, _ := json.Marshal(msg)

	return b
}

func hitBiller(msg []byte) []byte {
	resp, err := http.Post("http://localhost:6066/biller",
		"application/json", bytes.NewBuffer(msg))

	if err != nil {
		log.Println(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	}
	return body
}
