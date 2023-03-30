package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ServicesOrderList
// @id 1052
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询订单列表)
func (client *Client) ServicesOrderList(request *ServicesOrderListRequest) (response *ServicesOrderListResponse, err error) {
	rpcResponse := CreateServicesOrderListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ServicesOrderListRequest struct {
	*api.BaseRequest

	Page            ServicesOrderListRequestPage `json:"page,omitempty"`
	EndCreateTime   int64                        `json:"endCreateTime,omitempty"`
	OrderStatus     int64                        `json:"orderStatus,omitempty"`
	PageIndex       int64                        `json:"pageIndex,omitempty"`
	PageSize        int64                        `json:"pageSize,omitempty"`
	PayStatus       int64                        `json:"payStatus,omitempty"`
	StartCreateTime int64                        `json:"startCreateTime,omitempty"`
	StoreId         int64                        `json:"storeId,omitempty"`
}

type ServicesOrderListResponse struct {
}

func CreateServicesOrderListRequest() (request *ServicesOrderListRequest) {
	request = &ServicesOrderListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "servicesOrder/list", "api")
	request.Method = api.POST
	return
}

func CreateServicesOrderListResponse() (response *api.BaseResponse[ServicesOrderListResponse]) {
	response = api.CreateResponse[ServicesOrderListResponse](&ServicesOrderListResponse{})
	return
}

type ServicesOrderListRequestPage struct {
}
