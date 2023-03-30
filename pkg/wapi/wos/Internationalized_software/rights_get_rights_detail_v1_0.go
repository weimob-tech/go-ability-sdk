package Internationalized_software

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsGetRightsDetail
// @id 1197
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1197?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=售后订单详情)
func (client *Client) RightsGetRightsDetailV1_0(request *WeimobShopexpressRightsGetRightsDetailRequest) (response *WeimobShopexpressRightsGetRightsDetailResponse, err error) {
	rpcResponse := CreateWeimobShopexpressRightsGetRightsDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WeimobShopexpressRightsGetRightsDetailRequest struct {
	*api.BaseRequest

	AftersaleNo string `json:"aftersaleNo,omitempty"`
}

type WeimobShopexpressRightsGetRightsDetailResponse struct {
	ItemOutList      []WeimobShopexpressRightsGetRightsDetailResponseItemOutList      `json:"itemOutList,omitempty"`
	OperationOutList []WeimobShopexpressRightsGetRightsDetailResponseOperationOutList `json:"operationOutList,omitempty"`
	Uid              int64                                                            `json:"uid,omitempty"`
	NickName         string                                                           `json:"nickName,omitempty"`
	OrderNo          string                                                           `json:"orderNo,omitempty"`
	AftersaleNo      string                                                           `json:"aftersaleNo,omitempty"`
	RefundTypeText   string                                                           `json:"refundTypeText,omitempty"`
	RefundStatus     int64                                                            `json:"refundStatus,omitempty"`
	RefundStatusText string                                                           `json:"refundStatusText,omitempty"`
	RefundStatusDesc string                                                           `json:"refundStatusDesc,omitempty"`
	UserType         int64                                                            `json:"userType,omitempty"`
	TransactionId    string                                                           `json:"transactionId,omitempty"`
	RefundAmount     string                                                           `json:"refundAmount,omitempty"`
	DeliveryFee      string                                                           `json:"deliveryFee,omitempty"`
	LogisticalTax    string                                                           `json:"logisticalTax,omitempty"`
	RealAmount       string                                                           `json:"realAmount,omitempty"`
	ReasonTypeText   string                                                           `json:"reasonTypeText,omitempty"`
	RefundOriginText string                                                           `json:"refundOriginText,omitempty"`
	Reason           string                                                           `json:"reason,omitempty"`
	ProofPic         []int64                                                          `json:"proofPic,omitempty"`
	Refuse           string                                                           `json:"refuse,omitempty"`
	FailReason       string                                                           `json:"failReason,omitempty"`
	CreateTime       string                                                           `json:"createTime,omitempty"`
	Currency         string                                                           `json:"currency,omitempty"`
	TimeZone         string                                                           `json:"timeZone,omitempty"`
	IsRefundTax      bool                                                             `json:"isRefundTax,omitempty"`
	OfflineRefund    bool                                                             `json:"offlineRefund,omitempty"`
}

func CreateWeimobShopexpressRightsGetRightsDetailRequest() (request *WeimobShopexpressRightsGetRightsDetailRequest) {
	request = &WeimobShopexpressRightsGetRightsDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("Internationalized_software", "v1.0", "rights/getRightsDetail", "apigw")
	request.Method = api.POST
	return
}

func CreateWeimobShopexpressRightsGetRightsDetailResponse() (response *api.BaseResponse[WeimobShopexpressRightsGetRightsDetailResponse]) {
	response = api.CreateResponse[WeimobShopexpressRightsGetRightsDetailResponse](&WeimobShopexpressRightsGetRightsDetailResponse{})
	return
}

type WeimobShopexpressRightsGetRightsDetailResponseItemOutList struct {
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
	RealAmount        string `json:"realAmount,omitempty"`
	GoodsRealAmount   string `json:"goodsRealAmount,omitempty"`
	CouponAmount      string `json:"couponAmount,omitempty"`
	OrderDetailAmount string `json:"orderDetailAmount,omitempty"`
	FreightAmount     string `json:"freightAmount,omitempty"`
	ExciseTax         string `json:"exciseTax,omitempty"`
	LogisticalTax     string `json:"logisticalTax,omitempty"`
}

type WeimobShopexpressRightsGetRightsDetailResponseOperationOutList struct {
	OperationName string `json:"operationName,omitempty"`
	Description   string `json:"description,omitempty"`
	CreateTime    string `json:"createTime,omitempty"`
}
