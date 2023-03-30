package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// RightsReturnUpdate
// @id 1843
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1843?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=商家修改退货信息)
func (client *Client) RightsReturnUpdate(request *RightsReturnUpdateRequest) (response *RightsReturnUpdateResponse, err error) {
	rpcResponse := CreateRightsReturnUpdateResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type RightsReturnUpdateRequest struct {
	*api.BaseRequest

	DeliveryInfo   RightsReturnUpdateRequestDeliveryInfo   `json:"deliveryInfo,omitempty"`
	DoorPickupInfo RightsReturnUpdateRequestDoorPickupInfo `json:"doorPickupInfo,omitempty"`
	AutoCompleted  bool                                    `json:"autoCompleted,omitempty"`
	DeliveryType   int64                                   `json:"deliveryType,omitempty"`
	RightsId       int64                                   `json:"rightsId,omitempty"`
}

type RightsReturnUpdateResponse struct {
	Success bool `json:"success,omitempty"`
}

func CreateRightsReturnUpdateRequest() (request *RightsReturnUpdateRequest) {
	request = &RightsReturnUpdateRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "rights/return/update", "apigw")
	request.Method = api.POST
	return
}

func CreateRightsReturnUpdateResponse() (response *api.BaseResponse[RightsReturnUpdateResponse]) {
	response = api.CreateResponse[RightsReturnUpdateResponse](&RightsReturnUpdateResponse{})
	return
}

type RightsReturnUpdateRequestDeliveryInfo struct {
	CompanyCode string `json:"companyCode,omitempty"`
	CompanyName string `json:"companyName,omitempty"`
	DeliveryNo  string `json:"deliveryNo,omitempty"`
}

type RightsReturnUpdateRequestDoorPickupInfo struct {
	AddressInfo      RightsReturnUpdateRequestAddressInfo `json:"addressInfo,omitempty"`
	ExpectPickUpDate int64                                `json:"expectPickUpDate,omitempty"`
	ExpectPickupTime string                               `json:"expectPickupTime,omitempty"`
}

type RightsReturnUpdateRequestAddressInfo struct {
	Name         string `json:"name,omitempty"`
	Phone        string `json:"phone,omitempty"`
	ProvinceName string `json:"provinceName,omitempty"`
	CityName     string `json:"cityName,omitempty"`
	CountyName   string `json:"countyName,omitempty"`
	AreaName     string `json:"areaName,omitempty"`
	Address      string `json:"address,omitempty"`
	ProvinceCode string `json:"provinceCode,omitempty"`
	CityCode     string `json:"cityCode,omitempty"`
	CountyCode   string `json:"countyCode,omitempty"`
	AreaCode     string `json:"areaCode,omitempty"`
}
