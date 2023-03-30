package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ProductUp
// @id 2656
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=上架产品)
func (client *Client) ProductUp(request *ProductUpRequest) (response *ProductUpResponse, err error) {
	rpcResponse := CreateProductUpResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ProductUpRequest struct {
	*api.BaseRequest

	ProductUniqueRequestList []map[string]any `json:"productUniqueRequestList,omitempty"`
	UserWid                  int64            `json:"userWid,omitempty"`
	NewClassificationId      int64            `json:"newClassificationId,omitempty"`
}

type ProductUpResponse struct {
}

func CreateProductUpRequest() (request *ProductUpRequest) {
	request = &ProductUpRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "product/up", "api")
	request.Method = api.POST
	return
}

func CreateProductUpResponse() (response *api.BaseResponse[ProductUpResponse]) {
	response = api.CreateResponse[ProductUpResponse](&ProductUpResponse{})
	return
}
