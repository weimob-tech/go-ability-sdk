package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// TradeMerchantCreateOrder
// @id 1286
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=上传订单)
func (client *Client) TradeMerchantCreateOrder(request *TradeMerchantCreateOrderRequest) (response *TradeMerchantCreateOrderResponse, err error) {
	rpcResponse := CreateTradeMerchantCreateOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type TradeMerchantCreateOrderRequest struct {
	*api.BaseRequest

	ConsumerWid                    int64   `json:"consumerWid,omitempty"`
	OrderAmount                    float64 `json:"orderAmount,omitempty"`
	PayStatus                      int     `json:"payStatus,omitempty"`
	StoreId                        int64   `json:"storeId,omitempty"`
	ThirdOrderNo                   string  `json:"thirdOrderNo,omitempty"`
	StoreDepositCardDiscountAmount float64 `json:"storeDepositCardDiscountAmount,omitempty"`
	StoreDepositCardTempletId      string  `json:"storeDepositCardTempletId,omitempty"`
	StoreDepositCardCode           string  `json:"storeDepositCardCode,omitempty"`
	CouponDiscount                 float64 `json:"couponDiscount,omitempty"`
	CouponCode                     string  `json:"couponCode,omitempty"`
	BalanceDiscount                float64 `json:"balanceDiscount,omitempty"`
	MemberLevelDiscount            float64 `json:"memberLevelDiscount,omitempty"`
	PointDiscount                  float64 `json:"pointDiscount,omitempty"`
	PaidAmount                     float64 `json:"paidAmount,omitempty"`
}

type TradeMerchantCreateOrderResponse struct {
}

func CreateTradeMerchantCreateOrderRequest() (request *TradeMerchantCreateOrderRequest) {
	request = &TradeMerchantCreateOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "tradeMerchant/createOrder", "api")
	request.Method = api.POST
	return
}

func CreateTradeMerchantCreateOrderResponse() (response *api.BaseResponse[TradeMerchantCreateOrderResponse]) {
	response = api.CreateResponse[TradeMerchantCreateOrderResponse](&TradeMerchantCreateOrderResponse{})
	return
}
