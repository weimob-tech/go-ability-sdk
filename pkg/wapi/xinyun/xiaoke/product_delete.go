package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ProductDelete
// @id 2658
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除产品)
func (client *Client) ProductDelete(request *ProductDeleteRequest) (response *ProductDeleteResponse, err error) {
	rpcResponse := CreateProductDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ProductDeleteRequest struct {
	*api.BaseRequest

	UserWid         int64  `json:"userWid,omitempty"`
	ProductUniqueNo string `json:"productUniqueNo,omitempty"`
	Version         int    `json:"version,omitempty"`
}

type ProductDeleteResponse struct {
	SuccessCount string `json:"successCount,omitempty"`
	FailedCount  string `json:"failedCount,omitempty"`
}

func CreateProductDeleteRequest() (request *ProductDeleteRequest) {
	request = &ProductDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "product/delete", "api")
	request.Method = api.POST
	return
}

func CreateProductDeleteResponse() (response *api.BaseResponse[ProductDeleteResponse]) {
	response = api.CreateResponse[ProductDeleteResponse](&ProductDeleteResponse{})
	return
}
