package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ShopAddShopInfo
// @id 286
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=新增门店)
func (client *Client) ShopAddShopInfo(request *ShopAddShopInfoRequest) (response *ShopAddShopInfoResponse, err error) {
	rpcResponse := CreateShopAddShopInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ShopAddShopInfoRequest struct {
	*api.BaseRequest

	ShopAccount      string `json:"shop_account,omitempty"`
	ShopPassword     string `json:"shop_password,omitempty"`
	ShopName         string `json:"shop_name,omitempty"`
	Tel              string `json:"tel,omitempty"`
	BusinessTime     string `json:"business_time,omitempty"`
	BusinessWorkdate string `json:"business_workdate,omitempty"`
	ProvinceName     string `json:"province_name,omitempty"`
	DistrictName     string `json:"district_name,omitempty"`
	CityName         string `json:"city_name,omitempty"`
	LocateDetail     string `json:"locate_detail,omitempty"`
	Longitude        string `json:"longitude,omitempty"`
	Latitude         string `json:"latitude,omitempty"`
	IsPayonline      bool   `json:"is_payonline,omitempty"`
	IsPayoffline     bool   `json:"is_payoffline,omitempty"`
	ShopStore        int    `json:"shop_store,omitempty"`
	IsOnlyservice    bool   `json:"is_onlyservice,omitempty"`
	IsRecommend      bool   `json:"is_recommend,omitempty"`
	AreaCode         string `json:"area_code,omitempty"`
	Mobile           string `json:"mobile,omitempty"`
}

type ShopAddShopInfoResponse struct {
}

func CreateShopAddShopInfoRequest() (request *ShopAddShopInfoRequest) {
	request = &ShopAddShopInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "shop/AddShopInfo", "api")
	request.Method = api.POST
	return
}

func CreateShopAddShopInfoResponse() (response *api.BaseResponse[ShopAddShopInfoResponse]) {
	response = api.CreateResponse[ShopAddShopInfoResponse](&ShopAddShopInfoResponse{})
	return
}
