package Internationalized_software

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderGetOrderDetail
// @id 1192
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1192?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=订单详情)
func (client *Client) OrderGetOrderDetailV1_0(request *WeimobShopexpressOrderGetOrderDetailRequest) (response *WeimobShopexpressOrderGetOrderDetailResponse, err error) {
	rpcResponse := CreateWeimobShopexpressOrderGetOrderDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WeimobShopexpressOrderGetOrderDetailRequest struct {
	*api.BaseRequest

	Uid     int64  `json:"uid,omitempty"`
	OrderNo string `json:"orderNo,omitempty"`
}

type WeimobShopexpressOrderGetOrderDetailResponse struct {
	CountdownOut           WeimobShopexpressOrderGetOrderDetailResponseCountdownOut             `json:"countdownOut,omitempty"`
	OrderDeliveryInfoOut   WeimobShopexpressOrderGetOrderDetailResponseOrderDeliveryInfoOut     `json:"orderDeliveryInfoOut,omitempty"`
	OrderItemDetailOutList []WeimobShopexpressOrderGetOrderDetailResponseOrderItemDetailOutList `json:"orderItemDetailOutList,omitempty"`
	OrderRefundInfoOut     WeimobShopexpressOrderGetOrderDetailResponseOrderRefundInfoOut       `json:"orderRefundInfoOut,omitempty"`
	StoreId                int64                                                                `json:"storeId,omitempty"`
	Uid                    int64                                                                `json:"uid,omitempty"`
	OrderNo                string                                                               `json:"orderNo,omitempty"`
	UserName               string                                                               `json:"userName,omitempty"`
	Phone                  string                                                               `json:"phone,omitempty"`
	Email                  string                                                               `json:"email,omitempty"`
	UserBuySumCount        int64                                                                `json:"userBuySumCount,omitempty"`
	UserBuySumAmount       string                                                               `json:"userBuySumAmount,omitempty"`
	ItemCount              int64                                                                `json:"itemCount,omitempty"`
	Currency               string                                                               `json:"currency,omitempty"`
	TimeZone               string                                                               `json:"timeZone,omitempty"`
	Amount                 string                                                               `json:"amount,omitempty"`
	ExciseTax              string                                                               `json:"exciseTax,omitempty"`
	LogisticalTax          string                                                               `json:"logisticalTax,omitempty"`
	Freight                string                                                               `json:"freight,omitempty"`
	PayAmount              string                                                               `json:"payAmount,omitempty"`
	CouponAmount           string                                                               `json:"couponAmount,omitempty"`
	OrderStatus            int64                                                                `json:"orderStatus,omitempty"`
	OrderType              int64                                                                `json:"orderType,omitempty"`
	PayStatus              int64                                                                `json:"payStatus,omitempty"`
	PayStatusText          string                                                               `json:"payStatusText,omitempty"`
	PayTime                string                                                               `json:"payTime,omitempty"`
	ArchiveStatus          int64                                                                `json:"archiveStatus,omitempty"`
	CreateTime             string                                                               `json:"createTime,omitempty"`
	Remark                 string                                                               `json:"remark,omitempty"`
	AutoCloseTime          string                                                               `json:"autoCloseTime,omitempty"`
	PayChannelNo           string                                                               `json:"payChannelNo,omitempty"`
	AutoReceiptTime        string                                                               `json:"autoReceiptTime,omitempty"`
	IsSubmitRefund         bool                                                                 `json:"isSubmitRefund,omitempty"`
	TaxSum                 string                                                               `json:"taxSum,omitempty"`
	CouponCode             string                                                               `json:"couponCode,omitempty"`
	CanRefund              bool                                                                 `json:"canRefund,omitempty"`
	CanDelivery            bool                                                                 `json:"canDelivery,omitempty"`
	ExistRefunding         bool                                                                 `json:"existRefunding,omitempty"`
}

func CreateWeimobShopexpressOrderGetOrderDetailRequest() (request *WeimobShopexpressOrderGetOrderDetailRequest) {
	request = &WeimobShopexpressOrderGetOrderDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("Internationalized_software", "v1.0", "order/getOrderDetail", "apigw")
	request.Method = api.POST
	return
}

func CreateWeimobShopexpressOrderGetOrderDetailResponse() (response *api.BaseResponse[WeimobShopexpressOrderGetOrderDetailResponse]) {
	response = api.CreateResponse[WeimobShopexpressOrderGetOrderDetailResponse](&WeimobShopexpressOrderGetOrderDetailResponse{})
	return
}

type WeimobShopexpressOrderGetOrderDetailResponseCountdownOut struct {
	Start int64 `json:"start,omitempty"`
	End   int64 `json:"end,omitempty"`
}

type WeimobShopexpressOrderGetOrderDetailResponseOrderDeliveryInfoOut struct {
	OwnLogistics        bool   `json:"ownLogistics,omitempty"`
	LogisticsStatus     int64  `json:"logisticsStatus,omitempty"`
	LogisticsStatusText string `json:"logisticsStatusText,omitempty"`
	ExpressNo           string `json:"expressNo,omitempty"`
	ExpressCompanyCode  string `json:"expressCompanyCode,omitempty"`
	ExpressCompany      string `json:"expressCompany,omitempty"`
	ConsigneeName       string `json:"consigneeName,omitempty"`
	ConsigneeTel        string `json:"consigneeTel,omitempty"`
	ConsigneeAddress    string `json:"consigneeAddress,omitempty"`
	ReceiptTime         string `json:"receiptTime,omitempty"`
	DeliveryTime        string `json:"deliveryTime,omitempty"`
	ExpressUrl          string `json:"expressUrl,omitempty"`
}

type WeimobShopexpressOrderGetOrderDetailResponseOrderItemDetailOutList struct {
	OrderItemNo     string `json:"orderItemNo,omitempty"`
	GoodsCode       string `json:"goodsCode,omitempty"`
	GoodsName       string `json:"goodsName,omitempty"`
	SkuCode         string `json:"skuCode,omitempty"`
	SkuName         string `json:"skuName,omitempty"`
	ImageUrl        string `json:"imageUrl,omitempty"`
	Count           int64  `json:"count,omitempty"`
	Price           string `json:"price,omitempty"`
	Amount          string `json:"amount,omitempty"`
	PayAmount       string `json:"payAmount,omitempty"`
	FreightAmount   string `json:"freightAmount,omitempty"`
	ExciseTax       string `json:"exciseTax,omitempty"`
	LogisticalTax   string `json:"logisticalTax,omitempty"`
	CreateTime      string `json:"createTime,omitempty"`
	CustomSkuCode   string `json:"customSkuCode,omitempty"`
	RefundStatusStr string `json:"refundStatusStr,omitempty"`
	RefundStatus    int64  `json:"refundStatus,omitempty"`
}

type WeimobShopexpressOrderGetOrderDetailResponseOrderRefundInfoOut struct {
	OwnRefund    bool   `json:"ownRefund,omitempty"`
	RefundMoney  string `json:"refundMoney,omitempty"`
	RefundTime   string `json:"refundTime,omitempty"`
	RefundRemark string `json:"refundRemark,omitempty"`
}
