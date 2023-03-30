package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsGroupAdd
// @id 1972
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1972?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商品分组添加)
func (client *Client) GoodsGroupAdd(request *GoodsGroupAddRequest) (response *GoodsGroupAddResponse, err error) {
	rpcResponse := CreateGoodsGroupAddResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsGroupAddRequest struct {
	*api.BaseRequest

	GroupDescribe string `json:"groupDescribe,omitempty"`
	ImgUrl        string `json:"imgUrl,omitempty"`
	Name          string `json:"name,omitempty"`
	OperationName string `json:"operationName,omitempty"`
	ShowStatus    int64  `json:"showStatus,omitempty"`
}

type GoodsGroupAddResponse struct {
	GroupCode string `json:"groupCode,omitempty"`
}

func CreateGoodsGroupAddRequest() (request *GoodsGroupAddRequest) {
	request = &GoodsGroupAddRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "goods/group/add", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsGroupAddResponse() (response *api.BaseResponse[GoodsGroupAddResponse]) {
	response = api.CreateResponse[GoodsGroupAddResponse](&GoodsGroupAddResponse{})
	return
}
