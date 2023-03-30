package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderGet
// @id 1976
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1976?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=订单详情)
func (client *Client) OrderGet(request *OrderGetRequest) (response *OrderGetResponse, err error) {
	rpcResponse := CreateOrderGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderGetRequest struct {
	*api.BaseRequest

	OrderNo string `json:"orderNo,omitempty"`
	Uid     int64  `json:"uid,omitempty"`
}

type OrderGetResponse struct {
	OrderItemDetailOutList   []OrderGetResponseOrderItemDetailOutList   `json:"orderItemDetailOutList,omitempty"`
	AddressContext           OrderGetResponseAddressContext             `json:"addressContext,omitempty"`
	OrderDeliveryInfoOutList []OrderGetResponseOrderDeliveryInfoOutList `json:"orderDeliveryInfoOutList,omitempty"`
	Uid                      int64                                      `json:"uid,omitempty"`
	OrderNo                  string                                     `json:"orderNo,omitempty"`
	UserName                 string                                     `json:"userName,omitempty"`
	Phone                    string                                     `json:"phone,omitempty"`
	Email                    string                                     `json:"email,omitempty"`
	ItemCount                int64                                      `json:"itemCount,omitempty"`
	Currency                 string                                     `json:"currency,omitempty"`
	Amount                   string                                     `json:"amount,omitempty"`
	ExciseTax                string                                     `json:"exciseTax,omitempty"`
	LogisticalTax            string                                     `json:"logisticalTax,omitempty"`
	DiscountAmount           string                                     `json:"discountAmount,omitempty"`
	PayAmount                string                                     `json:"payAmount,omitempty"`
	OrderStatus              int64                                      `json:"orderStatus,omitempty"`
	OrderType                int64                                      `json:"orderType,omitempty"`
	PayStatus                int64                                      `json:"payStatus,omitempty"`
	PayStatusText            string                                     `json:"payStatusText,omitempty"`
	PayTime                  string                                     `json:"payTime,omitempty"`
	LogisticsStatus          int64                                      `json:"logisticsStatus,omitempty"`
	CreateTime               string                                     `json:"createTime,omitempty"`
	Remark                   string                                     `json:"remark,omitempty"`
	CloseTime                string                                     `json:"closeTime,omitempty"`
	RealFreight              string                                     `json:"realFreight,omitempty"`
	TaxSum                   string                                     `json:"taxSum,omitempty"`
}

func CreateOrderGetRequest() (request *OrderGetRequest) {
	request = &OrderGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "order/get", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderGetResponse() (response *api.BaseResponse[OrderGetResponse]) {
	response = api.CreateResponse[OrderGetResponse](&OrderGetResponse{})
	return
}

type OrderGetResponseOrderItemDetailOutList struct {
	CustomizeInfo  OrderGetResponseCustomizeInfo `json:"customizeInfo,omitempty"`
	Uid            int64                         `json:"uid,omitempty"`
	OrderNo        string                        `json:"orderNo,omitempty"`
	OrderItemNo    string                        `json:"orderItemNo,omitempty"`
	GoodsCode      string                        `json:"goodsCode,omitempty"`
	GoodsName      string                        `json:"goodsName,omitempty"`
	SkuCode        string                        `json:"skuCode,omitempty"`
	CustomSkuCode  string                        `json:"customSkuCode,omitempty"`
	BarCode        string                        `json:"barCode,omitempty"`
	Count          int64                         `json:"count,omitempty"`
	Price          string                        `json:"price,omitempty"`
	Amount         string                        `json:"amount,omitempty"`
	FreightAmount  string                        `json:"freightAmount,omitempty"`
	ExciseTax      string                        `json:"exciseTax,omitempty"`
	LogisticalTax  string                        `json:"logisticalTax,omitempty"`
	CreateTime     string                        `json:"createTime,omitempty"`
	ImageUrl       string                        `json:"imageUrl,omitempty"`
	SkuName        string                        `json:"skuName,omitempty"`
	PayAmount      string                        `json:"payAmount,omitempty"`
	DiscountAmount string                        `json:"discountAmount,omitempty"`
}

type OrderGetResponseCustomizeInfo struct {
	CustomizeContentList []OrderGetResponseCustomizeContentList `json:"customizeContentList,omitempty"`
	UserCustomizeCode    string                                 `json:"userCustomizeCode,omitempty"`
	Price                string                                 `json:"price,omitempty"`
}

type OrderGetResponseCustomizeContentList struct {
	Content OrderGetResponseContent `json:"content,omitempty"`
	Type    string                  `json:"type,omitempty"`
	Title   string                  `json:"title,omitempty"`
}

type OrderGetResponseContent struct {
	OptionName string `json:"optionName,omitempty"`
	PicShape   string `json:"picShape,omitempty"`
	Text       string `json:"text,omitempty"`
	Font       string `json:"font,omitempty"`
	Color      string `json:"color,omitempty"`
}

type OrderGetResponseAddressContext struct {
	Address         string `json:"address,omitempty"`
	ApartmentNumber string `json:"apartmentNumber,omitempty"`
	City            string `json:"city,omitempty"`
	Company         string `json:"company,omitempty"`
	Country         string `json:"country,omitempty"`
	CountryISO      string `json:"countryISO,omitempty"`
	PostCode        string `json:"postCode,omitempty"`
	Province        string `json:"province,omitempty"`
	ProvinceISO     string `json:"provinceISO,omitempty"`
	FirstName       string `json:"firstName,omitempty"`
	LastName        string `json:"lastName,omitempty"`
	Phone           string `json:"phone,omitempty"`
}

type OrderGetResponseOrderDeliveryInfoOutList struct {
	OrderDeliveryItemOutList []OrderGetResponseOrderDeliveryItemOutList `json:"orderDeliveryItemOutList,omitempty"`
	DeliveryOrderCode        string                                     `json:"deliveryOrderCode,omitempty"`
	ExpressNo                string                                     `json:"expressNo,omitempty"`
	ExpressCompanyCode       string                                     `json:"expressCompanyCode,omitempty"`
	ExpressCompany           string                                     `json:"expressCompany,omitempty"`
	TotalCount               int64                                      `json:"totalCount,omitempty"`
	DeliveryTime             string                                     `json:"deliveryTime,omitempty"`
}

type OrderGetResponseOrderDeliveryItemOutList struct {
	GoodsName     string `json:"goodsName,omitempty"`
	SkuName       string `json:"skuName,omitempty"`
	GoodsCode     string `json:"goodsCode,omitempty"`
	SkuCode       string `json:"skuCode,omitempty"`
	OrderItemNo   string `json:"orderItemNo,omitempty"`
	Price         string `json:"price,omitempty"`
	ImageUrl      string `json:"imageUrl,omitempty"`
	Count         int64  `json:"count,omitempty"`
	Amount        int64  `json:"amount,omitempty"`
	RefundStatus  int64  `json:"refundStatus,omitempty"`
	CustomSkuCode string `json:"customSkuCode,omitempty"`
}
