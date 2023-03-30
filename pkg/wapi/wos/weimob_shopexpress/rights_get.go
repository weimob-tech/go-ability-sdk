package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsGet
// @id 1962
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1962?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=售后订单详情)
func (client *Client) RightsGet(request *RightsGetRequest) (response *RightsGetResponse, err error) {
	rpcResponse := CreateRightsGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsGetRequest struct {
	*api.BaseRequest

	AftersaleNo string `json:"aftersaleNo,omitempty"`
}

type RightsGetResponse struct {
	ItemOutList      []RightsGetResponseItemOutList `json:"itemOutList,omitempty"`
	Uid              int64                          `json:"uid,omitempty"`
	NickName         string                         `json:"nickName,omitempty"`
	OrderNo          string                         `json:"orderNo,omitempty"`
	AftersaleNo      string                         `json:"aftersaleNo,omitempty"`
	RefundTypeText   string                         `json:"refundTypeText,omitempty"`
	RefundStatus     int64                          `json:"refundStatus,omitempty"`
	RefundStatusText string                         `json:"refundStatusText,omitempty"`
	RefundStatusDesc string                         `json:"refundStatusDesc,omitempty"`
	UserType         int64                          `json:"userType,omitempty"`
	RefundAmount     string                         `json:"refundAmount,omitempty"`
	DeliveryFee      string                         `json:"deliveryFee,omitempty"`
	LogisticalTax    string                         `json:"logisticalTax,omitempty"`
	ReasonTypeText   string                         `json:"reasonTypeText,omitempty"`
	Reason           string                         `json:"reason,omitempty"`
	ProofPic         []int64                        `json:"proofPic,omitempty"`
	Refuse           string                         `json:"refuse,omitempty"`
	FailReason       string                         `json:"failReason,omitempty"`
	CreateTime       string                         `json:"createTime,omitempty"`
	Currency         string                         `json:"currency,omitempty"`
	TimeZone         string                         `json:"timeZone,omitempty"`
	IsRefundTax      bool                           `json:"isRefundTax,omitempty"`
}

func CreateRightsGetRequest() (request *RightsGetRequest) {
	request = &RightsGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "rights/get", "apigw")
	request.Method = api.POST
	return
}

func CreateRightsGetResponse() (response *api.BaseResponse[RightsGetResponse]) {
	response = api.CreateResponse[RightsGetResponse](&RightsGetResponse{})
	return
}

type RightsGetResponseItemOutList struct {
	OrderItemNo       string `json:"orderItemNo,omitempty"`
	Count             int64  `json:"count,omitempty"`
	Price             string `json:"price,omitempty"`
	GoodsName         string `json:"goodsName,omitempty"`
	GoodsCode         string `json:"goodsCode,omitempty"`
	SkuCode           string `json:"skuCode,omitempty"`
	GoodsPic          string `json:"goodsPic,omitempty"`
	SkuName           string `json:"skuName,omitempty"`
	CustomSkuCode     string `json:"customSkuCode,omitempty"`
	TotalAmount       string `json:"totalAmount,omitempty"`
	GoodsRealAmount   string `json:"goodsRealAmount,omitempty"`
	CouponAmount      string `json:"couponAmount,omitempty"`
	OrderDetailAmount string `json:"orderDetailAmount,omitempty"`
	FreightAmount     string `json:"freightAmount,omitempty"`
	ExciseTax         string `json:"exciseTax,omitempty"`
	LogisticalTax     string `json:"logisticalTax,omitempty"`
}
