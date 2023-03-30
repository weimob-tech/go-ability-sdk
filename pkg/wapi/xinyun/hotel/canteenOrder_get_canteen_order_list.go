package hotel

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CanteenOrderGetCanteenOrderList
// @id 1435
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=堂食订单列表)
func (client *Client) CanteenOrderGetCanteenOrderList(request *CanteenOrderGetCanteenOrderListRequest) (response *CanteenOrderGetCanteenOrderListResponse, err error) {
	rpcResponse := CreateCanteenOrderGetCanteenOrderListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CanteenOrderGetCanteenOrderListRequest struct {
	*api.BaseRequest

	From      int    `json:"from,omitempty"`
	To        int    `json:"to,omitempty"`
	Status    int    `json:"status,omitempty"`
	Keyword   string `json:"keyword,omitempty"`
	PageIndex int    `json:"pageIndex,omitempty"`
	PageSize  int    `json:"PageSize,omitempty"`
}

type CanteenOrderGetCanteenOrderListResponse struct {
}

func CreateCanteenOrderGetCanteenOrderListRequest() (request *CanteenOrderGetCanteenOrderListRequest) {
	request = &CanteenOrderGetCanteenOrderListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("hotel", "1_0", "canteenOrder/getCanteenOrderList", "api")
	request.Method = api.POST
	return
}

func CreateCanteenOrderGetCanteenOrderListResponse() (response *api.BaseResponse[CanteenOrderGetCanteenOrderListResponse]) {
	response = api.CreateResponse[CanteenOrderGetCanteenOrderListResponse](&CanteenOrderGetCanteenOrderListResponse{})
	return
}
