package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderGetList
// @id 1973
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1973?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=订单列表)
func (client *Client) OrderGetList(request *OrderGetListRequest) (response *OrderGetListResponse, err error) {
	rpcResponse := CreateOrderGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderGetListRequest struct {
	*api.BaseRequest

	EndTime             string  `json:"endTime,omitempty"`
	LogisticsStatusList []int64 `json:"logisticsStatusList,omitempty"`
	PageIndex           int64   `json:"pageIndex,omitempty"`
	PageSize            int64   `json:"pageSize,omitempty"`
	PayStatus           int64   `json:"payStatus,omitempty"`
	StartTime           string  `json:"startTime,omitempty"`
}

type OrderGetListResponse struct {
	Items      []OrderGetListResponseItems `json:"items,omitempty"`
	TotalCount int64                       `json:"totalCount,omitempty"`
}

func CreateOrderGetListRequest() (request *OrderGetListRequest) {
	request = &OrderGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "order/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateOrderGetListResponse() (response *api.BaseResponse[OrderGetListResponse]) {
	response = api.CreateResponse[OrderGetListResponse](&OrderGetListResponse{})
	return
}

type OrderGetListResponseItems struct {
	Uid                 int64  `json:"uid,omitempty"`
	UserName            string `json:"userName,omitempty"`
	Phone               string `json:"phone,omitempty"`
	Email               string `json:"email,omitempty"`
	OrderNo             string `json:"orderNo,omitempty"`
	Currency            string `json:"currency,omitempty"`
	TimeZone            string `json:"timeZone,omitempty"`
	ItemCount           int64  `json:"itemCount,omitempty"`
	Amount              string `json:"amount,omitempty"`
	PayAmount           string `json:"payAmount,omitempty"`
	OrderStatus         int64  `json:"orderStatus,omitempty"`
	PayStatus           int64  `json:"payStatus,omitempty"`
	PayStatusText       string `json:"payStatusText,omitempty"`
	LogisticsStatus     int64  `json:"logisticsStatus,omitempty"`
	LogisticsStatusText string `json:"logisticsStatusText,omitempty"`
	PayTime             string `json:"payTime,omitempty"`
	CloseTime           string `json:"closeTime,omitempty"`
	Freight             string `json:"freight,omitempty"`
	DiscountAmount      string `json:"discountAmount,omitempty"`
	ExciseTax           string `json:"exciseTax,omitempty"`
	ConsigneeName       string `json:"consigneeName,omitempty"`
	ConsigneeAddress    string `json:"consigneeAddress,omitempty"`
	CreateTime          string `json:"createTime,omitempty"`
	LogisticalTax       string `json:"logisticalTax,omitempty"`
}
