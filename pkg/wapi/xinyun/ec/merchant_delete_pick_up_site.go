package ec

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// MerchantDeletePickUpSite
// @id 1186
// @author WeimobCloud
// @create 2023-3-21
// @doc [](https://doc.weimobcloud.com/search?key=删除自提点)
func (client *Client) MerchantDeletePickUpSite(request *MerchantDeletePickUpSiteRequest) (response *MerchantDeletePickUpSiteResponse, err error) {
	rpcResponse := CreateMerchantDeletePickUpSiteResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type MerchantDeletePickUpSiteRequest struct {
	*api.BaseRequest

	SiteId int64 `json:"siteId,omitempty"`
}

type MerchantDeletePickUpSiteResponse struct {
}

func CreateMerchantDeletePickUpSiteRequest() (request *MerchantDeletePickUpSiteRequest) {
	request = &MerchantDeletePickUpSiteRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("ec", "1_0", "merchant/deletePickUpSite", "api")
	request.Method = api.POST
	return
}

func CreateMerchantDeletePickUpSiteResponse() (response *api.BaseResponse[MerchantDeletePickUpSiteResponse]) {
	response = api.CreateResponse[MerchantDeletePickUpSiteResponse](&MerchantDeletePickUpSiteResponse{})
	return
}
