package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsTagAddList
// @id 1969
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1969?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商品标签列表新增)
func (client *Client) GoodsTagAddList(request *GoodsTagAddListRequest) (response *GoodsTagAddListResponse, err error) {
	rpcResponse := CreateGoodsTagAddListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsTagAddListRequest struct {
	*api.BaseRequest

	TagNames []int64 `json:"tagNames,omitempty"`
}

type GoodsTagAddListResponse struct {
	TagCode string `json:"tagCode,omitempty"`
	TagName string `json:"tagName,omitempty"`
}

func CreateGoodsTagAddListRequest() (request *GoodsTagAddListRequest) {
	request = &GoodsTagAddListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "goods/tag/addList", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsTagAddListResponse() (response *api.BaseResponse[GoodsTagAddListResponse]) {
	response = api.CreateResponse[GoodsTagAddListResponse](&GoodsTagAddListResponse{})
	return
}
