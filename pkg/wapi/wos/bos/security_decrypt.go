package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SecurityDecrypt
// @id 1709
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1709?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=数据解密)
func (client *Client) SecurityDecrypt(request *SecurityDecryptRequest) (response *SecurityDecryptResponse, err error) {
	rpcResponse := CreateSecurityDecryptResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SecurityDecryptRequest struct {
	*api.BaseRequest

	Source string `json:"source,omitempty"`
}

type SecurityDecryptResponse struct {
	Result string `json:"result,omitempty"`
}

func CreateSecurityDecryptRequest() (request *SecurityDecryptRequest) {
	request = &SecurityDecryptRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "security/decrypt", "apigw")
	request.Method = api.POST
	return
}

func CreateSecurityDecryptResponse() (response *api.BaseResponse[SecurityDecryptResponse]) {
	response = api.CreateResponse[SecurityDecryptResponse](&SecurityDecryptResponse{})
	return
}
