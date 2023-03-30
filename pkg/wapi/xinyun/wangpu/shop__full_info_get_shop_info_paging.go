package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ShopFullInfoGetShopInfoPaging
// @id 331
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=分页获取门店信息)
func (client *Client) ShopFullInfoGetShopInfoPaging(request *ShopFullInfoGetShopInfoPagingRequest) (response *ShopFullInfoGetShopInfoPagingResponse, err error) {
	rpcResponse := CreateShopFullInfoGetShopInfoPagingResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ShopFullInfoGetShopInfoPagingRequest struct {
	*api.BaseRequest

	LocateCityId     int    `json:"locateCityId,omitempty"`
	LocateProvinceId int    `json:"locateProvinceId,omitempty"`
	LocateDistrictId int    `json:"locateDistrictId,omitempty"`
	ShopLocateDetail string `json:"shopLocateDetail,omitempty"`
	PageSize         int    `json:"page_size,omitempty"`
	PageNo           int    `json:"page_no,omitempty"`
}

type ShopFullInfoGetShopInfoPagingResponse struct {
}

func CreateShopFullInfoGetShopInfoPagingRequest() (request *ShopFullInfoGetShopInfoPagingRequest) {
	request = &ShopFullInfoGetShopInfoPagingRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "shop/FullInfoGetShopInfoPaging", "api")
	request.Method = api.POST
	return
}

func CreateShopFullInfoGetShopInfoPagingResponse() (response *api.BaseResponse[ShopFullInfoGetShopInfoPagingResponse]) {
	response = api.CreateResponse[ShopFullInfoGetShopInfoPagingResponse](&ShopFullInfoGetShopInfoPagingResponse{})
	return
}
