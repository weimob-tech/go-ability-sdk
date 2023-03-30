package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantprivilegeUpdateAccountPhone
// @id 2733
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=修改账号的手机号)
func (client *Client) MerchantprivilegeUpdateAccountPhone(request *MerchantprivilegeUpdateAccountPhoneRequest) (response *MerchantprivilegeUpdateAccountPhoneResponse, err error) {
	rpcResponse := CreateMerchantprivilegeUpdateAccountPhoneResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantprivilegeUpdateAccountPhoneRequest struct {
	*api.BaseRequest

	OldZone  string `json:"oldZone,omitempty"`
	OldPhone string `json:"oldPhone,omitempty"`
	NewZone  string `json:"newZone,omitempty"`
	NewPhone string `json:"newPhone,omitempty"`
}

type MerchantprivilegeUpdateAccountPhoneResponse struct {
}

func CreateMerchantprivilegeUpdateAccountPhoneRequest() (request *MerchantprivilegeUpdateAccountPhoneRequest) {
	request = &MerchantprivilegeUpdateAccountPhoneRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchantprivilege/updateAccountPhone", "api")
	request.Method = api.POST
	return
}

func CreateMerchantprivilegeUpdateAccountPhoneResponse() (response *api.BaseResponse[MerchantprivilegeUpdateAccountPhoneResponse]) {
	response = api.CreateResponse[MerchantprivilegeUpdateAccountPhoneResponse](&MerchantprivilegeUpdateAccountPhoneResponse{})
	return
}
