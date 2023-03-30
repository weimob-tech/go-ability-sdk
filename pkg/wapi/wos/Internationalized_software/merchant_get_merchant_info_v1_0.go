package Internationalized_software

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantGetMerchantInfo
// @id 1229
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1229?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取商城信息)
func (client *Client) MerchantGetMerchantInfoV1_0(request *WeimobShopexpressMerchantGetMerchantInfoRequest) (response *WeimobShopexpressMerchantGetMerchantInfoResponse, err error) {
	rpcResponse := CreateWeimobShopexpressMerchantGetMerchantInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type WeimobShopexpressMerchantGetMerchantInfoRequest struct {
	*api.BaseRequest
}

type WeimobShopexpressMerchantGetMerchantInfoResponse struct {
	MerchantName string `json:"merchantName,omitempty"`
	Mail         string `json:"mail,omitempty"`
}

func CreateWeimobShopexpressMerchantGetMerchantInfoRequest() (request *WeimobShopexpressMerchantGetMerchantInfoRequest) {
	request = &WeimobShopexpressMerchantGetMerchantInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("Internationalized_software", "v1.0", "merchant/getMerchantInfo", "apigw")
	request.Method = api.POST
	return
}

func CreateWeimobShopexpressMerchantGetMerchantInfoResponse() (response *api.BaseResponse[WeimobShopexpressMerchantGetMerchantInfoResponse]) {
	response = api.CreateResponse[WeimobShopexpressMerchantGetMerchantInfoResponse](&WeimobShopexpressMerchantGetMerchantInfoResponse{})
	return
}
