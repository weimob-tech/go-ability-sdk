package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsTagDelete
// @id 1851
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1851?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=删除标签)
func (client *Client) GoodsTagDelete(request *GoodsTagDeleteRequest) (response *GoodsTagDeleteResponse, err error) {
	rpcResponse := CreateGoodsTagDeleteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsTagDeleteRequest struct {
	*api.BaseRequest

	GoodsTagIdList []GoodsTagDeleteRequestGoodsTagIdList `json:"goodsTagIdList,omitempty"`
	BasicInfo      GoodsTagDeleteRequestBasicInfo        `json:"basicInfo,omitempty"`
}

type GoodsTagDeleteResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateGoodsTagDeleteRequest() (request *GoodsTagDeleteRequest) {
	request = &GoodsTagDeleteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/tag/delete", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsTagDeleteResponse() (response *api.BaseResponse[GoodsTagDeleteResponse]) {
	response = api.CreateResponse[GoodsTagDeleteResponse](&GoodsTagDeleteResponse{})
	return
}

type GoodsTagDeleteRequestGoodsTagIdList struct {
}

type GoodsTagDeleteRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
