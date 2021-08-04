package main

type Header struct {
	TraceNum            string `json:"traceNum"`
	ProcessingCode      string `json:"pcode"`
	FromId              string `json:"fromId"`
	ToId                string `json:"toId"`
	BusinessMessageId   string `json:"businessMessageId"`
	MessageDefinitionId string `json:"messageDefinitionId"`
	BusinessServiceId   string `json:"businessServiceId,omitempty"`
	CreationDate        string `json:"creationDate"`
	CopyDuplicate       string `json:"copyDuplicate,omitempty"`
	PossibleDuplicate   string `json:"possibleDuplicate,omitempty"`
}

type PACS008AccEnq struct {
	TraceNum                  string `json:"traceNum"`
	ProcessingCode            string `json:"pcode"`
	FromId                    string `json:"fromId"`
	ToId                      string `json:"toId"`
	BusinessMessageId         string `json:"businessMessageId"`
	MessageDefinitionId       string `json:"messageDefinitionId"`
	BusinessServiceId         string `json:"businessServiceId,omitempty"`
	CreationDate              string `json:"creationDate"`
	CopyDuplicate             string `json:"copyDuplicate,omitempty"`
	PossibleDuplicate         string `json:"possibleDuplicate,omitempty"`
	Messageid                 string `json:"messageId,omitempty"`
	Creationdatetime          string `json:"creationDateTime,omitempty"`
	Numbertransaction         string `json:"NumberTransaction,omitempty"`
	Settlementmethod          string `json:"SettlementMethod,omitempty"`
	Endtoendid                string `json:"EndToEndID,omitempty"`
	Transactionid             string `json:"TransactionID,omitempty"`
	Interbanksettlementamount string `json:"InterBankSettlementAmount,omitempty"`
	Currencycode              string `json:"CurrencyCode,omitempty"`
	Chargebearer              string `json:"ChargeBearer,omitempty"`
	Debtorbankid              string `json:"DebtorBankID,omitempty"`
	Creditorbankid            string `json:"CreditorBankID,omitempty"`
	Customeraccountnumber     string `json:"CustomerAccountNumber,omitempty"`
}

type PACS028PayStatusReq struct {
	TraceNum            string `json:"traceNum"`
	ProcessingCode      string `json:"pcode"`
	FromId              string `json:"fromId"`
	ToId                string `json:"toId"`
	BusinessMessageId   string `json:"businessMessageId"`
	MessageDefinitionId string `json:"messageDefinitionId"`
	BusinessServiceId   string `json:"businessServiceId,omitempty"`
	CreationDate        string `json:"creationDate"`
	CopyDuplicate       string `json:"copyDuplicate,omitempty"`
	PossibleDuplicate   string `json:"possibleDuplicate,omitempty"`
	Messageid           string `json:"messageId,omitempty"`
	Creationdatetime    string `json:"creationDateTime,omitempty"`
	OrgnlEndToEndId     string `json:"OrgnlEndToEndId"`
}