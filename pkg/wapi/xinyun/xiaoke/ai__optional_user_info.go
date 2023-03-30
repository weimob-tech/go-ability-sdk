package xiaoke

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AiOptionalUserInfo
// @id 3738
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询一对一模式员工信息)
func (client *Client) AiOptionalUserInfo(request *AiOptionalUserInfoRequest) (response *AiOptionalUserInfoResponse, err error) {
	rpcResponse := CreateAiOptionalUserInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AiOptionalUserInfoRequest struct {
	*api.BaseRequest

	LoginUserWid int64 `json:"loginUserWid,omitempty"`
	SalesTalkId  int64 `json:"salesTalkId,omitempty"`
}

type AiOptionalUserInfoResponse struct {
	AiCount  int64  `json:"aiCount,omitempty"`
	UserName string `json:"userName,omitempty"`
	UserWid  int64  `json:"userWid,omitempty"`
}

func CreateAiOptionalUserInfoRequest() (request *AiOptionalUserInfoRequest) {
	request = &AiOptionalUserInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("xiaoke", "1_0", "ai/OptionalUserInfo", "api")
	request.Method = api.POST
	return
}

func CreateAiOptionalUserInfoResponse() (response *api.BaseResponse[AiOptionalUserInfoResponse]) {
	response = api.CreateResponse[AiOptionalUserInfoResponse](&AiOptionalUserInfoResponse{})
	return
}
