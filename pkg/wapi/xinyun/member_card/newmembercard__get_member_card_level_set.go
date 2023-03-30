package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// NewmembercardGetMemberCardLevelSet
// @id 105
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取会员卡等级配置信息)
func (client *Client) NewmembercardGetMemberCardLevelSet(request *NewmembercardGetLevelSetRequest) (response *NewmembercardGetLevelSetResponse, err error) {
	rpcResponse := CreateNewmembercardGetLevelSetResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type NewmembercardGetLevelSetRequest struct {
	*api.BaseRequest
}

type NewmembercardGetLevelSetResponse struct {
}

func CreateNewmembercardGetLevelSetRequest() (request *NewmembercardGetLevelSetRequest) {
	request = &NewmembercardGetLevelSetRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "newmembercard/GetMemberCardLevelSet", "api")
	request.Method = api.POST
	return
}

func CreateNewmembercardGetLevelSetResponse() (response *api.BaseResponse[NewmembercardGetLevelSetResponse]) {
	response = api.CreateResponse[NewmembercardGetLevelSetResponse](&NewmembercardGetLevelSetResponse{})
	return
}
