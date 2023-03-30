package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserCreate
// @id 2037
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2037?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=多渠道创建wid)
func (client *Client) UserCreate(request *UserCreateRequest) (response *UserCreateResponse, err error) {
	rpcResponse := CreateUserCreateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserCreateRequest struct {
	*api.BaseRequest

	SourceObjectList []UserCreateRequestSourceObjectList `json:"sourceObjectList,omitempty"`
	BosId            int64                               `json:"bosId,omitempty"`
}

type UserCreateResponse struct {
	BosId    int64 `json:"bosId,omitempty"`
	SuperWid int64 `json:"superWid,omitempty"`
}

func CreateUserCreateRequest() (request *UserCreateRequest) {
	request = &UserCreateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "user/create", "apigw")
	request.Method = api.POST
	return
}

func CreateUserCreateResponse() (response *api.BaseResponse[UserCreateResponse]) {
	response = api.CreateResponse[UserCreateResponse](&UserCreateResponse{})
	return
}

type UserCreateRequestSourceObjectList struct {
	SourceData UserCreateRequestSourceData `json:"sourceData,omitempty"`
	Source     int64                       `json:"source,omitempty"`
	AppId      string                      `json:"appId,omitempty"`
	OriginalId string                      `json:"originalId,omitempty"`
}

type UserCreateRequestSourceData struct {
	SourceType    string `json:"sourceType,omitempty"`
	Nickname      string `json:"nickname,omitempty"`
	HeadUrl       string `json:"headUrl,omitempty"`
	Sex           int64  `json:"sex,omitempty"`
	Country       string `json:"country,omitempty"`
	Province      string `json:"province,omitempty"`
	City          string `json:"city,omitempty"`
	Language      string `json:"language,omitempty"`
	Subscribe     int64  `json:"subscribe,omitempty"`
	SubscribeTime string `json:"subscribeTime,omitempty"`
}
