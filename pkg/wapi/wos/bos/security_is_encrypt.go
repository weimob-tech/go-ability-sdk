package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SecurityIsEncrypt
// @id 1810
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1810?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=数据是否加密)
func (client *Client) SecurityIsEncrypt(request *SecurityIsEncryptRequest) (response *SecurityIsEncryptResponse, err error) {
	rpcResponse := CreateSecurityIsEncryptResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SecurityIsEncryptRequest struct {
	*api.BaseRequest

	Source string `json:"source,omitempty"`
}

type SecurityIsEncryptResponse struct {
	IsEncrypt bool `json:"isEncrypt,omitempty"`
}

func CreateSecurityIsEncryptRequest() (request *SecurityIsEncryptRequest) {
	request = &SecurityIsEncryptRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "security/isEncrypt", "apigw")
	request.Method = api.POST
	return
}

func CreateSecurityIsEncryptResponse() (response *api.BaseResponse[SecurityIsEncryptResponse]) {
	response = api.CreateResponse[SecurityIsEncryptResponse](&SecurityIsEncryptResponse{})
	return
}
