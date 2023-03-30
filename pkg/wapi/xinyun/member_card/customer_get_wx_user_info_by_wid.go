package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CustomerGetWxUserInfoByWid
// @id 983
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=根据wid获取微信用户信息)
func (client *Client) CustomerGetWxUserInfoByWid(request *CustomerGetWxUserInfoByWidRequest) (response *CustomerGetWxUserInfoByWidResponse, err error) {
	rpcResponse := CreateCustomerGetWxUserInfoByWidResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CustomerGetWxUserInfoByWidRequest struct {
	*api.BaseRequest

	Wid int64 `json:"wid,omitempty"`
}

type CustomerGetWxUserInfoByWidResponse struct {
}

func CreateCustomerGetWxUserInfoByWidRequest() (request *CustomerGetWxUserInfoByWidRequest) {
	request = &CustomerGetWxUserInfoByWidRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "customer/getWxUserInfoByWid", "api")
	request.Method = api.POST
	return
}

func CreateCustomerGetWxUserInfoByWidResponse() (response *api.BaseResponse[CustomerGetWxUserInfoByWidResponse]) {
	response = api.CreateResponse[CustomerGetWxUserInfoByWidResponse](&CustomerGetWxUserInfoByWidResponse{})
	return
}
