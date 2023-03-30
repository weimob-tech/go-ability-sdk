package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderBatchDecrypt
// @id 1862
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量解密)
func (client *Client) OrderBatchDecrypt(request *OrderBatchDecryptRequest) (response *OrderBatchDecryptResponse, err error) {
	rpcResponse := CreateOrderBatchDecryptResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderBatchDecryptRequest struct {
	*api.BaseRequest

	CipherInfos []OrderBatchDecryptRequestCipherInfos `json:"cipherInfos,omitempty"`
}

type OrderBatchDecryptResponse struct {
}

func CreateOrderBatchDecryptRequest() (request *OrderBatchDecryptRequest) {
	request = &OrderBatchDecryptRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "order/batchDecrypt", "api")
	request.Method = api.POST
	return
}

func CreateOrderBatchDecryptResponse() (response *api.BaseResponse[OrderBatchDecryptResponse]) {
	response = api.CreateResponse[OrderBatchDecryptResponse](&OrderBatchDecryptResponse{})
	return
}

type OrderBatchDecryptRequestCipherInfos struct {
	OrderNo    int64  `json:"orderNo,omitempty"`
	CipherText string `json:"cipherText,omitempty"`
}
