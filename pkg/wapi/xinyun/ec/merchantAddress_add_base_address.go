package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantAddressAddBaseAddress
// @id 2442
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增地址库地址)
func (client *Client) MerchantAddressAddBaseAddress(request *MerchantAddressAddBaseAddressRequest) (response *MerchantAddressAddBaseAddressResponse, err error) {
	rpcResponse := CreateMerchantAddressAddBaseAddressResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantAddressAddBaseAddressRequest struct {
	*api.BaseRequest

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

type MerchantAddressAddBaseAddressResponse struct {
	IsSucceed string `json:"isSucceed,omitempty"`
}

func CreateMerchantAddressAddBaseAddressRequest() (request *MerchantAddressAddBaseAddressRequest) {
	request = &MerchantAddressAddBaseAddressRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchantAddress/addBaseAddress", "api")
	request.Method = api.POST
	return
}

func CreateMerchantAddressAddBaseAddressResponse() (response *api.BaseResponse[MerchantAddressAddBaseAddressResponse]) {
	response = api.CreateResponse[MerchantAddressAddBaseAddressResponse](&MerchantAddressAddBaseAddressResponse{})
	return
}
