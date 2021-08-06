package main

type kafkaMessage struct {
	Head    string `json:"head"`
	Content []byte `json:"content"`
}

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

type AccountEnquiryReq struct {
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
	MessageId                 string `json:"messageId,omitempty"`
	CreationDateTime          string `json:"creationDateTime,omitempty"`
	NumberTransaction         string `json:"numberTransaction,omitempty"`
	SettlementMethod          string `json:"settlementMethod,omitempty"`
	EndToEndId                string `json:"endToEndID,omitempty"`
	TransactionId             string `json:"transactionID,omitempty"`
	InterBankSettlementAmount string `json:"interBankSettlementAmount,omitempty"`
	CurrencyCode              string `json:"currencyCode,omitempty"`
	ChargeBearer              string `json:"chargeBearer,omitempty"`
	DebtorBankId              string `json:"debtorBankID,omitempty"`
	CreditorBankId            string `json:"creditorBankID,omitempty"`
	CustomerAccountNumber     string `json:"customerAccountNumber,omitempty"`
}

type AccountEnquiryResp struct {
	TraceNum               string `json:"traceNum"`
	ProcessingCode         string `json:"pcode"`
	FromId                 string `json:"fromId"`
	ToId                   string `json:"toId"`
	BusinessMessageId      string `json:"businessMessageId"`
	MessageDefinitionId    string `json:"messageDefinitionId"`
	BusinessServiceId      string `json:"businessServiceId,omitempty"`
	CreationDate           string `json:"creationDate"`
	CopyDuplicate          string `json:"copyDuplicate,omitempty"`
	PossibleDuplicate      string `json:"possibleDuplicate,omitempty"`
	MessageId              string `json:"messageId,omitempty"`
	CreationDateTime       string `json:"creationDateTime,omitempty"`
	OriginalEndToEndId     string `json:"originalEndtoEndId,omitempty"`
	OriginalTransactionId  string `json:"originalTransactionId,omitempty"`
	TransactionStatus      string `json:"transactionStatus,omitempty"`
	ReasonCode             string `json:"reasonCode,omitempty"`
	CreditorName           string `json:"creditorName,omitempty"`
	CreditorAccountId      string `json:"creditorAccountId,omitempty"`
	CreditorAccountType    string `json:"creditorAccountType,omitempty"`
	CreditorType           string `json:"creditorType,omitempty"`
	CreditorNationalId     string `json:"creditorBankID,omitempty"`
	CreditorResidentStatus string `json:"creditorResidentStatus,omitempty"`
	CreditorTownName       string `json:"creditorTownName,omitempty`
}
type PaymentStatusReq struct {
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
	MessageId           string `json:"messageId,omitempty"`
	CreationDateTime    string `json:"creationDateTime,omitempty"`
	OrgnlEndToEndId     string `json:"OrgnlEndToEndId"`
}

type PaymentStatusResp struct {
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
	OriginalMessageId         string `json:"originalMessageId,omitempty"`
	OriginalMessageNameId     string `json:"originalMessageNameId",omitempty`
	OriginalCreationdatetime  string `json:"originalCreationDateTime,omitempty"`
	OriginalEndToEndId        string `json:"originalEndToEndID,omitempty"`
	OriginalTransactionId     string `json:"originalTransactionID,omitempty"`
	TransactionStatus         string `json:"transactionStatus,omitempty"`
	ReasonCode                string `json:"reasonCode,omitempty"`
	AdditionalInfo            string `json:"additionalInfo",omitempty`
	InterBankSettlementData   string `json:"interBankSettlementDate",omitempty`
	DebtorName                string `json:"debtorName,omitempty"`
	DebtorAccountId           string `json:"debtorAccountId,omitempty`
	DebtorAccountType         string `json:"debtorAccountType,omitempty"`
	OriginalDebtorFiId        string `json:"originalDebtorFiId,omitempty"`
	CreditorFiId              string `json:"creditorFiId,omitempty"`
	CreditorName              string `json:"creditorName,omitempty"`
	CreditorAccountId         string `json:"creditorAccountId,omitempty"`
	CreditorAccountType       string `json:"creditorAccountType,omitempty"`
	DebtorType                string `json:"debtorType,omitempty"`
	DebtorNationalId          string `json:"debtorNationalId,omitempty"`
	DebtorResidentStatus      string `json:"debtorResidentStatus,omitempty"`
	DebtorTownName            string `json:"debtorTownName,omitempty"`
	CreditorType              string `json:"creditorType,omitempty"`
	CreditorNationalId        string `json:"creditorNationalId,omitempty"`
	CreditorResidentStatus    string `json:"creditorResidentStatus,omitempty"`
	CreditorTownName          string `json:"creditorTownName,omitempty"`
	DebtorSettlementAccount   string `json:"debtorSettlementAccount,omitempty"`
	CreditorSettlementAccount string `json:"creditorSettlementAccount,omitempty"`
}

type AliasRegistrationInquiryReq struct {
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
	Id                  string `json:"id,omitempty"`
	Registrationid      string `json:"registrationId,omitempty"`
	Type                string `json:"type,omitempty"`
	Value               string `json:"value,omitempty"`
}

type AliasRegistrationInquiryResp struct {
	TraceNum                 string `json:"traceNum"`
	ProcessingCode           string `json:"pcode"`
	FromId                   string `json:"fromId"`
	ToId                     string `json:"toId"`
	BusinessMessageId        string `json:"businessMessageId"`
	MessageDefinitionId      string `json:"messageDefinitionId"`
	BusinessServiceId        string `json:"businessServiceId,omitempty"`
	CreationDate             string `json:"creationDate"`
	CopyDuplicate            string `json:"copyDuplicate,omitempty"`
	PossibleDuplicate        string `json:"possibleDuplicate,omitempty"`
	Messageid                string `json:"messageId,omitempty"`
	Creationdatetime         string `json:"creationDateTime,omitempty"`
	Id                       string `json:"id,omitempty"`
	Originalmessageid        string `json:"originalMessageId,omitempty"`
	Originalcreationdatetime string `json:"originalCreationDateTime,omitempty"`
	Proxyresponsestatus      string `json:"proxyResponseStatus,omitempty"`
	Proprietary              string `json:"proprietary,omitempty"`
	Registrationid           string `json:"registrationId,omitempty"`
	Displayname              string `json:"displayName,omitempty"`
	Participantid            string `json:"participantId,omitempty"`
	Proxytype                string `json:"proxyType,omitempty"`
	Proxyvalue               string `json:"proxyValue,omitempty"`
	Status                   string `json:"status,omitempty"`
	Agentid                  string `json:"agentId,omitempty"`
	Accountid                string `json:"accountId,omitempty"`
	Accountproprietary       string `json:"accountProprietary,omitempty"`
	Accountname              string `json:"accountName,omitempty"`
	Customertype             string `json:"customerType,omitempty"`
	Customerid               string `json:"customerId,omitempty"`
	Residentstatus           string `json:"residentStatus,omitempty"`
	Townname                 string `json:"townName,omitempty"`
}

type AliasResolutionReq struct {
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
	Id                  string `json:"id,omitempty"`
	Proxytype           string `json:"proxyType,omitempty"`
	Proxyvalue          string `json:"proxyValue,omitempty"`
}

type AliasResolutionResp struct {
	TraceNum                 string `json:"traceNum"`
	ProcessingCode           string `json:"pcode"`
	FromId                   string `json:"fromId"`
	ToId                     string `json:"toId"`
	BusinessMessageId        string `json:"businessMessageId"`
	MessageDefinitionId      string `json:"messageDefinitionId"`
	BusinessServiceId        string `json:"businessServiceId,omitempty"`
	CreationDate             string `json:"creationDate"`
	CopyDuplicate            string `json:"copyDuplicate,omitempty"`
	PossibleDuplicate        string `json:"possibleDuplicate,omitempty"`
	Messageid                string `json:"messageId,omitempty"`
	Creationdatetime         string `json:"creationDateTime,omitempty"`
	Id                       string `json:"id,omitempty"`
	Originalmessageid        string `json:"originalMessageId,omitempty"`
	Originalmessagenameid    string `json:"originalMessageNameId,omitempty"`
	Originalcreationdatetime string `json:"originalCreationDateTime,omitempty"`
	Originalid               string `json:"originalId,omitempty"`
	Type                     string `json:"type,omitempty"`
	Value                    string `json:"value,omitempty"`
	Proxyresponsestatus      string `json:"proxyResponseStatus,omitempty"`
	Proprietary              string `json:"proprietary,omitempty"`
	Proxytype                string `json:"proxyType,omitempty"`
	Proxyvalue               string `json:"proxyValue,omitempty"`
	Registrationid           string `json:"registrationId,omitempty"`
	Displayname              string `json:"displayName,omitempty"`
	Agentid                  string `json:"agentId,omitempty"`
	Accountid                string `json:"accountId,omitempty"`
	Accountproprietary       string `json:"accountProprietary,omitempty"`
	Accountname              string `json:"accountName,omitempty"`
	Customertype             string `json:"customerType,omitempty"`
	Customerid               string `json:"customerId,omitempty"`
	Residentstatus           string `json:"residentStatus,omitempty"`
	Townname                 string `json:"townName,omitempty"`
}

type AliasNotification struct {
	TraceNum               string `json:"traceNum"`
	ProcessingCode         string `json:"pcode"`
	FromId                 string `json:"fromId"`
	ToId                   string `json:"toId"`
	BusinessMessageId      string `json:"businessMessageId"`
	MessageDefinitionId    string `json:"messageDefinitionId"`
	BusinessServiceId      string `json:"businessServiceId,omitempty"`
	CreationDate           string `json:"creationDate"`
	CopyDuplicate          string `json:"copyDuplicate,omitempty"`
	PossibleDuplicate      string `json:"possibleDuplicate,omitempty"`
	Messageid              string `json:"messageId,omitempty"`
	Creationdatetime       string `json:"creationDateTime,omitempty"`
	Id                     string `json:"id,omitempty"`
	Originalid             string `json:"originalId,omitempty"`
	Type                   string `json:"type,omitempty"`
	Value                  string `json:"value,omitempty"`
	OriginalRegistrationId string `json:"originalRegistrationId,omitempty"`
	DisplayName            string `json:"displayName,omitempty"`
	BankId                 string `json:"bankId,omitempty"`
	AccountId              string `json:"accountId,omitempty"`
	AccountProprietary     string `json:"accountProprietary,omitempty"`
	AccountName            string `json:"accountName,omitempty"`
	CustomerType           string `json:"customerType,omitempty"`
	CustomerId             string `json:"customerId,omitempty"`
	ResidentStatus         string `json:"residentStatus,omitempty"`
	TownName               string `json:"townName,omitempty"`
	NewRegistrationId      string `json:"newRegistrationId,omitempty"`
	NewDisplayName         string `json:"newDisplayName,omitempty"`
	NewBankId              string `json:"newBankId,omitempty"`
	NewAccountId           string `json:"newAccountId,omitempty"`
	NewAccountProprietary  string `json:"newAccountProprietary,omitempty"`
	NewAccountName         string `json:"newAccountName,omitempty"`
	NewCustomerType        string `json:"newCustomerType,omitempty"`
	NewCustomerId          string `json:"newCustomerId,omitempty"`
	NewResidentStatus      string `json:"newResidentStatus,omitempty"`
	NewTownName            string `json:"newTownName,omitempty"`
}

type AliasManagementReq struct {
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
	Id                  string `json:"id,omitempty"`
	Registrationtype    string `json:"registrationType,omitempty"`
	Proxytype           string `json:"proxyType,omitempty"`
	Proxyvalue          string `json:"proxyValue,omitempty"`
	Registrationid      string `json:"registrationId,omitempty"`
	Displayname         string `json:"displayName,omitempty"`
	Agentid             string `json:"agentId,omitempty"`
	Accountid           string `json:"accountId,omitempty"`
	Proprietary         string `json:"proprietary,omitempty"`
	Accountname         string `json:"accountName,omitempty"`
	Secondaryidtype     string `json:"secondaryIdType,omitempty"`
	Secondaryidval      string `json:"secondaryIdVal,omitempty"`
	Registrationstatus  string `json:"registrationStatus,omitempty"`
	Customertype        string `json:"customerType,omitempty"`
	Customerid          string `json:"customerId,omitempty"`
	Residentstatus      string `json:"residentStatus,omitempty"`
	Townname            string `json:"townName,omitempty"`
}

type AliasManagementRes struct {
	TraceNum                 string `json:"traceNum"`
	ProcessingCode           string `json:"pcode"`
	FromId                   string `json:"fromId"`
	ToId                     string `json:"toId"`
	BusinessMessageId        string `json:"businessMessageId"`
	MessageDefinitionId      string `json:"messageDefinitionId"`
	BusinessServiceId        string `json:"businessServiceId,omitempty"`
	CreationDate             string `json:"creationDate"`
	CopyDuplicate            string `json:"copyDuplicate,omitempty"`
	PossibleDuplicate        string `json:"possibleDuplicate,omitempty"`
	Messageid                string `json:"messageId,omitempty"`
	Creationdatetime         string `json:"creationDateTime,omitempty"`
	Id                       string `json:"id,omitempty"`
	Originalmessageid        string `json:"originalMessageId,omitempty"`
	Originalmessagenameid    string `json:"originalMessageNameId,omitempty"`
	Originalcreationdatetime string `json:"originalCreationDateTime,omitempty"`
	Proxyresponsestatus      string `json:"proxyResponseStatus,omitempty"`
	Proprietary              string `json:"proprietary,omitempty"`
	Originalregistrationtype string `json:"originalRegistrationType,omitempty"`
	Registrationid           string `json:"registrationId,omitempty"`
	Agentid                  string `json:"agentId,omitempty"`
	Customertype             string `json:"customerType,omitempty"`
	Customerid               string `json:"customerId,omitempty"`
	Residentstatus           string `json:"residentStatus,omitempty"`
	Townname                 string `json:"townName,omitempty"`
}
