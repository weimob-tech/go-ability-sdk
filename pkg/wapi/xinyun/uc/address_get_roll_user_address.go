package uc

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AddressGetRollUserAddress
// @id 1937
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=滚动获取收货地址列表)
func (client *Client) AddressGetRollUserAddress(request *AddressGetRollUserAddressRequest) (response *AddressGetRollUserAddressResponse, err error) {
	rpcResponse := CreateAddressGetRollUserAddressResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AddressGetRollUserAddressRequest struct {
	*api.BaseRequest
}

type AddressGetRollUserAddressResponse struct {
}

func CreateAddressGetRollUserAddressRequest() (request *AddressGetRollUserAddressRequest) {
	request = &AddressGetRollUserAddressRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("uc", "1_0", "address/getRollUserAddress", "api")
	request.Method = api.POST
	return
}

func CreateAddressGetRollUserAddressResponse() (response *api.BaseResponse[AddressGetRollUserAddressResponse]) {
	response = api.CreateResponse[AddressGetRollUserAddressResponse](&AddressGetRollUserAddressResponse{})
	return
}
