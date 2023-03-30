package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsTagUpdate
// @id 1852
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1852?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新标签)
func (client *Client) GoodsTagUpdate(request *GoodsTagUpdateRequest) (response *GoodsTagUpdateResponse, err error) {
	rpcResponse := CreateGoodsTagUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsTagUpdateRequest struct {
	*api.BaseRequest

	BasicInfo GoodsTagUpdateRequestBasicInfo `json:"basicInfo,omitempty"`
	Title     string                         `json:"title,omitempty"`
	TagId     int64                          `json:"tagId,omitempty"`
}

type GoodsTagUpdateResponse struct {
	Result bool `json:"result,omitempty"`
}

func CreateGoodsTagUpdateRequest() (request *GoodsTagUpdateRequest) {
	request = &GoodsTagUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/tag/update", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsTagUpdateResponse() (response *api.BaseResponse[GoodsTagUpdateResponse]) {
	response = api.CreateResponse[GoodsTagUpdateResponse](&GoodsTagUpdateResponse{})
	return
}

type GoodsTagUpdateRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
