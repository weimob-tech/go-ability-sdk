package wangpu

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// ShopUpdateShopInfo
// @id 284
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=更改门店信息)
func (client *Client) ShopUpdateShopInfo(request *ShopUpdateShopInfoRequest) (response *ShopUpdateShopInfoResponse, err error) {
	rpcResponse := CreateShopUpdateShopInfoResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type ShopUpdateShopInfoRequest struct {
	*api.BaseRequest

	ShopId           int64  `json:"shop_id,omitempty"`
	ShopAccount      string `json:"shop_account,omitempty"`
	ShopPassword     string `json:"shop_password,omitempty"`
	ShopName         string `json:"shop_name,omitempty"`
	Tel              string `json:"tel,omitempty"`
	BusinessTime     string `json:"business_time,omitempty"`
	BusinessWorkdate string `json:"business_workdate,omitempty"`
	LocateDetail     string `json:"locate_detail,omitempty"`
	IsDel            bool   `json:"is_del,omitempty"`
	IsOnbusiness     bool   `json:"is_onbusiness,omitempty"`
	Longitude        string `json:"longitude,omitempty"`
	Latitude         string `json:"latitude,omitempty"`
	IsPayonline      bool   `json:"is_payonline,omitempty"`
	IsPayoffline     bool   `json:"is_payoffline,omitempty"`
	ShopStore        int    `json:"shop_store,omitempty"`
	IsOnlyservice    bool   `json:"is_onlyservice,omitempty"`
	IsRecommend      bool   `json:"is_recommend,omitempty"`
}

type ShopUpdateShopInfoResponse struct {
}

func CreateShopUpdateShopInfoRequest() (request *ShopUpdateShopInfoRequest) {
	request = &ShopUpdateShopInfoRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("wangpu", "1_0", "shop/UpdateShopInfo", "api")
	request.Method = api.POST
	return
}

func CreateShopUpdateShopInfoResponse() (response *api.BaseResponse[ShopUpdateShopInfoResponse]) {
	response = api.CreateResponse[ShopUpdateShopInfoResponse](&ShopUpdateShopInfoResponse{})
	return
}
