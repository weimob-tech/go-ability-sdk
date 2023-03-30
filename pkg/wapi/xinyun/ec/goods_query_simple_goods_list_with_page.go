package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsQuerySimpleGoodsListWithPage
// @id 1161
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询商品列表-获取外部编码)
func (client *Client) GoodsQuerySimpleGoodsListWithPage(request *GoodsQuerySimpleGoodsListWithPageRequest) (response *GoodsQuerySimpleGoodsListWithPageResponse, err error) {
	rpcResponse := CreateGoodsQuerySimpleGoodsListWithPageResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsQuerySimpleGoodsListWithPageRequest struct {
	*api.BaseRequest

	PageSize     int   `json:"pageSize,omitempty"`
	StartGoodsId int64 `json:"startGoodsId,omitempty"`
	StoreId      int64 `json:"storeId,omitempty"`
}

type GoodsQuerySimpleGoodsListWithPageResponse struct {
}

func CreateGoodsQuerySimpleGoodsListWithPageRequest() (request *GoodsQuerySimpleGoodsListWithPageRequest) {
	request = &GoodsQuerySimpleGoodsListWithPageRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goods/querySimpleGoodsListWithPage", "api")
	request.Method = api.POST
	return
}

func CreateGoodsQuerySimpleGoodsListWithPageResponse() (response *api.BaseResponse[GoodsQuerySimpleGoodsListWithPageResponse]) {
	response = api.CreateResponse[GoodsQuerySimpleGoodsListWithPageResponse](&GoodsQuerySimpleGoodsListWithPageResponse{})
	return
}
