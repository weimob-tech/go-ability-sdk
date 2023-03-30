package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantAddressQueryDefaultAddress
// @id 2443
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=查询默认发货地址)
func (client *Client) MerchantAddressQueryDefaultAddress(request *MerchantAddressQueryDefaultAddressRequest) (response *MerchantAddressQueryDefaultAddressResponse, err error) {
	rpcResponse := CreateMerchantAddressQueryDefaultAddressResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantAddressQueryDefaultAddressRequest struct {
	*api.BaseRequest

	Pid     int64 `json:"pid,omitempty"`
	StoreId int64 `json:"storeId,omitempty"`
}

type MerchantAddressQueryDefaultAddressResponse struct {
	Id                    int64   `json:"id,omitempty"`
	Pid                   int64   `json:"pid,omitempty"`
	StoreId               int64   `json:"storeId,omitempty"`
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
}

func CreateMerchantAddressQueryDefaultAddressRequest() (request *MerchantAddressQueryDefaultAddressRequest) {
	request = &MerchantAddressQueryDefaultAddressRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchantAddress/queryDefaultAddress", "api")
	request.Method = api.POST
	return
}

func CreateMerchantAddressQueryDefaultAddressResponse() (response *api.BaseResponse[MerchantAddressQueryDefaultAddressResponse]) {
	response = api.CreateResponse[MerchantAddressQueryDefaultAddressResponse](&MerchantAddressQueryDefaultAddressResponse{})
	return
}
