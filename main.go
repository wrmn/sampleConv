package main

import (
	"bytes"
	"encoding/json"
	"errors"
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

type inMsg struct {
	MsgDefId string `json:"messageDefinitionId"`
	TraceNum string `json:"traceNum"`
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

	msgDefId, traceNum := getMsgDefId(content)

	bodyRequest, err := convJsonIso(content, msgDefId)
	if err != nil {
		log.Fatal("conv error:", err)
	}

	request := hitBiller(bodyRequest)

	response := ChannelInput{}
	err = json.Unmarshal(request, &response)

	if err != nil {
		log.Fatal("json unmarshall err:", err)
	}

	responseJson, err := convIsoJson(response, traceNum)
	if err != nil {
		log.Fatal("Conv error:", err)
	}

	fmt.Println(string(responseJson))

}

func convIsoJson(jsonRes ChannelInput, traceNum string) ([]byte, error) {
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
	var result []byte
	err = nil
	switch MsgDefIdn {
	case "pacs.002.001.10":
		jsonResConv := convertAccEnqIso(response, traceNum)
		result, err = json.Marshal(jsonResConv)
	default:
		err = errors.New("undefined message definition identifier")
	}
	return result, err
}

func convJsonIso(content []byte, msgDefId string) ([]byte, error) {
	var result []byte
	var err error
	switch msgDefId {
	case "pacs.008.001.08":
		jsonReq := AccountEnquiryReq{}
		err := json.Unmarshal(content, &jsonReq)
		if err != nil {
			fmt.Println("Err")
		}
		result = convertAccEnqJson(jsonReq)
		err = nil
	default:
		err = errors.New("undefined message definition identifier")
	}
	return result, err
}

// headerMsg *head.BusinessApplicationHeaderV01, documentMsg interface{}
func convertAccEnqIso(msg BusMsgAfterUnmarshal, traceNum string) AccountEnquiryResp {
	jsonResConv := AccountEnquiryResp{}
	bodyM := *msg.Document
	bodyMsg := bodyM.(*pacs.Document00200110)
	headerMsg := *msg.AppHdr
	jsonResConv.TraceNum = traceNum
	jsonResConv.FromId = string(*headerMsg.From.FinancialInstitutionIdentification.FinancialInstitutionIdentification.Other.Identification)
	jsonResConv.ToId = string(*headerMsg.To.FinancialInstitutionIdentification.FinancialInstitutionIdentification.Other.Identification)
	jsonResConv.BusinessMessageId = string(*headerMsg.BusinessMessageIdentifier)
	jsonResConv.MessageDefinitionId = string(*headerMsg.MessageDefinitionIdentifier)
	// jsonResConv.BusinessServiceId = string(*headerMsg.BusinessService)
	jsonResConv.CreationDate = string(*headerMsg.CreationDate)
	jsonResConv.MessageId = string(*bodyMsg.Message.GroupHeader.MsgId)
	jsonResConv.CreationDateTime = string(*bodyMsg.Message.GroupHeader.CreDtTm)
	jsonResConv.OriginalEndToEndId = string(*bodyMsg.Message.TransactionInformationAndStatus[0].OrgnlEndToEndId)
	jsonResConv.OriginalTransactionId = string(*bodyMsg.Message.TransactionInformationAndStatus[0].OrgnlTxId)
	jsonResConv.TransactionStatus = string(*bodyMsg.Message.TransactionInformationAndStatus[0].TxSts)
	jsonResConv.ReasonCode = string(*bodyMsg.Message.TransactionInformationAndStatus[0].StsRsnInf[0].Rsn.Proprietary)
	jsonResConv.CreditorName = string(*bodyMsg.Message.TransactionInformationAndStatus[0].OrgnlTxRef.Cdtr.Pty.Nm)
	jsonResConv.CreditorAccountId = string(*bodyMsg.Message.TransactionInformationAndStatus[0].OrgnlTxRef.CdtrAcct.Id.Other.Identification)
	jsonResConv.CreditorAccountType = string(*bodyMsg.Message.TransactionInformationAndStatus[0].OrgnlTxRef.CdtrAcct.Tp.Proprietary)
	jsonResConv.CreditorType = string(*bodyMsg.Message.TransactionInformationAndStatus[0].SplmtryData[0].Envlp.Cdtr.Tp)
	jsonResConv.CreditorNationalId = string(*bodyMsg.Message.TransactionInformationAndStatus[0].SplmtryData[0].Envlp.Cdtr.Id)
	jsonResConv.CreditorResidentStatus = string(*bodyMsg.Message.TransactionInformationAndStatus[0].SplmtryData[0].Envlp.Cdtr.RsdntSts)
	jsonResConv.CreditorTownName = string(*bodyMsg.Message.TransactionInformationAndStatus[0].SplmtryData[0].Envlp.Cdtr.TwnNm)

	return jsonResConv
}

func convertAccEnqJson(jsonRes AccountEnquiryReq) []byte {
	head := new(head.BusinessApplicationHeaderV01)
	head.SetBusinessMessageIdentifier(jsonRes.BusinessMessageId)
	head.SetMessageDefinitionIdentifier(jsonRes.MessageDefinitionId)
	head.SetCreationDate(jsonRes.CreationDate)
	head.SetBusinessService(jsonRes.BusinessServiceId)
	head.SetCopyDuplicate(jsonRes.CopyDuplicate)
	head.SetPossibleDuplicate(jsonRes.PossibleDuplicate)

	headFr := head.AddFrom().AddFinancialInstitutionIdentification().AddFinancialInstitutionIdentification().AddOther()
	headFr.SetIdentification(jsonRes.FromId)

	headTo := head.AddTo().AddFinancialInstitutionIdentification().AddFinancialInstitutionIdentification().AddOther()
	headTo.SetIdentification(jsonRes.ToId)

	res := new(pacs.Document00800108)
	resAdd := res.AddMessage()

	resGrpHeader := resAdd.AddGroupHeader()
	resGrpHeader.SetMsgId(jsonRes.MessageId)
	resGrpHeader.SetCreDtTm(jsonRes.CreationDateTime)
	resGrpHeader.SetNbOfTxs(jsonRes.NumberTransaction)

	resSttlntInfo := resGrpHeader.AddSttlmInf()
	resSttlntInfo.SetSttlmMtd(jsonRes.SettlementMethod)

	resCtf := resAdd.AddCreditTransferTransactionInformation()
	resCtf.AddDbtr()
	resCtf.AddCdtr()
	resCtf.SetChrgBr(jsonRes.ChargeBearer)

	resPayment := resCtf.AddPmtId()
	resPayment.SetEndToEndId(jsonRes.EndToEndId)
	resPayment.SetTxId(jsonRes.TransactionId)

	resDbtr := resCtf.AddDbtrAgt()
	resDbtrFinInstnId := resDbtr.AddFinInstnId()
	resDbtrOthr := resDbtrFinInstnId.AddOthr()
	resDbtrOthr.SetIdentification(jsonRes.DebtorBankId)

	resCdtr := resCtf.AddCdtrAgt()
	resCdtrFinInstnId := resCdtr.AddFinInstnId()
	resCdtrOthr := resCdtrFinInstnId.AddOthr()
	resCdtrOthr.SetIdentification(jsonRes.DebtorBankId)

	resCtf.SetInterbankSettlementAmount(jsonRes.InterBankSettlementAmount, jsonRes.CurrencyCode)

	body, _ := json.Marshal(res)

	msg := ChannelInput{}
	msg.BusMsg.AppHdr = head
	msg.BusMsg.Document = body

	fullMessage, _ := json.Marshal(msg)

	return fullMessage
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
		log.Println(err.Error())
	}

	return body
}

func getMsgDefId(content []byte) (string, string) {
	res := inMsg{}
	json.Unmarshal(content, &res)
	return res.MsgDefId, res.TraceNum
}
