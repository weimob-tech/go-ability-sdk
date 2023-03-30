package weimob_crm

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerAssetVerifyresultGet
// @id 1647
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1647?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取资产核销客户授权校验结果)
func (client *Client) CustomerAssetVerifyresultGet(request *CustomerAssetVerifyresultGetRequest) (response *CustomerAssetVerifyresultGetResponse, err error) {
	rpcResponse := CreateCustomerAssetVerifyresultGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerAssetVerifyresultGetRequest struct {
	*api.BaseRequest

	Wid        int64  `json:"wid,omitempty"`
	Vid        int64  `json:"vid,omitempty"`
	CheckToken string `json:"checkToken,omitempty"`
}

type CustomerAssetVerifyresultGetResponse struct {
	Status   int64  `json:"status,omitempty"`
	HasError int64  `json:"hasError,omitempty"`
	Message  string `json:"message,omitempty"`
}

func CreateCustomerAssetVerifyresultGetRequest() (request *CustomerAssetVerifyresultGetRequest) {
	request = &CustomerAssetVerifyresultGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_crm", "v2.0", "customer/asset/verifyresult/get", "apigw")
	request.Method = api.POST
	return
}

func CreateCustomerAssetVerifyresultGetResponse() (response *api.BaseResponse[CustomerAssetVerifyresultGetResponse]) {
	response = api.CreateResponse[CustomerAssetVerifyresultGetResponse](&CustomerAssetVerifyresultGetResponse{})
	return
}
