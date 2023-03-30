package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberGetMcBasicInfoListByWid
// @id 1796
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询会员卡基本信息)
func (client *Client) MemberGetMcBasicInfoListByWidV2_0(request *MemberGetMcBasicInfoListByWidRequestV2_0) (response *MemberGetMcBasicInfoListByWidResponseV2_0, err error) {
	rpcResponse := CreateMemberGetMcBasicInfoListByWidResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberGetMcBasicInfoListByWidRequestV2_0 struct {
	*api.BaseRequest

	Wid int64 `json:"wid,omitempty"`
}

type MemberGetMcBasicInfoListByWidResponseV2_0 struct {
	Items []MemberGetMcBasicInfoListByWidResponseV2_0Items `json:"items,omitempty"`
}

func CreateMemberGetMcBasicInfoListByWidRequestV2_0() (request *MemberGetMcBasicInfoListByWidRequestV2_0) {
	request = &MemberGetMcBasicInfoListByWidRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "2_0", "member/getMcBasicInfoListByWid", "api")
	request.Method = api.POST
	return
}

func CreateMemberGetMcBasicInfoListByWidResponseV2_0() (response *api.BaseResponse[MemberGetMcBasicInfoListByWidResponseV2_0]) {
	response = api.CreateResponse[MemberGetMcBasicInfoListByWidResponseV2_0](&MemberGetMcBasicInfoListByWidResponseV2_0{})
	return
}

type MemberGetMcBasicInfoListByWidResponseV2_0Items struct {
	Pid                  int64  `json:"pid,omitempty"`
	Code                 string `json:"code,omitempty"`
	Wid                  int64  `json:"wid,omitempty"`
	MemberCardTemplateId int64  `json:"memberCardTemplateId,omitempty"`
	StoreId              int64  `json:"storeId,omitempty"`
	Status               int    `json:"status,omitempty"`
	StartTime            int64  `json:"startTime,omitempty"`
	ExpireTime           int64  `json:"expireTime,omitempty"`
	GetChannel           int    `json:"getChannel,omitempty"`
	GetChannelId         int    `json:"getChannelId,omitempty"`
	CardPublishTime      int64  `json:"cardPublishTime,omitempty"`
	Growth               int64  `json:"growth,omitempty"`
	Discount             int    `json:"discount,omitempty"`
	Rank                 string `json:"rank,omitempty"`
	RankId               int64  `json:"rankId,omitempty"`
	RankExpType          int    `json:"rankExpType,omitempty"`
	RankStartDate        int64  `json:"rankStartDate,omitempty"`
	RankEndDate          int64  `json:"rankEndDate,omitempty"`
	MemberCardChannel    int    `json:"memberCardChannel,omitempty"`
}
