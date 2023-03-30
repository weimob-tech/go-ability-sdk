package weimob_cs

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MsgReceive
// @id 1880
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1880?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=接收消息和事件)
func (client *Client) MsgReceive(request *MsgReceiveRequest) (response *MsgReceiveResponse, err error) {
	rpcResponse := CreateMsgReceiveResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MsgReceiveRequest struct {
	*api.BaseRequest

	MsgList                 []MsgReceiveRequestMsgList `json:"msgList,omitempty"`
	Ext                     MsgReceiveRequestExt       `json:"ext,omitempty"`
	BosId                   int64                      `json:"bosId,omitempty"`
	TargetProductInstanceId int64                      `json:"targetProductInstanceId,omitempty"`
	Status                  int64                      `json:"status,omitempty"`
}

type MsgReceiveResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateMsgReceiveRequest() (request *MsgReceiveRequest) {
	request = &MsgReceiveRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_cs", "v2.0", "msg/receive", "apigw")
	request.Method = api.POST
	return
}

func CreateMsgReceiveResponse() (response *api.BaseResponse[MsgReceiveResponse]) {
	response = api.CreateResponse[MsgReceiveResponse](&MsgReceiveResponse{})
	return
}

type MsgReceiveRequestMsgList struct {
	Id        int64  `json:"id,omitempty"`
	MsgType   string `json:"msgType,omitempty"`
	Content   string `json:"content,omitempty"`
	Status    int64  `json:"status,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
}

type MsgReceiveRequestExt struct {
	CustomerId      int64  `json:"customerId,omitempty"`
	Channel         string `json:"channel,omitempty"`
	ClientId        string `json:"clientId,omitempty"`
	ClientName      string `json:"clientName,omitempty"`
	OriginProductId int64  `json:"originProductId,omitempty"`
	Wid             int64  `json:"wid,omitempty"`
}
