package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsGetFreightTemplateDetail
// @id 356
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取运费详情)
func (client *Client) GoodsGetFreightTemplateDetail(request *GoodsGetFreightTemplateDetailRequest) (response *GoodsGetFreightTemplateDetailResponse, err error) {
	rpcResponse := CreateGoodsGetFreightTemplateDetailResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsGetFreightTemplateDetailRequest struct {
	*api.BaseRequest

	TemplateId int64 `json:"templateId,omitempty"`
}

type GoodsGetFreightTemplateDetailResponse struct {
}

func CreateGoodsGetFreightTemplateDetailRequest() (request *GoodsGetFreightTemplateDetailRequest) {
	request = &GoodsGetFreightTemplateDetailRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "goods/getFreightTemplateDetail", "api")
	request.Method = api.POST
	return
}

func CreateGoodsGetFreightTemplateDetailResponse() (response *api.BaseResponse[GoodsGetFreightTemplateDetailResponse]) {
	response = api.CreateResponse[GoodsGetFreightTemplateDetailResponse](&GoodsGetFreightTemplateDetailResponse{})
	return
}
