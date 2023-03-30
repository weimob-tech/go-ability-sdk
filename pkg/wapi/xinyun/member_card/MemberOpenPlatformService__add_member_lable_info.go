package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberOpenPlatformServiceAddMemberLableInfo
// @id 173
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增客户标签)
func (client *Client) MemberOpenPlatformServiceAddMemberLableInfo(request *MemberOpenPlatformServiceAddMemberLableInfoRequest) (response *MemberOpenPlatformServiceAddMemberLableInfoResponse, err error) {
	rpcResponse := CreateMemberOpenPlatformServiceAddMemberLableInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberOpenPlatformServiceAddMemberLableInfoRequest struct {
	*api.BaseRequest

	LableIds    []int    `json:"lable_ids,omitempty"`
	CustomerNos []string `json:"customer_nos,omitempty"`
	Remark      string   `json:"remark,omitempty"`
}

type MemberOpenPlatformServiceAddMemberLableInfoResponse struct {
}

func CreateMemberOpenPlatformServiceAddMemberLableInfoRequest() (request *MemberOpenPlatformServiceAddMemberLableInfoRequest) {
	request = &MemberOpenPlatformServiceAddMemberLableInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "MemberOpenPlatformService/AddMemberLableInfo", "api")
	request.Method = api.POST
	return
}

func CreateMemberOpenPlatformServiceAddMemberLableInfoResponse() (response *api.BaseResponse[MemberOpenPlatformServiceAddMemberLableInfoResponse]) {
	response = api.CreateResponse[MemberOpenPlatformServiceAddMemberLableInfoResponse](&MemberOpenPlatformServiceAddMemberLableInfoResponse{})
	return
}
