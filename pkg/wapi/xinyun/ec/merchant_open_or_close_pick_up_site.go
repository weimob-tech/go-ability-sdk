package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantOpenOrClosePickUpSite
// @id 1183
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=启用/关闭自提点)
func (client *Client) MerchantOpenOrClosePickUpSite(request *MerchantOpenOrClosePickUpSiteRequest) (response *MerchantOpenOrClosePickUpSiteResponse, err error) {
	rpcResponse := CreateMerchantOpenOrClosePickUpSiteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantOpenOrClosePickUpSiteRequest struct {
	*api.BaseRequest

	SiteId int64 `json:"siteId,omitempty"`
	IsUse  int   `json:"isUse,omitempty"`
}

type MerchantOpenOrClosePickUpSiteResponse struct {
}

func CreateMerchantOpenOrClosePickUpSiteRequest() (request *MerchantOpenOrClosePickUpSiteRequest) {
	request = &MerchantOpenOrClosePickUpSiteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchant/openOrClosePickUpSite", "api")
	request.Method = api.POST
	return
}

func CreateMerchantOpenOrClosePickUpSiteResponse() (response *api.BaseResponse[MerchantOpenOrClosePickUpSiteResponse]) {
	response = api.CreateResponse[MerchantOpenOrClosePickUpSiteResponse](&MerchantOpenOrClosePickUpSiteResponse{})
	return
}
