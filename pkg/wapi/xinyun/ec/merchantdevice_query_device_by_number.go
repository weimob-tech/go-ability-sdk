package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantdeviceQueryDeviceByNumber
// @id 1419
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=根据设备编号批量查询设备信息)
func (client *Client) MerchantdeviceQueryDeviceByNumber(request *MerchantdeviceQueryDeviceByNumberRequest) (response *MerchantdeviceQueryDeviceByNumberResponse, err error) {
	rpcResponse := CreateMerchantdeviceQueryDeviceByNumberResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantdeviceQueryDeviceByNumberRequest struct {
	*api.BaseRequest
}

type MerchantdeviceQueryDeviceByNumberResponse struct {
}

func CreateMerchantdeviceQueryDeviceByNumberRequest() (request *MerchantdeviceQueryDeviceByNumberRequest) {
	request = &MerchantdeviceQueryDeviceByNumberRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchantdevice/queryDeviceByNumber", "api")
	request.Method = api.POST
	return
}

func CreateMerchantdeviceQueryDeviceByNumberResponse() (response *api.BaseResponse[MerchantdeviceQueryDeviceByNumberResponse]) {
	response = api.CreateResponse[MerchantdeviceQueryDeviceByNumberResponse](&MerchantdeviceQueryDeviceByNumberResponse{})
	return
}
