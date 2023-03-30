package tour

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderExport
// @id 1005
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=导出出行人列表)
func (client *Client) OrderExport(request *OrderExportRequest) (response *OrderExportResponse, err error) {
	rpcResponse := CreateOrderExportResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderExportRequest struct {
	*api.BaseRequest

	OrderStatus int    `json:"orderStatus,omitempty"`
	OrderSource int    `json:"orderSource,omitempty"`
	TimeType    int    `json:"timeType,omitempty"`
	StartTime   int    `json:"startTime,omitempty"`
	EndTime     int    `json:"endTime,omitempty"`
	Keyword     int    `json:"keyword,omitempty"`
	Value       string `json:"value,omitempty"`
	PageIndex   int    `json:"pageIndex,omitempty"`
	PageSize    int    `json:"pageSize,omitempty"`
}

type OrderExportResponse struct {
}

func CreateOrderExportRequest() (request *OrderExportRequest) {
	request = &OrderExportRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("tour", "1_0", "order/export", "api")
	request.Method = api.POST
	return
}

func CreateOrderExportResponse() (response *api.BaseResponse[OrderExportResponse]) {
	response = api.CreateResponse[OrderExportResponse](&OrderExportResponse{})
	return
}
