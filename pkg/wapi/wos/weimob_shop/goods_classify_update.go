package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsClassifyUpdate
// @id 1652
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1652?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=修改商品分组)
func (client *Client) GoodsClassifyUpdate(request *GoodsClassifyUpdateRequest) (response *GoodsClassifyUpdateResponse, err error) {
	rpcResponse := CreateGoodsClassifyUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsClassifyUpdateRequest struct {
	*api.BaseRequest

	BasicInfo  GoodsClassifyUpdateRequestBasicInfo `json:"basicInfo,omitempty"`
	ClassifyId int64                               `json:"classifyId,omitempty"`
	Name       string                              `json:"name,omitempty"`
	IsShow     bool                                `json:"isShow,omitempty"`
	ImageUrl   string                              `json:"imageUrl,omitempty"`
	Sort       int64                               `json:"sort,omitempty"`
	IsHot      bool                                `json:"isHot,omitempty"`
}

type GoodsClassifyUpdateResponse struct {
	ReturnResult bool `json:"returnResult,omitempty"`
}

func CreateGoodsClassifyUpdateRequest() (request *GoodsClassifyUpdateRequest) {
	request = &GoodsClassifyUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "goods/classify/update", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsClassifyUpdateResponse() (response *api.BaseResponse[GoodsClassifyUpdateResponse]) {
	response = api.CreateResponse[GoodsClassifyUpdateResponse](&GoodsClassifyUpdateResponse{})
	return
}

type GoodsClassifyUpdateRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}
