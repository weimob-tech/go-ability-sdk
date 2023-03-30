package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FulfillFindMerchantTemplateList
// @id 1237
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1237?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取商家配送模板列表)
func (client *Client) FulfillFindMerchantTemplateList(request *FulfillFindMerchantTemplateListRequest) (response *FulfillFindMerchantTemplateListResponse, err error) {
	rpcResponse := CreateFulfillFindMerchantTemplateListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FulfillFindMerchantTemplateListRequest struct {
	*api.BaseRequest

	BizStoreVid int64 `json:"bizStoreVid,omitempty"`
}

type FulfillFindMerchantTemplateListResponse struct {
	FreightTemplateList    []FulfillFindMerchantTemplateListResponseFreightTemplateList  `json:"freightTemplateList,omitempty"`
	DefaultFreightTemplate FulfillFindMerchantTemplateListResponseDefaultFreightTemplate `json:"defaultFreightTemplate,omitempty"`
}

func CreateFulfillFindMerchantTemplateListRequest() (request *FulfillFindMerchantTemplateListRequest) {
	request = &FulfillFindMerchantTemplateListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "fulfill/findMerchantTemplateList", "apigw")
	request.Method = api.POST
	return
}

func CreateFulfillFindMerchantTemplateListResponse() (response *api.BaseResponse[FulfillFindMerchantTemplateListResponse]) {
	response = api.CreateResponse[FulfillFindMerchantTemplateListResponse](&FulfillFindMerchantTemplateListResponse{})
	return
}

type FulfillFindMerchantTemplateListResponseFreightTemplateList struct {
	TemplateName string `json:"templateName,omitempty"`
	TemplateId   int64  `json:"templateId,omitempty"`
}

type FulfillFindMerchantTemplateListResponseDefaultFreightTemplate struct {
	TemplateName string `json:"templateName,omitempty"`
	TemplateId   int64  `json:"templateId,omitempty"`
}
