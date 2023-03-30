package uc

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserGetUserInfo
// @id 725
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=根据wid获取用户信息)
func (client *Client) UserGetUserInfo(request *UserGetUserInfoRequest) (response *UserGetUserInfoResponse, err error) {
	rpcResponse := CreateUserGetUserInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserGetUserInfoRequest struct {
	*api.BaseRequest

	Wid int64 `json:"wid,omitempty"`
}

type UserGetUserInfoResponse struct {
	SourceObjectList []UserGetUserInfoResponseSourceObjectList `json:"sourceObjectList,omitempty"`
	Pid              int64                                     `json:"pid,omitempty"`
	SuperWid         int64                                     `json:"superWid,omitempty"`
	MergeStatus      int64                                     `json:"mergeStatus,omitempty"`
}

func CreateUserGetUserInfoRequest() (request *UserGetUserInfoRequest) {
	request = &UserGetUserInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("uc", "1_0", "user/getUserInfo", "api")
	request.Method = api.POST
	return
}

func CreateUserGetUserInfoResponse() (response *api.BaseResponse[UserGetUserInfoResponse]) {
	response = api.CreateResponse[UserGetUserInfoResponse](&UserGetUserInfoResponse{})
	return
}

type UserGetUserInfoResponseSourceObjectList struct {
	SourceData   UserGetUserInfoResponseSourceData `json:"sourceData,omitempty"`
	SourceOpenid string                            `json:"sourceOpenid,omitempty"`
	Source       int64                             `json:"source,omitempty"`
	SourceAppid  string                            `json:"sourceAppid,omitempty"`
}

type UserGetUserInfoResponseSourceData struct {
	SourceType    string `json:"sourceType,omitempty"`
	Nickname      string `json:"nickname,omitempty"`
	Headurl       string `json:"headurl,omitempty"`
	Sex           int64  `json:"sex,omitempty"`
	Country       string `json:"country,omitempty"`
	Province      string `json:"province,omitempty"`
	City          string `json:"city,omitempty"`
	Language      string `json:"language,omitempty"`
	Subscribe     int64  `json:"subscribe,omitempty"`
	SubscribeTime int64  `json:"subscribeTime,omitempty"`
	Unionid       string `json:"unionid,omitempty"`
}
