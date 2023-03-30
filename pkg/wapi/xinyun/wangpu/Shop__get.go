package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ShopGet
// @id 17
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取店铺信息)
func (client *Client) ShopGet(request *ShopGetRequest) (response *ShopGetResponse, err error) {
	rpcResponse := CreateShopGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ShopGetRequest struct {
	*api.BaseRequest

	IncludeIntro bool `json:"include_intro,omitempty"`
}

type ShopGetResponse struct {
}

func CreateShopGetRequest() (request *ShopGetRequest) {
	request = &ShopGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "Shop/Get", "api")
	request.Method = api.POST
	return
}

func CreateShopGetResponse() (response *api.BaseResponse[ShopGetResponse]) {
	response = api.CreateResponse[ShopGetResponse](&ShopGetResponse{})
	return
}
