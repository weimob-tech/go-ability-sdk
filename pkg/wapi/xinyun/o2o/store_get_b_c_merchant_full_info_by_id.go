package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// StoreGetBCMerchantFullInfoById
// @id 2091
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询微盟商户信息)
func (client *Client) StoreGetBCMerchantFullInfoById(request *StoreGetBCMerchantFullInfoByIdRequest) (response *StoreGetBCMerchantFullInfoByIdResponse, err error) {
	rpcResponse := CreateStoreGetBCMerchantFullInfoByIdResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type StoreGetBCMerchantFullInfoByIdRequest struct {
	*api.BaseRequest
}

type StoreGetBCMerchantFullInfoByIdResponse struct {
}

func CreateStoreGetBCMerchantFullInfoByIdRequest() (request *StoreGetBCMerchantFullInfoByIdRequest) {
	request = &StoreGetBCMerchantFullInfoByIdRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "store/getBCMerchantFullInfoById", "api")
	request.Method = api.POST
	return
}

func CreateStoreGetBCMerchantFullInfoByIdResponse() (response *api.BaseResponse[StoreGetBCMerchantFullInfoByIdResponse]) {
	response = api.CreateResponse[StoreGetBCMerchantFullInfoByIdResponse](&StoreGetBCMerchantFullInfoByIdResponse{})
	return
}
