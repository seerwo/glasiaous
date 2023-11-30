package order

import (
	"encoding/json"
	"fmt"
	"github.com/seerwo/gasiaous/erp/context"
	"github.com/seerwo/gasiaous/util"
	"reflect"
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

//GetSupplierOrder get supplier order
func (order *Order) GetOrderListGet(req ReqOrderListGet) (res ResOrderListGet, err error) {

	var accessParam string
	accessParam, err = order.GetAccessParam(SQ_ORDER_LIST_GET, req)
	if err != nil {
		return
	}
	uri := fmt.Sprintf(util.HTTP_BASE_URL, order.AppID, SQ_ORDER_LIST_GET)
	var response []byte
	response, err = util.NewHTTPPost(uri, accessParam)

	if err != nil {
		return
	}
	err = json.Unmarshal(response, &res)
	if err != nil {
		return
	}
	if res.ErrCode != "0000" {
		err = fmt.Errorf("GetOrderListGet Error , errcode=%s , errmsg=%s", res.ErrCode, res.ErrMsg)
		return
	}
	return
}
