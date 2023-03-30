package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NotificationSendShortMsg
// @id 2545
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=短消息发送)
func (client *Client) NotificationSendShortMsg(request *NotificationSendShortMsgRequest) (response *NotificationSendShortMsgResponse, err error) {
	rpcResponse := CreateNotificationSendShortMsgResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NotificationSendShortMsgRequest struct {
	*api.BaseRequest

	Content  string   `json:"content,omitempty"`
	PhoneNos []string `json:"phoneNos,omitempty"`
	Remark   string   `json:"remark,omitempty"`
	Sign     string   `json:"sign,omitempty"`
}

type NotificationSendShortMsgResponse struct {
}

func CreateNotificationSendShortMsgRequest() (request *NotificationSendShortMsgRequest) {
	request = &NotificationSendShortMsgRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "notification/sendShortMsg", "api")
	request.Method = api.POST
	return
}

func CreateNotificationSendShortMsgResponse() (response *api.BaseResponse[NotificationSendShortMsgResponse]) {
	response = api.CreateResponse[NotificationSendShortMsgResponse](&NotificationSendShortMsgResponse{})
	return
}
