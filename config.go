package go_buy365

const (
	ACCESS_KEY = "hiahaihaihaihaihaiahiahihaih" //调用psp的签名key
	BACK_KEY   = "hiahaihaihaihaihaiahiahihaih" //回调的签名key

	//deposit
	BASE_URL = "https://swpapi.tpex.cc/UtInRecordApi/orderGateWay"

	//withdraw
	WITHDRAW_URL         = "https://mmapi.proxima131.com/AjaxOpen/saveOutOrder"
	WITHDRAW_CONFIRM_URL = "https://mmapi.proxima131.com/AjaxOpen/appealOutOrder"

	//orderlist
	ORDERLIST_URL = "https://mmapi.proxima131.com/AjaxOpen/getOutOrderList"
)
