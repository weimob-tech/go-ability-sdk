package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsTagList
// @id 75
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询菜品标签列表)
func (client *Client) GoodsTagList(request *GoodsTagListRequest) (response *GoodsTagListResponse, err error) {
	rpcResponse := CreateGoodsTagListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsTagListRequest struct {
	*api.BaseRequest

	StoreId    int64 `json:"storeId,omitempty"`
	GoodsTagId int64 `json:"goodsTagId,omitempty"`
	MerchantId int64 `json:"merchantId,omitempty"`
}

type GoodsTagListResponse struct {
}

func CreateGoodsTagListRequest() (request *GoodsTagListRequest) {
	request = &GoodsTagListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "goodsTag/list", "api")
	request.Method = api.POST
	return
}

func CreateGoodsTagListResponse() (response *api.BaseResponse[GoodsTagListResponse]) {
	response = api.CreateResponse[GoodsTagListResponse](&GoodsTagListResponse{})
	return
}
