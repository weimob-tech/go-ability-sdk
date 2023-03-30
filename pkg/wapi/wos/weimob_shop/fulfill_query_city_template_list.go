package weimob_shop

import "github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"

// FulfillQueryCityTemplateList
// @id 1238
// @author WeimobCloud
// @create 2023-3-20
// @meta [](https://pt.hsmob.com/app/bosops/abilities/detail/1/1238?tab=1)
// @doc [](https://doc.weimobcloud.com/search?key=获取同城配送模板列表)
func (client *Client) FulfillQueryCityTemplateList(request *FulfillQueryCityTemplateListRequest) (response *FulfillQueryCityTemplateListResponse, err error) {
	rpcResponse := CreateFulfillQueryCityTemplateListResponse()
	err = client.DoAction(request, rpcResponse)
	if err != nil {
		return
	}
	return rpcResponse.GetData()
}

type FulfillQueryCityTemplateListRequest struct {
	*api.BaseRequest

	BizStoreVid int64 `json:"bizStoreVid,omitempty"`
}

type FulfillQueryCityTemplateListResponse struct {
	FreightTemplateList []FulfillQueryCityTemplateListResponseFreightTemplateList `json:"freightTemplateList,omitempty"`
	DefaultTemplate     FulfillQueryCityTemplateListResponseDefaultTemplate       `json:"defaultTemplate,omitempty"`
	Success             bool                                                      `json:"success,omitempty"`
}

func CreateFulfillQueryCityTemplateListRequest() (request *FulfillQueryCityTemplateListRequest) {
	request = &FulfillQueryCityTemplateListRequest{
		BaseRequest: &api.BaseRequest{},
	}
	request.InitWithApiInfo("weimob_shop", "v2.0", "fulfill/queryCityTemplateList", "apigw")
	request.Method = api.POST
	return
}

func CreateFulfillQueryCityTemplateListResponse() (response *api.BaseResponse[FulfillQueryCityTemplateListResponse]) {
	response = api.CreateResponse[FulfillQueryCityTemplateListResponse](&FulfillQueryCityTemplateListResponse{})
	return
}

type FulfillQueryCityTemplateListResponseFreightTemplateList struct {
	IsDefault    FulfillQueryCityTemplateListResponseIsDefault `json:"isDefault,omitempty"`
	TemplateName string                                        `json:"templateName,omitempty"`
	Id           int64                                         `json:"id,omitempty"`
}

type FulfillQueryCityTemplateListResponseIsDefault struct {
}

type FulfillQueryCityTemplateListResponseDefaultTemplate struct {
	IsDefault    FulfillQueryCityTemplateListResponseIsDefault2 `json:"isDefault,omitempty"`
	TemplateName string                                         `json:"templateName,omitempty"`
	Id           int64                                          `json:"id,omitempty"`
}

type FulfillQueryCityTemplateListResponseIsDefault2 struct {
}
