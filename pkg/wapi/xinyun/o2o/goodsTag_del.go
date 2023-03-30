package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsTagDel
// @id 78
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除标签)
func (client *Client) GoodsTagDel(request *GoodsTagDelRequest) (response *GoodsTagDelResponse, err error) {
	rpcResponse := CreateGoodsTagDelResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsTagDelRequest struct {
	*api.BaseRequest

	GoodsTagId int64 `json:"goodsTagId,omitempty"`
	StoreId    int64 `json:"storeId,omitempty"`
	MerchantId int64 `json:"merchantId,omitempty"`
}

type GoodsTagDelResponse struct {
}

func CreateGoodsTagDelRequest() (request *GoodsTagDelRequest) {
	request = &GoodsTagDelRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goodsTag/del", "api")
	request.Method = api.POST
	return
}

func CreateGoodsTagDelResponse() (response *api.BaseResponse[GoodsTagDelResponse]) {
	response = api.CreateResponse[GoodsTagDelResponse](&GoodsTagDelResponse{})
	return
}
