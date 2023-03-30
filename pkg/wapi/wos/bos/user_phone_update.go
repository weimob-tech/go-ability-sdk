package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserPhoneUpdate
// @id 2099
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2099?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=更新手机号)
func (client *Client) UserPhoneUpdate(request *UserPhoneUpdateRequest) (response *UserPhoneUpdateResponse, err error) {
	rpcResponse := CreateUserPhoneUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserPhoneUpdateRequest struct {
	*api.BaseRequest

	BosId    int64  `json:"bosId,omitempty"`
	Wid      int64  `json:"wid	,omitempty"`
	OldZone  string `json:"oldZone,omitempty"`
	OldPhone string `json:"oldPhone,omitempty"`
	NewZone  string `json:"newZone,omitempty"`
	NewPhone string `json:"newPhone,omitempty"`
}

type UserPhoneUpdateResponse struct {
}

func CreateUserPhoneUpdateRequest() (request *UserPhoneUpdateRequest) {
	request = &UserPhoneUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "user/phone/update", "apigw")
	request.Method = api.POST
	return
}

func CreateUserPhoneUpdateResponse() (response *api.BaseResponse[UserPhoneUpdateResponse]) {
	response = api.CreateResponse[UserPhoneUpdateResponse](&UserPhoneUpdateResponse{})
	return
}
