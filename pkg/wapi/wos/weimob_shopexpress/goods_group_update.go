package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsGroupUpdate
// @id 1974
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1974?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商品分组更新)
func (client *Client) GoodsGroupUpdate(request *GoodsGroupUpdateRequest) (response *GoodsGroupUpdateResponse, err error) {
	rpcResponse := CreateGoodsGroupUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsGroupUpdateRequest struct {
	*api.BaseRequest

	GroupCode     string `json:"groupCode,omitempty"`
	GroupDescribe string `json:"groupDescribe,omitempty"`
	ImgUrl        string `json:"imgUrl,omitempty"`
	Name          string `json:"name,omitempty"`
	OperationName string `json:"operationName,omitempty"`
	ShowStatus    int64  `json:"showStatus,omitempty"`
}

type GoodsGroupUpdateResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateGoodsGroupUpdateRequest() (request *GoodsGroupUpdateRequest) {
	request = &GoodsGroupUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "goods/group/update", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsGroupUpdateResponse() (response *api.BaseResponse[GoodsGroupUpdateResponse]) {
	response = api.CreateResponse[GoodsGroupUpdateResponse](&GoodsGroupUpdateResponse{})
	return
}
