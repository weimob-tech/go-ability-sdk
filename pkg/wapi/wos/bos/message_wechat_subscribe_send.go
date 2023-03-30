package bos

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MessageWechatSubscribeSend
// @id 2328
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2328?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=发送微信小程序订阅消息)
func (client *Client) MessageWechatSubscribeSend(request *MessageWechatSubscribeSendRequest) (response *MessageWechatSubscribeSendResponse, err error) {
	rpcResponse := CreateMessageWechatSubscribeSendResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MessageWechatSubscribeSendRequest struct {
	*api.BaseRequest

	AppId        string  `json:"appId,omitempty"`
	TemplateId   string  `json:"templateId,omitempty"`
	TemplateData string  `json:"templateData,omitempty"`
	WidList      []int64 `json:"widList,omitempty"`
}

type MessageWechatSubscribeSendResponse struct {
	RequestId string `json:"requestId,omitempty"`
}

func CreateMessageWechatSubscribeSendRequest() (request *MessageWechatSubscribeSendRequest) {
	request = &MessageWechatSubscribeSendRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("bos", "v2.0", "message/wechat/subscribe/send", "apigw")
	request.Method = api.POST
	return
}

func CreateMessageWechatSubscribeSendResponse() (response *api.BaseResponse[MessageWechatSubscribeSendResponse]) {
	response = api.CreateResponse[MessageWechatSubscribeSendResponse](&MessageWechatSubscribeSendResponse{})
	return
}
