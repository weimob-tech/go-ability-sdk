package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NavigationExtGetMiniCode
// @id 3642
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取商品小程序码)
func (client *Client) NavigationExtGetMiniCode(request *NavigationExtGetMiniCodeRequest) (response *NavigationExtGetMiniCodeResponse, err error) {
	rpcResponse := CreateNavigationExtGetMiniCodeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NavigationExtGetMiniCodeRequest struct {
	*api.BaseRequest

	GoodsId      int64  `json:"goodsId,omitempty"`
	StoreId      int64  `json:"storeId,omitempty"`
	Pid          string `json:"pid,omitempty"`
	MiniUrlParam string `json:"miniUrlParam,omitempty"`
	Width        int64  `json:"width,omitempty"`
}

type NavigationExtGetMiniCodeResponse struct {
	MiniCodeUrl string `json:"miniCodeUrl,omitempty"`
}

func CreateNavigationExtGetMiniCodeRequest() (request *NavigationExtGetMiniCodeRequest) {
	request = &NavigationExtGetMiniCodeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "navigation/extGetMiniCode", "api")
	request.Method = api.POST
	return
}

func CreateNavigationExtGetMiniCodeResponse() (response *api.BaseResponse[NavigationExtGetMiniCodeResponse]) {
	response = api.CreateResponse[NavigationExtGetMiniCodeResponse](&NavigationExtGetMiniCodeResponse{})
	return
}
