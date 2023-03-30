package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembershipQueryUserRecord
// @id 1414
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=会员行为记录查询)
func (client *Client) MembershipQueryUserRecord(request *MembershipQueryUserRecordRequest) (response *MembershipQueryUserRecordResponse, err error) {
	rpcResponse := CreateMembershipQueryUserRecordResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembershipQueryUserRecordRequest struct {
	*api.BaseRequest

	Wid             int64  `json:"wid,omitempty"`
	StoreId         int64  `json:"storeId,omitempty"`
	CreateBeginDate string `json:"createBeginDate,omitempty"`
	CreateEndDate   string `json:"createEndDate,omitempty"`
	StartIndex      int    `json:"startIndex,omitempty"`
	PageSize        int    `json:"pageSize,omitempty"`
	BizType         int    `json:"bizType,omitempty"`
}

type MembershipQueryUserRecordResponse struct {
}

func CreateMembershipQueryUserRecordRequest() (request *MembershipQueryUserRecordRequest) {
	request = &MembershipQueryUserRecordRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "membership/queryUserRecord", "api")
	request.Method = api.POST
	return
}

func CreateMembershipQueryUserRecordResponse() (response *api.BaseResponse[MembershipQueryUserRecordResponse]) {
	response = api.CreateResponse[MembershipQueryUserRecordResponse](&MembershipQueryUserRecordResponse{})
	return
}
