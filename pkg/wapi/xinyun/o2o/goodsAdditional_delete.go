package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsAdditionalDelete
// @id 151
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除加值)
func (client *Client) GoodsAdditionalDelete(request *GoodsAdditionalDeleteRequest) (response *GoodsAdditionalDeleteResponse, err error) {
	rpcResponse := CreateGoodsAdditionalDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsAdditionalDeleteRequest struct {
	*api.BaseRequest

	MerchantId        int64 `json:"merchantId,omitempty"`
	StoreId           int64 `json:"storeId,omitempty"`
	GoodsAdditionalId int64 `json:"goodsAdditionalId,omitempty"`
}

type GoodsAdditionalDeleteResponse struct {
}

func CreateGoodsAdditionalDeleteRequest() (request *GoodsAdditionalDeleteRequest) {
	request = &GoodsAdditionalDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goodsAdditional/delete", "api")
	request.Method = api.POST
	return
}

func CreateGoodsAdditionalDeleteResponse() (response *api.BaseResponse[GoodsAdditionalDeleteResponse]) {
	response = api.CreateResponse[GoodsAdditionalDeleteResponse](&GoodsAdditionalDeleteResponse{})
	return
}
