package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderBatchSensitive
// @id 1863
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量脱敏)
func (client *Client) OrderBatchSensitive(request *OrderBatchSensitiveRequest) (response *OrderBatchSensitiveResponse, err error) {
	rpcResponse := CreateOrderBatchSensitiveResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderBatchSensitiveRequest struct {
	*api.BaseRequest

	CipherInfos []OrderBatchSensitiveRequestCipherInfos `json:"cipherInfos,omitempty"`
}

type OrderBatchSensitiveResponse struct {
}

func CreateOrderBatchSensitiveRequest() (request *OrderBatchSensitiveRequest) {
	request = &OrderBatchSensitiveRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "order/batchSensitive", "api")
	request.Method = api.POST
	return
}

func CreateOrderBatchSensitiveResponse() (response *api.BaseResponse[OrderBatchSensitiveResponse]) {
	response = api.CreateResponse[OrderBatchSensitiveResponse](&OrderBatchSensitiveResponse{})
	return
}

type OrderBatchSensitiveRequestCipherInfos struct {
	OrderNo    int64  `json:"orderNo,omitempty"`
	CipherText string `json:"cipherText,omitempty"`
}
