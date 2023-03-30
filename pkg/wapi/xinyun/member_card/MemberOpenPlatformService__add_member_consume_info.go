package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberOpenPlatformServiceAddMemberConsumeInfo
// @id 172
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增后台消费)
func (client *Client) MemberOpenPlatformServiceAddMemberConsumeInfo(request *MemberOpenPlatformServiceAddMemberConsumeInfoRequest) (response *MemberOpenPlatformServiceAddMemberConsumeInfoResponse, err error) {
	rpcResponse := CreateMemberOpenPlatformServiceAddMemberConsumeInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberOpenPlatformServiceAddMemberConsumeInfoRequest struct {
	*api.BaseRequest

	CustomerNos []string `json:"customer_nos,omitempty"`
	TotalAmount float64  `json:"total_amount,omitempty"`
	RealAmount  float64  `json:"real_amount,omitempty"`
	Balance     float64  `json:"balance,omitempty"`
	Cash        float64  `json:"cash,omitempty"`
	Point       int      `json:"point,omitempty"`
	PointAmount float64  `json:"point_amount,omitempty"`
	ReturnCash  float64  `json:"return_cash,omitempty"`
	DisCount    int      `json:"dis_count,omitempty"`
	Password    string   `json:"password,omitempty"`
	Datasource  int      `json:"datasource,omitempty"`
	Opreater    string   `json:"opreater,omitempty"`
	Remark      string   `json:"remark,omitempty"`
	Mulshopid   int64    `json:"mulshopid,omitempty"`
}

type MemberOpenPlatformServiceAddMemberConsumeInfoResponse struct {
}

func CreateMemberOpenPlatformServiceAddMemberConsumeInfoRequest() (request *MemberOpenPlatformServiceAddMemberConsumeInfoRequest) {
	request = &MemberOpenPlatformServiceAddMemberConsumeInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "MemberOpenPlatformService/AddMemberConsumeInfo", "api")
	request.Method = api.POST
	return
}

func CreateMemberOpenPlatformServiceAddMemberConsumeInfoResponse() (response *api.BaseResponse[MemberOpenPlatformServiceAddMemberConsumeInfoResponse]) {
	response = api.CreateResponse[MemberOpenPlatformServiceAddMemberConsumeInfoResponse](&MemberOpenPlatformServiceAddMemberConsumeInfoResponse{})
	return
}
