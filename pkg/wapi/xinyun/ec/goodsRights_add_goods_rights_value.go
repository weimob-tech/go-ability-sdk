package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsRightsAddGoodsRightsValue
// @id 1326
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加商品服务值)
func (client *Client) GoodsRightsAddGoodsRightsValue(request *GoodsRightsAddGoodsRightsValueRequest) (response *GoodsRightsAddGoodsRightsValueResponse, err error) {
	rpcResponse := CreateGoodsRightsAddGoodsRightsValueResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsRightsAddGoodsRightsValueRequest struct {
	*api.BaseRequest

	GoodsRightsId        int64    `json:"goodsRightsId,omitempty"`
	GoodsRightValueTitle []string `json:"goodsRightValueTitle,omitempty"`
}

type GoodsRightsAddGoodsRightsValueResponse struct {
}

func CreateGoodsRightsAddGoodsRightsValueRequest() (request *GoodsRightsAddGoodsRightsValueRequest) {
	request = &GoodsRightsAddGoodsRightsValueRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goodsRights/addGoodsRightsValue", "api")
	request.Method = api.POST
	return
}

func CreateGoodsRightsAddGoodsRightsValueResponse() (response *api.BaseResponse[GoodsRightsAddGoodsRightsValueResponse]) {
	response = api.CreateResponse[GoodsRightsAddGoodsRightsValueResponse](&GoodsRightsAddGoodsRightsValueResponse{})
	return
}
