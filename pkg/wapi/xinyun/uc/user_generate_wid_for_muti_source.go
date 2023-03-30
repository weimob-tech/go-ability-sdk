package uc

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// UserGenerateWidForMutiSource
// @id 1524
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=为多个来源生成wid)
func (client *Client) UserGenerateWidForMutiSource(request *UserGenerateWidForMutiSourceRequest) (response *UserGenerateWidForMutiSourceResponse, err error) {
	rpcResponse := CreateUserGenerateWidForMutiSourceResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type UserGenerateWidForMutiSourceRequest struct {
	*api.BaseRequest

	SourceObjectList []UserGenerateWidForMutiSourceRequestSourceObjectList `json:"sourceObjectList,omitempty"`
	Pid              int64                                                 `json:"pid,omitempty"`
}

type UserGenerateWidForMutiSourceResponse struct {
}

func CreateUserGenerateWidForMutiSourceRequest() (request *UserGenerateWidForMutiSourceRequest) {
	request = &UserGenerateWidForMutiSourceRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("uc", "1_0", "user/generateWidForMutiSource", "api")
	request.Method = api.POST
	return
}

func CreateUserGenerateWidForMutiSourceResponse() (response *api.BaseResponse[UserGenerateWidForMutiSourceResponse]) {
	response = api.CreateResponse[UserGenerateWidForMutiSourceResponse](&UserGenerateWidForMutiSourceResponse{})
	return
}

type UserGenerateWidForMutiSourceRequestSourceObjectList struct {
	SourceData   UserGenerateWidForMutiSourceRequestSourceData `json:"sourceData,omitempty"`
	SourceOpenid string                                        `json:"sourceOpenid,omitempty"`
	Source       int                                           `json:"source,omitempty"`
	SourceAppid  string                                        `json:"sourceAppid,omitempty"`
}

type UserGenerateWidForMutiSourceRequestSourceData struct {
	SourceType    string `json:"sourceType,omitempty"`
	Nickname      string `json:"nickname,omitempty"`
	Headurl       string `json:"headurl,omitempty"`
	Sex           int    `json:"sex,omitempty"`
	Country       string `json:"country,omitempty"`
	Province      string `json:"province,omitempty"`
	City          string `json:"city,omitempty"`
	Subscribe     int    `json:"subscribe,omitempty"`
	SubscribeTime int    `json:"subscribeTime,omitempty"`
}
