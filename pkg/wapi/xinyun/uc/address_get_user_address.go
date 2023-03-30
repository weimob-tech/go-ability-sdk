package uc

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// AddressGetUserAddress
// @id 1909
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=获取用户收货地址列表（新）)
func (client *Client) AddressGetUserAddress(request *AddressGetUserAddressRequest) (response *AddressGetUserAddressResponse, err error) {
	rpcResponse := CreateAddressGetUserAddressResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type AddressGetUserAddressRequest struct {
	*api.BaseRequest

	Pid   int64 `json:"pid,omitempty"`
	Wid   int64 `json:"wid,omitempty"`
	Scope int   `json:"scope,omitempty"`
}

type AddressGetUserAddressResponse struct {
	Items      AddressGetUserAddressResponseItems `json:"items,omitempty"`
	TotalCount int                                `json:"totalCount,omitempty"`
}

func CreateAddressGetUserAddressRequest() (request *AddressGetUserAddressRequest) {
	request = &AddressGetUserAddressRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("uc", "1_0", "address/getUserAddress", "api")
	request.Method = api.POST
	return
}

func CreateAddressGetUserAddressResponse() (response *api.BaseResponse[AddressGetUserAddressResponse]) {
	response = api.CreateResponse[AddressGetUserAddressResponse](&AddressGetUserAddressResponse{})
	return
}

type AddressGetUserAddressResponseItems struct {
	Id           string `json:"id,omitempty"`
	Wid          string `json:"wid,omitempty"`
	ProvinceId   string `json:"provinceId,omitempty"`
	ProvinceName string `json:"provinceName,omitempty"`
	CityId       string `json:"cityId,omitempty"`
	CityName     string `json:"cityName,omitempty"`
	DistrctId    string `json:"distrctId,omitempty"`
	DistrctName  string `json:"distrctName,omitempty"`
	Address      string `json:"address,omitempty"`
	IsDefault    string `json:"isDefault,omitempty"`
	Name         string `json:"name,omitempty"`
	Phone        string `json:"phone,omitempty"`
	PostalCode   string `json:"postalCode,omitempty"`
	StreetId     string `json:"streetId,omitempty"`
	StreetName   string `json:"streetName,omitempty"`
	IdentityCard string `json:"identityCard,omitempty"`
	MapType      string `json:"mapType,omitempty"`
	Longitude    string `json:"longitude,omitempty"`
	Latitude     string `json:"latitude,omitempty"`
}
