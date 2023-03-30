package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// CardPublishToUser
// @id 430
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=向特定用户发券)
func (client *Client) CardPublishToUser(request *CardPublishToUserRequest) (response *CardPublishToUserResponse, err error) {
	rpcResponse := CreateCardPublishToUserResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type CardPublishToUserRequest struct {
	*api.BaseRequest

	CardTemplateId int64  `json:"cardTemplateId,omitempty"`
	OpenId         string `json:"openId,omitempty"`
	Wid            string `json:"wid,omitempty"`
	Count          int    `json:"count,omitempty"`
}

type CardPublishToUserResponse struct {
}

func CreateCardPublishToUserRequest() (request *CardPublishToUserRequest) {
	request = &CardPublishToUserRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "card/publishToUser", "api")
	request.Method = api.POST
	return
}

func CreateCardPublishToUserResponse() (response *api.BaseResponse[CardPublishToUserResponse]) {
	response = api.CreateResponse[CardPublishToUserResponse](&CardPublishToUserResponse{})
	return
}
