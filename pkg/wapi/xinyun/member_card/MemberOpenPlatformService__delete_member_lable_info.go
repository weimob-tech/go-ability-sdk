package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberOpenPlatformServiceDeleteMemberLableInfo
// @id 170
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除客户标签)
func (client *Client) MemberOpenPlatformServiceDeleteMemberLableInfo(request *MemberOpenPlatformServiceDeleteMemberLableInfoRequest) (response *MemberOpenPlatformServiceDeleteMemberLableInfoResponse, err error) {
	rpcResponse := CreateMemberOpenPlatformServiceDeleteMemberLableInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberOpenPlatformServiceDeleteMemberLableInfoRequest struct {
	*api.BaseRequest

	LableIds    []int    `json:"lable_ids,omitempty"`
	CustomerNos []string `json:"customer_nos,omitempty"`
	Remark      string   `json:"remark,omitempty"`
}

type MemberOpenPlatformServiceDeleteMemberLableInfoResponse struct {
}

func CreateMemberOpenPlatformServiceDeleteMemberLableInfoRequest() (request *MemberOpenPlatformServiceDeleteMemberLableInfoRequest) {
	request = &MemberOpenPlatformServiceDeleteMemberLableInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "MemberOpenPlatformService/DeleteMemberLableInfo", "api")
	request.Method = api.POST
	return
}

func CreateMemberOpenPlatformServiceDeleteMemberLableInfoResponse() (response *api.BaseResponse[MemberOpenPlatformServiceDeleteMemberLableInfoResponse]) {
	response = api.CreateResponse[MemberOpenPlatformServiceDeleteMemberLableInfoResponse](&MemberOpenPlatformServiceDeleteMemberLableInfoResponse{})
	return
}
