package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// OrderBatchEncrypt
// @id 1861
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=批量加密)
func (client *Client) OrderBatchEncrypt(request *OrderBatchEncryptRequest) (response *OrderBatchEncryptResponse, err error) {
	rpcResponse := CreateOrderBatchEncryptResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type OrderBatchEncryptRequest struct {
	*api.BaseRequest

	BatchEncrypts OrderBatchEncryptRequestBatchEncrypts `json:"batchEncrypts,omitempty"`
}

type OrderBatchEncryptResponse struct {
}

func CreateOrderBatchEncryptRequest() (request *OrderBatchEncryptRequest) {
	request = &OrderBatchEncryptRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "order/batchEncrypt", "api")
	request.Method = api.POST
	return
}

func CreateOrderBatchEncryptResponse() (response *api.BaseResponse[OrderBatchEncryptResponse]) {
	response = api.CreateResponse[OrderBatchEncryptResponse](&OrderBatchEncryptResponse{})
	return
}

type OrderBatchEncryptRequestBatchEncrypts struct {
	PlainText      string `json:"plainText,omitempty"`
	OrderNo        int64  `json:"orderNo,omitempty"`
	IsSupportIndex bool   `json:"isSupportIndex,omitempty"`
	SensitiveType  int    `json:"sensitiveType,omitempty"`
}
