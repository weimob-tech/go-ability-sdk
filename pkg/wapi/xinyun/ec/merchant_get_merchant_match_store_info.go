package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantGetMerchantMatchStoreInfo
// @id 1455
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询商户主门店id)
func (client *Client) MerchantGetMerchantMatchStoreInfo(request *MerchantGetMerchantMatchStoreInfoRequest) (response *MerchantGetMerchantMatchStoreInfoResponse, err error) {
	rpcResponse := CreateMerchantGetMerchantMatchStoreInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantGetMerchantMatchStoreInfoRequest struct {
	*api.BaseRequest
}

type MerchantGetMerchantMatchStoreInfoResponse struct {
}

func CreateMerchantGetMerchantMatchStoreInfoRequest() (request *MerchantGetMerchantMatchStoreInfoRequest) {
	request = &MerchantGetMerchantMatchStoreInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchant/getMerchantMatchStoreInfo", "api")
	request.Method = api.POST
	return
}

func CreateMerchantGetMerchantMatchStoreInfoResponse() (response *api.BaseResponse[MerchantGetMerchantMatchStoreInfoResponse]) {
	response = api.CreateResponse[MerchantGetMerchantMatchStoreInfoResponse](&MerchantGetMerchantMatchStoreInfoResponse{})
	return
}
