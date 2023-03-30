package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberOpenPlatformServiceGetMemberInfoPageListAndTotal
// @id 184
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取客户列表)
func (client *Client) MemberOpenPlatformServiceGetMemberInfoPageListAndTotal(request *MemberOpenPlatformServiceGetMemberInfoPageListAndTotalRequest) (response *MemberOpenPlatformServiceGetMemberInfoPageListAndTotalResponse, err error) {
	rpcResponse := CreateMemberOpenPlatformServiceGetMemberInfoPageListAndTotalResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberOpenPlatformServiceGetMemberInfoPageListAndTotalRequest struct {
	*api.BaseRequest

	Sex                   int     `json:"sex,omitempty"`
	Sortby                int     `json:"sortby,omitempty"`
	AscOrDesc             int     `json:"asc_or_desc,omitempty"`
	MemberLevel           string  `json:"member_level,omitempty"`
	Province              string  `json:"province,omitempty"`
	City                  string  `json:"city,omitempty"`
	ConsumingCountMin     int     `json:"consuming_count_min,omitempty"`
	ConsumingCountMax     int     `json:"consuming_count_max,omitempty"`
	AllconsumingAmountMin float64 `json:"allconsuming_amount_min,omitempty"`
	AllconsumingAmountMax float64 `json:"allconsuming_amount_max,omitempty"`
	LastconsumingTimeMin  int64   `json:"lastconsuming_time_min,omitempty"`
	LastconsumingTimeMax  int64   `json:"lastconsuming_time_max,omitempty"`
	IsNeedTagsInfo        bool    `json:"is_need_tags_info,omitempty"`
	Pageindex             int     `json:"pageindex,omitempty"`
	Pagesize              int     `json:"pagesize,omitempty"`
	Nickname              string  `json:"nickname,omitempty"`
	MenbercardId          string  `json:"menbercard_id,omitempty"`
	Phone                 string  `json:"phone,omitempty"`
	Name                  string  `json:"name,omitempty"`
	CreateTime            int64   `json:"create_time,omitempty"`
	RegisterDate          int64   `json:"register_date,omitempty"`
	UpdateTime            int64   `json:"update_time,omitempty"`
	IsFans                bool    `json:"is_fans,omitempty"`
}

type MemberOpenPlatformServiceGetMemberInfoPageListAndTotalResponse struct {
}

func CreateMemberOpenPlatformServiceGetMemberInfoPageListAndTotalRequest() (request *MemberOpenPlatformServiceGetMemberInfoPageListAndTotalRequest) {
	request = &MemberOpenPlatformServiceGetMemberInfoPageListAndTotalRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "1_0", "MemberOpenPlatformService/GetMemberInfoPageListAndTotal", "api")
	request.Method = api.POST
	return
}

func CreateMemberOpenPlatformServiceGetMemberInfoPageListAndTotalResponse() (response *api.BaseResponse[MemberOpenPlatformServiceGetMemberInfoPageListAndTotalResponse]) {
	response = api.CreateResponse[MemberOpenPlatformServiceGetMemberInfoPageListAndTotalResponse](&MemberOpenPlatformServiceGetMemberInfoPageListAndTotalResponse{})
	return
}
