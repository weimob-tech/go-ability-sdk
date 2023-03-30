package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NewmembercardDeleteOfflineMemberInfoByPhoneList
// @id 119
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除实体会员)
func (client *Client) NewmembercardDeleteOfflineMemberInfoByPhoneList(request *NewmembercardDeleteOfflineMemberInfoByPhoneListRequest) (response *NewmembercardDeleteOfflineMemberInfoByPhoneListResponse, err error) {
	rpcResponse := CreateNewmembercardDeleteOfflineMemberInfoByPhoneListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NewmembercardDeleteOfflineMemberInfoByPhoneListRequest struct {
	*api.BaseRequest

	PhoneList []string `json:"phoneList,omitempty"`
}

type NewmembercardDeleteOfflineMemberInfoByPhoneListResponse struct {
}

func CreateNewmembercardDeleteOfflineMemberInfoByPhoneListRequest() (request *NewmembercardDeleteOfflineMemberInfoByPhoneListRequest) {
	request = &NewmembercardDeleteOfflineMemberInfoByPhoneListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "newmembercard/DeleteOfflineMemberInfoByPhoneList", "api")
	request.Method = api.POST
	return
}

func CreateNewmembercardDeleteOfflineMemberInfoByPhoneListResponse() (response *api.BaseResponse[NewmembercardDeleteOfflineMemberInfoByPhoneListResponse]) {
	response = api.CreateResponse[NewmembercardDeleteOfflineMemberInfoByPhoneListResponse](&NewmembercardDeleteOfflineMemberInfoByPhoneListResponse{})
	return
}
