package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrdersourceDayGetList
// @id 2276
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2276?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=江南布衣交易访问来源数据接口_日粒度)
func (client *Client) OrdersourceDayGetList(request *OrdersourceDayGetListRequest) (response *OrdersourceDayGetListResponse, err error) {
	rpcResponse := CreateOrdersourceDayGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrdersourceDayGetListRequest struct {
	*api.BaseRequest

	BosId    int64   `json:"bosId,omitempty"`
	EndDd    string  `json:"endDd,omitempty"`
	PageNum  int64   `json:"pageNum,omitempty"`
	PageSize int64   `json:"pageSize,omitempty"`
	StartDd  string  `json:"startDd,omitempty"`
	StatType []int64 `json:"statType,omitempty"`
}

type OrdersourceDayGetListResponse struct {
	List      []OrdersourceDayGetListResponselist `json:"list,omitempty"`
	PageNum   int64                               `json:"pageNum,omitempty"`
	PageSize  int64                               `json:"pageSize,omitempty"`
	TotalSize int64                               `json:"totalSize,omitempty"`
}

func CreateOrdersourceDayGetListRequest() (request *OrdersourceDayGetListRequest) {
	request = &OrdersourceDayGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "ordersource/day/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateOrdersourceDayGetListResponse() (response *api.BaseResponse[OrdersourceDayGetListResponse]) {
	response = api.CreateResponse[OrdersourceDayGetListResponse](&OrdersourceDayGetListResponse{})
	return
}

type OrdersourceDayGetListResponselist struct {
	Uv           int64  `json:"uv,omitempty"`
	StatType     int64  `json:"statType,omitempty"`
	StartDd      string `json:"startDd,omitempty"`
	StoreId      int64  `json:"storeId,omitempty"`
	PayCnt       int64  `json:"payCnt,omitempty"`
	PayAmt       string `json:"payAmt,omitempty"`
	PayVisitRate string `json:"payVisitRate,omitempty"`
	DateType     int64  `json:"dateType,omitempty"`
	AreaId       string `json:"areaId,omitempty"`
	PayUsrCnt    int64  `json:"payUsrCnt,omitempty"`
	AreaName     string `json:"areaName,omitempty"`
	OrdCnt       int64  `json:"ordCnt,omitempty"`
	BosId        int64  `json:"bosId,omitempty"`
	StatName     string `json:"statName,omitempty"`
	EndDd        string `json:"endDd,omitempty"`
	StoreName    string `json:"storeName,omitempty"`
	Brand        string `json:"brand,omitempty"`
	StoreCode    string `json:"storeCode,omitempty"`
}
