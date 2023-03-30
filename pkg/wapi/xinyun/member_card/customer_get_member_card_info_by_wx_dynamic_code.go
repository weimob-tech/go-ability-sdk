package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerGetMemberCardInfoByWxDynamicCode
// @id 970
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=根据微信动态码获取会员卡code)
func (client *Client) CustomerGetMemberCardInfoByWxDynamicCode(request *CustomerGetInfoByWxDynamicCodeRequest) (response *CustomerGetInfoByWxDynamicCodeResponse, err error) {
	rpcResponse := CreateCustomerGetInfoByWxDynamicCodeResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerGetInfoByWxDynamicCodeRequest struct {
	*api.BaseRequest

	Code string `json:"code,omitempty"`
}

type CustomerGetInfoByWxDynamicCodeResponse struct {
}

func CreateCustomerGetInfoByWxDynamicCodeRequest() (request *CustomerGetInfoByWxDynamicCodeRequest) {
	request = &CustomerGetInfoByWxDynamicCodeRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "customer/getMemberCardInfoByWxDynamicCode", "api")
	request.Method = api.POST
	return
}

func CreateCustomerGetInfoByWxDynamicCodeResponse() (response *api.BaseResponse[CustomerGetInfoByWxDynamicCodeResponse]) {
	response = api.CreateResponse[CustomerGetInfoByWxDynamicCodeResponse](&CustomerGetInfoByWxDynamicCodeResponse{})
	return
}
