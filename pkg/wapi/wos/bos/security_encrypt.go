package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// SecurityEncrypt
// @id 1708
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1708?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=数据加密)
func (client *Client) SecurityEncrypt(request *SecurityEncryptRequest) (response *SecurityEncryptResponse, err error) {
	rpcResponse := CreateSecurityEncryptResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type SecurityEncryptRequest struct {
	*api.BaseRequest

	Source string `json:"source,omitempty"`
}

type SecurityEncryptResponse struct {
	Result string `json:"result,omitempty"`
}

func CreateSecurityEncryptRequest() (request *SecurityEncryptRequest) {
	request = &SecurityEncryptRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "security/encrypt", "apigw")
	request.Method = api.POST
	return
}

func CreateSecurityEncryptResponse() (response *api.BaseResponse[SecurityEncryptResponse]) {
	response = api.CreateResponse[SecurityEncryptResponse](&SecurityEncryptResponse{})
	return
}
