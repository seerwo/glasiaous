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
	FieldId   string `json:"FieldId"`
	FieldName string `json:"FieldName"`
	Value     string `json:"Value"`
	Col       int    `json:"Col"`
}
//
type SoSaveKeyValue struct {
	InputDate string `json:"inputDate" comment:"計上日"`
	VoucherNo string `json:"voucherNo" comment:"伝票No."`
	CustomerNo string `json:"customerNo" comment:"客先注文No"`
	SucceedVoucherNo string `json:"succeedVoucherNo" comment:""`
	CorrespType string `json:"correspType" comment:"取引先区分"`
	CustomerCode string `json:"customerCode" comment:"取引先"`
	DelivToCode string `json:"delivToCode" comment:"納品先/納入元"`
	DeptCode string `json:"deptCode" comment:"部門"`
	InchargeCode string `json:"inchargeCode" comment:"担当者"`
	ArPatternCode string `json:"ar_PatternCode" comment:"取引パターン"`
	CurrencyCode string `json:"currencyCode" comment:"通貨"`
	RateTypeCode string `json:"rateTypeCode" comment:"レートタイプ"`
	Rate string `json:"rate" comment:"レート"`
	TaxableCode string `json:"taxableCode" comment:""`
	ForConditionList string `json:"forConditionList" comment:"ヘッダ課税区分"`
	TaxType string `json:"taxType" comment:"内税/外税区分"`
	HeaderDescriptCode string `json:"header_DescriptCode" comment:"伝票摘要コード"`
	HeaderDescriptDescirptName string `json:"Header_DescriptDescirptName" comment:"伝票摘要内容"`
	PostpadColor string `json:"postpadColor" comment:"付箋紙色"`
	PostpadText string `json:"postpadText" comment:"付箋紙内容"`
	SettleMethod string `json:"settleMethod" comment:"決済方法"`
	BankAcctOwn string `json:"bankAcct_Own" comment:"自社口座"`
	PaymentTerms string `json:"paymentTerms" comment:"回収条件"`
	DueDate string `json:"dueDate" comment:"決済予定日"`
	FileName string `json:"fileName" comment:"ファイル名"`
	HeaderUserMaster string `json:"header_UserMaster" comment:"ヘッダユーザマスタ"`
	HeaderUserText string `json:"header_UserText" comment:"ヘッダユーザテキスト"`
	HeaderFlexMaster1 string `json:"header_FlexMaster1" comment:"ヘッダフレックスマスタ1"`
	HeaderFlexMaster2 string `json:"header_FlexMaster2" comment:"ヘッダフレックスマスタ2"`
	HeaderFlexCode1 string `json:"header_FlexCode1" comment:"ヘッダフレックスコード1"`
	HeaderFlexCode2 string `json:"header_FlexCode2" comment:"ヘッダフレックスコード2"`
	HeaderFlexText1 string `json:"header_FlexText1" comment:"ヘッダフレックステキスト1"`
	HeaderFlexText2 string `json:"header_FlexText2" comment:"ヘッダフレックステキスト2"`
	HeaderFlexText3 string `json:"header_FlexText3" comment:"ヘッダフレックステキスト3"`
	HeaderFlexText4 string `json:"header_FlexText4" comment:"ヘッダフレックステキスト4"`
	HeaderFlexText5 string `json:"header_FlexText5" comment:"ヘッダフレックステキスト5"`
	HeaderFlexDate string `json:"header_FlexDate" comment:"ヘッダフレックス日付"`
	Remarks1 string `json:"remarks1" comment:"ﾒﾓ情報1"`
	Remarks2 string `json:"remarks2" comment:"ﾒﾓ情報2"`
	Remarks3 string `json:"remarks3" comment:"ﾒﾓ情報3"`
	Remarks4 string `json:"remarks4" comment:"ﾒﾓ情報4"`
	Remarks5 string `json:"remarks5" comment:"ﾒﾓ情報5"`
	Remarks6 string `json:"remarks6" comment:"ﾒﾓ情報6"`
	Remarks7 string `json:"remarks7" comment:"ﾒﾓ情報7"`
	Remarks8 string `json:"remarks8" comment:"ﾒﾓ情報8"`
	Remarks9 string `json:"remarks9" comment:"ﾒﾓ情報9"`
	Remarks10 string `json:"remarks10" comment:"ﾒﾓ情報10"`
	Remarks11 string `json:"remarks11" comment:"ﾒﾓ情報11"`
	Remarks12 string `json:"remarks12" comment:"ﾒﾓ情報12"`
	Remarks13 string `json:"remarks13" comment:"ﾒﾓ情報13"`
	Remarks14 string `json:"remarks14" comment:"ﾒﾓ情報14"`
	Remarks15 string `json:"remarks15" comment:"ﾒﾓ情報15"`
	Remarks16 string `json:"remarks16" comment:"ﾒﾓ情報16"`
	Remarks17 string `json:"remarks17" comment:"ﾒﾓ情報17"`
	Remarks18 string `json:"remarks18" comment:"ﾒﾓ情報18"`
	Remarks19 string `json:"remarks19" comment:"ﾒﾓ情報19"`
	Remarks20 string `json:"remarks20" comment:"ﾒﾓ情報20"`
	Remarks21 string `json:"remarks21" comment:"ﾒﾓ情報21"`
	Remarks22 string `json:"remarks22" comment:"ﾒﾓ情報22"`
	Remarks23 string `json:"remarks23" comment:"ﾒﾓ情報23"`
	Remarks24 string `json:"remarks24" comment:"ﾒﾓ情報24"`
	Remarks25 string `json:"remarks25" comment:"ﾒﾓ情報25"`
	Remarks26 string `json:"remarks26" comment:"ﾒﾓ情報26"`
	Remarks27 string `json:"remarks27" comment:"ﾒﾓ情報27"`
	Remarks28 string `json:"remarks28" comment:"ﾒﾓ情報28"`
	Remarks29 string `json:"remarks29" comment:"ﾒﾓ情報29"`
	Remarks30 string `json:"remarks30" comment:"ﾒﾓ情報30"`
	RowNo string `json:"rowNo" comment:""`
	DetailType string `json:"detailType" comment:"明細区分（品目・諸掛）"`
	DetailDeptCode string `json:"detail_DeptCode" comment:"部門"`
	Warehouse string `json:"warehouse" comment:"倉庫"`
	LocationCode string `json:"locationCode" comment:"ロケーション"`
	GoodsGroup string `json:"goodsGroup" comment:"品目グループ"`
	GoodsCode string `json:"goodsCode" comment:"品目"`
	ChargesCode string `json:"chargesCode" comment:"諸掛"`
	CorrespGoods string `json:"corresp_Goods" comment:"取引先品番"`
	UnitType string `json:"unitType" comment:"単位計上/荷姿単位計上"`
	DetailTaxableCode string `json:"detail_TaxableCode" comment:"課税区分"`
	DetailTaxType string `json:"detail_TaxType" comment:"内税/外税区分"`
	UnitPrice string `json:"unitPrice" comment:"単価"`
	Qty string `json:"qty" comment:"数量"`
	InpAmountF string `json:"inpAmount_Fc" comment:"入力金額"`
	TaxAmountFc string `json:"taxAmount_Fc" comment:"入力税額"`
	InpAmountSc string `json:"inpAmount_Sc" comment:"入力税対象金額"`
	TaxAmountSc string `json:"taxAmount_Sc" comment:"基準税額"`
	TaxableAmountSc string `json:"taxableAmount_Sc" comment:"基準税対象金額"`
	PackingUnitCode string `json:"packingUnitCode" comment:"荷姿コード"`
	PackingUnitPrice string `json:"packingUnitPrice" comment:"荷姿単価"`
	PackingUnitQty string `json:"packingUnitQty" comment:"荷姿数量"`
	PreshippingDate string `json:"preshippingDate" comment:"出荷予定日"`
	DelivDate string `json:"delivDate" comment:"納入期日"`
	DetailDescriptCode string `json:"detail_Descript_Code" comment:"明細摘要コード"`
	DetailDescriptName string `json:"detail_Descript_Name" comment:"明細摘要内容"`
	DetailUserMaster1 string `json:"detail_UserMaster1" comment:"明細ユーザマスタ1"`
	DetailUserMaster2 string `json:"detail_UserMaster2" comment:"明細ユーザマスタ2"`
	DetailUserMaster3 string `json:"detail_UserMaster3" comment:"明細ユーザマスタ3"`
	DetailUserMaster4 string `json:"detail_UserMaster4" comment:"明細ユーザマスタ4"`
	DetailUserMaster5 string `json:"detail_UserMaster5" comment:"明細ユーザマスタ5"`
	DetailUserCode string `json:"detail_UserCode" comment:"明細ユーザコード"`
	DetailUserText string `json:"detail_UserText" comment:"明細ユーザテキスト"`
	DetailUserDate string `json:"detail_UserDate" comment:"明細ユーザ日付"`
	DetailUserQty string `json:"detail_UserQty" comment:"明細ユーザ数量"`
	DetailFlexMaster string `json:"detail_FlexMaster" comment:"明細フレックスマスタ"`
	DetailFlexCode string `json:"detail_FlexCode" comment:"明細フレックスコード"`
	DetailFlexText1 string `json:"detail_FlexText1" comment:"明細フレックステキスト1"`
	DetailFlexText2 string `json:"detail_FlexText2" comment:"明細フレックステキスト2"`
	DetailFlexDate string `json:"detail_FlexDate" comment:"明細フレックス日付"`
	SubDescriptCode string `json:"sub_Descript_Code" comment:"明細摘要コード"`
	SubDescriptName string `json:"sub_Descript_Name" comment:"明細摘要内容"`
	SubUserMaster1 string `json:"sub_UserMaster1" comment:"明細ユーザマスタ1"`
	SubUserMaster2 string `json:"sub_UserMaster2" comment:"明細ユーザマスタ2"`
	SubUserMaster3 string `json:"sub_UserMaster3" comment:"明細ユーザマスタ3"`
	SubUserMaster4 string `json:"sub_UserMaster4" comment:"明細ユーザマスタ4"`
	SubUserMaster5 string `json:"sub_UserMaster5" comment:"明細ユーザマスタ5"`
	SubUserCode string `json:"sub_UserCode" comment:"明細ユーザコード"`
	SubUserDate string `json:"Sub_UserDate" comment:"明細ユーザ日付"`
	SpriceCode string `json:"spriceCode" comment:""`
	SpriceIndex string `json:"spriceIndex" comment:""`
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
	uri := fmt.Sprintf(util.HTTP_BASE_URL +  ORDER_SO_SEARCH)
	var response []byte

	response, err = util.NewHTTPPost(uri, accessParam)
	fmt.Println(uri, accessParam)
	if err != nil {
		return
	}
	fmt.Println(string(response))
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.Status != 0{
		err = fmt.Errorf("GetOrderSearch Error , errcode=%d , errmsg=%d", res.Status, res.Totals)
		return
	}
	return
}

//GetOrderSearch get so order
func (order *Order) GetOrderSave(req SoSaveReq) (res SoSaveRes, err error) {

	var accessParam string
	accessParam, err = order.GetAccessParam(ORDER_SO_SAVE, req)
	if err != nil {
		return
	}
	uri := fmt.Sprintf(util.HTTP_BASE_URL + ORDER_SO_SAVE)
	var response []byte
	fmt.Println(uri, accessParam)
	response, err = util.NewHTTPPost(uri, accessParam)

	if err != nil {
		return
	}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.Status != 0{
		err = fmt.Errorf("GetOrderSave Error , errcode=%d , errmsg=%s", res.Status, res.SaveStatus)
		return
	}
	return
}