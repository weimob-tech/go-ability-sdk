package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembershipRollSearchCustomer
// @id 1822
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=客户列表滚动接口)
func (client *Client) MembershipRollSearchCustomer(request *MembershipRollSearchCustomerRequest) (response *MembershipRollSearchCustomerResponse, err error) {
	rpcResponse := CreateMembershipRollSearchCustomerResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembershipRollSearchCustomerRequest struct {
	*api.BaseRequest

	Name                 string `json:"name,omitempty"`
	Phone                string `json:"phone,omitempty"`
	McInfoCode           string `json:"mcInfoCode,omitempty"`
	McInfoCardTemplateId int64  `json:"mcInfoCardTemplateId,omitempty"`
	BizUserId            string `json:"bizUserId,omitempty"`
	StoreId              int64  `json:"storeId,omitempty"`
	AppChannel           int    `json:"appChannel,omitempty"`
	HomeStoreId          int64  `json:"homeStoreId,omitempty"`
	McInfoRankId         int64  `json:"mcInfoRankId,omitempty"`
	IsMember             bool   `json:"isMember,omitempty"`
	Cursor               string `json:"cursor,omitempty"`
	Wid                  int64  `json:"wid,omitempty"`
	CustomerStartTime    int    `json:"customerStartTime,omitempty"`
	CustomerEndTime      int    `json:"customerEndTime,omitempty"`
	MembershipStartTime  int    `json:"membershipStartTime,omitempty"`
	MembershipEndTime    int    `json:"membershipEndTime,omitempty"`
	McInfoCardStartTime  int    `json:"mcInfoCardStartTime,omitempty"`
	McInfoCardEndTime    int    `json:"mcInfoCardEndTime,omitempty"`
	ScrollSize           int    `json:"scrollSize,omitempty"`
}

type MembershipRollSearchCustomerResponse struct {
}

func CreateMembershipRollSearchCustomerRequest() (request *MembershipRollSearchCustomerRequest) {
	request = &MembershipRollSearchCustomerRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "membership/rollSearchCustomer", "api")
	request.Method = api.POST
	return
}

func CreateMembershipRollSearchCustomerResponse() (response *api.BaseResponse[MembershipRollSearchCustomerResponse]) {
	response = api.CreateResponse[MembershipRollSearchCustomerResponse](&MembershipRollSearchCustomerResponse{})
	return
}
