package model

type JNTRate struct {
	Code    int    `json:code`
	Data    string `json:data`
	NewData []Data
	Desc    string `json:desc`
	Success bool   `json:success`
}

type Data struct {
	StandardFee   string `json:standardFee`
	CbFee         string `json:cbFee`
	ZkcusFee      string `json: zkcusFee`
	AfterZkcusFee string `json: afterzkcusFee`
	InsuranceFee  string `json: insuranceFee`
	YTotalFee     string `json: ytotalFee`
	TaxFee        string `json: taxFee`
	ServiceFees   string `json: serviceFees`
	TotalFee      string `json: totalFee`
}

type AreaAddress struct {
	Province 		string `json:Province`
	City 			string `json:City`
	CountyArea 		string `json:CountyArea`
	DestinationCode string `json:DestinationCode`
	SixWordCode 	string `json:SixWordCode`
}