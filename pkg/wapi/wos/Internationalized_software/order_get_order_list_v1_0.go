package Internationalized_software

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderGetOrderList
// @id 1193
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1193?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=订单列表)
func (client *Client) OrderGetOrderListV1_0(request *WeimobShopexpressOrderGetOrderListRequest) (response *WeimobShopexpressOrderGetOrderListResponse, err error) {
	rpcResponse := CreateWeimobShopexpressOrderGetOrderListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WeimobShopexpressOrderGetOrderListRequest struct {
	*api.BaseRequest

	ArchiveStatus   int64  `json:"archiveStatus,omitempty"`
	PayStatus       int64  `json:"payStatus,omitempty"`
	LogisticsStatus int64  `json:"logisticsStatus,omitempty"`
	StartTime       string `json:"startTime,omitempty"`
	EndTime         string `json:"endTime,omitempty"`
	SearchKeyType   int64  `json:"searchKeyType,omitempty"`
	SearchKeyValue  string `json:"searchKeyValue,omitempty"`
	Uid             int64  `json:"uid,omitempty"`
	PageSize        int64  `json:"pageSize,omitempty"`
	PageIndex       int64  `json:"pageIndex,omitempty"`
}

type WeimobShopexpressOrderGetOrderListResponse struct {
	OrderItemListOutList []WeimobShopexpressOrderGetOrderListResponseOrderItemListOutList `json:"orderItemListOutList,omitempty"`
	StoreId              int64                                                            `json:"storeId,omitempty"`
	Uid                  int64                                                            `json:"uid,omitempty"`
	UserName             string                                                           `json:"userName,omitempty"`
	Phone                string                                                           `json:"phone,omitempty"`
	Email                string                                                           `json:"email,omitempty"`
	OrderNo              string                                                           `json:"orderNo,omitempty"`
	Currency             string                                                           `json:"currency,omitempty"`
	TimeZone             string                                                           `json:"timeZone,omitempty"`
	ItemCount            int64                                                            `json:"itemCount,omitempty"`
	Amount               string                                                           `json:"amount,omitempty"`
	PayAmount            string                                                           `json:"payAmount,omitempty"`
	OrderStatus          int64                                                            `json:"orderStatus,omitempty"`
	PayStatus            int64                                                            `json:"payStatus,omitempty"`
	PayStatusText        string                                                           `json:"payStatusText,omitempty"`
	ArchiveStatus        int64                                                            `json:"archiveStatus,omitempty"`
	LogisticsStatus      int64                                                            `json:"logisticsStatus,omitempty"`
	LogisticsStatusText  string                                                           `json:"logisticsStatusText,omitempty"`
	OwnRefund            bool                                                             `json:"ownRefund,omitempty"`
	PayTime              string                                                           `json:"payTime,omitempty"`
	DeliveryTime         string                                                           `json:"deliveryTime,omitempty"`
	CloseTime            string                                                           `json:"closeTime,omitempty"`
	Freight              string                                                           `json:"freight,omitempty"`
	DiscountAmount       string                                                           `json:"discountAmount,omitempty"`
	ExciseTax            string                                                           `json:"exciseTax,omitempty"`
	ConsigneeName        string                                                           `json:"consigneeName,omitempty"`
	ConsigneeAddress     string                                                           `json:"consigneeAddress,omitempty"`
	CreateTime           string                                                           `json:"createTime,omitempty"`
	LogisticalTax        string                                                           `json:"logisticalTax,omitempty"`
	ExistRefunding       bool                                                             `json:"existRefunding,omitempty"`
	CanRefund            bool                                                             `json:"canRefund,omitempty"`
	CanDelivery          bool                                                             `json:"canDelivery,omitempty"`
}

func CreateWeimobShopexpressOrderGetOrderListRequest() (request *WeimobShopexpressOrderGetOrderListRequest) {
	request = &WeimobShopexpressOrderGetOrderListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("Internationalized_software", "v1.0", "order/getOrderList", "apigw")
	request.Method = api.POST
	return
}

func CreateWeimobShopexpressOrderGetOrderListResponse() (response *api.BaseResponse[WeimobShopexpressOrderGetOrderListResponse]) {
	response = api.CreateResponse[WeimobShopexpressOrderGetOrderListResponse](&WeimobShopexpressOrderGetOrderListResponse{})
	return
}

type WeimobShopexpressOrderGetOrderListResponseOrderItemListOutList struct {
	OrderItemNo    string `json:"orderItemNo,omitempty"`
	GoodsCode      string `json:"goodsCode,omitempty"`
	GoodsName      string `json:"goodsName,omitempty"`
	SkuCode        string `json:"skuCode,omitempty"`
	SkuName        string `json:"skuName,omitempty"`
	ImageUrl       string `json:"imageUrl,omitempty"`
	Count          int64  `json:"count,omitempty"`
	Price          string `json:"price,omitempty"`
	Amount         string `json:"amount,omitempty"`
	PayAmount      string `json:"payAmount,omitempty"`
	CreateTime     string `json:"createTime,omitempty"`
	CustomSkuCode  string `json:"customSkuCode,omitempty"`
	RefundStatus   int64  `json:"refundStatus,omitempty"`
	BarCode        string `json:"barCode,omitempty"`
	FreightAmount  string `json:"freightAmount,omitempty"`
	DiscountAmount string `json:"discountAmount,omitempty"`
	ExciseTax      string `json:"exciseTax,omitempty"`
	LogisticalTax  string `json:"logisticalTax,omitempty"`
}
