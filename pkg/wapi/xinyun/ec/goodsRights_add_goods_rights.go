package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsRightsAddGoodsRights
// @id 1327
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加商品服务)
func (client *Client) GoodsRightsAddGoodsRights(request *GoodsRightsAddGoodsRightsRequest) (response *GoodsRightsAddGoodsRightsResponse, err error) {
	rpcResponse := CreateGoodsRightsAddGoodsRightsResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsRightsAddGoodsRightsRequest struct {
	*api.BaseRequest

	Title            string   `json:"title,omitempty"`
	GoodsRightsValue []string `json:"goodsRightsValue,omitempty"`
}

type GoodsRightsAddGoodsRightsResponse struct {
}

func CreateGoodsRightsAddGoodsRightsRequest() (request *GoodsRightsAddGoodsRightsRequest) {
	request = &GoodsRightsAddGoodsRightsRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goodsRights/addGoodsRights", "api")
	request.Method = api.POST
	return
}

func CreateGoodsRightsAddGoodsRightsResponse() (response *api.BaseResponse[GoodsRightsAddGoodsRightsResponse]) {
	response = api.CreateResponse[GoodsRightsAddGoodsRightsResponse](&GoodsRightsAddGoodsRightsResponse{})
	return
}
