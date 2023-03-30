package kld

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// IOrderServiceQueryBills
// @id 224
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取收款记录)
func (client *Client) IOrderServiceQueryBills(request *IOrderServiceQueryBillsRequest) (response *IOrderServiceQueryBillsResponse, err error) {
	rpcResponse := CreateIOrderServiceQueryBillsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type IOrderServiceQueryBillsRequest struct {
	*api.BaseRequest

	Data          IOrderServiceQueryBillsRequestData `json:"data,omitempty"`
	CurrentUserId string                             `json:"currentUserId,omitempty"`
	StartTime     int64                              `json:"startTime,omitempty"`
	EndTime       int64                              `json:"endTime,omitempty"`
	StoreId       int64                              `json:"storeId,omitempty"`
	CheckstandId  int                                `json:"checkstandId,omitempty"`
	TagId         int                                `json:"tagId,omitempty"`
	Identity      int                                `json:"identity,omitempty"`
	Sort          []string                           `json:"sort,omitempty"`
	Order         string                             `json:"order,omitempty"`
	PageIndex     int                                `json:"pageIndex,omitempty"`
	PageSize      int                                `json:"pageSize,omitempty"`
}

type IOrderServiceQueryBillsResponse struct {
}

func CreateIOrderServiceQueryBillsRequest() (request *IOrderServiceQueryBillsRequest) {
	request = &IOrderServiceQueryBillsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("kld", "1_0", "IOrderService/queryBills", "api")
	request.Method = api.POST
	return
}

func CreateIOrderServiceQueryBillsResponse() (response *api.BaseResponse[IOrderServiceQueryBillsResponse]) {
	response = api.CreateResponse[IOrderServiceQueryBillsResponse](&IOrderServiceQueryBillsResponse{})
	return
}

type IOrderServiceQueryBillsRequestData struct {
}
