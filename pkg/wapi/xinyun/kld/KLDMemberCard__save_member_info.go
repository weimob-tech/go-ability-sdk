package kld

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// KLDMemberCardSaveMemberInfo
// @id 213
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=编辑会员信息)
func (client *Client) KLDMemberCardSaveMemberInfo(request *KLDMemberCardSaveMemberInfoRequest) (response *KLDMemberCardSaveMemberInfoResponse, err error) {
	rpcResponse := CreateKLDMemberCardSaveMemberInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type KLDMemberCardSaveMemberInfoRequest struct {
	*api.BaseRequest

	ListOther    []KLDMemberCardSaveMemberInfoRequestOther     `json:"listOther,omitempty"`
	AddressInfo  KLDMemberCardSaveMemberInfoRequestAddressInfo `json:"addressInfo,omitempty"`
	MemberCardNo string                                        `json:"memberCardNo,omitempty"`
	Name         string                                        `json:"name,omitempty"`
	Nickname     string                                        `json:"nickname,omitempty"`
	Sex          int                                           `json:"sex,omitempty"`
	Birthday     string                                        `json:"birthday,omitempty"`
	EMail        string                                        `json:"eMail,omitempty"`
	Degree       int                                           `json:"degree,omitempty"`
	Profession   int                                           `json:"profession,omitempty"`
	Income       int                                           `json:"income,omitempty"`
	HobbyList    []int                                         `json:"hobbyList,omitempty"`
	Operator     string                                        `json:"operator,omitempty"`
	ProvinceName string                                        `json:"provinceName,omitempty"`
	ProvinceId   string                                        `json:"provinceId,omitempty"`
	CityName     string                                        `json:"cityName,omitempty"`
	CityId       string                                        `json:"cityId,omitempty"`
	DistrictName string                                        `json:"districtName,omitempty"`
	DistrictId   string                                        `json:"districtId,omitempty"`
	Address      string                                        `json:"address,omitempty"`
	MapType      int                                           `json:"mapType,omitempty"`
	Longitude    float64                                       `json:"longitude,omitempty"`
	Latitude     float64                                       `json:"latitude,omitempty"`
	Code         string                                        `json:"code,omitempty"`
}

type KLDMemberCardSaveMemberInfoResponse struct {
}

func CreateKLDMemberCardSaveMemberInfoRequest() (request *KLDMemberCardSaveMemberInfoRequest) {
	request = &KLDMemberCardSaveMemberInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("kld", "1_0", "KLDMemberCard/SaveMemberInfo", "api")
	request.Method = api.POST
	return
}

func CreateKLDMemberCardSaveMemberInfoResponse() (response *api.BaseResponse[KLDMemberCardSaveMemberInfoResponse]) {
	response = api.CreateResponse[KLDMemberCardSaveMemberInfoResponse](&KLDMemberCardSaveMemberInfoResponse{})
	return
}

type KLDMemberCardSaveMemberInfoRequestListOther struct {
}

type KLDMemberCardSaveMemberInfoRequestAddressInfo struct {
}
