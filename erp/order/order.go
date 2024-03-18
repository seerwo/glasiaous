package order

import (
	"encoding/json"
	"fmt"
	"github.com/seerwo/glasiaous/erp/context"
	"github.com/seerwo/glasiaous/util"
)

const (
	ORDER_SO_SEARCH                  = "/order/so/search"                  //
	ORDER_SO_SEARCH_DETAIL           = "/order/so/search/detail"          //
	ORDER_SO_DELETE                	 = "/order/so/delete"                //
	ORDER_SO_APPROVE                 = "/order/so/approve"                 //
	ORDER_SO_SAVE             		 = "/order/so/save"             //
	ORDER_SOSUCCEED_SEARCH			 = "/order/sosucceed/search"		//
	ORDER_SO_CLOSE					 = "/order/so/close"
	ORDER_SO_CLOSE_CANCEL			 = "/order/so/close/cancel"
	ORDER_PO_SEARCH					 = "/order/po/search"
	ORDER_PO_SEARCH_DETAIL			 = "/order/po/search/detail"
	ORDER_PO_SAVE					 = "/order/po/save"
	ORDER_PO_APPROVE				 = "/order/po/approve"
	ORDER_PO_DELETE					 = "/order/po/delete"
	ORDER_POSUCCEED_SEARCH			 = "/order/posucceed/search"
	ORDER_POSUCCEED_SAVE			 = "/order/posucceed/save"
	ORDER_PO_CLOSE					 = "/order/po/close"
	ORDER_PO_CLOSE_CANCEL			 = "/order/po/close/cancel"
	ORDER_RESERVE_SEARCH			 = "/order/reserve/search"
	ORDER_RESERVE_RESERVEALL		 = "/order/reserve/reserveALL"
	ORDER_RESERVE_RESERVELND		 = "/order/reserve/reserveLnd"
)



//Order struct
type Order struct {
	*context.Context
}

//NewMenu instance
func NewOrder(context *context.Context) *Order {
	order := new(Order)
	order.Context = context
	return order
}

type SoSearchReq struct {
	FromInpDate string `json:"FromInpDate"`
	ToInpDate string `json:"toInpDate"`
	StatusH []string `json:"statusH"`
	PostpadColor string `json:"postpadColor"`
	ApproveStatus string `json:"approveStatus"`
	SlipOutputType string `json:"slipOutputType"`
	CorrespType string `json:"correspType"`
	ForConditionList []string `json:"forConditionList"`
	ReserveStatusList []string `json:"reserveStatusList"`
	DetailStatusList []string `json:"detail_StatusList"`
	ArrangeTypeList []string `json:"arrangeTypeList"`
	FromTaxAmountFc string `json:"fromTaxAmount_Fc"`
	ToTaxAmountFc string `json:"toTaxAmount_Fc"`
	FromTaxAmountSc string `json:"fromTaxAmount_Sc"`
	ToTaxAmountSc string `json:"toTaxAmount_Sc"`
	FromTotalAmountFc string `json:"fromTotalAmount_Fc"`
	TotTotalAmountFc string `json:"totTotalAmount_Fc"`
	FromTotalAmountSc string `json:"fromTotalAmount_Sc"`
	ToTotalAmountSc string `json:"toTotalAmount_Sc"`
	ImportFlagList []string `json:"importFlagList"`
	FromQty string `json:"fromQty"`
	ToQty string `json:"toQty"`
	FromPackingUnitQty string `json:"fromPackingUnitQty"`
	ToPackingUnitQty string `json:"toPackingUnitQty"`
	DetailFromTaxableAmountFc string `json:"detail_FromTaxableAmount_Fc"`
	DetailToTaxableAmountFc string `json:"detail_ToTaxableAmount_Fc"`
	DetailFromInpAmountFc string `json:"detail_FromInpAmount_Fc"`
	DetailFromTaxAmountFc string `json:"detail_FromTaxAmount_Fc"`
	DetailToInpAmountSc string `json:"Detail_ToInpAmount_Sc"`
	FromDetailUserQty string `json:"FromDetail_UserQty"`
	ToDetailUserQty string `json:"toDetail_UserQty"`
	FromSubUserQty string `json:"FromSub_UserQty"`
	ToSubUserQty string `json:"ToSub_UserQty"`
	IsExitsFileFlagList []string `json:"isExitsFileFlagList"`
	IsExistsCommentFlagList []string `json:"isExistsCommentFlagList"`
	Type string `json:"type"`
}

type SoSearchRes struct {
	Totals int64 `json:"totals"`
	Status int64 `json:"status"`
	Data struct {
		ChangeCount int `json:"ChangeCount"`
		InpDate string `json:"InpDate"`
		SlipType int `json:"SlipType"`
		SlipTypeName string `json:"SlipTypeName"`
		VoucherNo string `json:"VoucherNo"`
		CustomerNo interface{} `json:"CustomerNo"`
		CorrespType int `json:"CorrespType"`
		CorrespTypeName string `json:"CorrespTypeName"`
		CorrespCode string `json:"CorrespCode"`
		CorrespName string `json:"CorrespName"`
		DelivToCode string `json:"DelivToCode"`
		DelivToName interface{} `json:"DelivToName"`
		HeaderDeptCode string `json:"Header_DeptCode"`
		HeaderDeptName string `json:"Header_DeptName"`
		InchargeCode string `json:"InchargeCode"`
		TradePattern string `json:"TradePattern"`
		TradePatternName string `json:"TradePatternName"`
		HeaderStatus int `json:"HeaderStatus"`
		HeaderStatusName string `json:"HeaderStatusName"`
		Currency string `json:"Currency"`
		CurrencyName string `json:"CurrencyName"`
		RateType string `json:"RateType"`
		RateTypeName string `json:"RateTypeName"`
		Rate int `json:"Rate"`
		HeaderTaxableCode string `json:"Header_TaxableCode"`
		HeaderTaxableName string `json:"header_TaxableName"`
		HeaderTaxType int `json:"Header_TaxType"`
		HeaderTaxTypeName string `json:"Header_TaxTypeName"`
		TaxRate int `json:"TaxRate"`
		HeaderTaxAmountFC int `json:"Header_TaxAmount_FC"`
		HeaderTaxAmountSc int `json:"Header_TaxAmount_Sc"`
		TotalAmountFc int `json:"TotalAmount_Fc"`
		TotalAMountSc int `json:"TotalAMount_Sc"`
		HeaderDescriptCode string `json:"Header_DescriptCode"`
		HeaderDescriptName string `json:"Header_DescriptName"`
		PostpadColor string `json:"PostpadColor"`
		SettleMethod string `json:"SettleMethod"`
		SettleMethodName string `json:"SettleMethodName"`
		BankAcOwnCode string `json:"BankAc_OwnCode"`
		BankAcOwnName string `json:"BankAc_OwnName"`
		PaymentTerms string `json:"PaymentTerms"`
		DueDate string `json:"DueDate"`
		FileName string `json:"FileName"`
		HeaderUserMaster string `json:"Header_UserMaster"`
		HeaderUserMName string `json:"Header_UserMName"`
		HeaderUserText string `json:"Header_UserText"`
		HeaderFlexMaster1 string `json:"Header_FlexMaster1"`
		HeaderFlexMName1 string `json:"Header_FlexMName1"`
		HeaderFlexMaster2 string `json:"Header_FlexMaster2"`
		HeaderFlexMName2 string `json:"Header_FlexMName2"`
		HeaderFlexCode1 string `json:"Header_FlexCode1"`
		HeaderFlexCode2 string `json:"Header_FlexCode2"`
		HeaderFlexText1 string `json:"Header_FlexText1"`
		HeaderFlexText2 string `json:"Header_FlexText2"`
		HeaderFlexText3 string `json:"Header_FlexText3"`
		HeaderFlexText4 string `json:"Header_FlexText4"`
		HeaderFlexText5 string `json:"Header_FlexText5"`
		HeaderFlexDate string `json:"Header_FlexDate"`
		ForCondition string `json:"For Condition"`
		ForConditionName string `json:"ForConditionName"`
		ApproveStatus int `json:"approveStatus"`
		ApproveStatusName string `json:"ApproveStatusName"`
		ApproveOperator string `json:"ApproveOperator"`
		ApproveOperatorName string `json:"ApproveOperatorName"`
		ApproveDateTime string `json:"ApproveDateTime"`
		CreateOperator string `json:"CreateOperator"`
		CreateOperatorName string `json:"CreateOperatorName"`
		CreateDateTime string `json:"CreateDateTime"`
		ChangeOperator string `json:"Change Operator"`
		ChangeOperatorName string `json:"ChangeOperatorName"`
		ChangeDateTime string `json:"ChangeDateTime"`
		ImportFlag string `json:"ImportFlag"`
		ImportFlagName string `json:"ImportFlagName"`
		IsExitsFile string `json:"IsExitsFile"`
		IsExistsComment string `json:"IsExistsComment"`
		//BankAcOwnCode string `json:"BankAc_OwnCode"`
		//BankAcOwnName string `json:"BankAcOwnName"`
		BankAcCorrespCode string `json:"BankAc_CorrespCode"`
		BankAcCorrespName string `json:"BankAc_CorrespName"`
		MaxReserveQty int `json:"MaxReserveQty"`
		Sequence int `json:"Sequence"`
		PageId string `json:"PageId"`
		CompanyCode string `json:"CompanyCode"`
		DispId string `json:"DispId"`
		InternalNo string `json:"InternalNo"`
		AccountYear int `json:"Account Year"`
		AccountTerm int `json:"AccountTerm"`
		SucceedInternalNo string `json:"SucceedInternalNo"`
		SucceedVoucherNo string `json:"SucceedVoucherNo"`
		CopyInternalNo string `json:"CopyInternalNo"`
		Validity string `json:"validity"`
		VatSummarySet int `json:"VatSummarySet"`
		VoucherLinkUrl string `json:"VoucherLinkUrl"`
		PrintCount int `json:"PrintCount"`
		SettleFlag bool `json:"SettleFlag"`
		IsTaxable int `json:"IsTaxable"`
		DrCrType int `json:"DrCrType"`
		VatFlag bool `json:"VatFlag"`
		CrAcSubCode string `json:"Cr_Ac_SubCode"`
		CrAcCode string `json:"Cr_Ac_Code"`
		DrAcSubCode string `json:"Dr_Ac_SubCode"`
		DrAcCode string `json:"Dr_Ac_Code"`
		GoodsAcFlag string `json:"Goods_AcFlag"`
		SalVatSummarySet int `json:"Sal_VatSummarySet"`
		IsApproved bool `json:"IsApproved"`
		SucceedVoucherLinkUrl string `json:"SucceedVoucherLinkUrl"`
		CreditLimit int `json:"CreditLimit"`
		CreditLimitFlag int `json:"CreditLimitFlag"`
		Warning int `json:"Warning"`
		ReserveFlag int `json:"ReserveFlag"`
		CustomerNoFlagName string `json:"CustomerNoFlagName"`
		CustomerNoFlag int `json:"CustomerNoFlag"`
		TaxAmountFc int `json:"TaxAmount_Fc"`
		HeaderTaxRate int `json:"Header_TaxRate"`
		//TotalAmountFc int `json:"TotalAmountFc"`
		TotalAmountSc int `json:"totalAmountSc"`
		TaxAmountSc int `json:"TaxAmount_Sc"`
		Exchanged bool `json:"Exchanged"`
		TaxType int `json:"TaxType"`
		PurVatSummarySet int `json:"Pur_VatSummarySet"`
		CanReserveFlag int `json:"CanReserveFlag"`
		CanShpSucced int `json:"CanShpSucced"`
		CanPoSucceed int `json:"CanPoSucceed"`
		HasSucceed bool `json:"HasSucceed"`
		CanClose int `json:"CanClose"`
		CanDirSend int `json:"CanDirSend"`
	} `json:"data"`
}

type SoSearchDetailReq struct {

}

type SoSearchDetailRes struct {

}

type SoDeleteReq struct {

}

type SoDeleteRes struct {

}

type SoSaveReq struct{
	Data []SoSaveData `json:"Data"`
}

type SoSaveData struct {
	LineNo int `json:"LineNo"`
	Values []Value `json:"Values"`
}

type Value struct {
	FieldId string `json:"FieldId"`
	FieldName string `json:"FieldName"`
	Value string `json:"Value"`
	Col int `json:"Col"`
}

type SoSaveKeyValue struct {
	InputDate string `json:"inputDate"`
	VoucherNo string `json:"voucherNo"`
	CustomerNo string `json:"customerNo"`
	SucceedVoucherNo string `json:"succeedVoucherNo"`
	CorrespType string `json:"correspType"`
	CustomerCode string `json:"customerCode"`
	DelivToCode string `json:"delivToCode"`
	DeptCode string `json:"deptCode"`
	InchargeCode string `json:"inchargeCode"`
	ArPatternCode string `json:"ar_PatternCode"`
	CurrencyCode string `json:"currencyCode"`
	RateTypeCode string `json:"rateTypeCode"`
	Rate string `json:"rate"`
	TaxableCode string `json:"taxableCode"`
	ForConditionList string `json:"forConditionList"`
	TaxType string `json:"taxType"`
	HeaderDescriptCode string `json:"header_DescriptCode"`
	HeaderDescriptDescirptName string `json:"Header_DescriptDescirptName"`
	PostpadColor string `json:"postpadColor"`
	PostpadText string `json:"postpadText"`
	SettleMethod string `json:"settleMethod"`
	BankAcctOwn string `json:"bankAcct_Own"`
	PaymentTerms string `json:"paymentTerms"`
	DueDate string `json:"dueDate"`
	FileName string `json:"fileName"`
	HeaderUserMaster string `json:"header_UserMaster"`
	HeaderUserText string `json:"header_UserText"`
	HeaderFlexMaster1 string `json:"header_FlexMaster1"`
	HeaderFlexMaster2 string `json:"header_FlexMaster2"`
	HeaderFlexCode1 string `json:"header_FlexCode1"`
	HeaderFlexCode2 string `json:"header_FlexCode2"`
	HeaderFlexText1 string `json:"header_FlexText1"`
	HeaderFlexText2 string `json:"header_FlexText2"`
	HeaderFlexText3 string `json:"header_FlexText3"`
	HeaderFlexText4 string `json:"header_FlexText4"`
	HeaderFlexText5 string `json:"header_FlexText5"`
	HeaderFlexDate string `json:"header_FlexDate"`
	Remarks1 string `json:"remarks1"`
	Remarks2 string `json:"remarks2"`
	Remarks3 string `json:"remarks3"`
	Remarks4 string `json:"remarks4"`
	Remarks5 string `json:"remarks5"`
	Remarks6 string `json:"remarks6"`
	Remarks7 string `json:"remarks7"`
	Remarks8 string `json:"remarks8"`
	Remarks9 string `json:"remarks9"`
	Remarks10 string `json:"remarks10"`
	Remarks11 string `json:"remarks11"`
	Remarks12 string `json:"remarks12"`
	Remarks13 string `json:"remarks13"`
	Remarks14 string `json:"remarks14"`
	Remarks15 string `json:"remarks15"`
	Remarks16 string `json:"remarks16"`
	Remarks17 string `json:"remarks17"`
	Remarks18 string `json:"remarks18"`
	Remarks19 string `json:"remarks19"`
	Remarks20 string `json:"remarks20"`
	Remarks21 string `json:"remarks21"`
	Remarks22 string `json:"remarks22"`
	Remarks23 string `json:"remarks23"`
	Remarks24 string `json:"remarks24"`
	Remarks25 string `json:"remarks25"`
	Remarks26 string `json:"remarks26"`
	Remarks27 string `json:"remarks27"`
	Remarks28 string `json:"remarks28"`
	Remarks29 string `json:"remarks29"`
	Remarks30 string `json:"remarks30"`
	RowNo string `json:"rowNo"`
	DetailType string `json:"detailType"`
	DetailDeptCode string `json:"detail_DeptCode"`
	Warehouse string `json:"warehouse"`
	LocationCode string `json:"locationCode"`
	GoodsGroup string `json:"goodsGroup"`
	GoodsCode string `json:"goodsCode"`
	ChargesCode string `json:"chargesCode"`
	CorrespGoods string `json:"corresp_Goods"`
	UnitType string `json:"unitType"`
	DetailTaxableCode string `json:"detail_TaxableCode"`
	DetailTaxType string `json:"detail_TaxType"`
	UnitPrice string `json:"unitPrice"`
	Qty string `json:"qty"`
	InpAmountF string `json:"inpAmount_Fc"`
	TaxAmountFc string `json:"taxAmount_Fc"`
	InpAmountSc string `json:"inpAmount_Sc"`
	TaxAmountSc string `json:"taxAmount_Sc"`
	TaxableAmountSc string `json:"taxableAmount_Sc"`
	PackingUnitCode string `json:"packingUnitCode"`
	PackingUnitPrice string `json:"packingUnitPrice"`
	PackingUnitQty string `json:"packingUnitQty"`
	PreshippingDate string `json:"preshippingDate"`
	DelivDate string `json:"delivDate"`
	DetailDescriptCode string `json:"detail_Descript_Code"`
	DetailDescriptName string `json:"detail_Descript_Name"`
	DetailUserMaster1 string `json:"detail_UserMaster1"`
	DetailUserMaster2 string `json:"detail_UserMaster2"`
	DetailUserMaster3 string `json:"detail_UserMaster3"`
	DetailUserMaster4 string `json:"detail_UserMaster4"`
	DetailUserMaster5 string `json:"detail_UserMaster5"`
	DetailUserCode string `json:"detail_UserCode"`
	DetailUserText string `json:"detail_UserText"`
	DetailUserDate string `json:"detail_UserDate"`
	DetailUserQty string `json:"detail_UserQty"`
	DetailFlexMaster string `json:"detail_FlexMaster"`
	DetailFlexCode string `json:"detail_FlexCode"`
	DetailFlexText1 string `json:"detail_FlexText1"`
	DetailFlexText2 string `json:"detail_FlexText2"`
	DetailFlexDate string `json:"detail_FlexDate"`
	SubDescriptCode string `json:"sub_Descript_Code"`
	SubDescriptName string `json:"sub_Descript_Name"`
	SubUserMaster1 string `json:"sub_UserMaster1"`
	SubUserMaster2 string `json:"sub_UserMaster2"`
	SubUserMaster3 string `json:"sub_UserMaster3"`
	SubUserMaster4 string `json:"sub_UserMaster4"`
	SubUserMaster5 string `json:"sub_UserMaster5"`
	SubUserCode string `json:"sub_UserCode"`
	SubUserDate string `json:"Sub_UserDate"`
	SpriceCode string `json:"spriceCode"`
	SpriceIndex string `json:"spriceIndex"`
}

type SoSaveRes struct {
	Sequence float64 `json:"Sequence"`
	SaveStatus string `json:"SaveStatus"`
	Status int `json:"Status"`
}

type SosucceedSearchReq struct{

}

type SosucceedSearchRes struct{

}

type SoCloseReq struct {

}
type SoCloseRes struct{

}
type SoCloseCancelReq struct{

}
type SoCloseCancelRes struct{

}

//GetOrderSearch get so order
func (order *Order) GetOrderSearch(req SoSearchReq) (res SoSearchRes, err error) {

	var accessParam string
	accessParam, err = order.GetAccessParam(ORDER_SO_SEARCH, req)
	if err != nil {
		return
	}
	uri := fmt.Sprintf(util.HTTP_BASE_URL, order.AppID, ORDER_SO_SEARCH)
	var response []byte
	response, err = util.NewHTTPPost(uri, accessParam)

	if err != nil {
		return
	}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.Status != 0{
		err = fmt.Errorf("GetOrderListGet Error , errcode=%s , errmsg=%s", res.Status, res.Totals)
		return
	}
	return
}
