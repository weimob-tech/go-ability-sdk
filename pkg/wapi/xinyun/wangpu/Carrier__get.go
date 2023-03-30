package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CarrierGet
// @id 25
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取所有快递公司信息)
func (client *Client) CarrierGet(request *CarrierGetRequest) (response *CarrierGetResponse, err error) {
	rpcResponse := CreateCarrierGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CarrierGetRequest struct {
	*api.BaseRequest
}

type CarrierGetResponse struct {
}

func CreateCarrierGetRequest() (request *CarrierGetRequest) {
	request = &CarrierGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "Carrier/Get", "api")
	request.Method = api.POST
	return
}

func CreateCarrierGetResponse() (response *api.BaseResponse[CarrierGetResponse]) {
	response = api.CreateResponse[CarrierGetResponse](&CarrierGetResponse{})
	return
}
