package kld

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// KLDMemberCardAddOfflineMemberInfo
// @id 221
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=添加实体店会员)
func (client *Client) KLDMemberCardAddOfflineMemberInfo(request *KLDMemberCardAddOfflineMemberInfoRequest) (response *KLDMemberCardAddOfflineMemberInfoResponse, err error) {
	rpcResponse := CreateKLDMemberCardAddOfflineMemberInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type KLDMemberCardAddOfflineMemberInfoRequest struct {
	*api.BaseRequest

	Address            KLDMemberCardAddOfflineMemberInfoRequestAddress `json:"address,omitempty"`
	Name               string                                          `json:"name,omitempty"`
	Phone              string                                          `json:"phone,omitempty"`
	Sex                int                                             `json:"sex,omitempty"`
	Birthday           string                                          `json:"birthday,omitempty"`
	GrowthValue        int                                             `json:"growthValue,omitempty"`
	Points             int                                             `json:"points,omitempty"`
	AllPoints          int64                                           `json:"allPoints,omitempty"`
	Amount             float64                                         `json:"amount,omitempty"`
	AllConsumingAmount float64                                         `json:"allConsumingAmount,omitempty"`
}

type KLDMemberCardAddOfflineMemberInfoResponse struct {
}

func CreateKLDMemberCardAddOfflineMemberInfoRequest() (request *KLDMemberCardAddOfflineMemberInfoRequest) {
	request = &KLDMemberCardAddOfflineMemberInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("kld", "1_0", "KLDMemberCard/AddOfflineMemberInfo", "api")
	request.Method = api.POST
	return
}

func CreateKLDMemberCardAddOfflineMemberInfoResponse() (response *api.BaseResponse[KLDMemberCardAddOfflineMemberInfoResponse]) {
	response = api.CreateResponse[KLDMemberCardAddOfflineMemberInfoResponse](&KLDMemberCardAddOfflineMemberInfoResponse{})
	return
}

type KLDMemberCardAddOfflineMemberInfoRequestAddress struct {
	ProvinceName string `json:"provinceName,omitempty"`
	ProvinceId   string `json:"provinceId,omitempty"`
	CityName     string `json:"cityName,omitempty"`
	CityId       string `json:"cityId,omitempty"`
	DistrictName string `json:"districtName,omitempty"`
	DistrictId   string `json:"districtId,omitempty"`
	Address      string `json:"address,omitempty"`
	MapType      int64  `json:"mapType,omitempty"`
	Longitude    int64  `json:"longitude,omitempty"`
	Latitude     int64  `json:"latitude,omitempty"`
	Code         string `json:"code,omitempty"`
}
