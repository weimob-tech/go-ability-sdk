package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AddressGetAreasByLocation
// @id 533
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=根据经纬度查询四级地址)
func (client *Client) AddressGetAreasByLocation(request *AddressGetAreasByLocationRequest) (response *AddressGetAreasByLocationResponse, err error) {
	rpcResponse := CreateAddressGetAreasByLocationResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AddressGetAreasByLocationRequest struct {
	*api.BaseRequest

	Location string `json:"location,omitempty"`
}

type AddressGetAreasByLocationResponse struct {
}

func CreateAddressGetAreasByLocationRequest() (request *AddressGetAreasByLocationRequest) {
	request = &AddressGetAreasByLocationRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "address/getAreasByLocation", "api")
	request.Method = api.POST
	return
}

func CreateAddressGetAreasByLocationResponse() (response *api.BaseResponse[AddressGetAreasByLocationResponse]) {
	response = api.CreateResponse[AddressGetAreasByLocationResponse](&AddressGetAreasByLocationResponse{})
	return
}
