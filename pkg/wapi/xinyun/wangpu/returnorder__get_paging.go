package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ReturnorderGetPaging
// @id 35
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询维权单详情)
func (client *Client) ReturnorderGetPaging(request *ReturnorderGetPagingRequest) (response *ReturnorderGetPagingResponse, err error) {
	rpcResponse := CreateReturnorderGetPagingResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ReturnorderGetPagingRequest struct {
	*api.BaseRequest

	ReturnOrderStatus int    `json:"return_order_status,omitempty"`
	ShopId            string `json:"shop_id,omitempty"`
	ReturnOrderNo     string `json:"return_order_no,omitempty"`
	CreateBeginTime   string `json:"create_begin_time,omitempty"`
	CreateEndTime     string `json:"create_end_time,omitempty"`
	UpdateBeginTime   string `json:"update_begin_time,omitempty"`
	UpdateEndTime     string `json:"update_end_time,omitempty"`
	PageSize          int    `json:"page_size,omitempty"`
	PageNo            int    `json:"page_no,omitempty"`
}

type ReturnorderGetPagingResponse struct {
}

func CreateReturnorderGetPagingRequest() (request *ReturnorderGetPagingRequest) {
	request = &ReturnorderGetPagingRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "returnorder/GetPaging", "api")
	request.Method = api.POST
	return
}

func CreateReturnorderGetPagingResponse() (response *api.BaseResponse[ReturnorderGetPagingResponse]) {
	response = api.CreateResponse[ReturnorderGetPagingResponse](&ReturnorderGetPagingResponse{})
	return
}
