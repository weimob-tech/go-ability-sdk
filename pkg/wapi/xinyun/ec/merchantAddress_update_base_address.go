package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantAddressUpdateBaseAddress
// @id 2441
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更新地址库地址)
func (client *Client) MerchantAddressUpdateBaseAddress(request *MerchantAddressUpdateBaseAddressRequest) (response *MerchantAddressUpdateBaseAddressResponse, err error) {
	rpcResponse := CreateMerchantAddressUpdateBaseAddressResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantAddressUpdateBaseAddressRequest struct {
	*api.BaseRequest

	AddressId             int64   `json:"addressId,omitempty"`
	ContactPerson         string  `json:"contactPerson,omitempty"`
	Tel                   string  `json:"tel,omitempty"`
	PostCode              string  `json:"postCode,omitempty"`
	ProvinceCode          string  `json:"provinceCode,omitempty"`
	CityCode              string  `json:"cityCode,omitempty"`
	CountyCode            string  `json:"countyCode,omitempty"`
	AreaCode              string  `json:"areaCode,omitempty"`
	ProvinceName          string  `json:"provinceName,omitempty"`
	CityName              string  `json:"cityName,omitempty"`
	CountyName            string  `json:"countyName,omitempty"`
	AreaName              string  `json:"areaName,omitempty"`
	Address               string  `json:"address,omitempty"`
	Longitude             float64 `json:"longitude,omitempty"`
	Latitude              float64 `json:"latitude,omitempty"`
	IsDefaultReleaseGoods int     `json:"isDefaultReleaseGoods,omitempty"`
	IsDefaultBackGoods    int     `json:"isDefaultBackGoods,omitempty"`
	Pid                   int64   `json:"pid,omitempty"`
	StoreId               int64   `json:"storeId,omitempty"`
}

type MerchantAddressUpdateBaseAddressResponse struct {
	IsSucceed bool `json:"isSucceed,omitempty"`
}

func CreateMerchantAddressUpdateBaseAddressRequest() (request *MerchantAddressUpdateBaseAddressRequest) {
	request = &MerchantAddressUpdateBaseAddressRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchantAddress/updateBaseAddress", "api")
	request.Method = api.POST
	return
}

func CreateMerchantAddressUpdateBaseAddressResponse() (response *api.BaseResponse[MerchantAddressUpdateBaseAddressResponse]) {
	response = api.CreateResponse[MerchantAddressUpdateBaseAddressResponse](&MerchantAddressUpdateBaseAddressResponse{})
	return
}
