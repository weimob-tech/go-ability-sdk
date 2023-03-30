package o2o

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberPointList
// @id 424
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=该会员调整记录)
func (client *Client) MemberPointList(request *MemberPointListRequest) (response *MemberPointListResponse, err error) {
	rpcResponse := CreateMemberPointListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberPointListRequest struct {
	*api.BaseRequest

	MerchantId int64  `json:"merchantId,omitempty"`
	CustomerId int64  `json:"customerId,omitempty"`
	ChangeType int    `json:"changeType,omitempty"`
	StartTime  string `json:"startTime,omitempty"`
	EndTime    string `json:"endTime,omitempty"`
	PageNum    int    `json:"pageNum,omitempty"`
	PageSize   int    `json:"pageSize,omitempty"`
}

type MemberPointListResponse struct {
}

func CreateMemberPointListRequest() (request *MemberPointListRequest) {
	request = &MemberPointListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("o2o", "1_0", "memberPoint/list", "api")
	request.Method = api.POST
	return
}

func CreateMemberPointListResponse() (response *api.BaseResponse[MemberPointListResponse]) {
	response = api.CreateResponse[MemberPointListResponse](&MemberPointListResponse{})
	return
}
