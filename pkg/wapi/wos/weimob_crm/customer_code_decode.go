package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerCodeDecode
// @id 2072
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2072?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=识别客户码)
func (client *Client) CustomerCodeDecode(request *CustomerCodeDecodeRequest) (response *CustomerCodeDecodeResponse, err error) {
	rpcResponse := CreateCustomerCodeDecodeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerCodeDecodeRequest struct {
	*api.BaseRequest

	CodeType  int64  `json:"codeType,omitempty"`
	CodeValue string `json:"codeValue,omitempty"`
}

type CustomerCodeDecodeResponse struct {
	Wid int64 `json:"wid,omitempty"`
}

func CreateCustomerCodeDecodeRequest() (request *CustomerCodeDecodeRequest) {
	request = &CustomerCodeDecodeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "customer/code/decode", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerCodeDecodeResponse() (response *api.BaseResponse[CustomerCodeDecodeResponse]) {
	response = api.CreateResponse[CustomerCodeDecodeResponse](&CustomerCodeDecodeResponse{})
	return
}
