package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// LogsOrder
// @id 3069
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询订单跟进记录（下线）)
func (client *Client) LogsOrder(request *LogsOrderRequest) (response *LogsOrderResponse, err error) {
	rpcResponse := CreateLogsOrderResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type LogsOrderRequest struct {
	*api.BaseRequest
}

type LogsOrderResponse struct {
}

func CreateLogsOrderRequest() (request *LogsOrderRequest) {
	request = &LogsOrderRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "logs/order", "api")
	request.Method = api.POST
	return
}

func CreateLogsOrderResponse() (response *api.BaseResponse[LogsOrderResponse]) {
	response = api.CreateResponse[LogsOrderResponse](&LogsOrderResponse{})
	return
}
