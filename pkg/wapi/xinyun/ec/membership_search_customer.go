package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MembershipSearchCustomer
// @id 1821
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=客户列表搜索接口)
func (client *Client) MembershipSearchCustomer(request *MembershipSearchCustomerRequest) (response *MembershipSearchCustomerResponse, err error) {
	rpcResponse := CreateMembershipSearchCustomerResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MembershipSearchCustomerRequest struct {
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
	PageSize             string `json:"pageSize,omitempty"`
	PageNum              string `json:"pageNum,omitempty"`
	Wid                  int64  `json:"wid,omitempty"`
	CustomerStartTime    int    `json:"customerStartTime,omitempty"`
	CustomerEndTime      int    `json:"customerEndTime,omitempty"`
	MembershipStartTime  int    `json:"membershipStartTime,omitempty"`
	MembershipEndTime    int    `json:"membershipEndTime,omitempty"`
	McInfoCardStartTime  int    `json:"mcInfoCardStartTime,omitempty"`
	McInfoCardEndTime    int    `json:"mcInfoCardEndTime,omitempty"`
}

type MembershipSearchCustomerResponse struct {
}

func CreateMembershipSearchCustomerRequest() (request *MembershipSearchCustomerRequest) {
	request = &MembershipSearchCustomerRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "membership/searchCustomer", "api")
	request.Method = api.POST
	return
}

func CreateMembershipSearchCustomerResponse() (response *api.BaseResponse[MembershipSearchCustomerResponse]) {
	response = api.CreateResponse[MembershipSearchCustomerResponse](&MembershipSearchCustomerResponse{})
	return
}
