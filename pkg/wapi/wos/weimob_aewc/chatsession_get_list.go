package weimob_aewc

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ChatsessionGetList
// @id 2294
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/2294?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取员工会话)
func (client *Client) ChatsessionGetList(request *ChatsessionGetListRequest) (response *ChatsessionGetListResponse, err error) {
	rpcResponse := CreateChatsessionGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ChatsessionGetListRequest struct {
	*api.BaseRequest

	BasicInfo  ChatsessionGetListRequestBasicInfo `json:"basicInfo,omitempty"`
	UserId     string                             `json:"userId,omitempty"`
	ChatType   int64                              `json:"chatType,omitempty"`
	PersonType int64                              `json:"personType,omitempty"`
	PageIndex  int64                              `json:"pageIndex,omitempty"`
	PageSize   int64                              `json:"pageSize,omitempty"`
}

type ChatsessionGetListResponse struct {
	ContactsListOutList []ChatsessionGetListResponseContactsListOutList `json:"contactsListOutList,omitempty"`
	Total               int64                                           `json:"total,omitempty"`
}

func CreateChatsessionGetListRequest() (request *ChatsessionGetListRequest) {
	request = &ChatsessionGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_aewc", "v2.0", "chatsession/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateChatsessionGetListResponse() (response *api.BaseResponse[ChatsessionGetListResponse]) {
	response = api.CreateResponse[ChatsessionGetListResponse](&ChatsessionGetListResponse{})
	return
}

type ChatsessionGetListRequestBasicInfo struct {
	BosId int64 `json:"bosId,omitempty"`
}

type ChatsessionGetListResponseContactsListOutList struct {
	UserId      string `json:"userId,omitempty"`
	NickName    string `json:"nickName,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	ChatId      string `json:"chatId,omitempty"`
	ChatType    int64  `json:"chatType,omitempty"`
	LastMsg     string `json:"lastMsg,omitempty"`
	LastMsgTime string `json:"lastMsgTime,omitempty"`
}
