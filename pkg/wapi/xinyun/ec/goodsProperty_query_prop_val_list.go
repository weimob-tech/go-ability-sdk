package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsPropertyQueryPropValList
// @id 1171
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=根据属性Id列表查询属性值)
func (client *Client) GoodsPropertyQueryPropValList(request *GoodsPropertyQueryPropValListRequest) (response *GoodsPropertyQueryPropValListResponse, err error) {
	rpcResponse := CreateGoodsPropertyQueryPropValListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsPropertyQueryPropValListRequest struct {
	*api.BaseRequest

	PropIdList []int `json:"propIdList,omitempty"`
}

type GoodsPropertyQueryPropValListResponse struct {
}

func CreateGoodsPropertyQueryPropValListRequest() (request *GoodsPropertyQueryPropValListRequest) {
	request = &GoodsPropertyQueryPropValListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goodsProperty/queryPropValList", "api")
	request.Method = api.POST
	return
}

func CreateGoodsPropertyQueryPropValListResponse() (response *api.BaseResponse[GoodsPropertyQueryPropValListResponse]) {
	response = api.CreateResponse[GoodsPropertyQueryPropValListResponse](&GoodsPropertyQueryPropValListResponse{})
	return
}
