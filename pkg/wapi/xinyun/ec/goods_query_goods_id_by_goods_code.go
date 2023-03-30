package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsQueryGoodsIdByGoodsCode
// @id 1685
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=根据编码获取商品id)
func (client *Client) GoodsQueryGoodsIdByGoodsCode(request *GoodsQueryGoodsIdByGoodsCodeRequest) (response *GoodsQueryGoodsIdByGoodsCodeResponse, err error) {
	rpcResponse := CreateGoodsQueryGoodsIdByGoodsCodeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsQueryGoodsIdByGoodsCodeRequest struct {
	*api.BaseRequest

	OuterGoodsCodeList []map[string]any `json:"outerGoodsCodeList,omitempty"`
}

type GoodsQueryGoodsIdByGoodsCodeResponse struct {
}

func CreateGoodsQueryGoodsIdByGoodsCodeRequest() (request *GoodsQueryGoodsIdByGoodsCodeRequest) {
	request = &GoodsQueryGoodsIdByGoodsCodeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goods/queryGoodsIdByGoodsCode", "api")
	request.Method = api.POST
	return
}

func CreateGoodsQueryGoodsIdByGoodsCodeResponse() (response *api.BaseResponse[GoodsQueryGoodsIdByGoodsCodeResponse]) {
	response = api.CreateResponse[GoodsQueryGoodsIdByGoodsCodeResponse](&GoodsQueryGoodsIdByGoodsCodeResponse{})
	return
}
