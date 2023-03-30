package weimob_shopexpress

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// GoodsShellSwitch
// @id 1965
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1965?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商品上下架切换)
func (client *Client) GoodsShellSwitch(request *GoodsShellSwitchRequest) (response *GoodsShellSwitchResponse, err error) {
	rpcResponse := CreateGoodsShellSwitchResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type GoodsShellSwitchRequest struct {
	*api.BaseRequest

	GoodsCodeList []int64 `json:"goodsCodeList,omitempty"`
	Status        int64   `json:"status,omitempty"`
}

type GoodsShellSwitchResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateGoodsShellSwitchRequest() (request *GoodsShellSwitchRequest) {
	request = &GoodsShellSwitchRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shopexpress", "v2.0", "goods/shell/switch", "apigw")
	request.Method = api.POST
	return
}

func CreateGoodsShellSwitchResponse() (response *api.BaseResponse[GoodsShellSwitchResponse]) {
	response = api.CreateResponse[GoodsShellSwitchResponse](&GoodsShellSwitchResponse{})
	return
}
