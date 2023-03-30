package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ProductDown
// @id 2666
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=下架产品)
func (client *Client) ProductDown(request *ProductDownRequest) (response *ProductDownResponse, err error) {
	rpcResponse := CreateProductDownResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ProductDownRequest struct {
	*api.BaseRequest

	ProductUniqueRequestList []map[string]any `json:"productUniqueRequestList,omitempty"`
	UserWid                  int64            `json:"userWid,omitempty"`
	NewClassificationId      int64            `json:"newClassificationId,omitempty"`
}

type ProductDownResponse struct {
}

func CreateProductDownRequest() (request *ProductDownRequest) {
	request = &ProductDownRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "product/down", "api")
	request.Method = api.POST
	return
}

func CreateProductDownResponse() (response *api.BaseResponse[ProductDownResponse]) {
	response = api.CreateResponse[ProductDownResponse](&ProductDownResponse{})
	return
}
