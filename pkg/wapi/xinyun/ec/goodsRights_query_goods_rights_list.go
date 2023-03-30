package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsRightsQueryGoodsRightsList
// @id 1325
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询商品服务列表)
func (client *Client) GoodsRightsQueryGoodsRightsList(request *GoodsRightsQueryGoodsRightsListRequest) (response *GoodsRightsQueryGoodsRightsListResponse, err error) {
	rpcResponse := CreateGoodsRightsQueryGoodsRightsListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsRightsQueryGoodsRightsListRequest struct {
	*api.BaseRequest
}

type GoodsRightsQueryGoodsRightsListResponse struct {
}

func CreateGoodsRightsQueryGoodsRightsListRequest() (request *GoodsRightsQueryGoodsRightsListRequest) {
	request = &GoodsRightsQueryGoodsRightsListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goodsRights/queryGoodsRightsList", "api")
	request.Method = api.POST
	return
}

func CreateGoodsRightsQueryGoodsRightsListResponse() (response *api.BaseResponse[GoodsRightsQueryGoodsRightsListResponse]) {
	response = api.CreateResponse[GoodsRightsQueryGoodsRightsListResponse](&GoodsRightsQueryGoodsRightsListResponse{})
	return
}
