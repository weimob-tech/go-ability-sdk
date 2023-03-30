package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrdersourceMonthGetList
// @id 2274
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2274?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=江南布衣交易访问来源数据接口_月粒度)
func (client *Client) OrdersourceMonthGetList(request *OrdersourceMonthGetListRequest) (response *OrdersourceMonthGetListResponse, err error) {
	rpcResponse := CreateOrdersourceMonthGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrdersourceMonthGetListRequest struct {
	*api.BaseRequest

	BosId    int64   `json:"bosId,omitempty"`
	EndDd    string  `json:"endDd,omitempty"`
	PageNum  int64   `json:"pageNum,omitempty"`
	PageSize int64   `json:"pageSize,omitempty"`
	StartDd  string  `json:"startDd,omitempty"`
	StatType []int64 `json:"statType,omitempty"`
}

type OrdersourceMonthGetListResponse struct {
	List      []OrdersourceMonthGetListResponselist `json:"list,omitempty"`
	PageNum   int64                                 `json:"pageNum,omitempty"`
	PageSize  int64                                 `json:"pageSize,omitempty"`
	TotalSize int64                                 `json:"totalSize,omitempty"`
}

func CreateOrdersourceMonthGetListRequest() (request *OrdersourceMonthGetListRequest) {
	request = &OrdersourceMonthGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "ordersource/month/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateOrdersourceMonthGetListResponse() (response *api.BaseResponse[OrdersourceMonthGetListResponse]) {
	response = api.CreateResponse[OrdersourceMonthGetListResponse](&OrdersourceMonthGetListResponse{})
	return
}

type OrdersourceMonthGetListResponselist struct {
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
