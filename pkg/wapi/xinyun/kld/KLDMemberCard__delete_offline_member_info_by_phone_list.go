package kld

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// KLDMemberCardDeleteOfflineMemberInfoByPhoneList
// @id 232
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除实体店会员)
func (client *Client) KLDMemberCardDeleteOfflineMemberInfoByPhoneList(request *KLDMemberCardDeleteOfflineMemberInfoByPhoneListRequest) (response *KLDMemberCardDeleteOfflineMemberInfoByPhoneListResponse, err error) {
	rpcResponse := CreateKLDMemberCardDeleteOfflineMemberInfoByPhoneListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type KLDMemberCardDeleteOfflineMemberInfoByPhoneListRequest struct {
	*api.BaseRequest

	PhoneList []string `json:"phoneList,omitempty"`
}

type KLDMemberCardDeleteOfflineMemberInfoByPhoneListResponse struct {
}

func CreateKLDMemberCardDeleteOfflineMemberInfoByPhoneListRequest() (request *KLDMemberCardDeleteOfflineMemberInfoByPhoneListRequest) {
	request = &KLDMemberCardDeleteOfflineMemberInfoByPhoneListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("kld", "1_0", "KLDMemberCard/DeleteOfflineMemberInfoByPhoneList", "api")
	request.Method = api.POST
	return
}

func CreateKLDMemberCardDeleteOfflineMemberInfoByPhoneListResponse() (response *api.BaseResponse[KLDMemberCardDeleteOfflineMemberInfoByPhoneListResponse]) {
	response = api.CreateResponse[KLDMemberCardDeleteOfflineMemberInfoByPhoneListResponse](&KLDMemberCardDeleteOfflineMemberInfoByPhoneListResponse{})
	return
}
