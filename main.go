package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

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

func main() {

	content, err := ioutil.ReadFile("test.json") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
	}
	jsonRes := PACS008AccEnq{}
	err = json.Unmarshal(content, &jsonRes)
	if err != nil {
		fmt.Println("Err")
	}
	res := new(pacs.Document00800108)
	resAdd := res.AddMessage()

	resGrpHeader := resAdd.AddGroupHeader()
	resGrpHeader.SetMsgId(jsonRes.Messageid)
	resGrpHeader.SetCreDtTm(jsonRes.Creationdatetime)
	resGrpHeader.SetNbOfTxs(jsonRes.Numbertransaction)

	resSttlntInfo := resGrpHeader.AddSttlmInf()
	resSttlntInfo.SetSttlmMtd(jsonRes.Settlementmethod)

	resCtf := resAdd.AddCreditTransferTransactionInformation()
	resPayment := resCtf.AddPmtId()
	resPayment.SetEndToEndId(jsonRes.Endtoendid)
	resPayment.SetTxId(jsonRes.Transactionid)

	resCtf.SetChrgBr(jsonRes.Chargebearer)
	resDbtr := resCtf.AddDbtrAgt()
	resDbtrFinInstnId := resDbtr.AddFinInstnId()
	resDbtrOthr := resDbtrFinInstnId.AddOthr()
	resDbtrOthr.SetIdentification(jsonRes.Debtorbankid)

	resCdtr := resCtf.AddCdtrAgt()
	resCdtrFinInstnId := resCdtr.AddFinInstnId()
	resCdtrOthr := resCdtrFinInstnId.AddOthr()
	resCdtrOthr.SetIdentification(jsonRes.Debtorbankid)
	resCtf.AddDbtr()
	resCtf.AddCdtr()

	resCtf.SetInterbankSettlementAmount(jsonRes.Interbanksettlementamount, jsonRes.Currencycode)
	a, _ := json.Marshal(res)
	fmt.Println(string(a))
}
