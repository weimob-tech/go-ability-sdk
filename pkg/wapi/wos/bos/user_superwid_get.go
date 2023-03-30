package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserSuperwidGet
// @id 2034
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2034?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取主wid)
func (client *Client) UserSuperwidGet(request *UserSuperwidGetRequest) (response *UserSuperwidGetResponse, err error) {
	rpcResponse := CreateUserSuperwidGetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserSuperwidGetRequest struct {
	*api.BaseRequest

	BosId      int64  `json:"bosId,omitempty"`
	Source     int64  `json:"source,omitempty"`
	AppId      string `json:"appId,omitempty"`
	OriginalId string `json:"originalId,omitempty"`
	Wid        int64  `json:"wid,omitempty"`
}

type UserSuperwidGetResponse struct {
	BosId    int64 `json:"bosId,omitempty"`
	SuperWid int64 `json:"superWid,omitempty"`
}

func CreateUserSuperwidGetRequest() (request *UserSuperwidGetRequest) {
	request = &UserSuperwidGetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "user/superwid/get", "apigw")
	request.Method = api.POST
	return
}

func CreateUserSuperwidGetResponse() (response *api.BaseResponse[UserSuperwidGetResponse]) {
	response = api.CreateResponse[UserSuperwidGetResponse](&UserSuperwidGetResponse{})
	return
}
