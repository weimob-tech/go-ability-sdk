package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FulfillFreightMerchantTemplateGetList
// @id 1432
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1432?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=查询商家配送模板列表)
func (client *Client) FulfillFreightMerchantTemplateGetList(request *FulfillFreightMerchantTemplateGetListRequest) (response *FulfillFreightMerchantTemplateGetListResponse, err error) {
	rpcResponse := CreateFulfillFreightMerchantTemplateGetListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FulfillFreightMerchantTemplateGetListRequest struct {
	*api.BaseRequest

	BasicInfo FulfillFreightMerchantTemplateGetListRequestBasicInfo `json:"basicInfo,omitempty"`
}

type FulfillFreightMerchantTemplateGetListResponse struct {
	DefaultFreightTemplate FulfillFreightMerchantTemplateGetListResponseDefaultFreightTemplate `json:"defaultFreightTemplate,omitempty"`
	FreightTemplateList    []FulfillFreightMerchantTemplateGetListResponseFreightTemplateList  `json:"freightTemplateList,omitempty"`
}

func CreateFulfillFreightMerchantTemplateGetListRequest() (request *FulfillFreightMerchantTemplateGetListRequest) {
	request = &FulfillFreightMerchantTemplateGetListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "fulfill/freight/merchant/template/getList", "apigw")
	request.Method = api.POST
	return
}

func CreateFulfillFreightMerchantTemplateGetListResponse() (response *api.BaseResponse[FulfillFreightMerchantTemplateGetListResponse]) {
	response = api.CreateResponse[FulfillFreightMerchantTemplateGetListResponse](&FulfillFreightMerchantTemplateGetListResponse{})
	return
}

type FulfillFreightMerchantTemplateGetListRequestBasicInfo struct {
	Vid int64 `json:"vid,omitempty"`
}

type FulfillFreightMerchantTemplateGetListResponseDefaultFreightTemplate struct {
	TemplateId   int64  `json:"templateId,omitempty"`
	TemplateName string `json:"templateName,omitempty"`
}

type FulfillFreightMerchantTemplateGetListResponseFreightTemplateList struct {
	TemplateId   int64  `json:"templateId,omitempty"`
	TemplateName string `json:"templateName,omitempty"`
}
