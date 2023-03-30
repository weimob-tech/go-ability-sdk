package member_card

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MemberCreateLocalMember
// @id 1794
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=开通微盟会员卡)
func (client *Client) MemberCreateLocalMemberV2_0(request *MemberCreateLocalMemberRequestV2_0) (response *MemberCreateLocalMemberResponseV2_0, err error) {
	rpcResponse := CreateMemberCreateLocalMemberResponseV2_0()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MemberCreateLocalMemberRequestV2_0 struct {
	*api.BaseRequest

	ExtInfo              MemberCreateLocalMemberRequestV2_0ExtInfo `json:"extInfo,omitempty"`
	StoreId              int64                                     `json:"storeId,omitempty"`
	MemberCardTemplateId int64                                     `json:"memberCardTemplateId,omitempty"`
	Phone                string                                    `json:"phone,omitempty"`
	Wid                  int64                                     `json:"wid,omitempty"`
	AppChannel           int                                       `json:"appChannel,omitempty"`
	GetChannel           int                                       `json:"getChannel,omitempty"`
	GetChannelId         int64                                     `json:"getChannelId,omitempty"`
	MemberCardChannel    int                                       `json:"memberCardChannel,omitempty"`
	Name                 string                                    `json:"name,omitempty"`
	Birthday             int                                       `json:"birthday,omitempty"`
	IdCard               string                                    `json:"idCard,omitempty"`
	Address              string                                    `json:"address,omitempty"`
	Province             string                                    `json:"province,omitempty"`
	City                 string                                    `json:"city,omitempty"`
	Area                 string                                    `json:"area,omitempty"`
	ProvinceCode         string                                    `json:"provinceCode,omitempty"`
	CityCode             string                                    `json:"cityCode,omitempty"`
	AreaCode             string                                    `json:"areaCode,omitempty"`
	Mail                 string                                    `json:"mail,omitempty"`
	Sex                  int                                       `json:"sex,omitempty"`
	Education            string                                    `json:"education,omitempty"`
	Hobby                string                                    `json:"hobby,omitempty"`
	Income               string                                    `json:"income,omitempty"`
	Industry             string                                    `json:"industry,omitempty"`
	CustomerStartTime    int                                       `json:"customerStartTime,omitempty"`
	CustomerEndTime      int                                       `json:"customerEndTime,omitempty"`
}

type MemberCreateLocalMemberResponseV2_0 struct {
	Code string `json:"code,omitempty"`
}

func CreateMemberCreateLocalMemberRequestV2_0() (request *MemberCreateLocalMemberRequestV2_0) {
	request = &MemberCreateLocalMemberRequestV2_0{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("member_card", "2_0", "member/createLocalMember", "api")
	request.Method = api.POST
	return
}

func CreateMemberCreateLocalMemberResponseV2_0() (response *api.BaseResponse[MemberCreateLocalMemberResponseV2_0]) {
	response = api.CreateResponse[MemberCreateLocalMemberResponseV2_0](&MemberCreateLocalMemberResponseV2_0{})
	return
}

type MemberCreateLocalMemberRequestV2_0ExtInfo struct {
}
