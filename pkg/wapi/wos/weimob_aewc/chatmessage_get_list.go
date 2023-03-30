package weimob_aewc

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ChatmessageGetList
// @id 2295
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2295?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取会话消息)
func (client *Client) ChatmessageGetList(request *ChatmessageGetListRequest) (response *ChatmessageGetListResponse, err error) {
	rpcResponse := CreateChatmessageGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ChatmessageGetListRequest struct {
	*api.BaseRequest

	BasicInfo ChatmessageGetListRequestBasicInfo `json:"basicInfo,omitempty"`
	Recall    bool                               `json:"recall,omitempty"`
	From      []int64                            `json:"from,omitempty"`
	PageIndex int64                              `json:"pageIndex,omitempty"`
	Text      string                             `json:"text,omitempty"`
	StartTime string                             `json:"startTime,omitempty"`
	EndTime   string                             `json:"endTime,omitempty"`
	PageSize  int64                              `json:"pageSize,omitempty"`
}

type ChatmessageGetListResponse struct {
	ChatDataInfoList []ChatmessageGetListResponseChatDataInfoList `json:"chatDataInfoList,omitempty"`
	Total            int64                                        `json:"total,omitempty"`
}

func CreateChatmessageGetListRequest() (request *ChatmessageGetListRequest) {
	request = &ChatmessageGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_aewc", "v2.0", "chatmessage/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateChatmessageGetListResponse() (response *api.BaseResponse[ChatmessageGetListResponse]) {
	response = api.CreateResponse[ChatmessageGetListResponse](&ChatmessageGetListResponse{})
	return
}

type ChatmessageGetListRequestBasicInfo struct {
	BosId int64 `json:"bosId,omitempty"`
}

type ChatmessageGetListResponseChatDataInfoList struct {
	Content        ChatmessageGetListResponseContent         `json:"content,omitempty"`
	ChatMediaList  []ChatmessageGetListResponseChatMediaList `json:"chatMediaList,omitempty"`
	Id             int64                                     `json:"id,omitempty"`
	ChatId         string                                    `json:"chatId,omitempty"`
	RoomId         string                                    `json:"roomId,omitempty"`
	MsgType        string                                    `json:"msgType,omitempty"`
	MsgTime        string                                    `json:"msgTime,omitempty"`
	Text           string                                    `json:"text,omitempty"`
	FileName       string                                    `json:"fileName,omitempty"`
	MsgId          string                                    `json:"msgId,omitempty"`
	From           string                                    `json:"from,omitempty"`
	FromUserName   string                                    `json:"fromUserName,omitempty"`
	FromUserAvatar string                                    `json:"fromUserAvatar,omitempty"`
	FromType       int64                                     `json:"fromType,omitempty"`
	To             string                                    `json:"to,omitempty"`
	ToUserName     string                                    `json:"toUserName,omitempty"`
	ToUserAvatar   string                                    `json:"toUserAvatar,omitempty"`
	ToType         int64                                     `json:"toType,omitempty"`
	Seq            int64                                     `json:"seq,omitempty"`
	Recall         bool                                      `json:"recall,omitempty"`
}

type ChatmessageGetListResponseContent struct {
}

type ChatmessageGetListResponseChatMediaList struct {
	Sdkfileid string `json:"sdkfileid,omitempty"`
	Path      string `json:"path,omitempty"`
	FileName  string `json:"fileName,omitempty"`
}
