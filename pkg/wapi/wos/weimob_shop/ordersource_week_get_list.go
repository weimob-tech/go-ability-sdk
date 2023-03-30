package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrdersourceWeekGetList
// @id 2273
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2273?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=江南布衣交易访问来源数据接口_周粒度)
func (client *Client) OrdersourceWeekGetList(request *OrdersourceWeekGetListRequest) (response *OrdersourceWeekGetListResponse, err error) {
	rpcResponse := CreateOrdersourceWeekGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrdersourceWeekGetListRequest struct {
	*api.BaseRequest

	BosId    int64   `json:"bosId,omitempty"`
	EndDd    string  `json:"endDd,omitempty"`
	PageNum  int64   `json:"pageNum,omitempty"`
	PageSize int64   `json:"pageSize,omitempty"`
	StartDd  string  `json:"startDd,omitempty"`
	StatType []int64 `json:"statType,omitempty"`
}

type OrdersourceWeekGetListResponse struct {
	List      []OrdersourceWeekGetListResponselist `json:"list,omitempty"`
	PageNum   int64                                `json:"pageNum,omitempty"`
	PageSize  int64                                `json:"pageSize,omitempty"`
	TotalSize int64                                `json:"totalSize,omitempty"`
}

func CreateOrdersourceWeekGetListRequest() (request *OrdersourceWeekGetListRequest) {
	request = &OrdersourceWeekGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "ordersource/week/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateOrdersourceWeekGetListResponse() (response *api.BaseResponse[OrdersourceWeekGetListResponse]) {
	response = api.CreateResponse[OrdersourceWeekGetListResponse](&OrdersourceWeekGetListResponse{})
	return
}

type OrdersourceWeekGetListResponselist struct {
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
